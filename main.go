package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	baseURL           = "https://sing-box.sagernet.org"
	configurationURL  = "https://sing-box.sagernet.org/configuration/"
	outputFile        = "DOCUMENTATION.md"
	requestDelay      = 500 * time.Millisecond // Delay between requests to be respectful
	maxRetries        = 3
)

type DocPage struct {
	URL     string
	Title   string
	Content string
	Order   int // For maintaining hierarchy
}

var (
	visitedURLs = make(map[string]bool)
	verbose     = true
)

func main() {
	log.Println("🚀 Starting sing-box documentation scraper...")
	
	// Discover all documentation URLs
	log.Println("📋 Discovering documentation structure...")
	urls := discoverDocumentationURLs()
	
	log.Printf("✅ Found %d documentation pages to scrape\n", len(urls))
	
	// Fetch all documentation pages
	log.Println("📥 Fetching documentation content...")
	pages := fetchAllPages(urls)
	
	log.Printf("✅ Successfully fetched %d pages\n", len(pages))
	
	// Generate the consolidated documentation
	log.Println("📝 Generating DOCUMENTATION.md...")
	err := generateDocumentation(pages)
	if err != nil {
		log.Fatalf("❌ Failed to generate documentation: %v", err)
	}
	
	log.Printf("✅ Documentation successfully generated: %s\n", outputFile)
}

// discoverDocumentationURLs crawls the navigation to find all documentation links
func discoverDocumentationURLs() []string {
	urls := make(map[string]bool)
	
	// Start with the main configuration page
	discoverLinksRecursive(configurationURL, urls, 0)
	
	// Also check the manual section
	manualURL := "https://sing-box.sagernet.org/manual/proxy/server/"
	discoverLinksRecursive(manualURL, urls, 0)
	
	// Convert map to sorted slice
	result := make([]string, 0, len(urls))
	for url := range urls {
		result = append(result, url)
	}
	sort.Strings(result)
	
	return result
}

// discoverLinksRecursive recursively discovers documentation links
func discoverLinksRecursive(pageURL string, urls map[string]bool, depth int) {
	if depth > 3 { // Limit recursion depth
		return
	}
	
	if visitedURLs[pageURL] {
		return
	}
	visitedURLs[pageURL] = true
	
	// Add this URL to our collection
	if strings.Contains(pageURL, "sing-box.sagernet.org") {
		urls[pageURL] = true
	}
	
	// Fetch the page
	doc, err := fetchPage(pageURL)
	if err != nil {
		if verbose {
			log.Printf("⚠️  Failed to fetch %s: %v", pageURL, err)
		}
		return
	}
	
	// Find all navigation links in the sidebar
	doc.Find(".md-nav__link, .md-nav__item a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
			return
		}
		
		// Resolve relative URLs
		absoluteURL := resolveURL(pageURL, href)
		
		// Only process sing-box.sagernet.org URLs
		if strings.Contains(absoluteURL, "sing-box.sagernet.org") &&
			!strings.Contains(absoluteURL, "#") && // Skip anchors
			!urls[absoluteURL] {
			
			// Recursively discover links from this page
			discoverLinksRecursive(absoluteURL, urls, depth+1)
		}
	})
	
	time.Sleep(requestDelay)
}

// fetchPage fetches and parses an HTML page
func fetchPage(pageURL string) (*goquery.Document, error) {
	for attempt := 0; attempt < maxRetries; attempt++ {
		resp, err := http.Get(pageURL)
		if err != nil {
			if attempt < maxRetries-1 {
				time.Sleep(time.Second * time.Duration(attempt+1))
				continue
			}
			return nil, fmt.Errorf("failed to fetch page: %w", err)
		}
		
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			if attempt < maxRetries-1 {
				time.Sleep(time.Second * time.Duration(attempt+1))
				continue
			}
			return nil, fmt.Errorf("bad status code: %d", resp.StatusCode)
		}
		
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("failed to parse HTML: %w", err)
		}
		
		return doc, nil
	}
	
	return nil, fmt.Errorf("exhausted retries for %s", pageURL)
}

// fetchAllPages fetches content from all URLs
func fetchAllPages(urls []string) []DocPage {
	pages := make([]DocPage, 0, len(urls))
	
	for i, pageURL := range urls {
		if verbose {
			log.Printf("📄 [%d/%d] Fetching: %s", i+1, len(urls), pageURL)
		}
		
		doc, err := fetchPage(pageURL)
		if err != nil {
			log.Printf("⚠️  Failed to fetch %s: %v", pageURL, err)
			continue
		}
		
		// Extract the main content
		title := extractTitle(doc)
		content := extractContent(doc)
		
		if content != "" {
			pages = append(pages, DocPage{
				URL:     pageURL,
				Title:   title,
				Content: content,
				Order:   i,
			})
		}
		
		time.Sleep(requestDelay)
	}
	
	return pages
}

// extractTitle extracts the page title
func extractTitle(doc *goquery.Document) string {
	// Try to get the main heading
	title := doc.Find("h1").First().Text()
	if title == "" {
		title = doc.Find("title").First().Text()
	}
	return strings.TrimSpace(title)
}

// extractContent extracts the main documentation content
func extractContent(doc *goquery.Document) string {
	var content strings.Builder
	
	// Find the main article content
	article := doc.Find("article.md-content__inner, .md-content__inner")
	
	if article.Length() == 0 {
		return ""
	}
	
	// Process each element in the article
	article.Find("h1, h2, h3, h4, h5, h6, p, pre, code, ul, ol, table, blockquote").Each(func(i int, s *goquery.Selection) {
		tagName := goquery.NodeName(s)
		
		switch tagName {
		case "h1":
			content.WriteString(fmt.Sprintf("# %s\n\n", s.Text()))
		case "h2":
			content.WriteString(fmt.Sprintf("## %s\n\n", s.Text()))
		case "h3":
			content.WriteString(fmt.Sprintf("### %s\n\n", s.Text()))
		case "h4":
			content.WriteString(fmt.Sprintf("#### %s\n\n", s.Text()))
		case "h5":
			content.WriteString(fmt.Sprintf("##### %s\n\n", s.Text()))
		case "h6":
			content.WriteString(fmt.Sprintf("###### %s\n\n", s.Text()))
		case "p":
			text := strings.TrimSpace(s.Text())
			if text != "" {
				content.WriteString(text + "\n\n")
			}
		case "pre":
			// Code blocks
			codeText := s.Find("code").Text()
			if codeText == "" {
				codeText = s.Text()
			}
			// Try to detect language
			lang := ""
			if class, exists := s.Find("code").Attr("class"); exists {
				if strings.Contains(class, "language-") {
					lang = strings.TrimPrefix(strings.Split(class, " ")[0], "language-")
				}
			}
			content.WriteString(fmt.Sprintf("```%s\n%s\n```\n\n", lang, codeText))
		case "code":
			// Inline code (only if not inside pre)
			if s.Parent().Is("pre") {
				return
			}
			content.WriteString(fmt.Sprintf("`%s`", s.Text()))
		case "ul", "ol":
			// Lists
			s.Find("li").Each(func(j int, li *goquery.Selection) {
				prefix := "-"
				if tagName == "ol" {
					prefix = fmt.Sprintf("%d.", j+1)
				}
				content.WriteString(fmt.Sprintf("%s %s\n", prefix, strings.TrimSpace(li.Text())))
			})
			content.WriteString("\n")
		case "table":
			// Tables - simplified representation
			content.WriteString("| ")
			s.Find("thead th").Each(func(j int, th *goquery.Selection) {
				content.WriteString(th.Text() + " | ")
			})
			content.WriteString("\n|")
			s.Find("thead th").Each(func(j int, th *goquery.Selection) {
				content.WriteString(" --- |")
			})
			content.WriteString("\n")
			s.Find("tbody tr").Each(func(j int, tr *goquery.Selection) {
				content.WriteString("| ")
				tr.Find("td").Each(func(k int, td *goquery.Selection) {
					content.WriteString(td.Text() + " | ")
				})
				content.WriteString("\n")
			})
			content.WriteString("\n")
		case "blockquote":
			lines := strings.Split(s.Text(), "\n")
			for _, line := range lines {
				if trimmed := strings.TrimSpace(line); trimmed != "" {
					content.WriteString(fmt.Sprintf("> %s\n", trimmed))
				}
			}
			content.WriteString("\n")
		}
	})
	
	return content.String()
}

// generateDocumentation creates the final DOCUMENTATION.md file
func generateDocumentation(pages []DocPage) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()
	
	// Write header
	writeHeader(file)
	
	// Write table of contents
	writeTOC(file, pages)
	
	// Write all pages
	for i, page := range pages {
		writePageContent(file, page, i)
	}
	
	// Write footer
	writeFooter(file)
	
	return nil
}

// writeHeader writes the documentation header
func writeHeader(w io.Writer) {
	fmt.Fprintf(w, "# Sing-Box Configuration Documentation\n\n")
	fmt.Fprintf(w, "> **This documentation was generated automatically**\n")
	fmt.Fprintf(w, "> Generated on: %s\n", time.Now().Format("2006-01-02 15:04:05 MST"))
	fmt.Fprintf(w, "> Source: https://sing-box.sagernet.org\n\n")
	fmt.Fprintf(w, "---\n\n")
}

// writeTOC writes the table of contents
func writeTOC(w io.Writer, pages []DocPage) {
	fmt.Fprintf(w, "## Table of Contents\n\n")
	
	// Group pages by section
	sections := make(map[string][]DocPage)
	for _, page := range pages {
		section := categorizeURL(page.URL)
		sections[section] = append(sections[section], page)
	}
	
	// Write sections in order
	sectionOrder := []string{"Configuration", "Inbound", "Outbound", "Shared", "DNS", "Route", "Manual", "Other"}
	for _, section := range sectionOrder {
		if pages, ok := sections[section]; ok && len(pages) > 0 {
			fmt.Fprintf(w, "### %s\n\n", section)
			for _, page := range pages {
				anchor := strings.ToLower(strings.ReplaceAll(page.Title, " ", "-"))
				fmt.Fprintf(w, "- [%s](#%s)\n", page.Title, anchor)
			}
			fmt.Fprintf(w, "\n")
		}
	}
	
	fmt.Fprintf(w, "---\n\n")
}

// writePageContent writes a single page's content
func writePageContent(w io.Writer, page DocPage, index int) {
	fmt.Fprintf(w, "## %s\n\n", page.Title)
	fmt.Fprintf(w, "**Source URL**: <%s>\n\n", page.URL)
	fmt.Fprintf(w, "%s\n", page.Content)
	fmt.Fprintf(w, "---\n\n")
}

// writeFooter writes the documentation footer
func writeFooter(w io.Writer) {
	fmt.Fprintf(w, "\n---\n\n")
	fmt.Fprintf(w, "*End of Documentation*\n\n")
	fmt.Fprintf(w, "For the latest updates, visit: https://sing-box.sagernet.org\n")
}

// categorizeURL categorizes a URL into a section
func categorizeURL(pageURL string) string {
	switch {
	case strings.Contains(pageURL, "/inbound/"):
		return "Inbound"
	case strings.Contains(pageURL, "/outbound/"):
		return "Outbound"
	case strings.Contains(pageURL, "/shared/"):
		return "Shared"
	case strings.Contains(pageURL, "/dns/"):
		return "DNS"
	case strings.Contains(pageURL, "/route/"):
		return "Route"
	case strings.Contains(pageURL, "/manual/"):
		return "Manual"
	case strings.Contains(pageURL, "/configuration/"):
		return "Configuration"
	default:
		return "Other"
	}
}

// resolveURL resolves a relative URL against a base URL
func resolveURL(baseURL, href string) string {
	base, err := url.Parse(baseURL)
	if err != nil {
		return href
	}
	
	ref, err := url.Parse(href)
	if err != nil {
		return href
	}
	
	return base.ResolveReference(ref).String()
}
