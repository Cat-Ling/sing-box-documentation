# Sing-Box Configuration Documentation

> **This documentation was generated automatically**
> Generated on: 2026-02-09 02:07:12 UTC
> Source: https://sing-box.sagernet.org

---

## Table of Contents

### Configuration

- [Introduction](#introduction)
- [Certificate](#certificate)
- [Endpoint](#endpoint)
- [Tailscale](#tailscale)
- [WireGuard](#wireguard)
- [Experimental](#experimental)
- [Cache File](#cache-file)
- [Clash API](#clash-api)
- [V2Ray API](#v2ray-api)
- [Log](#log)
- [NTP](#ntp)
- [rule-set](#rule-set)
- [AdGuard DNS Filer](#adguard-dns-filer)
- [Headless Rule](#headless-rule)
- [Source Format](#source-format)
- [Service](#service)
- [CCM](#ccm)
- [DERP](#derp)
- [OCM](#ocm)
- [Resolved](#resolved)
- [SSM API](#ssm-api)

### Inbound

- [Inbound](#inbound)
- [AnyTLS](#anytls)
- [Direct](#direct)
- [HTTP](#http)
- [Hysteria](#hysteria)
- [Hysteria2](#hysteria2)
- [Mixed](#mixed)
- [Naive](#naive)
- [Redirect](#redirect)
- [Shadowsocks](#shadowsocks)
- [ShadowTLS](#shadowtls)
- [SOCKS](#socks)
- [TProxy](#tproxy)
- [Trojan](#trojan)
- [TUIC](#tuic)
- [Tun](#tun)
- [VLESS](#vless)
- [VMess](#vmess)

### Outbound

- [Outbound](#outbound)
- [AnyTLS](#anytls)
- [Block](#block)
- [Direct](#direct)
- [DNS](#dns)
- [HTTP](#http)
- [Hysteria](#hysteria)
- [Hysteria2](#hysteria2)
- [Naive](#naive)
- [Selector](#selector)
- [Shadowsocks](#shadowsocks)
- [ShadowTLS](#shadowtls)
- [SOCKS](#socks)
- [SSH](#ssh)
- [Tor](#tor)
- [Trojan](#trojan)
- [TUIC](#tuic)
- [URLTest](#urltest)
- [VLESS](#vless)
- [VMess](#vmess)
- [WireGuard](#wireguard)

### Shared

- [Dial Fields](#dial-fields)
- [DNS01 Challenge Fields](#dns01-challenge-fields)
- [Listen Fields](#listen-fields)
- [Multiplex](#multiplex)
- [Pre-match](#pre-match)
- [TCP Brutal](#tcp-brutal)
- [TLS](#tls)
- [UDP over TCP](#udp-over-tcp)
- [V2Ray Transport](#v2ray-transport)
- [Wi-Fi State](#wi-fi-state)

### DNS

- [DNS](#dns)
- [FakeIP](#fakeip)
- [DNS Rule](#dns-rule)
- [DNS Rule Action](#dns-rule-action)
- [DNS Server](#dns-server)
- [DHCP](#dhcp)
- [Fake IP](#fake-ip)
- [Hosts](#hosts)
- [DNS over HTTP3 (DoH3)](#dns-over-http3-(doh3))
- [DNS over HTTPS (DoH)](#dns-over-https-(doh))
- [Legacy](#legacy)
- [Local](#local)
- [DNS over QUIC (DoQ)](#dns-over-quic-(doq))
- [Resolved](#resolved)
- [Tailscale](#tailscale)
- [TCP](#tcp)
- [DNS over TLS (DoT)](#dns-over-tls-(dot))
- [UDP](#udp)

### Route

- [Route](#route)
- [GeoIP](#geoip)
- [Geosite](#geosite)
- [Route Rule](#route-rule)
- [Rule Action](#rule-action)
- [Protocol Sniff](#protocol-sniff)

### Manual

- [TunnelVision](#tunnelvision)
- [Hysteria 2](#hysteria-2)
- [Shadowsocks](#shadowsocks)
- [Trojan](#trojan)
- [Client](#client)
- [Server](#server)

### Other

- [Home](#home)
- [Change Log](#change-log)
- [Graphical Clients](#graphical-clients)
- [sing-box for Android](#sing-box-for-android)
- [Features](#features)
- [sing-box for Apple platforms](#sing-box-for-apple-platforms)
- [Features](#features)
- [General](#general)
- [Privacy policy](#privacy-policy)
- [Deprecated Feature List](#deprecated-feature-list)
- [Build from source](#build-from-source)
- [Docker](#docker)
- [Package Manager](#package-manager)
- [Migration](#migration)
- [Sponsors](#sponsors)
- [Support](#support)

---

## Home

**Source URL**: <https://sing-box.sagernet.org/>

#  Home

Welcome to the wiki page for the sing-box project.

The universal proxy platform.

## License

```
Copyright (C) 2022 by nekohasekai <[email protected]>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.

In addition, no derivative work may use the name or imply association
with this application without prior consent.

```


---

## Change Log

**Source URL**: <https://sing-box.sagernet.org/changelog/>

# Change Log

#### 1.13.0-beta.6

- Update uTLS to v1.8.2 1
- Fixes and improvements

1:

This update fixes missing padding extension for Chrome 120+ fingerprints.

Also, documentation has been updated with a warning about uTLS fingerprinting vulnerabilities.
uTLS is not recommended for censorship circumvention due to fundamental architectural limitations;
use NaiveProxy instead for TLS fingerprint resistance.

#### 1.12.17

- Update uTLS to v1.8.2 1
- Fixes and improvements

1:

This update fixes missing padding extension for Chrome 120+ fingerprints.

Also, documentation has been updated with a warning about uTLS fingerprinting vulnerabilities.
uTLS is not recommended for censorship circumvention due to fundamental architectural limitations;
use NaiveProxy instead for TLS fingerprint resistance.

#### 1.13.0-beta.5

- Fixes and improvements

#### 1.12.16

- Fixes and improvements

#### 1.13.0-beta.4

- Apple/Android: Add support for sharing configurations via QRS
- Android: Add support for resisting VPN detection via Xposed
- Update quic-go to v0.59.0
- Fixes and improvements

#### 1.13.0-beta.2

- Add bind_address_no_port option for dial fields 1
- Fixes and improvements

`bind_address_no_port`1:

Adds the Linux socket option IP_BIND_ADDRESS_NO_PORT support when explicitly binding to a source address.

`IP_BIND_ADDRESS_NO_PORT`This allows reusing the same source port for multiple connections, improving scalability for high-concurrency proxy scenarios.

See Dial Fields.

#### 1.13.0-beta.1

- Add system interface support for Tailscale endpoint 1
- Fixes and improvements

1:

Tailscale endpoint can now create a system TUN interface to handle traffic directly.

See Tailscale endpoint.

#### 1.12.15

- Fixes and improvements

#### 1.13.0-alpha.36

- Downgrade quic-go to v0.57.1
- Fixes and improvements

#### 1.13.0-alpha.35

- Add pre-match support for auto_redirect 1
- Fixes and improvements

`auto_redirect`1:

auto_redirect now allows you to bypass sing-box for connections based on routing rules.

`auto_redirect`A new rule action bypass is introduced to support this feature. When matched during pre-match, the connection will bypass sing-box and connect directly.

`bypass`This feature requires Linux with auto_redirect enabled.

`auto_redirect`See Pre-match and Rule Action.

#### 1.13.0-alpha.34

- Add Chrome Root Store certificate option 1
- Add new options for ACME DNS-01 challenge providers 2
- Add Wi-Fi state support for Linux and Windows 3
- Update naiveproxy to 143.0.7499.109
- Update quic-go to v0.58.0
- Update tailscale to v1.92.4
- Drop support for go1.23 4
- Drop support for Android 5.0 5

1:

Adds chrome as a new certificate store option alongside mozilla.
Both stores filter out China-based CA certificates.

`chrome``mozilla`See Certificate.

2:

See DNS-01 Challenge.

3:

sing-box can now monitor Wi-Fi state on Linux and Windows to enable routing rules based on wifi_ssid and wifi_bssid.

`wifi_ssid``wifi_bssid`See Wi-Fi State.

4:

Due to maintenance difficulties, sing-box 1.13.0 requires at least Go 1.24 to compile.

5:

Due to maintenance difficulties, sing-box 1.13.0 will be the last version to support Android 5.0,
and only through a separate legacy build (with -legacy-android-5 suffix).

`-legacy-android-5`For standalone binaries, the minimum Android version has been raised to Android 6.0,
since Termux requires Android 7.0 or later.

#### 1.12.14

- Fixes and improvements

#### 1.13.0-alpha.33

- Fixes and improvements

#### 1.13.0-alpha.32

- Remove certificate_public_key_sha256 option for NaiveProxy outbound 1
- Fixes and improvements

`certificate_public_key_sha256`1:

Self-signed certificates change traffic behavior significantly, which defeats the purpose of NaiveProxy's design to resist traffic analysis.
For this reason, and due to maintenance costs, there is no reason to continue supporting certificate_public_key_sha256, which was designed to simplify the use of self-signed certificates.

`certificate_public_key_sha256`#### 1.13.0-alpha.31

- Add QUIC support for NaiveProxy outbound 1
- Add QUIC congestion control option for NaiveProxy 2
- Fixes and improvements

1:

NaiveProxy outbound now supports QUIC.

See NaiveProxy outbound.

2:

NaiveProxy inbound and outbound now supports configurable QUIC congestion control algorithms, including BBR and BBRv2.

See NaiveProxy inbound and NaiveProxy outbound.

#### 1.13.0-alpha.30

- Add ECH support for NaiveProxy outbound 1
- Add tls.ech.query_server_name option 2
- Fix NaiveProxy outbound on Windows 3
- Add OpenAI Codex Multiplexer service 4
- Fixes and improvements

`tls.ech.query_server_name`1:

See NaiveProxy outbound.

2:

See TLS.

3:

Each Windows release now includes libcronet.dll.
Ensure this file is in the same directory as sing-box.exe or in a directory listed in PATH.

`libcronet.dll``sing-box.exe``PATH`4:

See OCM.

#### 1.13.0-alpha.29

- Add UDP over TCP support for naiveproxy outbound 1
- Fixes and improvements

1:

See NaiveProxy outbound.

#### 1.13.0-alpha.28

- Add naiveproxy outbound 1
- Add disable_tcp_keep_alive, tcp_keep_alive and tcp_keep_alive_interval options for dial fields 2
- Update default TCP keep-alive initial period from 10 minutes to 5 minutes
- Update quic-go to v0.57.1
- Fixes and improvements

`disable_tcp_keep_alive``tcp_keep_alive``tcp_keep_alive_interval`1:

Only available on Apple platforms, Android, Windows and some Linux architectures.

See NaiveProxy outbound.

2:

See Dial Fields.

- Unfortunately, for non-technical reasons, we are currently unable to notarize the standalone version of the macOS client:
because system extensions require signatures to function, we have had to temporarily halt its release.

We plan to fix the App Store release issue and launch a new standalone desktop client, but until then,
only clients on TestFlight will be available (unless you have an Apple Developer Program and compile from source code).

#### 1.12.13

- Fix naive inbound
- Fixes and improvements

Unfortunately, for non-technical reasons, we are currently unable to notarize the standalone version of the macOS client:
because system extensions require signatures to function, we have had to temporarily halt its release.

We plan to fix the App Store release issue and launch a new standalone desktop client, but until then,
only clients on TestFlight will be available (unless you have an Apple Developer Program and compile from source code).

#### 1.12.12

- Fixes and improvements

#### 1.13.0-alpha.26

- Update quic-go to v0.55.0
- Fix memory leak in hysteria2
- Fixes and improvements

#### 1.12.11

- Fixes and improvements

#### 1.13.0-alpha.24

- Add Claude Code Multiplexer service 1
- Fixes and improvements

1:

CCM (Claude Code Multiplexer) service allows you to access your local Claude Code subscription remotely through custom tokens, eliminating the need for OAuth authentication on remote clients.

See CCM.

#### 1.13.0-alpha.23

- Fix compatibility with MPTCP 1
- Fixes and improvements

1:

auto_redirect now rejects MPTCP connections by default to fix compatibility issues,
but you can change it to bypass the sing-box via the new exclude_mptcp option.

`auto_redirect``exclude_mptcp`See TUN.

#### 1.13.0-alpha.22

- Update uTLS to v1.8.1 1
- Fixes and improvements

1:

This update fixes an critical issue that could cause simulated Chrome fingerprints to be detected,
see https://github.com/refraction-networking/utls/pull/375.

#### 1.12.10

- Update uTLS to v1.8.1 1
- Fixes and improvements

1:

This update fixes an critical issue that could cause simulated Chrome fingerprints to be detected,
see https://github.com/refraction-networking/utls/pull/375.

#### 1.13.0-alpha.21

- Fix missing mTLS support in client options 1
- Fixes and improvements

See TLS.

#### 1.12.9

- Fixes and improvements

#### 1.13.0-alpha.16

- Add curve preferences, pinned public key SHA256 and mTLS for TLS options 1
- Fixes and improvements

See TLS.

#### 1.13.0-alpha.15

- Update quic-go to v0.54.0
- Update gVisor to v20250811
- Update Tailscale to v1.86.5
- Fixes and improvements

#### 1.12.8

- Fixes and improvements

#### 1.13.0-alpha.11

- Fixes and improvements

#### 1.12.5

- Fixes and improvements

#### 1.13.0-alpha.10

- Improve kTLS support 1
- Fixes and improvements

1:

kTLS is now compatible with custom TLS implementations other than uTLS.

#### 1.12.4

- Fixes and improvements

#### 1.12.3

- Fixes and improvements

#### 1.12.2

- Fixes and improvements

#### 1.12.1

- Fixes and improvements

#### 1.12.0

- Refactor DNS servers 1
- Add domain resolver options2
- Add TLS fragment/record fragment support to route options and outbound TLS options 3
- Add certificate options 4
- Add Tailscale endpoint and DNS server 5
- Drop support for go1.22 6
- Add AnyTLS protocol 7
- Migrate to stdlib ECH implementation 8
- Add NTP sniffer 9
- Add wildcard SNI support for ShadowTLS inbound 10
- Improve auto_redirect 11
- Add control options for listeners 12
- Add DERP service 13
- Add Resolved service and DNS server 14
- Add SSM API service 15
- Add loopback address support for tun 16
- Improve tun performance on Apple platforms 17
- Update quic-go to v0.52.0
- Update gVisor to 20250319.0
- Update the status of graphical clients in stores 18

`auto_redirect`1:

DNS servers are refactored for better performance and scalability.

See DNS server.

For migration, see Migrate to new DNS server formats.

Compatibility for old formats will be removed in sing-box 1.14.0.

2:

Legacy outbound DNS rules are deprecated
and can be replaced by the new domain_resolver option.

`outbound``domain_resolver`See Dial Fields and
Route.

For migration,
see Migrate outbound DNS rule items to domain resolver.

3:

See Route Action and TLS.

4:

New certificate options allow you to manage the default list of trusted X509 CA certificates.

For the system certificate list, fixed Go not reading Android trusted certificates correctly.

You can also use the Mozilla Included List instead, or add trusted certificates yourself.

See Certificate.

5:

See Tailscale.

6:

Due to maintenance difficulties, sing-box 1.12.0 requires at least Go 1.23 to compile.

For Windows 7 users, legacy binaries now continue to compile with Go 1.23 and patches
from MetaCubeX/go.

7:

The new AnyTLS protocol claims to mitigate TLS proxy traffic characteristics and comes with a new multiplexing scheme.

See AnyTLS Inbound and AnyTLS Outbound.

8:

See TLS.

The build tag with_ech is no longer needed and has been removed.

`with_ech`9:

See Protocol Sniff.

10:

See ShadowTLS.

11:

Now auto_redirect fixes compatibility issues between tun and Docker bridge networks,
see Tun.

`auto_redirect`12:

You can now set bind_interface, routing_mark and reuse_addr in Listen Fields.

`bind_interface``routing_mark``reuse_addr`See Listen Fields.

13:

DERP service is a Tailscale DERP server, similar to derper.

See DERP Service.

14:

Resolved service is a fake systemd-resolved DBUS service to receive DNS settings from other programs
(e.g. NetworkManager) and provide DNS resolution.

See Resolved Service and Resolved DNS Server.

15:

SSM API service is a RESTful API server for managing Shadowsocks servers.

See SSM API Service.

16:

TUN now implements SideStore's StosVPN.

See Tun.

17:

We have significantly improved the performance of tun inbound on Apple platforms, especially in the gVisor stack.

The following data was tested
using tun_bench on M4 MacBook pro.

| Version | Stack | MTU | Upload | Download | 
| --- | --- | --- | --- | --- |
| 1.11.15 | gvisor | 1500 | 852M | 2.57G | 
| 1.12.0-rc.4 | gvisor | 1500 | 2.90G | 4.68G | 
| 1.11.15 | gvisor | 4064 | 2.31G | 6.34G | 
| 1.12.0-rc.4 | gvisor | 4064 | 7.54G | 12.2G | 
| 1.11.15 | gvisor | 65535 | 27.6G | 18.1G | 
| 1.12.0-rc.4 | gvisor | 65535 | 39.8G | 34.7G | 
| 1.11.15 | system | 1500 | 664M | 706M | 
| 1.12.0-rc.4 | system | 1500 | 2.44G | 2.51G | 
| 1.11.15 | system | 4064 | 1.88G | 1.94G | 
| 1.12.0-rc.4 | system | 4064 | 6.45G | 6.27G | 
| 1.11.15 | system | 65535 | 26.2G | 17.4G | 
| 1.12.0-rc.4 | system | 65535 | 17.6G | 21.0G | 

18:

We continue to experience issues updating our sing-box apps on the App Store and Play Store.
Until we rewrite and resubmit the apps, they are considered irrecoverable.
Therefore, after this release, we will not be repeating this notice unless there is new information.

### 1.11.15

- Fixes and improvements

We are temporarily unable to update sing-box apps on the App Store because the reviewer mistakenly found that we
violated the rules (TestFlight users are not affected).

#### 1.12.0-beta.32

- Improve tun performance on Apple platforms 1
- Fixes and improvements

1:

We have significantly improved the performance of tun inbound on Apple platforms, especially in the gVisor stack.

### 1.11.14

- Fixes and improvements

We are temporarily unable to update sing-box apps on the App Store because the reviewer mistakenly found that we
violated the rules (TestFlight users are not affected).

#### 1.12.0-beta.24

- Allow tls_fragment and tls_record_fragment to be enabled together 1
- Also add fragment options for TLS client configuration 2
- Fixes and improvements

`tls_fragment``tls_record_fragment`1:

For debugging only, it is recommended to disable if record fragmentation works.

See Route Action.

2:

See TLS.

#### 1.12.0-beta.23

- Add loopback address support for tun 1
- Add cache support for ssm-api 2
- Fixes and improvements

1:

TUN now implements SideStore's StosVPN.

See Tun.

2:

See SSM API Service.

#### 1.12.0-beta.21

- Fix missing home option for DERP service 1
- Fixes and improvements

`home`1:

You can now choose what the DERP home page shows, just like with derper's -home flag.

`-home`See DERP.

### 1.11.13

- Fixes and improvements

We are temporarily unable to update sing-box apps on the App Store because the reviewer mistakenly found that we
violated the rules (TestFlight users are not affected).

#### 1.12.0-beta.17

- Update quic-go to v0.52.0
- Fixes and improvements

#### 1.12.0-beta.15

- Add DERP service 1
- Add Resolved service and DNS server 2
- Add SSM API service 3
- Fixes and improvements

1:

DERP service is a Tailscale DERP server, similar to derper.

See DERP Service.

2:

Resolved service is a fake systemd-resolved DBUS service to receive DNS settings from other programs
(e.g. NetworkManager) and provide DNS resolution.

See Resolved Service and Resolved DNS Server.

3:

SSM API service is a RESTful API server for managing Shadowsocks servers.

See SSM API Service.

### 1.11.11

- Fixes and improvements

We are temporarily unable to update sing-box apps on the App Store because the reviewer mistakenly found that we
violated the rules (TestFlight users are not affected).

#### 1.12.0-beta.13

- Add TLS record fragment route options 1
- Add missing accept_routes option for Tailscale 2
- Fixes and improvements

`accept_routes`1:

See Route Action.

2:

See Tailscale.

#### 1.12.0-beta.10

- Add control options for listeners 1
- Fixes and improvements

1:

You can now set bind_interface, routing_mark and reuse_addr in Listen Fields.

`bind_interface``routing_mark``reuse_addr`See Listen Fields.

### 1.11.10

- Undeprecate the block outbound 1
- Fixes and improvements

`block`1:

Since we don’t have a replacement for using the block outbound in selectors yet,
we decided to temporarily undeprecate the block outbound until a replacement is available in the future.

`block``block`We are temporarily unable to update sing-box apps on the App Store because the reviewer mistakenly found that we
violated the rules (TestFlight users are not affected).

#### 1.12.0-beta.9

- Update quic-go to v0.51.0
- Fixes and improvements

### 1.11.9

- Fixes and improvements

We are temporarily unable to update sing-box apps on the App Store because the reviewer mistakenly found that we
violated the rules (TestFlight users are not affected).

#### 1.12.0-beta.5

- Fixes and improvements

### 1.11.8

- Improve auto_redirect 1
- Fixes and improvements

`auto_redirect`1:

Now auto_redirect fixes compatibility issues between TUN and Docker bridge networks,
see Tun.

`auto_redirect`We are temporarily unable to update sing-box apps on the App Store because the reviewer mistakenly found that we
violated the rules (TestFlight users are not affected).

#### 1.12.0-beta.3

- Fixes and improvements

### 1.11.7

- Fixes and improvements

We are temporarily unable to update sing-box apps on the App Store because the reviewer mistakenly found that we
violated the rules (TestFlight users are not affected).

#### 1.12.0-beta.1

- Fixes and improvements

1:

Now auto_redirect fixes compatibility issues between tun and Docker bridge networks,
see Tun.

`auto_redirect`### 1.11.6

- Fixes and improvements

We are temporarily unable to update sing-box apps on the App Store because the reviewer mistakenly found that we
violated the rules (TestFlight users are not affected).

#### 1.12.0-alpha.19

- Update gVisor to 20250319.0
- Fixes and improvements

#### 1.12.0-alpha.18

- Add wildcard SNI support for ShadowTLS inbound 1
- Fixes and improvements

1:

See ShadowTLS.

#### 1.12.0-alpha.17

- Add NTP sniffer 1
- Fixes and improvements

1:

See Protocol Sniff.

#### 1.12.0-alpha.16

- Update domain_resolver behavior 1
- Fixes and improvements

`domain_resolver`1:

route.default_domain_resolver or outbound.domain_resolver is now optional when only one DNS server is configured.

`route.default_domain_resolver``outbound.domain_resolver`See Dial Fields.

### 1.11.5

- Fixes and improvements

We are temporarily unable to update sing-box apps on the App Store because the reviewer mistakenly found that we
violated the rules (TestFlight users are not affected).

#### 1.12.0-alpha.13

- Move predefined DNS server to DNS rule action 1
- Fixes and improvements

`predefined`1:

See DNS Rule Action.

### 1.11.4

- Fixes and improvements

#### 1.12.0-alpha.11

- Fixes and improvements

#### 1.12.0-alpha.10

- Add AnyTLS protocol 1
- Improve resolve route action 2
- Migrate to stdlib ECH implementation 3
- Fixes and improvements

`resolve`1:

The new AnyTLS protocol claims to mitigate TLS proxy traffic characteristics and comes with a new multiplexing scheme.

See AnyTLS Inbound and AnyTLS Outbound.

2:

resolve route action now accepts disable_cache and other options like in DNS route actions,
see Route Action.

`resolve``disable_cache`3:

See TLS.

The build tag with_ech is no longer needed and has been removed.

`with_ech`#### 1.12.0-alpha.7

- Add Tailscale DNS server 1
- Fixes and improvements

1:

See Tailscale.

#### 1.12.0-alpha.6

- Add Tailscale endpoint 1
- Drop support for go1.22 2
- Fixes and improvements

1:

See Tailscale.

2:

Due to maintenance difficulties, sing-box 1.12.0 requires at least Go 1.23 to compile.

For Windows 7 users, legacy binaries now continue to compile with Go 1.23 and patches
from MetaCubeX/go.

### 1.11.3

- Fixes and improvements

This version overwrites 1.11.2, as incorrect binaries were released due to a bug in the continuous integration
process.

#### 1.12.0-alpha.5

- Fixes and improvements

### 1.11.1

- Fixes and improvements

#### 1.12.0-alpha.2

- Update quic-go to v0.49.0
- Fixes and improvements

#### 1.12.0-alpha.1

- Refactor DNS servers 1
- Add domain resolver options2
- Add TLS fragment route options 3
- Add certificate options 4

1:

DNS servers are refactored for better performance and scalability.

See DNS server.

For migration, see Migrate to new DNS server formats.

Compatibility for old formats will be removed in sing-box 1.14.0.

2:

Legacy outbound DNS rules are deprecated
and can be replaced by the new domain_resolver option.

`outbound``domain_resolver`See Dial Fields and
Route.

For migration,
see Migrate outbound DNS rule items to domain resolver.

3:

The new TLS fragment route options allow you to fragment TLS handshakes to bypass firewalls.

This feature is intended to circumvent simple firewalls based on plaintext packet matching, and should not be used
to circumvent real censorship.

Since it is not designed for performance, it should not be applied to all connections, but only to server names that are
known to be blocked.

See Route Action.

4:

New certificate options allow you to manage the default list of trusted X509 CA certificates.

For the system certificate list, fixed Go not reading Android trusted certificates correctly.

You can also use the Mozilla Included List instead, or add trusted certificates yourself.

See Certificate.

### 1.11.0

Important changes since 1.10:

- Introducing rule actions 1
- Improve tun compatibility 3
- Merge route options to route actions 4
- Add network_type, network_is_expensive and network_is_constrainted rule items 5
- Add multi network dialing 6
- Add cache_capacity DNS option 7
- Add override_address and override_port route options 8
- Upgrade WireGuard outbound to endpoint 9
- Add UDP GSO support for WireGuard
- Make GSO adaptive 10
- Add UDP timeout route option 11
- Add more masquerade options for hysteria2 12
- Add rule-set merge command
- Add port hopping support for Hysteria2 13
- Hysteria2 ignore_client_bandwidth behavior update 14

`network_type``network_is_expensive``network_is_constrainted``cache_capacity``override_address``override_port``rule-set merge``ignore_client_bandwidth`1:

New rule actions replace legacy inbound fields and special outbound fields,
and can be used for pre-matching 2.

See Rule,
Rule Action,
DNS Rule and
DNS Rule Action.

For migration, see
Migrate legacy special outbounds to rule actions,
Migrate legacy inbound fields to rule actions
and Migrate legacy DNS route options to rule actions.

2:

Similar to Surge's pre-matching.

Specifically, new rule actions allow you to reject connections with
TCP RST (for TCP connections) and ICMP port unreachable (for UDP packets)
before connection established to improve tun's compatibility.

See Rule Action.

3:

When gvisor tun stack is enabled, even if the request passes routing,
if the outbound connection establishment fails,
the connection still does not need to be established and a TCP RST is replied.

`gvisor`4:

Route options in DNS route actions will no longer be considered deprecated,
see DNS Route Action.

Also, now udp_disable_domain_unmapping and udp_connect can also be configured in route action,
see Route Action.

`udp_disable_domain_unmapping``udp_connect`5:

When using in graphical clients, new routing rule items allow you to match on
network type (WIFI, cellular, etc.), whether the network is expensive, and whether Low Data Mode is enabled.

See Route Rule, DNS Route Rule
and Headless Rule.

6:

Similar to Surge's strategy.

New options allow you to connect using multiple network interfaces,
prefer or only use one type of interface,
and configure a timeout to fallback to other interfaces.

See Dial Fields,
Rule Action
and Route.

7:

See DNS.

8:

See Rule Action and
Migrate destination override fields to route options.

9:

The new WireGuard endpoint combines inbound and outbound capabilities,
and the old outbound will be removed in sing-box 1.13.0.

See Endpoint, WireGuard Endpoint
and Migrate WireGuard outbound fields to route options.

10:

For WireGuard outbound and endpoint, GSO will be automatically enabled when available,
see WireGuard Outbound.

For TUN, GSO has been removed,
see Deprecated.

11:

See Rule Action.

12:

See Hysteria2.

13:

See Hysteria2.

14:

When up_mbps and down_mbps are set, ignore_client_bandwidth instead denies clients from using BBR CC.

`up_mbps``down_mbps``ignore_client_bandwidth`### 1.10.7

- Fixes and improvements

#### 1.11.0-beta.20

- Hysteria2 ignore_client_bandwidth behavior update 1
- Fixes and improvements

`ignore_client_bandwidth`1:

When up_mbps and down_mbps are set, ignore_client_bandwidth instead denies clients from using BBR CC.

`up_mbps``down_mbps``ignore_client_bandwidth`See Hysteria2.

#### 1.11.0-beta.17

- Add port hopping support for Hysteria2 1
- Fixes and improvements

1:

See Hysteria2.

#### 1.11.0-beta.14

- Allow adding route (exclude) address sets to routes 1
- Fixes and improvements

1:

When auto_redirect is not enabled, directly add route[_exclude]_address_set
to tun routes (equivalent to route[_exclude]_address).

`auto_redirect``route[_exclude]_address_set``route[_exclude]_address`Note that it doesn't work on the Android graphical client due to
the Android VpnService not being able to handle a large number of routes (DeadSystemException),
but otherwise it works fine on all command line clients and Apple platforms.

See route_address_set and
route_exclude_address_set.

#### 1.11.0-beta.12

- Add rule-set merge command
- Fixes and improvements

`rule-set merge`#### 1.11.0-beta.3

- Add more masquerade options for hysteria2 1
- Fixes and improvements

1:

See Hysteria2.

#### 1.11.0-alpha.25

- Update quic-go to v0.48.2
- Fixes and improvements

#### 1.11.0-alpha.22

- Add UDP timeout route option 1
- Fixes and improvements

1:

See Rule Action.

#### 1.11.0-alpha.20

- Add UDP GSO support for WireGuard
- Make GSO adaptive 1

1:

For WireGuard outbound and endpoint, GSO will be automatically enabled when available,
see WireGuard Outbound.

For TUN, GSO has been removed,
see Deprecated.

#### 1.11.0-alpha.19

- Upgrade WireGuard outbound to endpoint 1
- Fixes and improvements

1:

The new WireGuard endpoint combines inbound and outbound capabilities,
and the old outbound will be removed in sing-box 1.13.0.

See Endpoint, WireGuard Endpoint
and Migrate WireGuard outbound fields to route options.

### 1.10.2

- Add deprecated warnings
- Fix proxying websocket connections in HTTP/mixed inbounds
- Fixes and improvements

#### 1.11.0-alpha.18

- Fixes and improvements

#### 1.11.0-alpha.16

- Add cache_capacity DNS option 1
- Add override_address and override_port route options 2
- Fixes and improvements

`cache_capacity``override_address``override_port`1:

See DNS.

2:

See Rule Action and
Migrate destination override fields to route options.

#### 1.11.0-alpha.15

- Improve multi network dialing 1
- Fixes and improvements

1:

New options allow you to configure the network strategy flexibly.

See Dial Fields,
Rule Action
and Route.

#### 1.11.0-alpha.14

- Add multi network dialing 1
- Fixes and improvements

1:

Similar to Surge's strategy.

New options allow you to connect using multiple network interfaces,
prefer or only use one type of interface,
and configure a timeout to fallback to other interfaces.

See Dial Fields,
Rule Action
and Route.

#### 1.11.0-alpha.13

- Fixes and improvements

#### 1.11.0-alpha.12

- Merge route options to route actions 1
- Add network_type, network_is_expensive and network_is_constrainted rule items 2
- Fixes and improvements

`network_type``network_is_expensive``network_is_constrainted`1:

Route options in DNS route actions will no longer be considered deprecated,
see DNS Route Action.

Also, now udp_disable_domain_unmapping and udp_connect can also be configured in route action,
see Route Action.

`udp_disable_domain_unmapping``udp_connect`2:

When using in graphical clients, new routing rule items allow you to match on
network type (WIFI, cellular, etc.), whether the network is expensive, and whether Low Data Mode is enabled.

See Route Rule, DNS Route Rule
and Headless Rule.

#### 1.11.0-alpha.9

- Improve tun compatibility 1
- Fixes and improvements

1:

When gvisor tun stack is enabled, even if the request passes routing,
if the outbound connection establishment fails,
the connection still does not need to be established and a TCP RST is replied.

`gvisor`#### 1.11.0-alpha.7

- Introducing rule actions 1

1:

New rule actions replace legacy inbound fields and special outbound fields,
and can be used for pre-matching 2.

See Rule,
Rule Action,
DNS Rule and
DNS Rule Action.

For migration, see
Migrate legacy special outbounds to rule actions,
Migrate legacy inbound fields to rule actions
and Migrate legacy DNS route options to rule actions.

2:

Similar to Surge's pre-matching.

Specifically, new rule actions allow you to reject connections with
TCP RST (for TCP connections) and ICMP port unreachable (for UDP packets)
before connection established to improve tun's compatibility.

See Rule Action.

#### 1.11.0-alpha.6

- Update quic-go to v0.48.1
- Set gateway for tun correctly
- Fixes and improvements

#### 1.11.0-alpha.2

- Add warnings for usage of deprecated features
- Fixes and improvements

#### 1.11.0-alpha.1

- Update quic-go to v0.48.0
- Fixes and improvements

### 1.10.1

- Fixes and improvements

### 1.10.0

Important changes since 1.9:

- Introducing auto-redirect 1
- Add AdGuard DNS Filter support 2
- TUN address fields are merged 3
- Add custom options for auto-route and auto-redirect 4
- Drop support for go1.18 and go1.19 5
- Add tailing comma support in JSON configuration
- Improve sniffers 6
- Add new inline rule-set type 7
- Add access control options for Clash API 8
- Add rule_set_ip_cidr_accept_empty DNS address filter rule item 9
- Add auto reload support for local rule-set
- Update fsnotify usages 10
- Add IP address support for rule-set match command
- Add rule-set decompile command
- Add process_path_regex rule item
- Update uTLS to v1.6.7 11
- Optimize memory usages of rule-sets 12

`auto-route``auto-redirect``inline``rule_set_ip_cidr_accept_empty``rule-set match``rule-set decompile``process_path_regex`1:

The new auto-redirect feature allows TUN to automatically
configure connection redirection to improve proxy performance.

When auto-redirect is enabled, new route address set options will allow you to
automatically configure destination IP CIDR rules from a specified rule set to the firewall.

Specified or unspecified destinations will bypass the sing-box routes to get better performance
(for example, keep hardware offloading of direct traffics on the router).

See TUN.

2:

The new feature allows you to use AdGuard DNS Filter lists in a sing-box without AdGuard Home.

See AdGuard DNS Filter.

3:

See Migration.

4:

See iproute2_table_index,
iproute2_rule_index,
auto_redirect_input_mark and
auto_redirect_output_mark.

5:

Due to maintenance difficulties, sing-box 1.10.0 requires at least Go 1.20 to compile.

6:

BitTorrent, DTLS, RDP, SSH sniffers are added.

Now the QUIC sniffer can correctly extract the server name from Chromium requests and
can identify common QUIC clients, including
Chromium, Safari, Firefox, quic-go (including uquic disguised as Chrome).

7:

The new rule-set type inline (which also becomes the default type)
allows you to write headless rules directly without creating a rule-set file.

8:

With new access control options, not only can you allow Clash dashboards
to access the Clash API on your local network,
you can also manually limit the websites that can access the API instead of allowing everyone.

See Clash API.

9:

See DNS Rule.

10:

sing-box now uses fsnotify correctly and will not cancel watching
if the target file is deleted or recreated via rename (e.g. mv).

`mv`This affects all path options that support reload, including
tls.certificate_path, tls.key_path, tls.ech.key_path and rule_set.path.

`tls.certificate_path``tls.key_path``tls.ech.key_path``rule_set.path`11:

Some legacy chrome fingerprints have been removed and will fallback to chrome,
see utls.

12:

See Source Format.

### 1.9.7

- Fixes and improvements

#### 1.10.0-beta.11

- Update uTLS to v1.6.7 1

1:

Some legacy chrome fingerprints have been removed and will fallback to chrome,
see utls.

#### 1.10.0-beta.10

- Add process_path_regex rule item
- Fixes and improvements

`process_path_regex`The macOS standalone versions of sing-box (>=1.9.5/<1.10.0-beta.11) now silently fail and require manual granting of
the Full Disk Access permission to system extension to start, probably due to Apple's changed security policy. We
will prompt users about this in feature versions.

### 1.9.6

- Fixes and improvements

### 1.9.5

- Update quic-go to v0.47.0
- Fix direct dialer not resolving domain
- Fix no error return when empty DNS cache retrieved
- Fix build with go1.23
- Fix stream sniffer
- Fix bad redirect in clash-api
- Fix wireguard events chan leak
- Fix cached conn eats up read deadlines
- Fix disconnected interface selected as default in windows
- Update Bundle Identifiers for Apple platform clients 1

1:

See Migration.

We are still working on getting all sing-box apps back on the App Store, which should be completed within a week
(SFI on the App Store and others on TestFlight are already available).

#### 1.10.0-beta.8

- Fixes and improvements

With the help of a netizen, we are in the process of getting sing-box apps back on the App Store, which should be
completed within a month (TestFlight is already available).

#### 1.10.0-beta.7

- Update quic-go to v0.47.0
- Fixes and improvements

#### 1.10.0-beta.6

- Add RDP sniffer
- Fixes and improvements

#### 1.10.0-beta.5

- Add PNA support for Clash API
- Fixes and improvements

#### 1.10.0-beta.3

- Add SSH sniffer
- Fixes and improvements

#### 1.10.0-beta.2

- Build with go1.23
- Fixes and improvements

### 1.9.4

- Update quic-go to v0.46.0
- Update Hysteria2 BBR congestion control
- Filter HTTPS ipv4hint/ipv6hint with domain strategy
- Fix crash on Android when using process rules
- Fix non-IP queries accepted by address filter rules
- Fix UDP server for shadowsocks AEAD multi-user inbounds
- Fix default next protos for v2ray QUIC transport
- Fix default end value of port range configuration options
- Fix reset v2ray transports
- Fix panic caused by rule-set generation of duplicate keys for domain_suffix
- Fix UDP connnection leak when sniffing
- Fixes and improvements

`domain_suffix`Due to problems with our Apple developer account,
sing-box apps on Apple platforms are temporarily unavailable for download or update.
If your company or organization is willing to help us return to the App Store,
please contact us.

#### 1.10.0-alpha.29

- Update quic-go to v0.46.0
- Fixes and improvements

#### 1.10.0-alpha.25

- Add AdGuard DNS Filter support 1

1:

The new feature allows you to use AdGuard DNS Filter lists in a sing-box without AdGuard Home.

See AdGuard DNS Filter.

#### 1.10.0-alpha.23

- Add Chromium support for QUIC sniffer
- Add client type detect support for QUIC sniffer 1
- Fixes and improvements

1:

Now the QUIC sniffer can correctly extract the server name from Chromium requests and
can identify common QUIC clients, including
Chromium, Safari, Firefox, quic-go (including uquic disguised as Chrome).

See Protocol Sniff and Route Rule.

#### 1.10.0-alpha.22

- Optimize memory usages of rule-sets 1
- Fixes and improvements

1:

See Source Format.

#### 1.10.0-alpha.20

- Add DTLS sniffer
- Fixes and improvements

#### 1.10.0-alpha.19

- Add rule-set decompile command
- Add IP address support for rule-set match command
- Fixes and improvements

`rule-set decompile``rule-set match`#### 1.10.0-alpha.18

- Add new inline rule-set type 1
- Add auto reload support for local rule-set
- Update fsnotify usages 2
- Fixes and improvements

`inline`1:

The new rule-set type inline (which also becomes the default type)
allows you to write headless rules directly without creating a rule-set file.

2:

sing-box now uses fsnotify correctly and will not cancel watching
if the target file is deleted or recreated via rename (e.g. mv).

`mv`This affects all path options that support reload, including
tls.certificate_path, tls.key_path, tls.ech.key_path and rule_set.path.

`tls.certificate_path``tls.key_path``tls.ech.key_path``rule_set.path`#### 1.10.0-alpha.17

- Some chaotic changes 1
- rule_set_ipcidr_match_source rule items are renamed 2
- Add rule_set_ip_cidr_accept_empty DNS address filter rule item 3
- Update quic-go to v0.45.1
- Fixes and improvements

`rule_set_ipcidr_match_source``rule_set_ip_cidr_accept_empty`1:

Something may be broken, please actively report problems with this version.

2:

rule_set_ipcidr_match_source route and DNS rule items are renamed to
rule_set_ip_cidr_match_source and will be remove in sing-box 1.11.0.

`rule_set_ipcidr_match_source``rule_set_ip_cidr_match_source`3:

See DNS Rule.

#### 1.10.0-alpha.16

- Add custom options for auto-route and auto-redirect 1
- Fixes and improvements

`auto-route``auto-redirect`1:

See iproute2_table_index,
iproute2_rule_index,
auto_redirect_input_mark and
auto_redirect_output_mark.

#### 1.10.0-alpha.13

- TUN address fields are merged 1
- Add route address set support for auto-redirect 2

1:

See Migration.

2:

The new feature will allow you to configure the destination IP CIDR rules
in the specified rule-sets to the firewall automatically.

Specified or unspecified destinations will bypass the sing-box routes to get better performance
(for example, keep hardware offloading of direct traffics on the router).

See route_address_set
and route_exclude_address_set.

#### 1.10.0-alpha.12

- Fix auto-redirect not configuring nftables forward chain correctly
- Fixes and improvements

### 1.9.3

- Fixes and improvements

#### 1.10.0-alpha.10

- Fixes and improvements

### 1.9.2

- Fixes and improvements

#### 1.10.0-alpha.8

- Drop support for go1.18 and go1.19 1
- Update quic-go to v0.45.0
- Update Hysteria2 BBR congestion control
- Fixes and improvements

1:

Due to maintenance difficulties, sing-box 1.10.0 requires at least Go 1.20 to compile.

### 1.9.1

- Fixes and improvements

#### 1.10.0-alpha.7

- Fixes and improvements

#### 1.10.0-alpha.5

- Improve auto-redirect 1

1:

nftables support and DNS hijacking has been added.

Tun inbounds with auto_route and auto_redirect now works as expected on routers without intervention.

`auto_route``auto_redirect`#### 1.10.0-alpha.4

- Fix auto-redirect 1
- Improve auto-route on linux 2

1:

Tun inbounds with auto_route and auto_redirect now works as expected on routers.

`auto_route``auto_redirect`2:

Tun inbounds with auto_route and strict_route now works as expected on routers and servers,
but the usages of exclude_interface need to be updated.

`auto_route``strict_route`#### 1.10.0-alpha.2

- Move auto-redirect to Tun 1
- Fixes and improvements

1:

Linux support are added.

See Tun.

#### 1.10.0-alpha.1

- Add tailing comma support in JSON configuration
- Add simple auto-redirect for Android 1
- Add BitTorrent sniffer 2

1:

It allows you to use redirect inbound in the sing-box Android client
and automatically configures IPv4 TCP redirection via su.

This may alleviate the symptoms of some OCD patients who think that
redirect can effectively save power compared to the system HTTP Proxy.

See Redirect.

2:

See Protocol Sniff.

### 1.9.0

- Fixes and improvements

Important changes since 1.8:

- domain_suffix behavior update 1
- process_path format update on Windows 2
- Add address filter DNS rule items 3
- Add support for client-subnet DNS options 4
- Add rejected DNS response cache support 5
- Add bypass_domain and search_domain platform HTTP proxy options 6
- Fix missing rule_set_ipcidr_match_source item in DNS rules 7
- Handle Windows power events
- Always disable cache for fake-ip DNS transport if dns.independent_cache disabled
- Improve DNS truncate behavior
- Update Hysteria protocol
- Update quic-go to v0.43.1
- Update gVisor to 20240422.0
- Mitigating TunnelVision attacks 8

`domain_suffix``process_path``client-subnet``bypass_domain``search_domain``rule_set_ipcidr_match_source``dns.independent_cache`1:

See Migration.

2:

See Migration.

3:

The new DNS feature allows you to more precisely bypass Chinese websites via DNS leaks. Do not use plain local DNS
if using this method.

See Address Filter Fields.

Client example updated.

4:

See DNS, DNS Server and DNS Rules.

Since this feature makes the scenario mentioned in alpha.1 no longer leak DNS requests,
the Client example has been updated.

`alpha.1`5:

The new feature allows you to cache the check results of
Address filter DNS rule items until expiration.

6:

See TUN inbound.

7:

See DNS Rule.

8:

See TunnelVision.

#### 1.9.0-rc.22

- Fixes and improvements

#### 1.9.0-rc.20

- Prioritize *_route_address in linux auto-route
- Fix *_route_address in darwin auto-route

`*_route_address``*_route_address`#### 1.8.14

- Fix hysteria2 panic
- Fixes and improvements

#### 1.9.0-rc.18

- Add custom prefix support in EDNS0 client subnet options
- Fix hysteria2 crash
- Fix store_rdrc corrupted
- Update quic-go to v0.43.1
- Fixes and improvements

`store_rdrc`#### 1.9.0-rc.16

- Mitigating TunnelVision attacks 1
- Fixes and improvements

1:

See TunnelVision.

#### 1.9.0-rc.15

- Fixes and improvements

#### 1.8.13

- Fix fake-ip mapping
- Fixes and improvements

#### 1.9.0-rc.14

- Fixes and improvements

#### 1.9.0-rc.13

- Update Hysteria protocol
- Update quic-go to v0.43.0
- Update gVisor to 20240422.0
- Fixes and improvements

#### 1.8.12

- Now we have official APT and DNF repositories 1
- Fix packet MTU for QUIC protocols
- Fixes and improvements

1:

Including stable and beta versions, see https://sing-box.sagernet.org/installation/package-manager/

#### 1.9.0-rc.11

- Fixes and improvements

#### 1.8.11

- Fixes and improvements

#### 1.8.10

- Fixes and improvements

#### 1.9.0-beta.17

- Update quic-go to v0.42.0
- Fixes and improvements

`quic-go`#### 1.9.0-beta.16

- Fixes and improvements

Our Testflight distribution has been temporarily blocked by Apple (possibly due to too many beta versions)
and you cannot join the test, install or update the sing-box beta app right now.
Please wait patiently for processing.

#### 1.9.0-beta.14

- Update gVisor to 20240212.0-65-g71212d503
- Fixes and improvements

#### 1.8.9

- Fixes and improvements

#### 1.8.8

- Fixes and improvements

#### 1.9.0-beta.7

- Fixes and improvements

#### 1.9.0-beta.6

- Fix address filter DNS rule items 1
- Fix DNS outbound responding with wrong data
- Fixes and improvements

1:

Fixed an issue where address filter DNS rule was incorrectly rejected under certain circumstances.
If you have enabled store_rdrc to save results, consider clearing the cache file.

`store_rdrc`#### 1.8.7

- Fixes and improvements

#### 1.9.0-alpha.15

- Fixes and improvements

#### 1.9.0-alpha.14

- Improve DNS truncate behavior
- Fixes and improvements

#### 1.9.0-alpha.13

- Fixes and improvements

#### 1.8.6

- Fixes and improvements

#### 1.9.0-alpha.12

- Handle Windows power events
- Always disable cache for fake-ip DNS transport if dns.independent_cache disabled
- Fixes and improvements

`dns.independent_cache`#### 1.9.0-alpha.11

- Fix missing rule_set_ipcidr_match_source item in DNS rules 1
- Fixes and improvements

`rule_set_ipcidr_match_source`1:

See DNS Rule.

#### 1.9.0-alpha.10

- Add bypass_domain and search_domain platform HTTP proxy options 1
- Fixes and improvements

`bypass_domain``search_domain`1:

See TUN inbound.

#### 1.9.0-alpha.8

- Add rejected DNS response cache support 1
- Fixes and improvements

1:

The new feature allows you to cache the check results of
Address filter DNS rule items until expiration.

#### 1.9.0-alpha.7

- Update gVisor to 20240206.0
- Fixes and improvements

#### 1.9.0-alpha.6

- Fixes and improvements

#### 1.9.0-alpha.3

- Update quic-go to v0.41.0
- Fixes and improvements

`quic-go`#### 1.9.0-alpha.2

- Add support for client-subnet DNS options 1
- Fixes and improvements

`client-subnet`1:

See DNS, DNS Server and DNS Rules.

Since this feature makes the scenario mentioned in alpha.1 no longer leak DNS requests,
the Client example has been updated.

`alpha.1`#### 1.9.0-alpha.1

- domain_suffix behavior update 1
- process_path format update on Windows 2
- Add address filter DNS rule items 3

`domain_suffix``process_path`1:

See Migration.

2:

See Migration.

3:

The new DNS feature allows you to more precisely bypass Chinese websites via DNS leaks. Do not use plain local DNS
if using this method.

See Address Filter Fields.

Client example updated.

#### 1.8.5

- Fixes and improvements

#### 1.8.4

- Fixes and improvements

#### 1.8.2

- Fixes and improvements

#### 1.8.1

- Fixes and improvements

### 1.8.0

- Fixes and improvements

Important changes since 1.7:

- Migrate cache file from Clash API to independent options 1
- Introducing rule-set 2
- Add sing-box geoip, sing-box geosite and sing-box rule-set commands 3
- Allow nested logical rules 4
- Independent source_ip_is_private and ip_is_private rules 5
- Add context to JSON decode error message 6
- Reject internal fake-ip queries 7
- Add GSO support for TUN and WireGuard system interface 8
- Add idle_timeout for URLTest outbound 9
- Add simple loopback detect
- Optimize memory usage of idle connections
- Update uTLS to 1.5.4 10
- Update dependencies 11

`sing-box geoip``sing-box geosite``sing-box rule-set``source_ip_is_private``ip_is_private``idle_timeout`1:

See Cache File and
Migration.

2:

rule-set is independent collections of rules that can be compiled into binaries to improve performance.
Compared to legacy GeoIP and Geosite resources,
it can include more types of rules, load faster,
use less memory, and update automatically.

See Route#rule_set,
Route Rule,
DNS Rule,
rule-set,
Source Format and
Headless Rule.

For GEO resources migration, see Migrate GeoIP to rule-sets and
Migrate Geosite to rule-sets.

3:

New commands manage GeoIP, Geosite and rule-set resources, and help you migrate GEO resources to rule-sets.

4:

Logical rules in route rules, DNS rules, and the new headless rule now allow nesting of logical rules.

5:

The private GeoIP country never existed and was actually implemented inside V2Ray.
Since GeoIP was deprecated, we made this rule independent, see Migration.

`private`6:

JSON parse errors will now include the current key path.
Only takes effect when compiled with Go 1.21+.

7:

All internal DNS queries now skip DNS rules with server type fakeip,
and the default DNS server can no longer be fakeip.

`server``fakeip``fakeip`This change is intended to break incorrect usage and essentially requires no action.

8:

See TUN inbound and WireGuard outbound.

9:

When URLTest is idle for a certain period of time, the scheduled delay test will be paused.

10:

Added some new fingerprints.
Also, starting with this release, uTLS requires at least Go 1.20.

11:

Updated cloudflare-tls, gomobile, smux, tfo-go and wireguard-go to latest, quic-go to 0.40.1 and  gvisor
to 20231204.0

`cloudflare-tls``gomobile``smux``tfo-go``wireguard-go``quic-go``0.40.1``gvisor``20231204.0`#### 1.8.0-rc.11

- Fixes and improvements

#### 1.7.8

- Fixes and improvements

#### 1.8.0-rc.10

- Fixes and improvements

#### 1.7.7

- Fix V2Ray transport path validation behavior 1
- Fixes and improvements

`path`1:

See V2Ray transport.

#### 1.8.0-rc.7

- Fixes and improvements

#### 1.8.0-rc.3

- Fix V2Ray transport path validation behavior 1
- Fixes and improvements

`path`1:

See V2Ray transport.

#### 1.7.6

- Fixes and improvements

#### 1.8.0-rc.1

- Fixes and improvements

#### 1.8.0-beta.9

- Add simple loopback detect
- Fixes and improvements

#### 1.7.5

- Fixes and improvements

#### 1.8.0-alpha.17

- Add GSO support for TUN and WireGuard system interface 1
- Update uTLS to 1.5.4 2
- Update dependencies 3
- Fixes and improvements

1:

See TUN inbound and WireGuard outbound.

2:

Added some new fingerprints.
Also, starting with this release, uTLS requires at least Go 1.20.

3:

Updated cloudflare-tls, gomobile, smux, tfo-go and wireguard-go to latest, and gvisor to 20231204.0

`cloudflare-tls``gomobile``smux``tfo-go``wireguard-go``gvisor``20231204.0`This may break something, good luck!

#### 1.7.4

- Fixes and improvements

Due to the long waiting time, this version is no longer waiting for approval
by the Apple App Store, so updates to Apple Platforms will be delayed.

#### 1.8.0-alpha.16

- Fixes and improvements

#### 1.8.0-alpha.15

- Some chaotic changes 1
- Fixes and improvements

1:

Designed to optimize memory usage of idle connections, may take effect on the following protocols:

| Protocol | TCP | UDP | 
| --- | --- | --- |
| HTTP proxy server |  | / | 
| SOCKS5 |  |  | 
| Shadowsocks none/AEAD/AEAD2022 |  |  | 
| Trojan | / |  | 
| TUIC/Hysteria/Hysteria2 |  |  | 
| Multiplex |  |  | 
| Plain TLS (Trojan/VLESS without extra sub-protocols) |  | / | 
| Other protocols |  |  | 

At the same time, everything existing may be broken, please actively report problems with this version.

#### 1.8.0-alpha.13

- Fixes and improvements

#### 1.8.0-alpha.10

- Add idle_timeout for URLTest outbound 1
- Fixes and improvements

`idle_timeout`1:

When URLTest is idle for a certain period of time, the scheduled delay test will be paused.

#### 1.7.2

- Fixes and improvements

#### 1.8.0-alpha.8

- Add context to JSON decode error message 1
- Reject internal fake-ip queries 2
- Fixes and improvements

1:

JSON parse errors will now include the current key path.
Only takes effect when compiled with Go 1.21+.

2:

All internal DNS queries now skip DNS rules with server type fakeip,
and the default DNS server can no longer be fakeip.

`server``fakeip``fakeip`This change is intended to break incorrect usage and essentially requires no action.

#### 1.8.0-alpha.7

- Fixes and improvements

#### 1.7.1

- Fixes and improvements

#### 1.8.0-alpha.6

- Fix rule-set matching logic 1
- Fixes and improvements

1:

Now the rules in the rule_set rule item can be logically considered to be merged into the rule using rule-sets,
rather than completely following the AND logic.

`rule_set`#### 1.8.0-alpha.5

- Parallel rule-set initialization
- Independent source_ip_is_private and ip_is_private rules 1

`source_ip_is_private``ip_is_private`1:

The private GeoIP country never existed and was actually implemented inside V2Ray.
Since GeoIP was deprecated, we made this rule independent, see Migration.

`private`#### 1.8.0-alpha.1

- Migrate cache file from Clash API to independent options 1
- Introducing rule-set 2
- Add sing-box geoip, sing-box geosite and sing-box rule-set commands 3
- Allow nested logical rules 4

`sing-box geoip``sing-box geosite``sing-box rule-set`1:

See Cache File and
Migration.

2:

rule-set is independent collections of rules that can be compiled into binaries to improve performance.
Compared to legacy GeoIP and Geosite resources,
it can include more types of rules, load faster,
use less memory, and update automatically.

See Route#rule_set,
Route Rule,
DNS Rule,
rule-set,
Source Format and
Headless Rule.

For GEO resources migration, see Migrate GeoIP to rule-sets and
Migrate Geosite to rule-sets.

3:

New commands manage GeoIP, Geosite and rule-set resources, and help you migrate GEO resources to rule-sets.

4:

Logical rules in route rules, DNS rules, and the new headless rule now allow nesting of logical rules.

### 1.7.0

- Fixes and improvements

Important changes since 1.6:

- Add exclude route support for TUN inbound
- Add udp_disable_domain_unmapping inbound listen option 1
- Add HTTPUpgrade V2Ray transport support 2
- Migrate multiplex and UoT server to inbound 3
- Add TCP Brutal support for multiplex 4
- Add wifi_ssid and wifi_bssid route and DNS rules 5
- Update quic-go to v0.40.0
- Update gVisor to 20231113.0

`udp_disable_domain_unmapping``wifi_ssid``wifi_bssid`1:

If enabled, for UDP proxy requests addressed to a domain,
the original packet address will be sent in the response instead of the mapped domain.

This option is used for compatibility with clients that
do not support receiving UDP packets with domain addresses, such as Surge.

2:

Introduced in V2Ray 5.10.0.

The new HTTPUpgrade transport has better performance than WebSocket and is better suited for CDN abuse.

3:

Starting in 1.7.0, multiplexing support is no longer enabled by default
and needs to be turned on explicitly in inbound
options.

4

Hysteria Brutal Congestion Control Algorithm in TCP. A kernel module needs to be installed on the Linux server,
see TCP Brutal for details.

5:

Only supported in graphical clients on Android and Apple platforms.

#### 1.7.0-rc.3

- Fixes and improvements

#### 1.6.7

- macOS: Add button for uninstall SystemExtension in the standalone graphical client
- Fix missing UDP user context on TUIC/Hysteria2 inbounds
- Fixes and improvements

#### 1.7.0-rc.2

- Fix missing UDP user context on TUIC/Hysteria2 inbounds
- macOS: Add button for uninstall SystemExtension in the standalone graphical client

#### 1.6.6

- Fixes and improvements

#### 1.7.0-rc.1

- Fixes and improvements

#### 1.7.0-beta.5

- Update gVisor to 20231113.0
- Fixes and improvements

#### 1.7.0-beta.4

- Add wifi_ssid and wifi_bssid route and DNS rules 1
- Fixes and improvements

`wifi_ssid``wifi_bssid`1:

Only supported in graphical clients on Android and Apple platforms.

#### 1.7.0-beta.3

- Fix zero TTL was incorrectly reset
- Fixes and improvements

#### 1.6.5

- Fix crash if TUIC inbound authentication failed
- Fixes and improvements

#### 1.7.0-beta.2

- Fix crash if TUIC inbound authentication failed
- Update quic-go to v0.40.0
- Fixes and improvements

#### 1.6.4

- Fixes and improvements

#### 1.7.0-beta.1

- Fixes and improvements

#### 1.6.3

- iOS/Android: Fix profile auto update
- Fixes and improvements

#### 1.7.0-alpha.11

- iOS/Android: Fix profile auto update
- Fixes and improvements

#### 1.7.0-alpha.10

- Fix tcp-brutal not working with TLS
- Fix Android client not closing in some cases
- Fixes and improvements

#### 1.6.2

- Fixes and improvements

#### 1.6.1

- Our Android client is now available in the Google Play Store ▶️
- Fixes and improvements

#### 1.7.0-alpha.6

- Fixes and improvements

#### 1.7.0-alpha.4

- Migrate multiplex and UoT server to inbound 1
- Add TCP Brutal support for multiplex 2

1:

Starting in 1.7.0, multiplexing support is no longer enabled by default and needs to be turned on explicitly in inbound
options.

2

Hysteria Brutal Congestion Control Algorithm in TCP. A kernel module needs to be installed on the Linux server,
see TCP Brutal for details.

#### 1.7.0-alpha.3

- Add HTTPUpgrade V2Ray transport support 1
- Fixes and improvements

1:

Introduced in V2Ray 5.10.0.

The new HTTPUpgrade transport has better performance than WebSocket and is better suited for CDN abuse.

### 1.6.0

- Fixes and improvements

Important changes since 1.5:

- Our Apple tvOS client is now available in the App Store 🍎
- Update BBR congestion control for TUIC and Hysteria2 1
- Update brutal congestion control for Hysteria2
- Add brutal_debug option for Hysteria2
- Update legacy Hysteria protocol 2
- Add TLS self sign key pair generate command
- Remove Deprecated Features by agreement

`brutal_debug`1:

None of the existing Golang BBR congestion control implementations have been reviewed or unit tested.
This update is intended to address the multi-send defects of the old implementation and may introduce new issues.

2

Based on discussions with the original author, the brutal CC and QUIC protocol parameters of
the old protocol (Hysteria 1) have been updated to be consistent with Hysteria 2

#### 1.7.0-alpha.2

- Fix bugs introduced in 1.7.0-alpha.1

#### 1.7.0-alpha.1

- Add exclude route support for TUN inbound
- Add udp_disable_domain_unmapping inbound listen option 1
- Fixes and improvements

`udp_disable_domain_unmapping`1:

If enabled, for UDP proxy requests addressed to a domain,
the original packet address will be sent in the response instead of the mapped domain.

This option is used for compatibility with clients that
do not support receiving UDP packets with domain addresses, such as Surge.

#### 1.5.5

- Fix IPv6 auto_route for Linux 1
- Add legacy builds for old Windows and macOS systems 2
- Fixes and improvements

`auto_route`1:

When auto_route is enabled and strict_route is disabled, the device can now be reached from external IPv6 addresses.

`auto_route``strict_route`2:

Built using Go 1.20, the last version that will run on
Windows 7, 8, Server 2008, Server 2012 and macOS 10.13 High
Sierra, 10.14 Mojave.

#### 1.6.0-rc.4

- Fixes and improvements

#### 1.6.0-rc.1

- Add legacy builds for old Windows and macOS systems 1
- Fixes and improvements

1:

Built using Go 1.20, the last version that will run on
Windows 7, 8, Server 2008, Server 2012 and macOS 10.13 High
Sierra, 10.14 Mojave.

#### 1.6.0-beta.4

- Fix IPv6 auto_route for Linux 1
- Fixes and improvements

`auto_route`1:

When auto_route is enabled and strict_route is disabled, the device can now be reached from external IPv6 addresses.

`auto_route``strict_route`#### 1.5.4

- Fix Clash cache crash on arm32 devices
- Fixes and improvements

#### 1.6.0-beta.3

- Update the legacy Hysteria protocol 1
- Fixes and improvements

1

Based on discussions with the original author, the brutal CC and QUIC protocol parameters of
the old protocol (Hysteria 1) have been updated to be consistent with Hysteria 2

#### 1.6.0-beta.2

- Add TLS self sign key pair generate command
- Update brutal congestion control for Hysteria2
- Fix Clash cache crash on arm32 devices
- Update golang.org/x/net to v0.17.0
- Fixes and improvements

#### 1.6.0-beta.3

- Update the legacy Hysteria protocol 1
- Fixes and improvements

1

Based on discussions with the original author, the brutal CC and QUIC protocol parameters of
the old protocol (Hysteria 1) have been updated to be consistent with Hysteria 2

#### 1.6.0-beta.2

- Add TLS self sign key pair generate command
- Update brutal congestion control for Hysteria2
- Fix Clash cache crash on arm32 devices
- Update golang.org/x/net to v0.17.0
- Fixes and improvements

#### 1.5.3

- Fix compatibility with Android 14
- Fixes and improvements

#### 1.6.0-beta.1

- Fixes and improvements

#### 1.6.0-alpha.5

- Fix compatibility with Android 14
- Update BBR congestion control for TUIC and Hysteria2 1
- Fixes and improvements

1:

None of the existing Golang BBR congestion control implementations have been reviewed or unit tested.
This update is intended to fix a memory leak flaw in the new implementation introduced in 1.6.0-alpha.1 and may
introduce new issues.

#### 1.6.0-alpha.4

- Add brutal_debug option for Hysteria2
- Fixes and improvements

`brutal_debug`#### 1.5.2

- Our Apple tvOS client is now available in the App Store 🍎
- Fixes and improvements

#### 1.6.0-alpha.3

- Fixes and improvements

#### 1.6.0-alpha.2

- Fixes and improvements

#### 1.5.1

- Fixes and improvements

#### 1.6.0-alpha.1

- Update BBR congestion control for TUIC and Hysteria2 1
- Update quic-go to v0.39.0
- Update gVisor to 20230814.0
- Remove Deprecated Features by agreement
- Fixes and improvements

1:

None of the existing Golang BBR congestion control implementations have been reviewed or unit tested.
This update is intended to address the multi-send defects of the old implementation and may introduce new issues.

### 1.5.0

- Fixes and improvements

Important changes since 1.4:

- Add TLS ECH server support
- Improve TLS TCH client configuration
- Add TLS ECH key pair generator 1
- Add TLS ECH support for QUIC based protocols 2
- Add KDE support for the set_system_proxy option in HTTP inbound
- Add Hysteria2 protocol support 3
- Add interrupt_exist_connections option for Selector and URLTest outbounds 4
- Add DNS01 challenge support for ACME TLS certificate issuer 5
- Add merge command 6
- Mark Deprecated Features

`set_system_proxy``interrupt_exist_connections``Selector``URLTest``merge`1:

Command: sing-box generate ech-keypair <plain_server_name> [--pq-signature-schemes-enabled]

`sing-box generate ech-keypair <plain_server_name> [--pq-signature-schemes-enabled]`2:

All inbounds and outbounds are supported, including Naiveproxy, Hysteria[/2], TUIC and V2ray QUIC transport.

`Naiveproxy``Hysteria[/2]``TUIC``V2ray QUIC transport`3:

See Hysteria2 inbound and Hysteria2 outbound

For protocol description, please refer to https://v2.hysteria.network

4:

Interrupt existing connections when the selected outbound has changed.

Only inbound connections are affected by this setting, internal connections will always be interrupted.

5:

Only Alibaba Cloud DNS and Cloudflare are supported, see ACME Fields
and DNS01 Challenge Fields.

`Alibaba Cloud DNS``Cloudflare`6:

This command also parses path resources that appear in the configuration file and replaces them with embedded
configuration, such as TLS certificates or SSH private keys.

#### 1.5.0-rc.6

- Fixes and improvements

#### 1.4.6

- Fixes and improvements

#### 1.5.0-rc.5

- Fixed an improper authentication vulnerability in the SOCKS5 inbound
- Fixes and improvements

Security Advisory

This update fixes an improper authentication vulnerability in the sing-box SOCKS inbound. This vulnerability allows an
attacker to craft special requests to bypass user authentication. All users exposing SOCKS servers with user
authentication in an insecure environment are advised to update immediately.

此更新修复了 sing-box SOCKS 入站中的一个不正确身份验证漏洞。 该漏洞允许攻击者制作特殊请求来绕过用户身份验证。建议所有将使用用户认证的
SOCKS 服务器暴露在不安全环境下的用户立更新。

#### 1.4.5

- Fixed an improper authentication vulnerability in the SOCKS5 inbound
- Fixes and improvements

Security Advisory

This update fixes an improper authentication vulnerability in the sing-box SOCKS inbound. This vulnerability allows an
attacker to craft special requests to bypass user authentication. All users exposing SOCKS servers with user
authentication in an insecure environment are advised to update immediately.

此更新修复了 sing-box SOCKS 入站中的一个不正确身份验证漏洞。 该漏洞允许攻击者制作特殊请求来绕过用户身份验证。建议所有将使用用户认证的
SOCKS 服务器暴露在不安全环境下的用户立更新。

#### 1.5.0-rc.3

- Fixes and improvements

#### 1.5.0-beta.12

- Add merge command 1
- Fixes and improvements

`merge`1:

This command also parses path resources that appear in the configuration file and replaces them with embedded
configuration, such as TLS certificates or SSH private keys.

```
Merge configurations

Usage:
  sing-box merge [output] [flags]

Flags:
  -h, --help   help for merge

Global Flags:
  -c, --config stringArray             set configuration file path
  -C, --config-directory stringArray   set configuration directory path
  -D, --directory string               set working directory
      --disable-color                  disable color output

```

#### 1.5.0-beta.11

- Add DNS01 challenge support for ACME TLS certificate issuer 1
- Fixes and improvements

1:

Only Alibaba Cloud DNS and Cloudflare are supported,
see ACME Fields
and DNS01 Challenge Fields.

`Alibaba Cloud DNS``Cloudflare`#### 1.5.0-beta.10

- Add interrupt_exist_connections option for Selector and URLTest outbounds 1
- Fixes and improvements

`interrupt_exist_connections``Selector``URLTest`1:

Interrupt existing connections when the selected outbound has changed.

Only inbound connections are affected by this setting, internal connections will always be interrupted.

#### 1.4.3

- Fixes and improvements

#### 1.5.0-beta.8

- Fixes and improvements

#### 1.4.2

- Fixes and improvements

#### 1.5.0-beta.6

- Fix compatibility issues with official Hysteria2 server and client
- Fixes and improvements
- Mark deprecated features

#### 1.5.0-beta.3

- Fixes and improvements
- Updated Hysteria2 documentation 1

1:

Added notes indicating compatibility issues with the official
Hysteria2 server and client when using fastOpen=false or UDP MTU >= 1200.

`fastOpen=false`#### 1.5.0-beta.2

- Add hysteria2 protocol support 1
- Fixes and improvements

1:

See Hysteria2 inbound and Hysteria2 outbound

For protocol description, please refer to https://v2.hysteria.network

#### 1.5.0-beta.1

- Add TLS ECH server support
- Improve TLS TCH client configuration
- Add TLS ECH key pair generator 1
- Add TLS ECH support for QUIC based protocols 2
- Add KDE support for the set_system_proxy option in HTTP inbound

`set_system_proxy`1:

Command: sing-box generate ech-keypair <plain_server_name> [--pq-signature-schemes-enabled]

`sing-box generate ech-keypair <plain_server_name> [--pq-signature-schemes-enabled]`2:

All inbounds and outbounds are supported, including Naiveproxy, Hysteria, TUIC and V2ray QUIC transport.

`Naiveproxy``Hysteria``TUIC``V2ray QUIC transport`#### 1.4.1

- Fixes and improvements

### 1.4.0

- Fix bugs and update dependencies

Important changes since 1.3:

- Add TUIC support 1
- Add udp_over_stream option for TUIC client 2
- Add MultiPath TCP support 3
- Add include_interface and exclude_interface options for tun inbound
- Pause recurring tasks when no network or device idle
- Improve Android and Apple platform clients

`udp_over_stream``include_interface``exclude_interface`1:

See TUIC inbound
and TUIC outbound

2:

This is the TUIC port of the UDP over TCP protocol, designed to provide a QUIC
stream based UDP relay mode that TUIC does not provide. Since it is an add-on protocol, you will need to use sing-box or
another program compatible with the protocol as a server.

This mode has no positive effect in a proper UDP proxy scenario and should only be applied to relay streaming UDP
traffic (basically QUIC streams).

3:

Requires sing-box to be compiled with Go 1.21.

#### 1.4.0-rc.3

- Fixes and improvements

#### 1.4.0-rc.2

- Fixes and improvements

#### 1.4.0-rc.1

- Fix TUIC UDP

#### 1.4.0-beta.6

- Add udp_over_stream option for TUIC client 1
- Add include_interface and exclude_interface options for tun inbound
- Fixes and improvements

`udp_over_stream``include_interface``exclude_interface`1:

This is the TUIC port of the UDP over TCP protocol, designed to provide a QUIC
stream based UDP relay mode that TUIC does not provide. Since it is an add-on protocol, you will need to use sing-box or
another program compatible with the protocol as a server.

This mode has no positive effect in a proper UDP proxy scenario and should only be applied to relay streaming UDP
traffic (basically QUIC streams).

#### 1.4.0-beta.5

- Fixes and improvements

#### 1.4.0-beta.4

- Graphical clients: Persistence group expansion state
- Fixes and improvements

#### 1.4.0-beta.3

- Fixes and improvements

#### 1.4.0-beta.2

- Add MultiPath TCP support 1
- Drop QUIC support for Go 1.18 and 1.19 due to upstream changes
- Fixes and improvements

1:

Requires sing-box to be compiled with Go 1.21.

#### 1.4.0-beta.1

- Add TUIC support 1
- Pause recurring tasks when no network or device idle
- Fixes and improvements

1:

See TUIC inbound
and TUIC outbound

#### 1.3.6

- Fixes and improvements

#### 1.3.5

- Fixes and improvements
- Introducing our Apple tvOS client applications 1
- Add per app proxy and app installed/updated trigger support for Android client
- Add profile sharing support for Android/iOS/macOS clients

1:

Due to the requirement of tvOS 17, the app cannot be submitted to the App Store for the time being, and can only be
downloaded through TestFlight.

#### 1.3.4

- Fixes and improvements
- We're now on the App Store, always free! It should be noted
  that due to stricter and slower review, the release of Store versions will be delayed.
- We've made a standalone version of the macOS client (the original Application Extension relies on App Store
  distribution), which you can download as SFM-version-universal.zip in the release artifacts.

#### 1.3.3

- Fixes and improvements

#### 1.3.1-rc.1

- Fix bugs and update dependencies

#### 1.3.1-beta.3

- Introducing our new iOS and macOS client applications **1
  **
- Fixes and improvements

1:

The old testflight link and app are no longer valid.

#### 1.3.1-beta.2

- Fix bugs and update dependencies

#### 1.3.1-beta.1

- Fixes and improvements

### 1.3.0

- Fix bugs and update dependencies

Important changes since 1.2:

- Add FakeIP support 1
- Improve multiplex 2
- Add DNS reverse mapping support
- Add rewrite_ttl DNS rule action
- Add store_fakeip Clash API option
- Add multi-peer support for WireGuard outbound
- Add loopback detect
- Add Clash.Meta API compatibility for Clash API
- Download Yacd-meta by default if the specified Clash external_ui directory is empty
- Add path and headers option for HTTP outbound
- Perform URLTest recheck after network changes
- Fix system tun stack for ios
- Fix network monitor for android/ios
- Update VLESS and XUDP protocol
- Make splice work with traffic statistics systems like Clash API
- Significantly reduces memory usage of idle connections
- Improve DNS caching
- Add independent_cache option for DNS
- Reimplemented shadowsocks client
- Add multiplex support for VLESS outbound
- Automatically add Windows firewall rules in order for the system tun stack to work
- Fix TLS 1.2 support for shadow-tls client
- Add cache_id option for Clash cache file
- Fix local DNS transport for Android

`rewrite_ttl``store_fakeip``external_ui``system``independent_cache``cache_id``local`1:

See FAQ for more information.

2:

Added new h2mux multiplex protocol and padding multiplex option, see Multiplex.

`h2mux``padding`#### 1.3-rc2

- Fix local DNS transport for Android
- Fix bugs and update dependencies

`local`#### 1.3-rc1

- Fix bugs and update dependencies

#### 1.3-beta14

- Fixes and improvements

#### 1.3-beta13

- Fix resolving fakeip domains  1
- Deprecate L3 routing
- Fix bugs and update dependencies

1:

If the destination address of the connection is obtained from fakeip, dns rules with server type fakeip will be skipped.

#### 1.3-beta12

- Automatically add Windows firewall rules in order for the system tun stack to work
- Fix TLS 1.2 support for shadow-tls client
- Add cache_id option for Clash cache file
- Fixes and improvements

`cache_id`#### 1.3-beta11

- Fix bugs and update dependencies

#### 1.3-beta10

- Improve direct copy 1
- Improve DNS caching
- Add independent_cache option for DNS
- Reimplemented shadowsocks client 2
- Add multiplex support for VLESS outbound
- Set TCP keepalive for WireGuard gVisor TCP connections
- Fixes and improvements

`independent_cache`1:

- Make splice work with traffic statistics systems like Clash API
- Significantly reduces memory usage of idle connections

2:

Improved performance and reduced memory usage.

#### 1.3-beta9

- Improve multiplex 1
- Fixes and improvements

1:

Added new h2mux multiplex protocol and padding multiplex option, see Multiplex.

`h2mux``padding`#### 1.2.6

- Fix bugs and update dependencies

#### 1.3-beta8

- Fix system tun stack for ios
- Fix network monitor for android/ios
- Update VLESS and XUDP protocol 1
- Fixes and improvements

`system`*1:

This is an incompatible update for XUDP in VLESS if vision flow is enabled.

#### 1.3-beta7

- Add path and headers options for HTTP outbound
- Add multi-user support for Shadowsocks legacy AEAD inbound
- Fixes and improvements

`path``headers`#### 1.2.4

- Fixes and improvements

#### 1.3-beta6

- Fix WireGuard reconnect
- Perform URLTest recheck after network changes
- Fix bugs and update dependencies

#### 1.3-beta5

- Add Clash.Meta API compatibility for Clash API
- Download Yacd-meta by default if the specified Clash external_ui directory is empty
- Add path and headers option for HTTP outbound
- Fixes and improvements

`external_ui`#### 1.3-beta4

- Fix bugs

#### 1.3-beta2

- Download clash-dashboard if the specified Clash external_ui directory is empty
- Fix bugs and update dependencies

`external_ui`#### 1.3-beta1

- Add DNS reverse mapping support
- Add L3 routing support 1
- Add rewrite_ttl DNS rule action
- Add FakeIP support 2
- Add store_fakeip Clash API option
- Add multi-peer support for WireGuard outbound
- Add loopback detect

`rewrite_ttl``store_fakeip`1:

It can currently be used to route connections directly to WireGuard or block connections
at the IP layer.

2:

See FAQ for more information.

#### 1.2.3

- Introducing our new Android client application
- Improve UDP domain destination NAT
- Update reality protocol
- Fix TTL calculation for DNS response
- Fix v2ray HTTP transport compatibility
- Fix bugs and update dependencies

#### 1.2.2

- Accept any outbound in dns rule 1
- Fix bugs and update dependencies

`any`1:

Now you can use the any outbound rule to match server address queries instead of filling in all server domains
to domain rule.

`any``domain`#### 1.2.1

- Fix missing default host in v2ray http transport`s request
- Flush DNS cache for macOS when tun start/close
- Fix tun's DNS hijacking compatibility with systemd-resolved

### 1.2.0

- Fix bugs and update dependencies

Important changes since 1.1:

- Introducing our new iOS client application
- Introducing UDP over TCP protocol version 2
- Add platform options for tun inbound
- Add ShadowTLS protocol v3
- Add VLESS server and vision support
- Add reality TLS support
- Add NTP service
- Add DHCP DNS server support
- Add SSH host key validation support
- Add query_type DNS rule item
- Add fallback support for v2ray transport
- Add custom TLS server support for http based v2ray transports
- Add health check support for http-based v2ray transports
- Add multiple configuration support

#### 1.2-rc1

- Fix bugs and update dependencies

#### 1.2-beta10

- Add multiple configuration support 1
- Fix bugs and update dependencies

1:

Now you can pass the parameter --config or -c multiple times, or use the new parameter --config-directory or -C
to load all configuration files in a directory.

`--config``-c``--config-directory``-C`Loaded configuration files are sorted by name. If you want to control the merge order, add a numeric prefix to the file
name.

#### 1.1.7

- Improve the stability of the VMESS server
- Fix auto_detect_interface incorrectly identifying the default interface on Windows
- Fix bugs and update dependencies

`auto_detect_interface`#### 1.2-beta9

- Introducing the UDP over TCP protocol version 2
- Add health check support for http-based v2ray transports
- Remove length limit on short_id for reality TLS config
- Fix bugs and update dependencies

#### 1.2-beta8

- Update reality and uTLS libraries
- Fix auto_detect_interface incorrectly identifying the default interface on Windows

`auto_detect_interface`#### 1.2-beta7

- Fix the compatibility issue between VLESS's vision sub-protocol and the Xray-core client
- Improve the stability of the VMESS server

#### 1.2-beta6

- Introducing our new iOS client application
- Add platform options for tun inbound
- Add custom TLS server support for http based v2ray transports
- Add generate commands
- Enable XUDP by default in VLESS
- Update reality server
- Update vision protocol
- Fixed user flow in vless server
- Bug fixes
- Update dependencies

#### 1.2-beta5

- Add VLESS server and vision support
- Add reality TLS support
- Fix match private address

#### 1.1.6

- Improve vmess request
- Fix ipv6 redirect on Linux
- Fix match geoip private
- Fix parse hysteria UDP message
- Fix socks connect response
- Disable vmess header protection if transport enabled
- Update QUIC v2 version number and initial salt

#### 1.2-beta4

- Add NTP service
- Add Add multiple server names and multi-user support for shadowtls
- Add strict mode support for shadowtls v3
- Add uTLS support for shadowtls v3

#### 1.2-beta3

- Update QUIC v2 version number and initial salt
- Fix shadowtls v3 implementation

#### 1.2-beta2

- Add ShadowTLS protocol v3
- Add fallback support for v2ray transport
- Fix parse hysteria UDP message
- Fix socks connect response
- Disable vmess header protection if transport enabled

#### 1.2-beta1

- Add DHCP DNS server support
- Add SSH host key validation support
- Add query_type DNS rule item
- Add v2ray user stats api
- Add new clash DNS query api
- Improve vmess request
- Fix ipv6 redirect on Linux
- Fix match geoip private

#### 1.1.5

- Add Go 1.20 support
- Fix inbound default DF value
- Fix auth_user route for naive inbound
- Fix gRPC lite header
- Ignore domain case in route rules

#### 1.1.4

- Fix DNS log
- Fix write to h2 conn after closed
- Fix create UDP DNS transport from plain IPv6 address

#### 1.1.2

- Fix http proxy auth
- Fix user from stream packet conn
- Fix DNS response TTL
- Fix override packet conn
- Skip override system proxy bypass list
- Improve DNS log

#### 1.1.1

- Fix acme config
- Fix vmess packet conn
- Suppress quic-go set DF error

#### 1.1

- Fix close clash cache

Important changes since 1.0:

- Add support for use with android VPNService
- Add tun support for WireGuard outbound
- Add system tun stack
- Add comment filter for config
- Add option for allow optional proxy protocol header
- Add Clash mode and persistence support
- Add TLS ECH and uTLS support for outbound TLS options
- Add internal simple-obfs and v2ray-plugin
- Add ShadowsocksR outbound
- Add VLESS outbound and XUDP
- Skip wait for hysteria tcp handshake response
- Add v2ray mux support for all inbound
- Add XUDP support for VMess
- Improve websocket writer
- Refine tproxy write back
- Fix DNS leak caused by
  Windows' ordinary multihomed DNS resolution behavior
- Add sniff_timeout listen option
- Add custom route support for tun
- Add option for custom wireguard reserved bytes
- Split bind_address into ipv4 and ipv6
- Add ShadowTLS v1 and v2 support

#### 1.1-rc1

- Fix TLS config for h2 server
- Fix crash when input bad method in shadowsocks multi-user inbound
- Fix listen UDP
- Fix check invalid packet on macOS

#### 1.1-beta18

- Enhance defense against active probe for shadowtls server 1

1:

The fallback_after option has been removed.

`fallback_after`#### 1.1-beta17

- Fix shadowtls server 1

1:

Added fallback_after option.

#### 1.0.7

- Add support for new x/h2 deadline
- Fix copy pipe
- Fix decrypt xplus packet
- Fix macOS Ventura process name match
- Fix smux keepalive
- Fix vmess request buffer
- Fix h2c transport
- Fix tor geoip
- Fix udp connect for mux client
- Fix default dns transport strategy

#### 1.1-beta16

- Improve shadowtls server
- Fix default dns transport strategy
- Update uTLS to v1.2.0

#### 1.1-beta15

- Add support for new x/h2 deadline
- Fix udp connect for mux client
- Fix dns buffer
- Fix quic dns retry
- Fix create TLS config
- Fix websocket alpn
- Fix tor geoip

#### 1.1-beta14

- Add multi-user support for hysteria inbound 1
- Add custom tls client support for std grpc
- Fix smux keep alive
- Fix vmess request buffer
- Fix default local DNS server behavior
- Fix h2c transport

1:

The auth and auth_str fields have been replaced by the users field.

`auth``auth_str``users`#### 1.1-beta13

- Add custom worker count option for WireGuard outbound
- Split bind_address into ipv4 and ipv6
- Move WFP manipulation to strict route
- Fix WireGuard outbound panic when close
- Fix macOS Ventura process name match
- Fix QUIC connection migration by @HyNetwork
- Fix handling QUIC client SNI by @HyNetwork

#### 1.1-beta12

- Fix uTLS config
- Update quic-go to v0.30.0
- Update cloudflare-tls to go1.18.7

#### 1.1-beta11

- Add option for custom wireguard reserved bytes
- Fix shadowtls v2
- Fix h3 dns transport
- Fix copy pipe
- Fix decrypt xplus packet
- Fix v2ray api
- Suppress no network error
- Improve local dns transport

#### 1.1-beta10

- Add sniff_timeout listen option
- Add custom route support for tun 1
- Fix interface monitor
- Fix websocket headroom
- Fix uTLS handshake
- Fix ssh outbound
- Fix sniff fragmented quic client hello
- Fix DF for hysteria
- Fix naive overflow
- Check destination before udp connect
- Update uTLS to v1.1.5
- Update tfo-go to v2.0.2
- Update fsnotify to v1.6.0
- Update grpc to v1.50.1

1:

The strict_route on windows is removed.

`strict_route`#### 1.0.6

- Fix ssh outbound
- Fix sniff fragmented quic client hello
- Fix naive overflow
- Check destination before udp connect

#### 1.1-beta9

- Fix windows route 1
- Add v2ray statistics api
- Add ShadowTLS v2 support 2
- Fixes and improvements

1:

- Fix DNS leak caused by
  Windows' ordinary multihomed DNS resolution behavior
- Flush Windows DNS cache when start/close

2:

See ShadowTLS inbound
and ShadowTLS outbound

#### 1.1-beta8

- Fix leaks on close
- Improve websocket writer
- Refine tproxy write back
- Refine 4in6 processing
- Fix shadowsocks plugins
- Fix missing source address from transport connection
- Fix fqdn socks5 outbound connection
- Fix read source address from grpc-go

#### 1.0.5

- Fix missing source address from transport connection
- Fix fqdn socks5 outbound connection
- Fix read source address from grpc-go

#### 1.1-beta7

- Add v2ray mux and XUDP support for VMess inbound
- Add XUDP support for VMess outbound
- Disable DF on direct outbound by default
- Fix bugs in 1.1-beta6

#### 1.1-beta6

- Add URLTest outbound
- Fix bugs in 1.1-beta5

#### 1.1-beta5

- Print tags in version command
- Redirect clash hello to external ui
- Move shadowsocksr implementation to clash
- Make gVisor optional 1
- Refactor to miekg/dns
- Refactor bind control
- Fix build on go1.18
- Fix clash store-selected
- Fix close grpc conn
- Fix port rule match logic
- Fix clash api proxy type

1:

The build tag no_gvisor is replaced by with_gvisor.

`no_gvisor``with_gvisor`The default tun stack is changed to system.

#### 1.0.4

- Fix close grpc conn
- Fix port rule match logic
- Fix clash api proxy type

#### 1.1-beta4

- Add internal simple-obfs and v2ray-plugin Shadowsocks plugins
- Add ShadowsocksR outbound
- Add VLESS outbound and XUDP
- Skip wait for hysteria tcp handshake response
- Fix socks4 client
- Fix hysteria inbound
- Fix concurrent write

#### 1.0.3

- Fix socks4 client
- Fix hysteria inbound
- Fix concurrent write

#### 1.1-beta3

- Fix using custom TLS client in http2 client
- Fix bugs in 1.1-beta2

#### 1.1-beta2

- Add Clash mode and persistence support 1
- Add TLS ECH and uTLS support for outbound TLS options 2
- Fix socks4 request
- Fix processing empty dns result

1:

Switching modes using the Clash API, and store-selected are now supported,
see Experimental.

`store-selected`2:

ECH (Encrypted Client Hello) is a TLS extension that allows a client to encrypt the first part of its ClientHello
message, see TLS#ECH.

uTLS is a fork of "crypto/tls", which provides ClientHello fingerprinting resistance,
see TLS#uTLS.

#### 1.0.2

- Fix socks4 request
- Fix processing empty dns result

#### 1.1-beta1

- Add support for use with android VPNService 1
- Add tun support for WireGuard outbound 2
- Add system tun stack 3
- Add comment filter for config 4
- Add option for allow optional proxy protocol header
- Add half close for smux
- Set UDP DF by default 5
- Set default tun mtu to 9000
- Update gVisor to 20220905.0

1:

In previous versions, Android VPN would not work with tun enabled.

The usage of tun over VPN and VPN over tun is now supported, see Tun Inbound.

2:

In previous releases, WireGuard outbound support was backed by the lower performance gVisor virtual interface.

It achieves the same performance as wireguard-go by providing automatic system interface support.

3:

It does not depend on gVisor and has better performance in some cases.

It is less compatible and may not be available in some environments.

4:

Annotated json configuration files are now supported.

5:

UDP fragmentation is now blocked by default.

Including shadowsocks-libev, shadowsocks-rust and quic-go all disable segmentation by default.

See Dial Fields
and Listen Fields.

#### 1.0.1

- Fix match 4in6 address in ip_cidr
- Fix clash api log level format error
- Fix clash api unknown proxy type

#### 1.0

- Fix wireguard reconnect
- Fix naive inbound
- Fix json format error message
- Fix processing vmess termination signal
- Fix hysteria stream error
- Fix listener close when proxyproto failed

#### 1.0-rc1

- Fix write log timestamp
- Fix write zero
- Fix dial parallel in direct outbound
- Fix write trojan udp
- Fix DNS routing
- Add attribute support for geosite
- Update documentation for Dial Fields

#### 1.0-beta3

- Add chained inbound support
- Add process_path rule item
- Add macOS redirect support
- Add ShadowTLS Inbound, Outbound
  and Examples
- Fix search android package in non-owner users
- Fix socksaddr type condition
- Fix smux session status
- Refactor inbound and outbound documentation
- Minor fixes

#### 1.0-beta2

- Add strict_route option for Tun inbound
- Add packetaddr support for VMess outbound
- Add better performing alternative gRPC implementation
- Add docker image
- Fix sniff override destination

#### 1.0-beta1

- Initial release

##### 2022/08/26

- Fix ipv6 route on linux
- Fix read DNS message

##### 2022/08/25

- Let vmess use zero instead of auto if TLS enabled
- Add trojan fallback for ALPN
- Improve ip_cidr rule
- Fix format bind_address
- Fix http proxy with compressed response
- Fix route connections

##### 2022/08/24

- Fix naive padding
- Fix unix search path
- Fix close non-duplex connections
- Add ACME EAB support
- Fix early close on windows and catch any
- Initial zh-CN document translation

##### 2022/08/23

- Add V2Ray Transport support for VMess and Trojan
- Allow plain http request in Naive inbound (It can now be used with nginx)
- Add proxy protocol support
- Free memory after start
- Parse X-Forward-For in HTTP requests
- Handle SIGHUP signal

##### 2022/08/22

- Add strategy setting for each DNS server
- Add bind address to outbound options

##### 2022/08/21

- Add Tor outbound
- Add SSH outbound

##### 2022/08/20

- Attempt to unwrap ip-in-fqdn socksaddr
- Fix read packages in android 12
- Fix route on some android devices
- Improve linux process searcher
- Fix write socks5 username password auth request
- Skip bind connection with private destination to interface
- Add Trojan connection fallback

##### 2022/08/19

- Add Hysteria Inbound and Outbund
- Add ACME TLS certificate issuer
- Allow read config from stdin (-c stdin)
- Update gVisor to 20220815.0

##### 2022/08/18

- Fix find process with lwip stack
- Fix crash on shadowsocks server
- Fix crash on darwin tun
- Fix write log to file

##### 2022/08/17

- Improve async dns transports

##### 2022/08/16

- Add ip_version (route/dns) rule item
- Add WireGuard outbound

##### 2022/08/15

- Add uid, android user and package rules support in Tun routing.

##### 2022/08/13

- Fix dns concurrent write

##### 2022/08/12

- Performance improvements
- Add UoT option for SOCKS outbound

##### 2022/08/11

- Add UoT option for Shadowsocks outbound, UoT support for all inbounds

##### 2022/08/10

- Add full-featured Naive inbound
- Fix default dns server option #9 by iKirby

##### 2022/08/09

No changelog before.


---

## Graphical Clients

**Source URL**: <https://sing-box.sagernet.org/clients/>

#  Graphical Clients

Maintained by Project S to provide a unified experience and platform-specific functionality.

| Platform | Client | 
| --- | --- |
|  Android | sing-box for Android | 
|  iOS/macOS/Apple tvOS | sing-box for Apple platforms | 
|  Desktop | Working in progress | 

Some third-party projects that claim to use sing-box or use sing-box as a selling point are not listed here. The core
motivation of the maintainers of such projects is to acquire more users, and even though they provide friendly VPN
client features, the code is usually of poor quality and contains ads.


---

## sing-box for Android

**Source URL**: <https://sing-box.sagernet.org/clients/android/>

# sing-box for Android

SFA allows users to manage and run local or remote sing-box configuration files, and provides
platform-specific function implementation, such as TUN transparent proxy implementation.

##  Requirements

- Android 5.0+

##  Download

- Play Store
- Play Store (Beta)
- GitHub Releases
- F-Droid (Unified signature via reproducible builds)

##  Source code

- GitHub


---

## Features

**Source URL**: <https://sing-box.sagernet.org/clients/android/features/>

#  Features

#### UI options

- Display realtime network speed in the notification

#### Service

SFA allows you to run sing-box through ForegroundService or VpnService (when TUN is required).

#### TUN

SFA provides an unprivileged TUN implementation through Android VpnService.

| TUN inbound option | Available | Note | 
| --- | --- | --- |
| interface_name |  | Managed by Android | 
| inet4_address |  | / | 
| inet6_address |  | / | 
| mtu |  | / | 
| gso |  | No permission | 
| auto_route |  | / | 
| strict_route |  | Not implemented | 
| inet4_route_address |  | / | 
| inet6_route_address |  | / | 
| inet4_route_exclude_address |  | / | 
| inet6_route_exclude_address |  | / | 
| endpoint_independent_nat |  | / | 
| stack |  | / | 
| include_interface |  | No permission | 
| exclude_interface |  | No permission | 
| include_uid |  | No permission | 
| exclude_uid |  | No permission | 
| include_android_user |  | No permission | 
| include_package |  | / | 
| exclude_package |  | / | 
| platform |  | / | 

`interface_name``inet4_address``inet6_address``mtu``gso``auto_route``strict_route``inet4_route_address``inet6_route_address``inet4_route_exclude_address``inet6_route_exclude_address``endpoint_independent_nat``stack``include_interface``exclude_interface``include_uid``exclude_uid``include_android_user``include_package``exclude_package``platform`| Route/DNS rule option | Available | Note | 
| --- | --- | --- |
| process_name |  | No permission | 
| process_path |  | No permission | 
| process_path_regex |  | No permission | 
| package_name |  | / | 
| user |  | Use package_name instead | 
| user_id |  | Use package_name instead | 
| wifi_ssid |  | Fine location permission required | 
| wifi_bssid |  | Fine location permission required | 

`process_name``process_path``process_path_regex``package_name``user``package_name``user_id``package_name``wifi_ssid``wifi_bssid`### Override

Overrides profile configuration items with platform-specific values.

#### Per-app proxy

SFA allows you to select a list of Android apps that require proxying or bypassing in the graphical interface to
override the include_package and exclude_package configuration items.

`include_package``exclude_package`In particular, the selector also provides the “China apps” scanning feature, providing Chinese users with an excellent
experience to bypass apps that do not require a proxy. Specifically, by scanning China application or SDK
characteristics through dex class path and other means, there will be almost no missed reports.

### Chore

- The working directory is located at /sdcard/Android/data/io.nekohasekai.sfa/files (External files directory)
- Crash logs is located in $working_directory/stderr.log

`/sdcard/Android/data/io.nekohasekai.sfa/files``$working_directory/stderr.log`
---

## sing-box for Apple platforms

**Source URL**: <https://sing-box.sagernet.org/clients/apple/>

# sing-box for Apple platforms

SFI/SFM/SFT allows users to manage and run local or remote sing-box configuration files, and provides
platform-specific function implementation, such as TUN transparent proxy implementation.

Due to non-technical reasons, we are temporarily unable to update the sing-box app on the App Store and release the standalone version of the macOS client (TestFlight users are not affected)

##  Requirements

- iOS 15.0+ / macOS 13.0+ / Apple tvOS 17.0+
- An Apple account outside of mainland China

##  Download

- App Store
- TestFlight (Beta)

TestFlight quota is only available to sponsors
(one-time sponsorships are accepted).
Once you donate, you can get an invitation by join our Telegram group for sponsors from @yet_another_sponsor_bot
or sending us your Apple ID via email.

##  Download (macOS standalone version)

- Homebrew Cask

```
# brew install sfm

```

- GitHub Releases

##  Source code

- GitHub


---

## Features

**Source URL**: <https://sing-box.sagernet.org/clients/apple/features/>

#  Features

#### UI options

- Always On
- Include All Networks (Proxy traffic for LAN and cellular services)
- (Apple tvOS) Import profile from iPhone/iPad

#### Service

SFI/SFM/SFT allows you to run sing-box through NetworkExtension with Application Extension or System Extension.

#### TUN

SFI/SFM/SFT provides an unprivileged TUN implementation through NetworkExtension.

| TUN inbound option | Available | Note | 
| --- | --- | --- |
| interface_name | ️ | Managed by Darwin | 
| inet4_address |  | / | 
| inet6_address |  | / | 
| mtu |  | / | 
| gso |  | Not implemented | 
| auto_route |  | / | 
| strict_route | ️ | Not implemented | 
| inet4_route_address |  | / | 
| inet6_route_address |  | / | 
| inet4_route_exclude_address |  | / | 
| inet6_route_exclude_address |  | / | 
| endpoint_independent_nat |  | / | 
| stack |  | / | 
| include_interface | ️ | Not implemented | 
| exclude_interface | ️ | Not implemented | 
| include_uid | ️ | Not implemented | 
| exclude_uid | ️ | Not implemented | 
| include_android_user | ️ | Not implemented | 
| include_package | ️ | Not implemented | 
| exclude_package | ️ | Not implemented | 
| platform |  | / | 

`interface_name``inet4_address``inet6_address``mtu``gso``auto_route``strict_route``inet4_route_address``inet6_route_address``inet4_route_exclude_address``inet6_route_exclude_address``endpoint_independent_nat``stack``include_interface``exclude_interface``include_uid``exclude_uid``include_android_user``include_package``exclude_package``platform`| Route/DNS rule option | Available | Note | 
| --- | --- | --- |
| process_name |  | No permission | 
| process_path |  | No permission | 
| process_path_regex |  | No permission | 
| package_name |  | / | 
| user |  | No permission | 
| user_id |  | No permission | 
| wifi_ssid |  | Only supported on iOS | 
| wifi_bssid |  | Only supported on iOS | 

`process_name``process_path``process_path_regex``package_name``user``user_id``wifi_ssid``wifi_bssid`### Chore

- Crash logs is located in Settings -> View Service Log

`Settings``View Service Log`
---

## General

**Source URL**: <https://sing-box.sagernet.org/clients/general/>

# General

Describes and explains the functions implemented uniformly by sing-box graphical clients.

### Profile

Profile describes a sing-box configuration file and its state.

#### Local

- Local Profile represents a local sing-box configuration with minimal state
- The graphical client must provide an editor to modify configuration content

#### iCloud (on iOS and macOS)

- iCloud Profile represents a remote sing-box configuration with iCloud as the update source
- The configuration file is stored in the sing-box folder under iCloud
- The graphical client must provide an editor to modify configuration content

#### Remote

- Remote Profile represents a remote sing-box configuration with a URL as the update source.
- The graphical client should provide a configuration content viewer
- The graphical client must implement automatic profile update (default interval is 60 minutes) and HTTP Basic
  authorization.

At the same time, the graphical client must provide support for importing remote profiles
through a specific URL Scheme. The URL is defined as follows:

```
sing-box://import-remote-profile?url=urlEncodedURL#urlEncodedName

```

### Dashboard

While the sing-box service is running, the graphical client should provide a Dashboard interface to manage the service.

#### Status

Dashboard should display status information such as memory, connection, and traffic.

#### Mode

Dashboard should provide a Mode selector for switching when the configuration uses at least two clash_mode values.

`clash_mode`#### Groups

When the configuration includes group outbounds (specifically, Selector or URLTest),
the dashboard should provide a Group selector for status display or switching.

### Chore

#### Core

Graphical clients should provide a Core region:

- Display the current sing-box version
- Provides a button to clean the working directory
- Provides a memory limiter switch


---

## Privacy policy

**Source URL**: <https://sing-box.sagernet.org/clients/privacy/>

# Privacy policy

sing-box and official graphics clients do not collect or share personal data,
and the data generated by the software is always on your device.

## Android

The broad package (App) visibility (QUERY_ALL_PACKAGES) permission
is used to provide per-application proxy features for VPN,
sing-box will not collect your app list.

If your configuration contains wifi_ssid or wifi_bssid routing rules,
sing-box uses the location permission in the background
to get information about the connected Wi-Fi network to make them work.

`wifi_ssid``wifi_bssid`
---

## Introduction

**Source URL**: <https://sing-box.sagernet.org/configuration/>

# Introduction

sing-box uses JSON for configuration files.

### Structure

```
{
  "log": {},
  "dns": {},
  "ntp": {},
  "certificate": {},
  "endpoints": [],
  "inbounds": [],
  "outbounds": [],
  "route": {},
  "services": [],
  "experimental": {}
}

```

### Fields

| Key | Format | 
| --- | --- |
| log | Log | 
| dns | DNS | 
| ntp | NTP | 
| certificate | Certificate | 
| endpoints | Endpoint | 
| inbounds | Inbound | 
| outbounds | Outbound | 
| route | Route | 
| services | Service | 
| experimental | Experimental | 

`log``dns``ntp``certificate``endpoints``inbounds``outbounds``route``services``experimental`### Check

```
sing-box check

```

### Format

```
sing-box format -w -c config.json -D config_directory

```

### Merge

```
sing-box merge output.json -c config.json -D config_directory

```


---

## Certificate

**Source URL**: <https://sing-box.sagernet.org/configuration/certificate/>

Since sing-box 1.12.0

Changes in sing-box 1.13.0

Chrome Root Store

# Certificate

### Structure

```
{
  "store": "",
  "certificate": [],
  "certificate_path": [],
  "certificate_directory_path": []
}

```

You can ignore the JSON Array [] tag when the content is only one item

### Fields

#### store

The default X509 trusted CA certificate list.

| Type | Description | 
| --- | --- |
| system (default) | System trusted CA certificates | 
| mozilla | Mozilla Included List with China CA certificates removed | 
| chrome | Chrome Root Store with China CA certificates removed | 
| none | Empty list | 

`system``mozilla``chrome``none`#### certificate

The certificate line array to trust, in PEM format.

#### certificate_path

Will be automatically reloaded if file modified.

The paths to certificates to trust, in PEM format.

#### certificate_directory_path

Will be automatically reloaded if file modified.

The directory path to search for certificates to trust,in PEM format.


---

## DNS

**Source URL**: <https://sing-box.sagernet.org/configuration/dns/>

Changes in sing-box 1.12.0

servers

Changes in sing-box 1.11.0

cache_capacity

# DNS

### Structure

```
{
  "dns": {
    "servers": [],
    "rules": [],
    "final": "",
    "strategy": "",
    "disable_cache": false,
    "disable_expire": false,
    "independent_cache": false,
    "cache_capacity": 0,
    "reverse_mapping": false,
    "client_subnet": "",
    "fakeip": {}
  }
}

```

### Fields

| Key | Format | 
| --- | --- |
| server | List of DNS Server | 
| rules | List of DNS Rule | 
| fakeip | FakeIP | 

`server``rules``fakeip`#### final

Default dns server tag.

The first server will be used if empty.

#### strategy

Default domain strategy for resolving the domain names.

One of prefer_ipv4 prefer_ipv6 ipv4_only ipv6_only.

`prefer_ipv4``prefer_ipv6``ipv4_only``ipv6_only`#### disable_cache

Disable dns cache.

#### disable_expire

Disable dns cache expire.

#### independent_cache

Make each DNS server's cache independent for special purposes. If enabled, will slightly degrade performance.

#### cache_capacity

Since sing-box 1.11.0

LRU cache capacity.

Value less than 1024 will be ignored.

#### reverse_mapping

Stores a reverse mapping of IP addresses after responding to a DNS query in order to provide domain names when routing.

Since this process relies on the act of resolving domain names by an application before making a request, it can be
problematic in environments such as macOS, where DNS is proxied and cached by the system.

#### client_subnet

Since sing-box 1.9.0

Append a edns0-subnet OPT extra record with the specified IP prefix to every query by default.

`edns0-subnet`If value is an IP address instead of prefix, /32 or /128 will be appended automatically.

`/32``/128`Can be overrides by servers.[].client_subnet or rules.[].client_subnet.

`servers.[].client_subnet``rules.[].client_subnet`
---

## FakeIP

**Source URL**: <https://sing-box.sagernet.org/configuration/dns/fakeip/>

# FakeIP

Deprecated in sing-box 1.12.0

Legacy fake-ip configuration is deprecated and will be removed in sing-box 1.14.0, check Migration.

### Structure

```
{
  "enabled": true,
  "inet4_range": "198.18.0.0/15",
  "inet6_range": "fc00::/18"
}

```

### Fields

#### enabled

Enable FakeIP service.

#### inet4_range

IPv4 address range for FakeIP.

#### inet6_address

IPv6 address range for FakeIP.


---

## DNS Rule

**Source URL**: <https://sing-box.sagernet.org/configuration/dns/rule/>

# DNS Rule

Changes in sing-box 1.13.0

interface_address
 network_interface_address
 default_interface_address

Changes in sing-box 1.12.0

ip_accept_any
 outbound

Changes in sing-box 1.11.0

action
 server
 disable_cache
 rewrite_ttl
 client_subnet
 network_type
 network_is_expensive
 network_is_constrained

Changes in sing-box 1.10.0

rule_set_ipcidr_match_source
 rule_set_ip_cidr_match_source
 rule_set_ip_cidr_accept_empty
 process_path_regex

Changes in sing-box 1.9.0

geoip
 ip_cidr
 ip_is_private
 client_subnet
 rule_set_ipcidr_match_source

Changes in sing-box 1.8.0

rule_set
 source_ip_is_private
 geoip
 geosite

### Structure

```
{
  "dns": {
    "rules": [
      {
        "inbound": [
          "mixed-in"
        ],
        "ip_version": 6,
        "query_type": [
          "A",
          "HTTPS",
          32768
        ],
        "network": "tcp",
        "auth_user": [
          "usera",
          "userb"
        ],
        "protocol": [
          "tls",
          "http",
          "quic"
        ],
        "domain": [
          "test.com"
        ],
        "domain_suffix": [
          ".cn"
        ],
        "domain_keyword": [
          "test"
        ],
        "domain_regex": [
          "^stun\\..+"
        ],
        "source_ip_cidr": [
          "10.0.0.0/24",
          "192.168.0.1"
        ],
        "source_ip_is_private": false,
        "ip_cidr": [
          "10.0.0.0/24",
          "192.168.0.1"
        ],
        "ip_is_private": false,
        "ip_accept_any": false,
        "source_port": [
          12345
        ],
        "source_port_range": [
          "1000:2000",
          ":3000",
          "4000:"
        ],
        "port": [
          80,
          443
        ],
        "port_range": [
          "1000:2000",
          ":3000",
          "4000:"
        ],
        "process_name": [
          "curl"
        ],
        "process_path": [
          "/usr/bin/curl"
        ],
        "process_path_regex": [
          "^/usr/bin/.+"
        ],
        "package_name": [
          "com.termux"
        ],
        "user": [
          "sekai"
        ],
        "user_id": [
          1000
        ],
        "clash_mode": "direct",
        "network_type": [
          "wifi"
        ],
        "network_is_expensive": false,
        "network_is_constrained": false,
        "interface_address": {
          "en0": [
            "2000::/3"
          ]
        },
        "network_interface_address": {
          "wifi": [
            "2000::/3"
          ]
        },
        "default_interface_address": [
          "2000::/3"
        ],
        "wifi_ssid": [
          "My WIFI"
        ],
        "wifi_bssid": [
          "00:00:00:00:00:00"
        ],
        "rule_set": [
          "geoip-cn",
          "geosite-cn"
        ],
        "rule_set_ip_cidr_match_source": false,
        "rule_set_ip_cidr_accept_empty": false,
        "invert": false,
        "outbound": [
          "direct"
        ],
        "action": "route",
        "server": "local",

        // Deprecated

        "rule_set_ipcidr_match_source": false,
        "geosite": [
          "cn"
        ],
        "source_geoip": [
          "private"
        ],
        "geoip": [
          "cn"
        ]
      },
      {
        "type": "logical",
        "mode": "and",
        "rules": [],
        "action": "route",
        "server": "local"
      }
    ]
  }
}

```

You can ignore the JSON Array [] tag when the content is only one item

### Default Fields

The default rule uses the following matching logic:
(domain || domain_suffix || domain_keyword || domain_regex || geosite) &&
(port || port_range) &&
(source_geoip || source_ip_cidr ｜｜ source_ip_is_private) &&
(source_port || source_port_range) &&
other fields

`domain``domain_suffix``domain_keyword``domain_regex``geosite``port``port_range``source_geoip``source_ip_cidr``source_ip_is_private``source_port``source_port_range``other fields`Additionally, included rule-sets can be considered merged rather than as a single rule sub-item.

#### inbound

Tags of Inbound.

#### ip_version

4 (A DNS query) or 6 (AAAA DNS query).

Not limited if empty.

#### query_type

DNS query type. Values can be integers or type name strings.

#### network

tcp or udp.

`tcp``udp`#### auth_user

Username, see each inbound for details.

#### protocol

Sniffed protocol, see Sniff for details.

#### domain

Match full domain.

#### domain_suffix

Match domain suffix.

#### domain_keyword

Match domain using keyword.

#### domain_regex

Match domain using regular expression.

#### geosite

Deprecated in sing-box 1.8.0

Geosite is deprecated and will be removed in sing-box 1.12.0, check Migration.

Match geosite.

#### source_geoip

Deprecated in sing-box 1.8.0

GeoIP is deprecated and will be removed in sing-box 1.12.0, check Migration.

Match source geoip.

#### source_ip_cidr

Match source IP CIDR.

#### source_ip_is_private

Since sing-box 1.8.0

Match non-public source IP.

#### source_port

Match source port.

#### source_port_range

Match source port range.

#### port

Match port.

#### port_range

Match port range.

#### process_name

Only supported on Linux, Windows, and macOS.

Match process name.

#### process_path

Only supported on Linux, Windows, and macOS.

Match process path.

#### process_path_regex

Since sing-box 1.10.0

Only supported on Linux, Windows, and macOS.

Match process path using regular expression.

#### package_name

Match android package name.

#### user

Only supported on Linux.

Match user name.

#### user_id

Only supported on Linux.

Match user id.

#### clash_mode

Match Clash mode.

#### network_type

Since sing-box 1.11.0

Only supported in graphical clients on Android and Apple platforms.

Match network type.

Available values: wifi, cellular, ethernet and other.

`wifi``cellular``ethernet``other`#### network_is_expensive

Since sing-box 1.11.0

Only supported in graphical clients on Android and Apple platforms.

Match if network is considered Metered (on Android) or considered expensive,
such as Cellular or a Personal Hotspot (on Apple platforms).

#### network_is_constrained

Since sing-box 1.11.0

Only supported in graphical clients on Apple platforms.

Match if network is in Low Data Mode.

#### interface_address

Since sing-box 1.13.0

Only supported on Linux, Windows, and macOS.

Match interface address.

#### network_interface_address

Since sing-box 1.13.0

Only supported in graphical clients on Android and Apple platforms.

Matches network interface (same values as network_type) address.

`network_type`#### default_interface_address

Since sing-box 1.13.0

Only supported on Linux, Windows, and macOS.

Match default interface address.

#### wifi_ssid

Only supported in graphical clients on Android and Apple platforms, or on Linux.

Match WiFi SSID.

#### wifi_bssid

Only supported in graphical clients on Android and Apple platforms, or on Linux.

Match WiFi BSSID.

#### rule_set

Since sing-box 1.8.0

Match rule-set.

#### rule_set_ipcidr_match_source

Since sing-box 1.9.0

Deprecated in sing-box 1.10.0

rule_set_ipcidr_match_source is renamed to rule_set_ip_cidr_match_source and will be remove in sing-box 1.11.0.

`rule_set_ipcidr_match_source``rule_set_ip_cidr_match_source`Make ip_cidr rule items in rule-sets match the source IP.

`ip_cidr`#### rule_set_ip_cidr_match_source

Since sing-box 1.10.0

Make ip_cidr rule items in rule-sets match the source IP.

`ip_cidr`#### invert

Invert match result.

#### outbound

Deprecated in sing-box 1.12.0

outbound rule items are deprecated and will be removed in sing-box 1.14.0, check Migration.

`outbound`Match outbound.

any can be used as a value to match any outbound.

`any`#### action

Required

See DNS Rule Actions for details.

#### server

Deprecated in sing-box 1.11.0

Moved to DNS Rule Action.

#### disable_cache

Deprecated in sing-box 1.11.0

Moved to DNS Rule Action.

#### rewrite_ttl

Deprecated in sing-box 1.11.0

Moved to DNS Rule Action.

#### client_subnet

Deprecated in sing-box 1.11.0

Moved to DNS Rule Action.

### Address Filter Fields

Only takes effect for address requests (A/AAAA/HTTPS). When the query results do not match the address filtering rule items, the current rule will be skipped.

ip_cidr items in included rule-sets also takes effect as an address filtering field.

`ip_cidr`Enable experimental.cache_file.store_rdrc to cache results.

`experimental.cache_file.store_rdrc`#### geoip

Removed in sing-box 1.12.0

GeoIP is deprecated in sing-box 1.8.0 and removed in sing-box 1.12.0, check Migration.

Match GeoIP with query response.

#### ip_cidr

Since sing-box 1.9.0

Match IP CIDR with query response.

#### ip_is_private

Since sing-box 1.9.0

Match private IP with query response.

#### rule_set_ip_cidr_accept_empty

Since sing-box 1.10.0

Make ip_cidr rules in rule-sets accept empty query response.

`ip_cidr`#### ip_accept_any

Since sing-box 1.12.0

Match any IP with query response.

### Logical Fields

#### type

logical

`logical`#### mode

and or or

`and``or`#### rules

Included rules.


---

## DNS Rule Action

**Source URL**: <https://sing-box.sagernet.org/configuration/dns/rule_action/>

# DNS Rule Action

Changes in sing-box 1.12.0

strategy
 predefined

Since sing-box 1.11.0

### route

```
{
  "action": "route",  // default
  "server": "",
  "strategy": "",
  "disable_cache": false,
  "rewrite_ttl": null,
  "client_subnet": null
}

```

route inherits the classic rule behavior of routing DNS requests to the specified server.

`route`#### server

Required

Tag of target server.

#### strategy

Since sing-box 1.12.0

Set domain strategy for this query.

One of prefer_ipv4 prefer_ipv6 ipv4_only ipv6_only.

`prefer_ipv4``prefer_ipv6``ipv4_only``ipv6_only`#### disable_cache

Disable cache and save cache in this query.

#### rewrite_ttl

Rewrite TTL in DNS responses.

#### client_subnet

Append a edns0-subnet OPT extra record with the specified IP prefix to every query by default.

`edns0-subnet`If value is an IP address instead of prefix, /32 or /128 will be appended automatically.

`/32``/128`Will overrides dns.client_subnet.

`dns.client_subnet`### route-options

```
{
  "action": "route-options",
  "disable_cache": false,
  "rewrite_ttl": null,
  "client_subnet": null
}

```

route-options set options for routing.

`route-options`### reject

```
{
  "action": "reject",
  "method": "",
  "no_drop": false
}

```

reject reject DNS requests.

`reject`#### method

- default: Reply with REFUSED.
- drop: Drop the request.

`default``drop`default will be used by default.

`default`#### no_drop

If not enabled, method will be temporarily overwritten to drop after 50 triggers in 30s.

`method``drop`Not available when method is set to drop.

`method`### predefined

Since sing-box 1.12.0

```
{
  "action": "predefined",
  "rcode": "",
  "answer": [],
  "ns": [],
  "extra": []
}

```

predefined responds with predefined DNS records.

`predefined`#### rcode

The response code.

| Value | Value in the legacy rcode server | Description | 
| --- | --- | --- |
| NOERROR | success | Ok | 
| FORMERR | format_error | Bad request | 
| SERVFAIL | server_failure | Server failure | 
| NXDOMAIN | name_error | Not found | 
| NOTIMP | not_implemented | Not implemented | 
| REFUSED | refused | Refused | 

`NOERROR``success``FORMERR``format_error``SERVFAIL``server_failure``NXDOMAIN``name_error``NOTIMP``not_implemented``REFUSED``refused`NOERROR will be used by default.

`NOERROR`#### answer

List of text DNS record to respond as answers.

Examples:

| Record Type | Example | 
| --- | --- |
| A | localhost. IN A 127.0.0.1 | 
| AAAA | localhost. IN AAAA ::1 | 
| TXT | localhost. IN TXT \"Hello\" | 

`A``localhost. IN A 127.0.0.1``AAAA``localhost. IN AAAA ::1``TXT``localhost. IN TXT \"Hello\"`#### ns

List of text DNS record to respond as name servers.

#### extra

List of text DNS record to respond as extra records.


---

## DNS Server

**Source URL**: <https://sing-box.sagernet.org/configuration/dns/server/>

Changes in sing-box 1.12.0

type

# DNS Server

### Structure

```
{
  "dns": {
    "servers": [
      {
        "type": "",
        "tag": ""
      }
    ]
  }
}

```

#### type

The type of the DNS server.

| Type | Format | 
| --- | --- |
| empty (default) | Legacy | 
| local | Local | 
| hosts | Hosts | 
| tcp | TCP | 
| udp | UDP | 
| tls | TLS | 
| quic | QUIC | 
| https | HTTPS | 
| h3 | HTTP/3 | 
| dhcp | DHCP | 
| fakeip | Fake IP | 
| tailscale | Tailscale | 
| resolved | Resolved | 

`local``hosts``tcp``udp``tls``quic``https``h3``dhcp``fakeip``tailscale``resolved`#### tag

The tag of the DNS server.


---

## DHCP

**Source URL**: <https://sing-box.sagernet.org/configuration/dns/server/dhcp/>

Since sing-box 1.12.0

# DHCP

### Structure

```
{
  "dns": {
    "servers": [
      {
        "type": "dhcp",
        "tag": "",

        "interface": "",

        // Dial Fields
      }
    ]
  }
}

```

### Fields

#### interface

Interface name to listen on.

Tge default interface will be used by default.

### Dial Fields

See Dial Fields for details.


---

## Fake IP

**Source URL**: <https://sing-box.sagernet.org/configuration/dns/server/fakeip/>

Since sing-box 1.12.0

# Fake IP

### Structure

```
{
  "dns": {
    "servers": [
      {
        "type": "fakeip",
        "tag": "",

        "inet4_range": "198.18.0.0/15",
        "inet6_range": "fc00::/18"
      }
    ]
  }
}

```

### Fields

#### inet4_range

IPv4 address range for FakeIP.

#### inet6_address

IPv6 address range for FakeIP.


---

## Hosts

**Source URL**: <https://sing-box.sagernet.org/configuration/dns/server/hosts/>

Since sing-box 1.12.0

# Hosts

### Structure

```
{
  "dns": {
    "servers": [
      {
        "type": "hosts",
        "tag": "",

        "path": [],
        "predefined": {}
      }
    ]
  }
}

```

You can ignore the JSON Array [] tag when the content is only one item

### Fields

#### path

List of paths to hosts files.

/etc/hosts is used by default.

`/etc/hosts`C:\Windows\System32\Drivers\etc\hosts is used by default on Windows.

`C:\Windows\System32\Drivers\etc\hosts`Example:

```
{
  // "path": "/etc/hosts"

  "path": [
    "/etc/hosts",
    "$HOME/.hosts"
  ]
}

```

#### predefined

Predefined hosts.

Example:

```
{
  "predefined": {
    "www.google.com": "127.0.0.1",
    "localhost": [
      "127.0.0.1",
      "::1"
    ]
  }
}

```

### Examples

```
{
  "dns": {
    "servers": [
      {
        ...
      },
      {
        "type": "hosts",
        "tag": "hosts"
      }
    ],
    "rules": [
      {
        "ip_accept_any": true,
        "server": "hosts"
      }
    ]
  }
}

```


---

## DNS over HTTP3 (DoH3)

**Source URL**: <https://sing-box.sagernet.org/configuration/dns/server/http3/>

Since sing-box 1.12.0

# DNS over HTTP3 (DoH3)

### Structure

```
{
  "dns": {
    "servers": [
      {
        "type": "h3",
        "tag": "",

        "server": "",
        "server_port": 443,

        "path": "",
        "headers": {},

        "tls": {},

        // Dial Fields
      }
    ]
  }
}

```

Difference from legacy H3 server

- The old server uses default outbound by default unless detour is specified; the new one uses dialer just like outbound, which is equivalent to using an empty direct outbound by default.
- The old server uses address_resolver and address_strategy to resolve the domain name in the server; the new one uses domain_resolver and domain_strategy in Dial Fields instead.

`address_resolver``address_strategy``domain_resolver``domain_strategy`### Fields

#### server

Required

The address of the DNS server.

If domain name is used, domain_resolver must also be set to resolve IP address.

`domain_resolver`#### server_port

The port of the DNS server.

443 will be used by default.

`443`#### path

The path of the DNS server.

/dns-query will be used by default.

`/dns-query`#### headers

Additional headers to be sent to the DNS server.

#### tls

TLS configuration, see TLS.

### Dial Fields

See Dial Fields for details.


---

## DNS over HTTPS (DoH)

**Source URL**: <https://sing-box.sagernet.org/configuration/dns/server/https/>

Since sing-box 1.12.0

# DNS over HTTPS (DoH)

### Structure

```
{
  "dns": {
    "servers": [
      {
        "type": "https",
        "tag": "",

        "server": "",
        "server_port": 443,

        "path": "",
        "headers": {},

        "tls": {},

        // Dial Fields
      }
    ]
  }
}

```

Difference from legacy HTTPS server

- The old server uses default outbound by default unless detour is specified; the new one uses dialer just like outbound, which is equivalent to using an empty direct outbound by default.
- The old server uses address_resolver and address_strategy to resolve the domain name in the server; the new one uses domain_resolver and domain_strategy in Dial Fields instead.

`address_resolver``address_strategy``domain_resolver``domain_strategy`### Fields

#### server

Required

The address of the DNS server.

If domain name is used, domain_resolver must also be set to resolve IP address.

`domain_resolver`#### server_port

The port of the DNS server.

443 will be used by default.

`443`#### path

The path of the DNS server.

/dns-query will be used by default.

`/dns-query`#### headers

Additional headers to be sent to the DNS server.

#### tls

TLS configuration, see TLS.

### Dial Fields

See Dial Fields for details.


---

## Legacy

**Source URL**: <https://sing-box.sagernet.org/configuration/dns/server/legacy/>

# Legacy

Deprecated in sing-box 1.12.0

Legacy DNS servers is deprecated and will be removed in sing-box 1.14.0, check Migration.

Changes in sing-box 1.9.0

client_subnet

### Structure

```
{
  "dns": {
    "servers": [
      {
        "tag": "",
        "address": "",
        "address_resolver": "",
        "address_strategy": "",
        "strategy": "",
        "detour": "",
        "client_subnet": ""
      }
    ]
  }
}

```

### Fields

#### tag

The tag of the dns server.

#### address

Required

The address of the dns server.

| Protocol | Format | 
| --- | --- |
| System | local | 
| TCP | tcp://1.0.0.1 | 
| UDP | 8.8.8.8 udp://8.8.4.4 | 
| TLS | tls://dns.google | 
| HTTPS | https://1.1.1.1/dns-query | 
| QUIC | quic://dns.adguard.com | 
| HTTP3 | h3://8.8.8.8/dns-query | 
| RCode | rcode://refused | 
| DHCP | dhcp://auto or dhcp://en0 | 
| FakeIP | fakeip | 

`System``local``TCP``tcp://1.0.0.1``UDP``8.8.8.8``udp://8.8.4.4``TLS``tls://dns.google``HTTPS``https://1.1.1.1/dns-query``QUIC``quic://dns.adguard.com``HTTP3``h3://8.8.8.8/dns-query``RCode``rcode://refused``DHCP``dhcp://auto``dhcp://en0``fakeip`To ensure that Android system DNS is in effect, rather than Go's built-in default resolver, enable CGO at compile time.

the RCode transport is often used to block queries. Use with rules and the disable_cache rule option.

`disable_cache`| RCode | Description | 
| --- | --- |
| success | No error | 
| format_error | Format error | 
| server_failure | Server failure | 
| name_error | Non-existent domain | 
| not_implemented | Not implemented | 
| refused | Query refused | 

`success``No error``format_error``Format error``server_failure``Server failure``name_error``Non-existent domain``not_implemented``Not implemented``refused``Query refused`#### address_resolver

Required if address contains domain

Tag of a another server to resolve the domain name in the address.

#### address_strategy

The domain strategy for resolving the domain name in the address.

One of prefer_ipv4 prefer_ipv6 ipv4_only ipv6_only.

`prefer_ipv4``prefer_ipv6``ipv4_only``ipv6_only`dns.strategy will be used if empty.

`dns.strategy`#### strategy

Default domain strategy for resolving the domain names.

One of prefer_ipv4 prefer_ipv6 ipv4_only ipv6_only.

`prefer_ipv4``prefer_ipv6``ipv4_only``ipv6_only`Take no effect if overridden by other settings.

#### detour

Tag of an outbound for connecting to the dns server.

Default outbound will be used if empty.

#### client_subnet

Since sing-box 1.9.0

Append a edns0-subnet OPT extra record with the specified IP prefix to every query by default.

`edns0-subnet`If value is an IP address instead of prefix, /32 or /128 will be appended automatically.

`/32``/128`Can be overrides by rules.[].client_subnet.

`rules.[].client_subnet`Will overrides dns.client_subnet.

`dns.client_subnet`
---

## Local

**Source URL**: <https://sing-box.sagernet.org/configuration/dns/server/local/>

Changes in sing-box 1.13.0

prefer_go

Since sing-box 1.12.0

# Local

### Structure

```
{
  "dns": {
    "servers": [
      {
        "type": "local",
        "tag": "",
        "prefer_go": false

        // Dial Fields
      }
    ]
  }
}

```

Difference from legacy local server

- The old legacy local server only handles IP requests; the new one handles all types of requests and supports concurrent for IP requests.
- The old local server uses default outbound by default unless detour is specified; the new one uses dialer just like outbound, which is equivalent to using an empty direct outbound by default.

### Fields

#### prefer_go

Since sing-box 1.13.0

When enabled, local DNS server will resolve DNS by dialing itself whenever possible.

`local`Specifically, it disables following behaviors which was added as features in sing-box 1.13.0:

1. On Apple platforms: Attempt to resolve A/AAAA requests using getaddrinfo in NetworkExtension.
2. On Linux: Resolve through systemd-resolvd's DBus interface when available.

`getaddrinfo``systemd-resolvd`As a sole exception, it cannot disable the following behavior:

1. In the Android graphical client,
local will always resolve DNS through the platform interface,
as there is no other way to obtain upstream DNS servers;
On devices running Android versions lower than 10, this interface can only resolve A/AAAA requests.
2. On macOS, local will try DHCP first in Network Extension, since DHCP respects DIal Fields,
it will not be disabled by prefer_go.

In the Android graphical client,
local will always resolve DNS through the platform interface,
as there is no other way to obtain upstream DNS servers;
On devices running Android versions lower than 10, this interface can only resolve A/AAAA requests.

`local`On macOS, local will try DHCP first in Network Extension, since DHCP respects DIal Fields,
it will not be disabled by prefer_go.

`local``prefer_go`### Dial Fields

See Dial Fields for details.


---

## DNS over QUIC (DoQ)

**Source URL**: <https://sing-box.sagernet.org/configuration/dns/server/quic/>

Since sing-box 1.12.0

# DNS over QUIC (DoQ)

### Structure

```
{
  "dns": {
    "servers": [
      {
        "type": "quic",
        "tag": "",

        "server": "",
        "server_port": 853,

        "tls": {},

        // Dial Fields
      }
    ]
  }
}

```

Difference from legacy QUIC server

- The old server uses default outbound by default unless detour is specified; the new one uses dialer just like outbound, which is equivalent to using an empty direct outbound by default.
- The old server uses address_resolver and address_strategy to resolve the domain name in the server; the new one uses domain_resolver and domain_strategy in Dial Fields instead.

`address_resolver``address_strategy``domain_resolver``domain_strategy`### Fields

#### server

Required

The address of the DNS server.

If domain name is used, domain_resolver must also be set to resolve IP address.

`domain_resolver`#### server_port

The port of the DNS server.

853 will be used by default.

`853`#### tls

TLS configuration, see TLS.

### Dial Fields

See Dial Fields for details.


---

## Resolved

**Source URL**: <https://sing-box.sagernet.org/configuration/dns/server/resolved/>

Since sing-box 1.12.0

# Resolved

```
{
  "dns": {
    "servers": [
      {
        "type": "resolved",
        "tag": "",

        "service": "resolved",
        "accept_default_resolvers": false
      }
    ]
  }
}

```

### Fields

#### service

Required

The tag of the Resolved Service.

#### accept_default_resolvers

Indicates whether the default DNS resolvers should be accepted for fallback queries in addition to matching domains.

Specifically, default DNS resolvers are DNS servers that have SetLinkDefaultRoute or SetLinkDomains ~. set.

`SetLinkDefaultRoute``SetLinkDomains ~.`If not enabled, NXDOMAIN will be returned for requests that do not match search or match domains.

`NXDOMAIN`### Examples

```
{
  "dns": {
    "servers": [
      {
        "type": "local",
        "tag": "local"
      },
      {
        "type": "resolved",
        "tag": "resolved",
        "service": "resolved"
      }
    ],
    "rules": [
      {
        "ip_accept_any": true,
        "server": "resolved"
      }
    ]
  }
}

```

```
{
  "dns": {
    "servers": [
      {
        "type": "resolved",
        "service": "resolved",
        "accept_default_resolvers": true
      }
    ]
  }
}

```


---

## Tailscale

**Source URL**: <https://sing-box.sagernet.org/configuration/dns/server/tailscale/>

Since sing-box 1.12.0

# Tailscale

### Structure

```
{
  "dns": {
    "servers": [
      {
        "type": "tailscale",
        "tag": "",

        "endpoint": "ts-ep",
        "accept_default_resolvers": false
      }
    ]
  }
}

```

### Fields

#### endpoint

Required

The tag of the Tailscale Endpoint.

#### accept_default_resolvers

Indicates whether default DNS resolvers should be accepted for fallback queries in addition to MagicDNS。

if not enabled, NXDOMAIN will be returned for non-Tailscale domain queries.

`NXDOMAIN`### Examples

```
{
  "dns": {
    "servers": [
      {
        "type": "local",
        "tag": "local"
      },
      {
        "type": "tailscale",
        "tag": "ts",
        "endpoint": "ts-ep"
      }
    ],
    "rules": [
      {
        "ip_accept_any": true,
        "server": "ts"
      }
    ]
  }
}

```

```
{
  "dns": {
    "servers": [
      {
        "type": "tailscale",
        "endpoint": "ts-ep",
        "accept_default_resolvers": true
      }
    ]
  }
}

```


---

## TCP

**Source URL**: <https://sing-box.sagernet.org/configuration/dns/server/tcp/>

Since sing-box 1.12.0

# TCP

### Structure

```
{
  "dns": {
    "servers": [
      {
        "type": "tcp",
        "tag": "",

        "server": "",
        "server_port": 53,

        // Dial Fields
      }
    ]
  }
}

```

Difference from legacy TCP server

- The old server uses default outbound by default unless detour is specified; the new one uses dialer just like outbound, which is equivalent to using an empty direct outbound by default.
- The old server uses address_resolver and address_strategy to resolve the domain name in the server; the new one uses domain_resolver and domain_strategy in Dial Fields instead.

`address_resolver``address_strategy``domain_resolver``domain_strategy`### Fields

#### server

Required

The address of the DNS server.

If domain name is used, domain_resolver must also be set to resolve IP address.

`domain_resolver`#### server_port

The port of the DNS server.

53 will be used by default.

`53`### Dial Fields

See Dial Fields for details.


---

## DNS over TLS (DoT)

**Source URL**: <https://sing-box.sagernet.org/configuration/dns/server/tls/>

Since sing-box 1.12.0

# DNS over TLS (DoT)

### Structure

```
{
  "dns": {
    "servers": [
      {
        "type": "tls",
        "tag": "",

        "server": "",
        "server_port": 853,

        "tls": {},

        // Dial Fields
      }
    ]
  }
}

```

Difference from legacy TLS server

- The old server uses default outbound by default unless detour is specified; the new one uses dialer just like outbound, which is equivalent to using an empty direct outbound by default.
- The old server uses address_resolver and address_strategy to resolve the domain name in the server; the new one uses domain_resolver and domain_strategy in Dial Fields instead.

`address_resolver``address_strategy``domain_resolver``domain_strategy`### Fields

#### server

Required

The address of the DNS server.

If domain name is used, domain_resolver must also be set to resolve IP address.

`domain_resolver`#### server_port

The port of the DNS server.

853 will be used by default.

`853`#### tls

TLS configuration, see TLS.

### Dial Fields

See Dial Fields for details.


---

## UDP

**Source URL**: <https://sing-box.sagernet.org/configuration/dns/server/udp/>

Since sing-box 1.12.0

# UDP

### Structure

```
{
  "dns": {
    "servers": [
      {
        "type": "udp",
        "tag": "",

        "server": "",
        "server_port": 53,

        // Dial Fields
      }
    ]
  }
}

```

Difference from legacy UDP server

- The old server uses default outbound by default unless detour is specified; the new one uses dialer just like outbound, which is equivalent to using an empty direct outbound by default.
- The old server uses address_resolver and address_strategy to resolve the domain name in the server; the new one uses domain_resolver and domain_strategy in Dial Fields instead.

`address_resolver``address_strategy``domain_resolver``domain_strategy`### Fields

#### server

Required

The address of the DNS server.

If domain name is used, domain_resolver must also be set to resolve IP address.

`domain_resolver`#### server_port

The port of the DNS server.

53 will be used by default.

`53`### Dial Fields

See Dial Fields for details.


---

## Endpoint

**Source URL**: <https://sing-box.sagernet.org/configuration/endpoint/>

Since sing-box 1.11.0

# Endpoint

An endpoint is a protocol with inbound and outbound behavior.

### Structure

```
{
  "endpoints": [
    {
      "type": "",
      "tag": ""
    }
  ]
}

```

### Fields

| Type | Format | 
| --- | --- |
| wireguard | WireGuard | 
| tailscale | Tailscale | 

`wireguard``tailscale`#### tag

The tag of the endpoint.


---

## Tailscale

**Source URL**: <https://sing-box.sagernet.org/configuration/endpoint/tailscale/>

# Tailscale

Changes in sing-box 1.13.0

relay_server_port
 relay_server_static_endpoints
 system_interface
 system_interface_name
 system_interface_mtu

Since sing-box 1.12.0

### Structure

```
{
  "type": "tailscale",
  "tag": "ts-ep",
  "state_directory": "",
  "auth_key": "",
  "control_url": "",
  "ephemeral": false,
  "hostname": "",
  "accept_routes": false,
  "exit_node": "",
  "exit_node_allow_lan_access": false,
  "advertise_routes": [],
  "advertise_exit_node": false,
  "relay_server_port": 0,
  "relay_server_static_endpoints": [],
  "system_interface": false,
  "system_interface_name": "",
  "system_interface_mtu": 0,
  "udp_timeout": "5m",

  ... // Dial Fields
}

```

### Fields

#### state_directory

The directory where the Tailscale state is stored.

tailscale is used by default.

`tailscale`Example: $HOME/.tailscale

`$HOME/.tailscale`#### auth_key

Note

Auth key is not required. By default, sing-box will log the login URL (or popup a notification on graphical clients).

The auth key to create the node. If the node is already created (from state previously stored), then this field is not
used.

#### control_url

The coordination server URL.

https://controlplane.tailscale.com is used by default.

`https://controlplane.tailscale.com`#### ephemeral

Indicates whether the instance should register as an Ephemeral node (https://tailscale.com/s/ephemeral-nodes).

#### hostname

The hostname of the node.

System hostname is used by default.

Example: localhost

`localhost`#### accept_routes

Indicates whether the node should accept routes advertised by other nodes.

#### exit_node

The exit node name or IP address to use.

#### exit_node_allow_lan_access

Note

When the exit node does not have a corresponding advertised route, private traffics cannot be routed to the exit node even if exit_node_allow_lan_access is set.

`exit_node_allow_lan_access is`Indicates whether locally accessible subnets should be routed directly or via the exit node.

#### advertise_routes

CIDR prefixes to advertise into the Tailscale network as reachable through the current node.

Example: ["192.168.1.1/24"]

`["192.168.1.1/24"]`#### advertise_exit_node

Indicates whether the node should advertise itself as an exit node.

#### relay_server_port

Since sing-box 1.13.0

The port to listen on for incoming relay connections from other Tailscale nodes.

#### relay_server_static_endpoints

Since sing-box 1.13.0

Static endpoints to advertise for the relay server.

#### system_interface

Since sing-box 1.13.0

Create a system TUN interface for Tailscale.

#### system_interface_name

Since sing-box 1.13.0

Custom TUN interface name. By default, tailscale (or utun on macOS) will be used.

`tailscale``utun`#### system_interface_mtu

Since sing-box 1.13.0

Override the TUN MTU. By default, Tailscale's own MTU is used.

#### udp_timeout

UDP NAT expiration time.

5m will be used by default.

`5m`### Dial Fields

Note

Dial Fields in Tailscale endpoints only control how it connects to the control plane and have nothing to do with actual connections.

See Dial Fields for details.


---

## WireGuard

**Source URL**: <https://sing-box.sagernet.org/configuration/endpoint/wireguard/>

# WireGuard

Since sing-box 1.11.0

### Structure

```
{
  "type": "wireguard",
  "tag": "wg-ep",

  "system": false,
  "name": "",
  "mtu": 1408,
  "address": [],
  "private_key": "",
  "listen_port": 10000,
  "peers": [
    {
      "address": "127.0.0.1",
      "port": 10001,
      "public_key": "",
      "pre_shared_key": "",
      "allowed_ips": [],
      "persistent_keepalive_interval": 0,
      "reserved": [0, 0, 0]
    }
  ],
  "udp_timeout": "",
  "workers": 0,

  ... // Dial Fields
}

```

You can ignore the JSON Array [] tag when the content is only one item

### Fields

#### system

Use system interface.

Requires privilege and cannot conflict with exists system interfaces.

#### name

Custom interface name for system interface.

#### mtu

WireGuard MTU.

1408 will be used by default.

`1408`#### address

Required

List of IP (v4 or v6) address prefixes to be assigned to the interface.

#### private_key

Required

WireGuard requires base64-encoded public and private keys. These can be generated using the wg(8) utility:

```
wg genkey
echo "private key" || wg pubkey

```

or sing-box generate wg-keypair.

`sing-box generate wg-keypair`#### peers

Required

List of WireGuard peers.

#### peers.address

WireGuard peer address.

#### peers.port

WireGuard peer port.

#### peers.public_key

Required

WireGuard peer public key.

#### peers.pre_shared_key

WireGuard peer pre-shared key.

#### peers.allowed_ips

Required

WireGuard allowed IPs.

#### peers.persistent_keepalive_interval

WireGuard persistent keepalive interval, in seconds.

Disabled by default.

#### peers.reserved

WireGuard reserved field bytes.

#### udp_timeout

UDP NAT expiration time.

5m will be used by default.

`5m`#### workers

WireGuard worker count.

CPU count is used by default.

### Dial Fields

See Dial Fields for details.


---

## Experimental

**Source URL**: <https://sing-box.sagernet.org/configuration/experimental/>

# Experimental

Changes in sing-box 1.8.0

cache_file
 clash_api

### Structure

```
{
  "experimental": {
    "cache_file": {},
    "clash_api": {},
    "v2ray_api": {}
  }
}

```

### Fields

| Key | Format | 
| --- | --- |
| cache_file | Cache File | 
| clash_api | Clash API | 
| v2ray_api | V2Ray API | 

`cache_file``clash_api``v2ray_api`
---

## Cache File

**Source URL**: <https://sing-box.sagernet.org/configuration/experimental/cache-file/>

# Cache File

Since sing-box 1.8.0

Changes in sing-box 1.9.0

store_rdrc
 rdrc_timeout

### Structure

```
{
  "enabled": true,
  "path": "",
  "cache_id": "",
  "store_fakeip": false,
  "store_rdrc": false,
  "rdrc_timeout": ""
}

```

### Fields

#### enabled

Enable cache file.

#### path

Path to the cache file.

cache.db will be used if empty.

`cache.db`#### cache_id

Identifier in the cache file

If not empty, configuration specified data will use a separate store keyed by it.

#### store_fakeip

Store fakeip in the cache file

#### store_rdrc

Store rejected DNS response cache in the cache file

The check results of Address filter DNS rule items
will be cached until expiration.

#### rdrc_timeout

Timeout of rejected DNS response cache.

7d is used by default.

`7d`
---

## Clash API

**Source URL**: <https://sing-box.sagernet.org/configuration/experimental/clash-api/>

# Clash API

Changes in sing-box 1.10.0

access_control_allow_origin
 access_control_allow_private_network

Changes in sing-box 1.8.0

store_mode
 store_selected
 store_fakeip
 cache_file
 cache_id

### Structure

```
{
  "external_controller": "127.0.0.1:9090",
  "external_ui": "",
  "external_ui_download_url": "",
  "external_ui_download_detour": "",
  "secret": "",
  "default_mode": "",
  "access_control_allow_origin": [],
  "access_control_allow_private_network": false,

  // Deprecated

  "store_mode": false,
  "store_selected": false,
  "store_fakeip": false,
  "cache_file": "",
  "cache_id": ""
}

```

Since sing-box 1.10.0

```
{
  "external_controller": "127.0.0.1:9090",
  "access_control_allow_origin": [
    "http://127.0.0.1",
    "http://yacd.haishan.me"
  ],
  "access_control_allow_private_network": true
}

```

Since sing-box 1.10.0

```
{
  "external_controller": "0.0.0.0:9090",
  "external_ui": "dashboard"
  // "external_ui_download_detour": "direct"
}

```

You can ignore the JSON Array [] tag when the content is only one item

### Fields

#### external_controller

RESTful web API listening address. Clash API will be disabled if empty.

#### external_ui

A relative path to the configuration directory or an absolute path to a
directory in which you put some static web resource. sing-box will then
serve it at http://{{external-controller}}/ui.

`http://{{external-controller}}/ui`#### external_ui_download_url

ZIP download URL for the external UI, will be used if the specified external_ui directory is empty.

`external_ui`https://github.com/MetaCubeX/Yacd-meta/archive/gh-pages.zip will be used if empty.

`https://github.com/MetaCubeX/Yacd-meta/archive/gh-pages.zip`#### external_ui_download_detour

The tag of the outbound to download the external UI.

Default outbound will be used if empty.

#### secret

Secret for the RESTful API (optional)
Authenticate by spedifying HTTP header Authorization: Bearer ${secret}
ALWAYS set a secret if RESTful API is listening on 0.0.0.0

`Authorization: Bearer ${secret}`#### default_mode

Default mode in clash, Rule will be used if empty.

`Rule`This setting has no direct effect, but can be used in routing and DNS rules via the clash_mode rule item.

`clash_mode`#### access_control_allow_origin

Since sing-box 1.10.0

CORS allowed origins, * will be used if empty.

`*`To access the Clash API on a private network from a public website, you must explicitly specify it in access_control_allow_origin instead of using *.

`access_control_allow_origin``*`#### access_control_allow_private_network

Since sing-box 1.10.0

Allow access from private network.

To access the Clash API on a private network from a public website, access_control_allow_private_network must be enabled.

`access_control_allow_private_network`#### store_mode

Deprecated in sing-box 1.8.0

store_mode is deprecated in Clash API and enabled by default if cache_file.enabled.

`store_mode``cache_file.enabled`Store Clash mode in cache file.

#### store_selected

Deprecated in sing-box 1.8.0

store_selected is deprecated in Clash API and enabled by default if cache_file.enabled.

`store_selected``cache_file.enabled`The tag must be set for target outbounds.

Store selected outbound for the Selector outbound in cache file.

`Selector`#### store_fakeip

Deprecated in sing-box 1.8.0

store_selected is deprecated in Clash API and migrated to cache_file.store_fakeip.

`store_selected``cache_file.store_fakeip`Store fakeip in cache file.

#### cache_file

Deprecated in sing-box 1.8.0

cache_file is deprecated in Clash API and migrated to cache_file.enabled and cache_file.path.

`cache_file``cache_file.enabled``cache_file.path`Cache file path, cache.db will be used if empty.

`cache.db`#### cache_id

Deprecated in sing-box 1.8.0

cache_id is deprecated in Clash API and migrated to cache_file.cache_id.

`cache_id``cache_file.cache_id`Identifier in cache file.

If not empty, configuration specified data will use a separate store keyed by it.


---

## V2Ray API

**Source URL**: <https://sing-box.sagernet.org/configuration/experimental/v2ray-api/>

# V2Ray API

V2Ray API is not included by default, see Installation.

### Structure

```
{
  "listen": "127.0.0.1:8080",
  "stats": {
    "enabled": true,
    "inbounds": [
      "socks-in"
    ],
    "outbounds": [
      "proxy",
      "direct"
    ],
    "users": [
      "sekai"
    ]
  }
}

```

### Fields

#### listen

gRPC API listening address. V2Ray API will be disabled if empty.

#### stats

Traffic statistics service settings.

#### stats.enabled

Enable statistics service.

#### stats.inbounds

Inbound list to count traffic.

#### stats.outbounds

Outbound list to count traffic.

#### stats.users

User list to count traffic.


---

## Inbound

**Source URL**: <https://sing-box.sagernet.org/configuration/inbound/>

# Inbound

### Structure

```
{
  "inbounds": [
    {
      "type": "",
      "tag": ""
    }
  ]
}

```

### Fields

| Type | Format | Injectable | 
| --- | --- | --- |
| direct | Direct |  | 
| mixed | Mixed | TCP | 
| socks | SOCKS | TCP | 
| http | HTTP | TCP | 
| shadowsocks | Shadowsocks | TCP | 
| vmess | VMess | TCP | 
| trojan | Trojan | TCP | 
| naive | Naive |  | 
| hysteria | Hysteria |  | 
| shadowtls | ShadowTLS | TCP | 
| tuic | TUIC |  | 
| hysteria2 | Hysteria2 |  | 
| vless | VLESS | TCP | 
| anytls | AnyTLS | TCP | 
| tun | Tun |  | 
| redirect | Redirect |  | 
| tproxy | TProxy |  | 

`direct``mixed``socks``http``shadowsocks``vmess``trojan``naive``hysteria``shadowtls``tuic``hysteria2``vless``anytls``tun``redirect``tproxy`#### tag

The tag of the inbound.


---

## AnyTLS

**Source URL**: <https://sing-box.sagernet.org/configuration/inbound/anytls/>

# AnyTLS

Since sing-box 1.12.0

### Structure

```
{
  "type": "anytls",
  "tag": "anytls-in",

  ... // Listen Fields

  "users": [
    {
      "name": "sekai",
      "password": "8JCsPssfgS8tiRwiMlhARg=="
    }
  ],
  "padding_scheme": [],
  "tls": {}
}

```

### Listen Fields

See Listen Fields for details.

### Fields

#### users

Required

AnyTLS users.

#### padding_scheme

AnyTLS padding scheme line array.

Default padding scheme:

```
[
  "stop=8",
  "0=30-30",
  "1=100-400",
  "2=400-500,c,500-1000,c,500-1000,c,500-1000,c,500-1000",
  "3=9-9,500-1000",
  "4=500-1000",
  "5=500-1000",
  "6=500-1000",
  "7=500-1000"
]

```

#### tls

TLS configuration, see TLS.


---

## Direct

**Source URL**: <https://sing-box.sagernet.org/configuration/inbound/direct/>

# Direct

direct inbound is a tunnel server.

`direct`### Structure

```
{
  "type": "direct",
  "tag": "direct-in",

  ... // Listen Fields

  "network": "udp",
  "override_address": "1.0.0.1",
  "override_port": 53
}

```

### Listen Fields

See Listen Fields for details.

### Fields

#### network

Listen network, one of tcp udp.

`tcp``udp`Both if empty.

#### override_address

Override the connection destination address.

#### override_port

Override the connection destination port.


---

## HTTP

**Source URL**: <https://sing-box.sagernet.org/configuration/inbound/http/>

# HTTP

### Structure

```
{
  "type": "http",
  "tag": "http-in",

  ... // Listen Fields

  "users": [
    {
      "username": "admin",
      "password": "admin"
    }
  ],
  "tls": {},
  "set_system_proxy": false
}

```

### Listen Fields

See Listen Fields for details.

### Fields

#### tls

TLS configuration, see TLS.

#### users

HTTP users.

No authentication required if empty.

#### set_system_proxy

Only supported on Linux, Android, Windows, and macOS.

To work on Android and Apple platforms without privileges, use tun.platform.http_proxy instead.

Automatically set system proxy configuration when start and clean up when stop.


---

## Hysteria

**Source URL**: <https://sing-box.sagernet.org/configuration/inbound/hysteria/>

# Hysteria

### Structure

```
{
  "type": "hysteria",
  "tag": "hysteria-in",

  ... // Listen Fields

  "up": "100 Mbps",
  "up_mbps": 100,
  "down": "100 Mbps",
  "down_mbps": 100,
  "obfs": "fuck me till the daylight",

  "users": [
    {
      "name": "sekai",
      "auth": "",
      "auth_str": "password"
    }
  ],

  "recv_window_conn": 0,
  "recv_window_client": 0,
  "max_conn_client": 0,
  "disable_mtu_discovery": false,
  "tls": {}
}

```

### Listen Fields

See Listen Fields for details.

### Fields

#### up, down

Required

Format: [Integer] [Unit] e.g. 100 Mbps, 640 KBps, 2 Gbps

`[Integer] [Unit]``100 Mbps, 640 KBps, 2 Gbps`Supported units (case sensitive, b = bits, B = bytes, 8b=1B):

```
bps (bits per second)
Bps (bytes per second)
Kbps (kilobits per second)
KBps (kilobytes per second)
Mbps (megabits per second)
MBps (megabytes per second)
Gbps (gigabits per second)
GBps (gigabytes per second)
Tbps (terabits per second)
TBps (terabytes per second)

```

#### up_mbps, down_mbps

Required

up, down in Mbps.

`up, down`#### obfs

Obfuscated password.

#### users

Hysteria users

#### users.auth

Authentication password, in base64.

#### users.auth_str

Authentication password.

#### recv_window_conn

The QUIC stream-level flow control window for receiving data.

15728640 (15 MB/s) will be used if empty.

`15728640 (15 MB/s)`#### recv_window_client

The QUIC connection-level flow control window for receiving data.

67108864 (64 MB/s) will be used if empty.

`67108864 (64 MB/s)`#### max_conn_client

The maximum number of QUIC concurrent bidirectional streams that a peer is allowed to open.

1024 will be used if empty.

`1024`#### disable_mtu_discovery

Disables Path MTU Discovery (RFC 8899). Packets will then be at most 1252 (IPv4) / 1232 (IPv6) bytes in size.

Force enabled on for systems other than Linux and Windows (according to upstream).

#### tls

Required

TLS configuration, see TLS.


---

## Hysteria2

**Source URL**: <https://sing-box.sagernet.org/configuration/inbound/hysteria2/>

# Hysteria2

Changes in sing-box 1.11.0

masquerade
 ignore_client_bandwidth

### Structure

```
{
  "type": "hysteria2",
  "tag": "hy2-in",

  ... // Listen Fields

  "up_mbps": 100,
  "down_mbps": 100,
  "obfs": {
    "type": "salamander",
    "password": "cry_me_a_r1ver"
  },
  "users": [
    {
      "name": "tobyxdd",
      "password": "goofy_ahh_password"
    }
  ],
  "ignore_client_bandwidth": false,
  "tls": {},
  "masquerade": "", // or {}
  "brutal_debug": false
}

```

Difference from official Hysteria2

The official program supports an authentication method called userpass,
which essentially uses a combination of <username>:<password> as the actual password,
while sing-box does not provide this alias.
To use sing-box with the official program, you need to fill in that combination as the actual password.

`<username>:<password>`### Listen Fields

See Listen Fields for details.

### Fields

#### up_mbps, down_mbps

Max bandwidth, in Mbps.

Not limited if empty.

Conflict with ignore_client_bandwidth.

`ignore_client_bandwidth`#### obfs.type

QUIC traffic obfuscator type, only available with salamander.

`salamander`Disabled if empty.

#### obfs.password

QUIC traffic obfuscator password.

#### users

Hysteria2 users

#### users.password

Authentication password

#### ignore_client_bandwidth

When up_mbps and down_mbps are not set:

`up_mbps``down_mbps`Commands clients to use the BBR CC instead of Hysteria CC.

When up_mbps and down_mbps are set:

`up_mbps``down_mbps`Deny clients to use the BBR CC.

#### tls

Required

TLS configuration, see TLS.

#### masquerade

HTTP3 server behavior (URL string configuration) when authentication fails.

| Scheme | Example | Description | 
| --- | --- | --- |
| file | file:///var/www | As a file server | 
| http/https | http://127.0.0.1:8080 | As a reverse proxy | 

`file``file:///var/www``http/https``http://127.0.0.1:8080`Conflict with masquerade.type.

`masquerade.type`A 404 page will be returned if masquerade is not configured.

#### masquerade.type

HTTP3 server behavior (Object configuration) when authentication fails.

| Type | Description | Fields | 
| --- | --- | --- |
| file | As a file server | directory | 
| proxy | As a reverse proxy | url, rewrite_host | 
| string | Reply with a fixed response | status_code, headers, content | 

`file``directory``proxy``url``rewrite_host``string``status_code``headers``content`Conflict with masquerade.

`masquerade`A 404 page will be returned if masquerade is not configured.

#### masquerade.directory

File server root directory.

#### masquerade.url

Reverse proxy target URL.

#### masquerade.rewrite_host

Rewrite the Host header to the target URL.

`Host`#### masquerade.status_code

Fixed response status code.

#### masquerade.headers

Fixed response headers.

#### masquerade.content

Fixed response content.

#### brutal_debug

Enable debug information logging for Hysteria Brutal CC.


---

## Mixed

**Source URL**: <https://sing-box.sagernet.org/configuration/inbound/mixed/>

# Mixed

mixed inbound is a socks4, socks4a, socks5 and http server.

`mixed`### Structure

```
{
  "type": "mixed",
  "tag": "mixed-in",

  ... // Listen Fields

  "users": [
    {
      "username": "admin",
      "password": "admin"
    }
  ],
  "set_system_proxy": false
}

```

### Listen Fields

See Listen Fields for details.

### Fields

#### users

SOCKS and HTTP users.

No authentication required if empty.

#### set_system_proxy

Only supported on Linux, Android, Windows, and macOS.

To work on Android and Apple platforms without privileges, use tun.platform.http_proxy instead.

Automatically set system proxy configuration when start and clean up when stop.


---

## Naive

**Source URL**: <https://sing-box.sagernet.org/configuration/inbound/naive/>

# Naive

Changes in sing-box 1.13.0

quic_congestion_control

### Structure

```
{
"type": "naive",
"tag": "naive-in",
"network": "udp",
...
// Listen Fields

"users": [
{
"username": "sekai",
"password": "password"
}
],
"quic_congestion_control": "",
"tls": {}
}

```

### Listen Fields

See Listen Fields for details.

### Fields

#### network

Listen network, one of tcp udp.

`tcp``udp`Both if empty.

#### users

Required

Naive users.

#### quic_congestion_control

Since sing-box 1.13.0

QUIC congestion control algorithm.

| Algorithm | Description | 
| --- | --- |
| bbr | BBR | 
| bbr_standard | BBR (Standard version) | 
| bbr2 | BBRv2 | 
| bbr2_variant | BBRv2 (An experimental variant) | 
| cubic | CUBIC | 
| reno | New Reno | 

`bbr``bbr_standard``bbr2``bbr2_variant``cubic``reno`bbr is used by default (the default of QUICHE, used by Chromium which NaiveProxy is based on).

`bbr`#### tls

TLS configuration, see TLS.


---

## Redirect

**Source URL**: <https://sing-box.sagernet.org/configuration/inbound/redirect/>

# Redirect

Only supported on Linux and macOS.

### Structure

```
{
  "type": "redirect",
  "tag": "redirect-in",

  ... // Listen Fields
}

```

### Listen Fields

See Listen Fields for details.


---

## Shadowsocks

**Source URL**: <https://sing-box.sagernet.org/configuration/inbound/shadowsocks/>

# Shadowsocks

### Structure

```
{
  "type": "shadowsocks",
  "tag": "ss-in",

  ... // Listen Fields

  "method": "2022-blake3-aes-128-gcm",
  "password": "8JCsPssfgS8tiRwiMlhARg==",
  "managed": false,
  "multiplex": {}
}

```

### Multi-User Structure

```
{
  "method": "2022-blake3-aes-128-gcm",
  "password": "8JCsPssfgS8tiRwiMlhARg==",
  "users": [
    {
      "name": "sekai",
      "password": "PCD2Z4o12bKUoFa3cC97Hw=="
    }
  ],
  "multiplex": {}
}

```

### Relay Structure

```
{
  "type": "shadowsocks",
  "method": "2022-blake3-aes-128-gcm",
  "password": "8JCsPssfgS8tiRwiMlhARg==",
  "destinations": [
    {
      "name": "test",
      "server": "example.com",
      "server_port": 8080,
      "password": "PCD2Z4o12bKUoFa3cC97Hw=="
    }
  ],
  "multiplex": {}
}

```

### Listen Fields

See Listen Fields for details.

### Fields

#### network

Listen network, one of tcp udp.

`tcp``udp`Both if empty.

#### method

Required

| Method | Key Length | 
| --- | --- |
| 2022-blake3-aes-128-gcm | 16 | 
| 2022-blake3-aes-256-gcm | 32 | 
| 2022-blake3-chacha20-poly1305 | 32 | 
| none | / | 
| aes-128-gcm | / | 
| aes-192-gcm | / | 
| aes-256-gcm | / | 
| chacha20-ietf-poly1305 | / | 
| xchacha20-ietf-poly1305 | / | 

#### password

Required

| Method | Password Format | 
| --- | --- |
| none | / | 
| 2022 methods | sing-box generate rand --base64 <Key Length> | 
| other methods | any string | 

`sing-box generate rand --base64 <Key Length>`#### managed

Defaults to false. Enable this when the inbound is managed by the SSM API for dynamic user.

`false`#### multiplex

See Multiplex for details.


---

## ShadowTLS

**Source URL**: <https://sing-box.sagernet.org/configuration/inbound/shadowtls/>

# ShadowTLS

Changes in sing-box 1.12.0

wildcard_sni

### Structure

```
{
  "type": "shadowtls",
  "tag": "st-in",

  ... // Listen Fields

  "version": 3,
  "password": "fuck me till the daylight",
  "users": [
    {
      "name": "sekai",
      "password": "8JCsPssfgS8tiRwiMlhARg=="
    }
  ],
  "handshake": {
    "server": "google.com",
    "server_port": 443,

    ... // Dial Fields
  },
  "handshake_for_server_name": {
    "example.com": {
      "server": "example.com",
      "server_port": 443,

      ... // Dial Fields
    }
  },
  "strict_mode": false,
  "wildcard_sni": ""
}

```

### Listen Fields

See Listen Fields for details.

### Fields

#### version

ShadowTLS protocol version.

| Value | Protocol Version | 
| --- | --- |
| 1 (default) | ShadowTLS v1 | 
| 2 | ShadowTLS v2 | 
| 3 | ShadowTLS v3 | 

`1``2``3`#### password

ShadowTLS password.

Only available in the ShadowTLS protocol 2.

#### users

ShadowTLS users.

Only available in the ShadowTLS protocol 3.

#### handshake

Required

When wildcard_sni is configured to all, the server address is optional.

`wildcard_sni``all`Handshake server address and Dial Fields.

#### handshake_for_server_name

Handshake server address and Dial Fields for specific server name.

Only available in the ShadowTLS protocol 2/3.

#### strict_mode

ShadowTLS strict mode.

Only available in the ShadowTLS protocol 3.

#### wildcard_sni

Since sing-box 1.12.0

ShadowTLS wildcard SNI mode.

Available values are:

- off: (default) Disabled.
- authed: Authenticated connections will have their destination overwritten to (servername):443
- all: All connections will have their destination overwritten to (servername):443

`off``authed``(servername):443``all``(servername):443`Additionally, connections matching handshake_for_server_name are not affected.

`handshake_for_server_name`Only available in the ShadowTLS protocol 3.


---

## SOCKS

**Source URL**: <https://sing-box.sagernet.org/configuration/inbound/socks/>

# SOCKS

socks inbound is a socks4, socks4a, socks5 server.

`socks`### Structure

```
{
  "type": "socks",
  "tag": "socks-in",

  ... // Listen Fields

  "users": [
    {
      "username": "admin",
      "password": "admin"
    }
  ]
}

```

### Listen Fields

See Listen Fields for details.

### Fields

#### users

SOCKS users.

No authentication required if empty.


---

## TProxy

**Source URL**: <https://sing-box.sagernet.org/configuration/inbound/tproxy/>

# TProxy

Only supported on Linux.

### Structure

```
{
  "type": "tproxy",
  "tag": "tproxy-in",

  ... // Listen Fields

  "network": "udp"
}

```

### Listen Fields

See Listen Fields for details.

### Fields

#### network

Listen network, one of tcp udp.

`tcp``udp`Both if empty.


---

## Trojan

**Source URL**: <https://sing-box.sagernet.org/configuration/inbound/trojan/>

# Trojan

### Structure

```
{
  "type": "trojan",
  "tag": "trojan-in",

  ... // Listen Fields

  "users": [
    {
      "name": "sekai",
      "password": "8JCsPssfgS8tiRwiMlhARg=="
    }
  ],
  "tls": {},
  "fallback": {
    "server": "127.0.0.1",
    "server_port": 8080
  },
  "fallback_for_alpn": {
    "http/1.1": {
      "server": "127.0.0.1",
      "server_port": 8081
    }
  },
  "multiplex": {},
  "transport": {}
}

```

### Listen Fields

See Listen Fields for details.

### Fields

#### users

Required

Trojan users.

#### tls

TLS configuration, see TLS.

#### fallback

There is no evidence that GFW detects and blocks Trojan servers based on HTTP responses, and opening the standard http/s port on the server is a much bigger signature.

Fallback server configuration. Disabled if fallback and fallback_for_alpn are empty.

`fallback``fallback_for_alpn`#### fallback_for_alpn

Fallback server configuration for specified ALPN.

If not empty, TLS fallback requests with ALPN not in this table will be rejected.

#### multiplex

See Multiplex for details.

#### transport

V2Ray Transport configuration, see V2Ray Transport.


---

## TUIC

**Source URL**: <https://sing-box.sagernet.org/configuration/inbound/tuic/>

# TUIC

### Structure

```
{
  "type": "tuic",
  "tag": "tuic-in",

  ... // Listen Fields

  "users": [
    {
      "name": "sekai",
      "uuid": "059032A9-7D40-4A96-9BB1-36823D848068",
      "password": "hello"
    }
  ],
  "congestion_control": "cubic",
  "auth_timeout": "3s",
  "zero_rtt_handshake": false,
  "heartbeat": "10s",
  "tls": {}
}

```

### Listen Fields

See Listen Fields for details.

### Fields

#### users

TUIC users

#### users.uuid

Required

TUIC user uuid

#### users.password

TUIC user password

#### congestion_control

QUIC congestion control algorithm

One of: cubic, new_reno, bbr

`cubic``new_reno``bbr`cubic is used by default.

`cubic`#### auth_timeout

How long the server should wait for the client to send the authentication command

3s is used by default.

`3s`#### zero_rtt_handshake

Enable 0-RTT QUIC connection handshake on the client side
This is not impacting much on the performance, as the protocol is fully multiplexed

Disabling this is highly recommended, as it is vulnerable to replay attacks.
See Attack of the clones

#### heartbeat

Interval for sending heartbeat packets for keeping the connection alive

10s is used by default.

`10s`#### tls

Required

TLS configuration, see TLS.


---

## Tun

**Source URL**: <https://sing-box.sagernet.org/configuration/inbound/tun/>

# Tun

Changes in sing-box 1.13.0

auto_redirect_reset_mark
 auto_redirect_nfqueue
 exclude_mptcp

Changes in sing-box 1.12.0

loopback_address

Changes in sing-box 1.11.0

gso
 route_address_set
 route_exclude_address_set

Changes in sing-box 1.10.0

address
 inet4_address
 inet6_address
 route_address
 inet4_route_address
 inet6_route_address
 route_exclude_address
 inet4_route_exclude_address
 inet6_route_exclude_address
 iproute2_table_index
 iproute2_rule_index
 auto_redirect
 auto_redirect_input_mark
 auto_redirect_output_mark
 route_address_set
 route_exclude_address_set

Changes in sing-box 1.9.0

platform.http_proxy.bypass_domain
 platform.http_proxy.match_domain

Changes in sing-box 1.8.0

gso
 stack

Only supported on Linux, Windows and macOS.

### Structure

```
{
  "type": "tun",
  "tag": "tun-in",
  "interface_name": "tun0",
  "address": [
    "172.18.0.1/30",
    "fdfe:dcba:9876::1/126"
  ],
  "mtu": 9000,
  "auto_route": true,
  "iproute2_table_index": 2022,
  "iproute2_rule_index": 9000,
  "auto_redirect": true,
  "auto_redirect_input_mark": "0x2023",
  "auto_redirect_output_mark": "0x2024",
  "auto_redirect_reset_mark": "0x2025",
  "auto_redirect_nfqueue": 100,
  "exclude_mptcp": false,
  "loopback_address": [
    "10.7.0.1"
  ],
  "strict_route": true,
  "route_address": [
    "0.0.0.0/1",
    "128.0.0.0/1",
    "::/1",
    "8000::/1"
  ],
  "route_exclude_address": [
    "192.168.0.0/16",
    "fc00::/7"
  ],
  "route_address_set": [
    "geoip-cloudflare"
  ],
  "route_exclude_address_set": [
    "geoip-cn"
  ],
  "endpoint_independent_nat": false,
  "udp_timeout": "5m",
  "stack": "system",
  "include_interface": [
    "lan0"
  ],
  "exclude_interface": [
    "lan1"
  ],
  "include_uid": [
    0
  ],
  "include_uid_range": [
    "1000:99999"
  ],
  "exclude_uid": [
    1000
  ],
  "exclude_uid_range": [
    "1000:99999"
  ],
  "include_android_user": [
    0,
    10
  ],
  "include_package": [
    "com.android.chrome"
  ],
  "exclude_package": [
    "com.android.captiveportallogin"
  ],
  "platform": {
    "http_proxy": {
      "enabled": false,
      "server": "127.0.0.1",
      "server_port": 8080,
      "bypass_domain": [],
      "match_domain": []
    }
  },
  // Deprecated
  "gso": false,
  "inet4_address": [
    "172.19.0.1/30"
  ],
  "inet6_address": [
    "fdfe:dcba:9876::1/126"
  ],
  "inet4_route_address": [
    "0.0.0.0/1",
    "128.0.0.0/1"
  ],
  "inet6_route_address": [
    "::/1",
    "8000::/1"
  ],
  "inet4_route_exclude_address": [
    "192.168.0.0/16"
  ],
  "inet6_route_exclude_address": [
    "fc00::/7"
  ],
  ...
  // Listen Fields
}

```

You can ignore the JSON Array [] tag when the content is only one item

If tun is running in non-privileged mode, addresses and MTU will not be configured automatically, please make sure the settings are accurate.

### Fields

#### interface_name

Virtual device name, automatically selected if empty.

#### address

Since sing-box 1.10.0

IPv4 and IPv6 prefix for the tun interface.

#### inet4_address

Deprecated in sing-box 1.10.0

inet4_address is merged to address and will be removed in sing-box 1.12.0.

`inet4_address``address`IPv4 prefix for the tun interface.

#### inet6_address

Deprecated in sing-box 1.10.0

inet6_address is merged to address and will be removed in sing-box 1.12.0.

`inet6_address``address`IPv6 prefix for the tun interface.

#### mtu

The maximum transmission unit.

#### gso

Deprecated in sing-box 1.11.0

GSO has no advantages for transparent proxy scenarios, is deprecated and no longer works, and will be removed in sing-box 1.12.0.

Since sing-box 1.8.0

Only supported on Linux with auto_route enabled.

`auto_route`Enable generic segmentation offload.

#### auto_route

Set the default route to the Tun.

To avoid traffic loopback, set route.auto_detect_interface or route.default_interface or outbound.bind_interface

`route.auto_detect_interface``route.default_interface``outbound.bind_interface`Use with Android VPN

By default, VPN takes precedence over tun. To make tun go through VPN, enable route.override_android_vpn.

`route.override_android_vpn`Also enable auto_redirect

`auto_redirect`auto_redirect is always recommended on Linux, it provides better routing, higher performance (better than tproxy), and avoids conflicts between TUN and Docker bridge networks.

`auto_redirect`#### iproute2_table_index

Since sing-box 1.10.0

Linux iproute2 table index generated by auto_route.

`auto_route`2022 is used by default.

`2022`#### iproute2_rule_index

Since sing-box 1.10.0

Linux iproute2 rule start index generated by auto_route.

`auto_route`9000 is used by default.

`9000`#### auto_redirect

Since sing-box 1.10.0

Only supported on Linux with auto_route enabled.

`auto_route`Improve TUN routing and performance using nftables.

auto_redirect is always recommended on Linux, it provides better routing,
higher performance (better than tproxy),
and avoids conflicts between TUN and Docker bridge networks.

`auto_redirect`Note that auto_redirect also works on Android, 
but due to the lack of nftables and ip6tables,
only simple IPv4 TCP forwarding is performed.
To share your VPN connection over hotspot or repeater on Android,
use VPNHotspot.

`auto_redirect``nftables``ip6tables`auto_redirect also automatically inserts compatibility rules
into the OpenWrt fw4 table, i.e. 
it will work on routers without any extra configuration.

`auto_redirect`Conflict with route.default_mark and [dialOptions].routing_mark.

`route.default_mark``[dialOptions].routing_mark`#### auto_redirect_input_mark

Since sing-box 1.10.0

Connection input mark used by auto_redirect.

`auto_redirect`0x2023 is used by default.

`0x2023`#### auto_redirect_output_mark

Since sing-box 1.10.0

Connection output mark used by auto_redirect.

`auto_redirect`0x2024 is used by default.

`0x2024`#### auto_redirect_reset_mark

Since sing-box 1.13.0

Connection reset mark used by auto_redirect pre-matching.

`auto_redirect`0x2025 is used by default.

`0x2025`#### auto_redirect_nfqueue

Since sing-box 1.13.0

NFQueue number used by auto_redirect pre-matching.

`auto_redirect`100 is used by default.

`100`#### exclude_mptcp

Since sing-box 1.13.0

Only supported on Linux with nftables and requires auto_route and auto_redirect enabled.

`auto_route``auto_redirect`MPTCP cannot be transparently proxied due to protocol limitations.

Such traffic is usually created by Apple systems.

When enabled, MPTCP connections will bypass sing-box and connect directly, otherwise, will be rejected to avoid errors by default.

#### loopback_address

Since sing-box 1.12.0

Loopback addresses make TCP connections to the specified address connect to the source address.

Setting option value to 10.7.0.1 achieves the same behavior as SideStore/StosVPN.

`10.7.0.1`When auto_redirect is enabled, the same behavior can be achieved for LAN devices (not just local) as a gateway.

`auto_redirect`#### strict_route

Enforce strict routing rules when auto_route is enabled:

`auto_route`In Linux:

- Let unsupported network unreachable
- For legacy reasons, when neither strict_route nor auto_redirect are enabled, all ICMP traffic will not go through TUN.

`strict_route``auto_redirect`In Windows:

- Let unsupported network unreachable
- prevent DNS leak caused by
  Windows' ordinary multihomed DNS resolution behavior

It may prevent some Windows applications (such as VirtualBox) from working properly in certain situations.

#### route_address

Since sing-box 1.10.0

Use custom routes instead of default when auto_route is enabled.

`auto_route`#### inet4_route_address

Deprecated in sing-box 1.10.0

inet4_route_address is deprecated and will be removed in sing-box 1.12.0, please use route_address
instead.

`inet4_route_address`Use custom routes instead of default when auto_route is enabled.

`auto_route`#### inet6_route_address

Deprecated in sing-box 1.10.0

inet6_route_address is deprecated and will be removed in sing-box 1.12.0, please use route_address
instead.

`inet6_route_address`Use custom routes instead of default when auto_route is enabled.

`auto_route`#### route_exclude_address

Since sing-box 1.10.0

Exclude custom routes when auto_route is enabled.

`auto_route`#### inet4_route_exclude_address

Deprecated in sing-box 1.10.0

inet4_route_exclude_address is deprecated and will be removed in sing-box 1.12.0, please
use route_exclude_address instead.

`inet4_route_exclude_address`Exclude custom routes when auto_route is enabled.

`auto_route`#### inet6_route_exclude_address

Deprecated in sing-box 1.10.0

inet6_route_exclude_address is deprecated and will be removed in sing-box 1.12.0, please
use route_exclude_address instead.

`inet6_route_exclude_address`Exclude custom routes when auto_route is enabled.

`auto_route`#### route_address_set

`auto_redirect``auto_redirect`Since sing-box 1.10.0

Only supported on Linux with nftables and requires auto_route and auto_redirect enabled.

`auto_route``auto_redirect`Add the destination IP CIDR rules in the specified rule-sets to the firewall.
Unmatched traffic will bypass the sing-box routes.

Conflict with route.default_mark and [dialOptions].routing_mark.

`route.default_mark``[dialOptions].routing_mark`Since sing-box 1.11.0

Add the destination IP CIDR rules in the specified rule-sets to routes, equivalent to adding to route_address.
Unmatched traffic will bypass the sing-box routes.

`route_address`Note that it doesn't work on the Android graphical client due to
the Android VpnService not being able to handle a large number of routes (DeadSystemException),
but otherwise it works fine on all command line clients and Apple platforms.

#### route_exclude_address_set

`auto_redirect``auto_redirect`Since sing-box 1.10.0

Only supported on Linux with nftables and requires auto_route and auto_redirect enabled.

`auto_route``auto_redirect`Add the destination IP CIDR rules in the specified rule-sets to the firewall.
Matched traffic will bypass the sing-box routes.

Since sing-box 1.11.0

Add the destination IP CIDR rules in the specified rule-sets to routes, equivalent to adding to route_exclude_address.
Matched traffic will bypass the sing-box routes.

`route_exclude_address`Note that it doesn't work on the Android graphical client due to
the Android VpnService not being able to handle a large number of routes (DeadSystemException),
but otherwise it works fine on all command line clients and Apple platforms.

#### endpoint_independent_nat

This item is only available on the gvisor stack, other stacks are endpoint-independent NAT by default.

Enable endpoint-independent NAT.

Performance may degrade slightly, so it is not recommended to enable on when it is not needed.

#### udp_timeout

UDP NAT expiration time.

5m will be used by default.

`5m`#### stack

Changes in sing-box 1.8.0

The legacy LWIP stack has been deprecated and removed.

TCP/IP stack.

| Stack | Description | 
| --- | --- |
| system | Perform L3 to L4 translation using the system network stack | 
| gvisor | Perform L3 to L4 translation using gVisor's virtual network stack | 
| mixed | Mixed system TCP stack and gvisor UDP stack | 

`system``gvisor``mixed``system``gvisor`Defaults to the mixed stack if the gVisor build tag is enabled, otherwise defaults to the system stack.

`mixed``system`#### include_interface

Interface rules are only supported on Linux and require auto_route.

Limit interfaces in route. Not limited by default.

Conflict with exclude_interface.

`exclude_interface`#### exclude_interface

When strict_route enabled, return traffic to excluded interfaces will not be automatically excluded, so add them as well (example: br-lan and pppoe-wan).

`strict_route``br-lan``pppoe-wan`Exclude interfaces in route.

Conflict with include_interface.

`include_interface`#### include_uid

UID rules are only supported on Linux and require auto_route.

Limit users in route. Not limited by default.

#### include_uid_range

Limit users in route, but in range.

#### exclude_uid

Exclude users in route.

#### exclude_uid_range

Exclude users in route, but in range.

#### include_android_user

Android user and package rules are only supported on Android and require auto_route.

Limit android users in route.

| Common user | ID | 
| --- | --- |
| Main | 0 | 
| Work Profile | 10 | 

#### include_package

Limit android packages in route.

#### exclude_package

Exclude android packages in route.

#### platform

Platform-specific settings, provided by client applications.

#### platform.http_proxy

System HTTP proxy settings.

#### platform.http_proxy.enabled

Enable system HTTP proxy.

#### platform.http_proxy.server

Required

HTTP proxy server address.

#### platform.http_proxy.server_port

Required

HTTP proxy server port.

#### platform.http_proxy.bypass_domain

On Apple platforms, bypass_domain items matches hostname suffixes.

`bypass_domain`Hostnames that bypass the HTTP proxy.

#### platform.http_proxy.match_domain

Only supported in graphical clients on Apple platforms.

Hostnames that use the HTTP proxy.

### Listen Fields

See Listen Fields for details.


---

## VLESS

**Source URL**: <https://sing-box.sagernet.org/configuration/inbound/vless/>

# VLESS

### Structure

```
{
  "type": "vless",
  "tag": "vless-in",

  ... // Listen Fields

  "users": [
    {
      "name": "sekai",
      "uuid": "bf000d23-0752-40b4-affe-68f7707a9661",
      "flow": ""
    }
  ],
  "tls": {},
  "multiplex": {},
  "transport": {}
}

```

### Listen Fields

See Listen Fields for details.

### Fields

#### users

Required

VLESS users.

#### users.uuid

Required

VLESS user id.

#### users.flow

VLESS Sub-protocol.

Available values:

- xtls-rprx-vision

`xtls-rprx-vision`#### tls

TLS configuration, see TLS.

#### multiplex

See Multiplex for details.

#### transport

V2Ray Transport configuration, see V2Ray Transport.


---

## VMess

**Source URL**: <https://sing-box.sagernet.org/configuration/inbound/vmess/>

# VMess

### Structure

```
{
  "type": "vmess",
  "tag": "vmess-in",

  ... // Listen Fields

  "users": [
    {
      "name": "sekai",
      "uuid": "bf000d23-0752-40b4-affe-68f7707a9661",
      "alterId": 0
    }
  ],
  "tls": {},
  "multiplex": {},
  "transport": {}
}

```

### Listen Fields

See Listen Fields for details.

### Fields

#### users

Required

VMess users.

| Alter ID | Description | 
| --- | --- |
| 0 | Disable legacy protocol | 
| > 0 | Enable legacy protocol | 

Legacy protocol support (VMess MD5 Authentication) is provided for compatibility purposes only, use of alterId > 1 is not recommended.

#### tls

TLS configuration, see TLS.

#### multiplex

See Multiplex for details.

#### transport

V2Ray Transport configuration, see V2Ray Transport.


---

## Log

**Source URL**: <https://sing-box.sagernet.org/configuration/log/>

# Log

### Structure

```
{
  "log": {
    "disabled": false,
    "level": "info",
    "output": "box.log",
    "timestamp": true
  }
}

```

### Fields

#### disabled

Disable logging, no output after start.

#### level

Log level. One of: trace debug info warn error fatal panic.

`trace``debug``info``warn``error``fatal``panic`#### output

Output file path. Will not write log to console after enable.

#### timestamp

Add time to each line.


---

## NTP

**Source URL**: <https://sing-box.sagernet.org/configuration/ntp/>

# NTP

Built-in NTP client service.

If enabled, it will provide time for protocols like TLS/Shadowsocks/VMess, which is useful for environments where time
synchronization is not possible.

### Structure

```
{
  "ntp": {
    "enabled": false,
    "server": "time.apple.com",
    "server_port": 123,
    "interval": "30m",

    ... // Dial Fields
  }
}

```

### Fields

#### enabled

Enable NTP service.

#### server

Required

NTP server address.

#### server_port

NTP server port.

123 is used by default.

#### interval

Time synchronization interval.

30 minutes is used by default.

### Dial Fields

See Dial Fields for details.


---

## Outbound

**Source URL**: <https://sing-box.sagernet.org/configuration/outbound/>

# Outbound

### Structure

```
{
  "outbounds": [
    {
      "type": "",
      "tag": ""
    }
  ]
}

```

### Fields

| Type | Format | 
| --- | --- |
| direct | Direct | 
| block | Block | 
| socks | SOCKS | 
| http | HTTP | 
| shadowsocks | Shadowsocks | 
| vmess | VMess | 
| trojan | Trojan | 
| wireguard | Wireguard | 
| hysteria | Hysteria | 
| vless | VLESS | 
| shadowtls | ShadowTLS | 
| tuic | TUIC | 
| hysteria2 | Hysteria2 | 
| anytls | AnyTLS | 
| tor | Tor | 
| ssh | SSH | 
| dns | DNS | 
| selector | Selector | 
| urltest | URLTest | 
| naive | NaiveProxy | 

`direct``block``socks``http``shadowsocks``vmess``trojan``wireguard``hysteria``vless``shadowtls``tuic``hysteria2``anytls``tor``ssh``dns``selector``urltest``naive`#### tag

The tag of the outbound.

### Features

#### Outbounds that support IP connection

- WireGuard

`WireGuard`
---

## AnyTLS

**Source URL**: <https://sing-box.sagernet.org/configuration/outbound/anytls/>

# AnyTLS

Since sing-box 1.12.0

### Structure

```
{
  "type": "anytls",
  "tag": "anytls-out",

  "server": "127.0.0.1",
  "server_port": 1080,
  "password": "8JCsPssfgS8tiRwiMlhARg==",
  "idle_session_check_interval": "30s",
  "idle_session_timeout": "30s",
  "min_idle_session": 5,
  "tls": {},

  ... // Dial Fields
}

```

### Fields

#### server

Required

The server address.

#### server_port

Required

The server port.

#### password

Required

The AnyTLS password.

#### idle_session_check_interval

Interval checking for idle sessions. Default: 30s.

#### idle_session_timeout

In the check, close sessions that have been idle for longer than this. Default: 30s.

#### min_idle_session

In the check, at least the first n idle sessions are kept open. Default value: n=0

`n``n`#### tls

Required

TLS configuration, see TLS.

### Dial Fields

See Dial Fields for details.


---

## Block

**Source URL**: <https://sing-box.sagernet.org/configuration/outbound/block/>

# Block

### Structure

```
{
  "type": "block",
  "tag": "block"
}

```

### Fields

No fields.


---

## Direct

**Source URL**: <https://sing-box.sagernet.org/configuration/outbound/direct/>

# Direct

Changes in sing-box 1.11.0

override_address
 override_port

direct outbound send requests directly.

`direct`### Structure

```
{
  "type": "direct",
  "tag": "direct-out",

  "override_address": "1.0.0.1",
  "override_port": 53,

  ... // Dial Fields
}

```

### Fields

#### override_address

Deprecated in sing-box 1.11.0

Destination override fields are deprecated in sing-box 1.11.0 and will be removed in sing-box 1.13.0, see Migration.

Override the connection destination address.

#### override_port

Deprecated in sing-box 1.11.0

Destination override fields are deprecated in sing-box 1.11.0 and will be removed in sing-box 1.13.0, see Migration.

Override the connection destination port.

Protocol value can be 1 or 2.

`1``2`### Dial Fields

See Dial Fields for details.


---

## DNS

**Source URL**: <https://sing-box.sagernet.org/configuration/outbound/dns/>

# DNS

Deprecated in sing-box 1.11.0

Legacy special outbounds are deprecated and will be removed in sing-box 1.13.0, check Migration.

dns outbound is a internal DNS server.

`dns`### Structure

```
{
  "type": "dns",
  "tag": "dns-out"
}

```

There are no outbound connections by the DNS outbound, all requests are handled internally.

### Fields

No fields.


---

## HTTP

**Source URL**: <https://sing-box.sagernet.org/configuration/outbound/http/>

# HTTP

http outbound is a HTTP CONNECT proxy client.

`http`### Structure

```
{
  "type": "http",
  "tag": "http-out",

  "server": "127.0.0.1",
  "server_port": 1080,
  "username": "sekai",
  "password": "admin",
  "path": "",
  "headers": {},
  "tls": {},

  ... // Dial Fields
}

```

### Fields

#### server

Required

The server address.

#### server_port

Required

The server port.

#### username

Basic authorization username.

#### password

Basic authorization password.

#### path

Path of HTTP request.

#### headers

Extra headers of HTTP request.

#### tls

TLS configuration, see TLS.

### Dial Fields

See Dial Fields for details.


---

## Hysteria

**Source URL**: <https://sing-box.sagernet.org/configuration/outbound/hysteria/>

# Hysteria

Changes in sing-box 1.12.0

server_ports
 hop_interval

### Structure

```
{
  "type": "hysteria",
  "tag": "hysteria-out",

  "server": "127.0.0.1",
  "server_port": 1080,
  "server_ports": [
    "2080:3000"
  ],
  "hop_interval": "",
  "up": "100 Mbps",
  "up_mbps": 100,
  "down": "100 Mbps",
  "down_mbps": 100,
  "obfs": "fuck me till the daylight",
  "auth": "",
  "auth_str": "password",
  "recv_window_conn": 0,
  "recv_window": 0,
  "disable_mtu_discovery": false,
  "network": "tcp",
  "tls": {},

  ... // Dial Fields
}

```

### Fields

#### server

Required

The server address.

#### server_port

Required

The server port.

#### server_ports

Since sing-box 1.12.0

Server port range list.

Conflicts with server_port.

`server_port`#### hop_interval

Since sing-box 1.12.0

Port hopping interval.

30s is used by default.

`30s`#### up, down

Required

Format: [Integer] [Unit] e.g. 100 Mbps, 640 KBps, 2 Gbps

`[Integer] [Unit]``100 Mbps, 640 KBps, 2 Gbps`Supported units (case sensitive, b = bits, B = bytes, 8b=1B):

```
bps (bits per second)
Bps (bytes per second)
Kbps (kilobits per second)
KBps (kilobytes per second)
Mbps (megabits per second)
MBps (megabytes per second)
Gbps (gigabits per second)
GBps (gigabytes per second)
Tbps (terabits per second)
TBps (terabytes per second)

```

#### up_mbps, down_mbps

Required

up, down in Mbps.

`up, down`#### obfs

Obfuscated password.

#### auth

Authentication password, in base64.

#### auth_str

Authentication password.

#### recv_window_conn

The QUIC stream-level flow control window for receiving data.

15728640 (15 MB/s) will be used if empty.

`15728640 (15 MB/s)`#### recv_window

The QUIC connection-level flow control window for receiving data.

67108864 (64 MB/s) will be used if empty.

`67108864 (64 MB/s)`#### disable_mtu_discovery

Disables Path MTU Discovery (RFC 8899). Packets will then be at most 1252 (IPv4) / 1232 (IPv6) bytes in size.

Force enabled on for systems other than Linux and Windows (according to upstream).

#### network

Enabled network

One of tcp udp.

`tcp``udp`Both is enabled by default.

#### tls

Required

TLS configuration, see TLS.

### Dial Fields

See Dial Fields for details.


---

## Hysteria2

**Source URL**: <https://sing-box.sagernet.org/configuration/outbound/hysteria2/>

# Hysteria2

Changes in sing-box 1.11.0

server_ports
 hop_interval

### Structure

```
{
  "type": "hysteria2",
  "tag": "hy2-out",

  "server": "127.0.0.1",
  "server_port": 1080,
  "server_ports": [
    "2080:3000"
  ],
  "hop_interval": "",
  "up_mbps": 100,
  "down_mbps": 100,
  "obfs": {
    "type": "salamander",
    "password": "cry_me_a_r1ver"
  },
  "password": "goofy_ahh_password",
  "network": "tcp",
  "tls": {},
  "brutal_debug": false,

  ... // Dial Fields
}

```

You can ignore the JSON Array [] tag when the content is only one item

Difference from official Hysteria2

The official Hysteria2 supports an authentication method called userpass,
which essentially uses a combination of <username>:<password> as the actual password,
while sing-box does not provide this alias.
If you are planning to use sing-box with the official program,
please note that you will need to fill the combination as the password.

`<username>:<password>`### Fields

#### server

Required

The server address.

#### server_port

Required

The server port.

Ignored if server_ports is set.

`server_ports`#### server_ports

Since sing-box 1.11.0

Server port range list.

Conflicts with server_port.

`server_port`#### hop_interval

Since sing-box 1.11.0

Port hopping interval.

30s is used by default.

`30s`#### up_mbps, down_mbps

Max bandwidth, in Mbps.

If empty, the BBR congestion control algorithm will be used instead of Hysteria CC.

#### obfs.type

QUIC traffic obfuscator type, only available with salamander.

`salamander`Disabled if empty.

#### obfs.password

QUIC traffic obfuscator password.

#### password

Authentication password.

#### network

Enabled network

One of tcp udp.

`tcp``udp`Both is enabled by default.

#### tls

Required

TLS configuration, see TLS.

#### brutal_debug

Enable debug information logging for Hysteria Brutal CC.

### Dial Fields

See Dial Fields for details.


---

## Naive

**Source URL**: <https://sing-box.sagernet.org/configuration/outbound/naive/>

# Naive

Since sing-box 1.13.0

### Structure

```
{
  "type": "naive",
  "tag": "naive-out",

  "server": "127.0.0.1",
  "server_port": 443,
  "username": "sekai",
  "password": "password",
  "insecure_concurrency": 0,
  "extra_headers": {},
  "udp_over_tcp": false | {},
  "quic": false,
  "quic_congestion_control": "",
  "tls": {},

  ... // Dial Fields
}

```

Platform Support

NaiveProxy outbound is only available on Apple platforms, Android, Windows and certain Linux builds.

Official Release Build Variants:

| Build Variant | Platforms | Description | 
| --- | --- | --- |
| (default) | Linux amd64/arm64 | purego build with libcronet.so included | 
| -glibc | Linux 386/amd64/arm/arm64 | CGO build dynamically linked with glibc, requires glibc >= 2.31 | 
| -musl | Linux 386/amd64/arm/arm64 | CGO build statically linked with musl, no system requirements | 
| (default) | Windows amd64/arm64 | purego build with libcronet.dll included | 

`libcronet.so``-glibc``-musl``libcronet.dll`Runtime Requirements:

- Linux purego: libcronet.so must be in the same directory as the sing-box binary or in system library path
- Windows: libcronet.dll must be in the same directory as sing-box.exe or in a directory listed in PATH

`libcronet.so``libcronet.dll``sing-box.exe``PATH`For self-built binaries, see Build from source.

### Fields

#### server

Required

The server address.

#### server_port

Required

The server port.

#### username

Authentication username.

#### password

Authentication password.

#### insecure_concurrency

Number of concurrent tunnel connections. Multiple connections make the tunneling easier to detect through traffic analysis, which defeats the purpose of NaiveProxy's design to resist traffic analysis.

#### extra_headers

Extra headers to send in HTTP requests.

#### udp_over_tcp

UDP over TCP protocol settings.

See UDP Over TCP for details.

#### quic

Use QUIC instead of HTTP/2.

#### quic_congestion_control

QUIC congestion control algorithm.

| Algorithm | Description | 
| --- | --- |
| bbr | BBR | 
| bbr2 | BBRv2 | 
| cubic | CUBIC | 
| reno | New Reno | 

`bbr``bbr2``cubic``reno`bbr is used by default (the default of QUICHE, used by Chromium which NaiveProxy is based on).

`bbr`#### tls

Required

TLS configuration, see TLS.

Only server_name, certificate, certificate_path and ech are supported.

`server_name``certificate``certificate_path``ech`Self-signed certificates change traffic behavior significantly, which defeats the purpose of NaiveProxy's design to resist traffic analysis, and should not be used in production.

### Dial Fields

See Dial Fields for details.


---

## Selector

**Source URL**: <https://sing-box.sagernet.org/configuration/outbound/selector/>

# Selector

### Structure

```
{
  "type": "selector",
  "tag": "select",

  "outbounds": [
    "proxy-a",
    "proxy-b",
    "proxy-c"
  ],
  "default": "proxy-c",
  "interrupt_exist_connections": false
}

```

The selector can only be controlled through the Clash API currently.

### Fields

#### outbounds

Required

List of outbound tags to select.

#### default

The default outbound tag. The first outbound will be used if empty.

#### interrupt_exist_connections

Interrupt existing connections when the selected outbound has changed.

Only inbound connections are affected by this setting, internal connections will always be interrupted.


---

## Shadowsocks

**Source URL**: <https://sing-box.sagernet.org/configuration/outbound/shadowsocks/>

# Shadowsocks

### Structure

```
{
  "type": "shadowsocks",
  "tag": "ss-out",

  "server": "127.0.0.1",
  "server_port": 1080,
  "method": "2022-blake3-aes-128-gcm",
  "password": "8JCsPssfgS8tiRwiMlhARg==",
  "plugin": "",
  "plugin_opts": "",
  "network": "udp",
  "udp_over_tcp": false | {},
  "multiplex": {},

  ... // Dial Fields
}

```

### Fields

#### server

Required

The server address.

#### server_port

Required

The server port.

#### method

Required

Encryption methods:

- 2022-blake3-aes-128-gcm
- 2022-blake3-aes-256-gcm
- 2022-blake3-chacha20-poly1305
- none
- aes-128-gcm
- aes-192-gcm
- aes-256-gcm
- chacha20-ietf-poly1305
- xchacha20-ietf-poly1305

`2022-blake3-aes-128-gcm``2022-blake3-aes-256-gcm``2022-blake3-chacha20-poly1305``none``aes-128-gcm``aes-192-gcm``aes-256-gcm``chacha20-ietf-poly1305``xchacha20-ietf-poly1305`Legacy encryption methods:

- aes-128-ctr
- aes-192-ctr
- aes-256-ctr
- aes-128-cfb
- aes-192-cfb
- aes-256-cfb
- rc4-md5
- chacha20-ietf
- xchacha20

`aes-128-ctr``aes-192-ctr``aes-256-ctr``aes-128-cfb``aes-192-cfb``aes-256-cfb``rc4-md5``chacha20-ietf``xchacha20`#### password

Required

The shadowsocks password.

#### plugin

Shadowsocks SIP003 plugin, implemented in internal.

Only obfs-local and v2ray-plugin are supported.

`obfs-local``v2ray-plugin`#### plugin_opts

Shadowsocks SIP003 plugin options.

#### network

Enabled network

One of tcp udp.

`tcp``udp`Both is enabled by default.

#### udp_over_tcp

UDP over TCP configuration.

See UDP Over TCP for details.

Conflict with multiplex.

`multiplex`#### multiplex

See Multiplex for details.

### Dial Fields

See Dial Fields for details.


---

## ShadowTLS

**Source URL**: <https://sing-box.sagernet.org/configuration/outbound/shadowtls/>

# ShadowTLS

### Structure

```
{
  "type": "shadowtls",
  "tag": "st-out",

  "server": "127.0.0.1",
  "server_port": 1080,
  "version": 3,
  "password": "fuck me till the daylight",
  "tls": {},

  ... // Dial Fields
}

```

### Fields

#### server

Required

The server address.

#### server_port

Required

The server port.

#### version

ShadowTLS protocol version.

| Value | Protocol Version | 
| --- | --- |
| 1 (default) | ShadowTLS v1 | 
| 2 | ShadowTLS v2 | 
| 3 | ShadowTLS v3 | 

`1``2``3`#### password

Set password.

Only available in the ShadowTLS v2/v3 protocol.

#### tls

Required

TLS configuration, see TLS.

### Dial Fields

See Dial Fields for details.


---

## SOCKS

**Source URL**: <https://sing-box.sagernet.org/configuration/outbound/socks/>

# SOCKS

socks outbound is a socks4/socks4a/socks5 client.

`socks`### Structure

```
{
  "type": "socks",
  "tag": "socks-out",

  "server": "127.0.0.1",
  "server_port": 1080,
  "version": "5",
  "username": "sekai",
  "password": "admin",
  "network": "udp",
  "udp_over_tcp": false | {},

  ... // Dial Fields
}

```

### Fields

#### server

Required

The server address.

#### server_port

Required

The server port.

#### version

The SOCKS version, one of 4 4a 5.

`4``4a``5`SOCKS5 used by default.

#### username

SOCKS username.

#### password

SOCKS5 password.

#### network

Enabled network

One of tcp udp.

`tcp``udp`Both is enabled by default.

#### udp_over_tcp

UDP over TCP protocol settings.

See UDP Over TCP for details.

### Dial Fields

See Dial Fields for details.


---

## SSH

**Source URL**: <https://sing-box.sagernet.org/configuration/outbound/ssh/>

# SSH

### Structure

```
{
  "type": "ssh",
  "tag": "ssh-out",

  "server": "127.0.0.1",
  "server_port": 22,
  "user": "root",
  "password": "admin",
  "private_key": "",
  "private_key_path": "$HOME/.ssh/id_rsa",
  "private_key_passphrase": "",
  "host_key": [
    "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdH..."
  ],
  "host_key_algorithms": [],
  "client_version": "SSH-2.0-OpenSSH_7.4p1",

  ... // Dial Fields
}

```

### Fields

#### server

Required

Server address.

#### server_port

Server port. 22 will be used if empty.

#### user

SSH user, root will be used if empty.

#### password

Password.

#### private_key

Private key.

#### private_key_path

Private key path.

#### private_key_passphrase

Private key passphrase.

#### host_key

Host key. Accept any if empty.

#### host_key_algorithms

Host key algorithms.

#### client_version

Client version. Random version will be used if empty.

### Dial Fields

See Dial Fields for details.


---

## Tor

**Source URL**: <https://sing-box.sagernet.org/configuration/outbound/tor/>

# Tor

### Structure

```
{
  "type": "tor",
  "tag": "tor-out",

  "executable_path": "/usr/bin/tor",
  "extra_args": [],
  "data_directory": "$HOME/.cache/tor",
  "torrc": {
    "ClientOnly": 1
  },

  ... // Dial Fields
}

```

Embedded Tor is not included by default, see Installation.

### Fields

#### executable_path

The path to the Tor executable.

Embedded Tor will be ignored if set.

#### extra_args

List of extra arguments passed to the Tor instance when started.

#### data_directory

Recommended

The data directory of Tor.

Each start will be very slow if not specified.

#### torrc

Map of torrc options.

See tor(1) for details.

### Dial Fields

See Dial Fields for details.


---

## Trojan

**Source URL**: <https://sing-box.sagernet.org/configuration/outbound/trojan/>

# Trojan

### Structure

```
{
  "type": "trojan",
  "tag": "trojan-out",

  "server": "127.0.0.1",
  "server_port": 1080,
  "password": "8JCsPssfgS8tiRwiMlhARg==",
  "network": "tcp",
  "tls": {},
  "multiplex": {},
  "transport": {},

  ... // Dial Fields
}

```

### Fields

#### server

Required

The server address.

#### server_port

Required

The server port.

#### password

Required

The Trojan password.

#### network

Enabled network

One of tcp udp.

`tcp``udp`Both is enabled by default.

#### tls

TLS configuration, see TLS.

#### multiplex

See Multiplex for details.

#### transport

V2Ray Transport configuration, see V2Ray Transport.

### Dial Fields

See Dial Fields for details.


---

## TUIC

**Source URL**: <https://sing-box.sagernet.org/configuration/outbound/tuic/>

# TUIC

### Structure

```
{
  "type": "tuic",
  "tag": "tuic-out",

  "server": "127.0.0.1",
  "server_port": 1080,
  "uuid": "2DD61D93-75D8-4DA4-AC0E-6AECE7EAC365",
  "password": "hello",
  "congestion_control": "cubic",
  "udp_relay_mode": "native",
  "udp_over_stream": false,
  "zero_rtt_handshake": false,
  "heartbeat": "10s",
  "network": "tcp",
  "tls": {},

  ... // Dial Fields
}

```

### Fields

#### server

Required

The server address.

#### server_port

Required

The server port.

#### uuid

Required

TUIC user uuid

#### password

TUIC user password

#### congestion_control

QUIC congestion control algorithm

One of: cubic, new_reno, bbr

`cubic``new_reno``bbr`cubic is used by default.

`cubic`#### udp_relay_mode

UDP packet relay mode

| Mode | Description | 
| --- | --- |
| native | native UDP characteristics | 
| quic | lossless UDP relay using QUIC streams, additional overhead is introduced | 

native is used by default.

`native`Conflict with udp_over_stream.

`udp_over_stream`#### udp_over_stream

This is the TUIC port of the UDP over TCP protocol, designed to provide a QUIC
stream based UDP relay mode that TUIC does not provide. Since it is an add-on protocol, you will need to use sing-box or
another program compatible with the protocol as a server.

This mode has no positive effect in a proper UDP proxy scenario and should only be applied to relay streaming UDP
traffic (basically QUIC streams).

Conflict with udp_relay_mode.

`udp_relay_mode`#### network

Enabled network

One of tcp udp.

`tcp``udp`Both is enabled by default.

#### tls

Required

TLS configuration, see TLS.

### Dial Fields

See Dial Fields for details.


---

## URLTest

**Source URL**: <https://sing-box.sagernet.org/configuration/outbound/urltest/>

# URLTest

### Structure

```
{
  "type": "urltest",
  "tag": "auto",

  "outbounds": [
    "proxy-a",
    "proxy-b",
    "proxy-c"
  ],
  "url": "",
  "interval": "",
  "tolerance": 0,
  "idle_timeout": "",
  "interrupt_exist_connections": false
}

```

### Fields

#### outbounds

Required

List of outbound tags to test.

#### url

The URL to test. https://www.gstatic.com/generate_204 will be used if empty.

`https://www.gstatic.com/generate_204`#### interval

The test interval. 3m will be used if empty.

`3m`#### tolerance

The test tolerance in milliseconds. 50 will be used if empty.

`50`#### idle_timeout

The idle timeout. 30m will be used if empty.

`30m`#### interrupt_exist_connections

Interrupt existing connections when the selected outbound has changed.

Only inbound connections are affected by this setting, internal connections will always be interrupted.


---

## VLESS

**Source URL**: <https://sing-box.sagernet.org/configuration/outbound/vless/>

# VLESS

### Structure

```
{
  "type": "vless",
  "tag": "vless-out",

  "server": "127.0.0.1",
  "server_port": 1080,
  "uuid": "bf000d23-0752-40b4-affe-68f7707a9661",
  "flow": "xtls-rprx-vision",
  "network": "tcp",
  "tls": {},
  "packet_encoding": "",
  "multiplex": {},
  "transport": {},

  ... // Dial Fields
}

```

### Fields

#### server

Required

The server address.

#### server_port

Required

The server port.

#### uuid

Required

VLESS user id.

#### flow

VLESS Sub-protocol.

Available values:

- xtls-rprx-vision

`xtls-rprx-vision`#### network

Enabled network

One of tcp udp.

`tcp``udp`Both is enabled by default.

#### tls

TLS configuration, see TLS.

#### packet_encoding

UDP packet encoding, xudp is used by default.

| Encoding | Description | 
| --- | --- |
| (none) | Disabled | 
| packetaddr | Supported by v2ray 5+ | 
| xudp | Supported by xray | 

#### multiplex

See Multiplex for details.

#### transport

V2Ray Transport configuration, see V2Ray Transport.

### Dial Fields

See Dial Fields for details.


---

## VMess

**Source URL**: <https://sing-box.sagernet.org/configuration/outbound/vmess/>

# VMess

### Structure

```
{
  "type": "vmess",
  "tag": "vmess-out",

  "server": "127.0.0.1",
  "server_port": 1080,
  "uuid": "bf000d23-0752-40b4-affe-68f7707a9661",
  "security": "auto",
  "alter_id": 0,
  "global_padding": false,
  "authenticated_length": true,
  "network": "tcp",
  "tls": {},
  "packet_encoding": "",
  "transport": {},
  "multiplex": {},

  ... // Dial Fields
}

```

### Fields

#### server

Required

The server address.

#### server_port

Required

The server port.

#### uuid

Required

The VMess user id.

#### security

Encryption methods:

- auto
- none
- zero
- aes-128-gcm
- chacha20-poly1305

`auto``none``zero``aes-128-gcm``chacha20-poly1305`Legacy encryption methods:

- aes-128-ctr

`aes-128-ctr`#### alter_id

| Alter ID | Description | 
| --- | --- |
| 0 | Use AEAD protocol | 
| 1 | Use legacy protocol | 
| > 1 | Unused, same as 1 | 

#### global_padding

Protocol parameter. Will waste traffic randomly if enabled (enabled by default in v2ray and cannot be disabled).

#### authenticated_length

Protocol parameter. Enable length block encryption.

#### network

Enabled network

One of tcp udp.

`tcp``udp`Both is enabled by default.

#### tls

TLS configuration, see TLS.

#### packet_encoding

UDP packet encoding.

| Encoding | Description | 
| --- | --- |
| (none) | Disabled | 
| packetaddr | Supported by v2ray 5+ | 
| xudp | Supported by xray | 

#### multiplex

See Multiplex for details.

#### transport

V2Ray Transport configuration, see V2Ray Transport.

### Dial Fields

See Dial Fields for details.


---

## WireGuard

**Source URL**: <https://sing-box.sagernet.org/configuration/outbound/wireguard/>

# WireGuard

Deprecated in sing-box 1.11.0

WireGuard outbound is deprecated and will be removed in sing-box 1.13.0, check Migration.

Changes in sing-box 1.11.0

gso

Changes in sing-box 1.8.0

gso

### Structure

```
{
  "type": "wireguard",
  "tag": "wireguard-out",

  "server": "127.0.0.1",
  "server_port": 1080,
  "system_interface": false,
  "interface_name": "wg0",
  "local_address": [
    "10.0.0.1/32"
  ],
  "private_key": "YNXtAzepDqRv9H52osJVDQnznT5AM11eCK3ESpwSt04=",
  "peers": [
    {
      "server": "127.0.0.1",
      "server_port": 1080,
      "public_key": "Z1XXLsKYkYxuiYjJIkRvtIKFepCYHTgON+GwPq7SOV4=",
      "pre_shared_key": "31aIhAPwktDGpH4JDhA8GNvjFXEf/a6+UaQRyOAiyfM=",
      "allowed_ips": [
        "0.0.0.0/0"
      ],
      "reserved": [0, 0, 0]
    }
  ],
  "peer_public_key": "Z1XXLsKYkYxuiYjJIkRvtIKFepCYHTgON+GwPq7SOV4=",
  "pre_shared_key": "31aIhAPwktDGpH4JDhA8GNvjFXEf/a6+UaQRyOAiyfM=",
  "reserved": [0, 0, 0],
  "workers": 4,
  "mtu": 1408,
  "network": "tcp",

  // Deprecated

  "gso": false,

  ... // Dial Fields
}

```

### Fields

#### server

Required if multi-peer disabled

The server address.

#### server_port

Required if multi-peer disabled

The server port.

#### system_interface

Use system interface.

Requires privilege and cannot conflict with exists system interfaces.

Forced if gVisor not included in the build.

#### interface_name

Custom interface name for system interface.

#### gso

Deprecated in sing-box 1.11.0

GSO will be automatically enabled when available since sing-box 1.11.0.

Since sing-box 1.8.0

Only supported on Linux.

Try to enable generic segmentation offload.

#### local_address

Required

List of IP (v4 or v6) address prefixes to be assigned to the interface.

#### private_key

Required

WireGuard requires base64-encoded public and private keys. These can be generated using the wg(8) utility:

```
wg genkey
echo "private key" || wg pubkey

```

#### peers

Multi-peer support.

If enabled, server, server_port, peer_public_key, pre_shared_key will be ignored.

`server, server_port, peer_public_key, pre_shared_key`#### peers.allowed_ips

WireGuard allowed IPs.

#### peers.reserved

WireGuard reserved field bytes.

$outbound.reserved will be used if empty.

`$outbound.reserved`#### peer_public_key

Required if multi-peer disabled

WireGuard peer public key.

#### pre_shared_key

WireGuard pre-shared key.

#### reserved

WireGuard reserved field bytes.

#### workers

WireGuard worker count.

CPU count is used by default.

#### mtu

WireGuard MTU.

1408 will be used if empty.

#### network

Enabled network

One of tcp udp.

`tcp``udp`Both is enabled by default.

### Dial Fields

See Dial Fields for details.


---

## Route

**Source URL**: <https://sing-box.sagernet.org/configuration/route/>

# Route

Changes in sing-box 1.12.0

default_domain_resolver
 geoip
 geosite

Changes in sing-box 1.11.0

default_network_strategy
 default_network_type
 default_fallback_network_type
 default_fallback_delay

Changes in sing-box 1.8.0

rule_set
 geoip
 geosite

### Structure

```
{
  "route": {
    "rules": [],
    "rule_set": [],
    "final": "",
    "auto_detect_interface": false,
    "override_android_vpn": false,
    "default_interface": "",
    "default_mark": 0,
    "default_domain_resolver": "", // or {}
    "default_network_strategy": "",
    "default_network_type": [],
    "default_fallback_network_type": [],
    "default_fallback_delay": "",

    // Removed

    "geoip": {},
    "geosite": {}
  }
}

```

You can ignore the JSON Array [] tag when the content is only one item

### Fields

#### rules

List of Route Rule

#### rule_set

Since sing-box 1.8.0

List of rule-set

#### final

Default outbound tag. the first outbound will be used if empty.

#### auto_detect_interface

Only supported on Linux, Windows and macOS.

Bind outbound connections to the default NIC by default to prevent routing loops under tun.

Takes no effect if outbound.bind_interface is set.

`outbound.bind_interface`#### override_android_vpn

Only supported on Android.

Accept Android VPN as upstream NIC when auto_detect_interface enabled.

`auto_detect_interface`#### default_interface

Only supported on Linux, Windows and macOS.

Bind outbound connections to the specified NIC by default to prevent routing loops under tun.

Takes no effect if auto_detect_interface is set.

`auto_detect_interface`#### default_mark

Only supported on Linux.

Set routing mark by default.

Takes no effect if outbound.routing_mark is set.

`outbound.routing_mark`#### default_domain_resolver

Since sing-box 1.12.0

See Dial Fields for details.

Can be overrides by outbound.domain_resolver.

`outbound.domain_resolver`#### default_network_strategy

Since sing-box 1.11.0

See Dial Fields for details.

Takes no effect if outbound.bind_interface, outbound.inet4_bind_address or outbound.inet6_bind_address is set.

`outbound.bind_interface``outbound.inet4_bind_address``outbound.inet6_bind_address`Can be overrides by outbound.network_strategy.

`outbound.network_strategy`Conflicts with default_interface.

`default_interface`#### default_network_type

Since sing-box 1.11.0

See Dial Fields for details.

#### default_fallback_network_type

Since sing-box 1.11.0

See Dial Fields for details.

#### default_fallback_delay

Since sing-box 1.11.0

See Dial Fields for details.


---

## GeoIP

**Source URL**: <https://sing-box.sagernet.org/configuration/route/geoip/>

# GeoIP

Removed in sing-box 1.12.0

GeoIP is deprecated in sing-box 1.8.0 and removed in sing-box 1.12.0, check Migration.

### Structure

```
{
  "route": {
    "geoip": {
      "path": "",
      "download_url": "",
      "download_detour": ""
    }
  }
}

```

### Fields

#### path

The path to the sing-geoip database.

geoip.db will be used if empty.

`geoip.db`#### download_url

The download URL of the sing-geoip database.

Default is https://github.com/SagerNet/sing-geoip/releases/latest/download/geoip.db.

`https://github.com/SagerNet/sing-geoip/releases/latest/download/geoip.db`#### download_detour

The tag of the outbound to download the database.

Default outbound will be used if empty.


---

## Geosite

**Source URL**: <https://sing-box.sagernet.org/configuration/route/geosite/>

# Geosite

Removed in sing-box 1.12.0

Geosite is deprecated in sing-box 1.8.0 and removed in sing-box 1.12.0, check Migration.

### Structure

```
{
  "route": {
    "geosite": {
      "path": "",
      "download_url": "",
      "download_detour": ""
    }
  }
}

```

### Fields

#### path

The path to the sing-geosite database.

geosite.db will be used if empty.

`geosite.db`#### download_url

The download URL of the sing-geoip database.

Default is https://github.com/SagerNet/sing-geosite/releases/latest/download/geosite.db.

`https://github.com/SagerNet/sing-geosite/releases/latest/download/geosite.db`#### download_detour

The tag of the outbound to download the database.

Default outbound will be used if empty.


---

## Route Rule

**Source URL**: <https://sing-box.sagernet.org/configuration/route/rule/>

# Route Rule

Changes in sing-box 1.13.0

interface_address
 network_interface_address
 default_interface_address
 preferred_by
 network

Changes in sing-box 1.11.0

action
 outbound
 network_type
 network_is_expensive
 network_is_constrained

Changes in sing-box 1.10.0

client
 rule_set_ipcidr_match_source
 rule_set_ip_cidr_match_source
 process_path_regex

Changes in sing-box 1.8.0

rule_set
 rule_set_ipcidr_match_source
 source_ip_is_private
 ip_is_private
 source_geoip
 geoip
 geosite

### Structure

```
{
  "route": {
    "rules": [
      {
        "inbound": [
          "mixed-in"
        ],
        "ip_version": 6,
        "network": [
          "tcp"
        ],
        "auth_user": [
          "usera",
          "userb"
        ],
        "protocol": [
          "tls",
          "http",
          "quic"
        ],
        "client": [
          "chromium",
          "safari",
          "firefox",
          "quic-go"
        ],
        "domain": [
          "test.com"
        ],
        "domain_suffix": [
          ".cn"
        ],
        "domain_keyword": [
          "test"
        ],
        "domain_regex": [
          "^stun\\..+"
        ],
        "geosite": [
          "cn"
        ],
        "source_geoip": [
          "private"
        ],
        "geoip": [
          "cn"
        ],
        "source_ip_cidr": [
          "10.0.0.0/24",
          "192.168.0.1"
        ],
        "source_ip_is_private": false,
        "ip_cidr": [
          "10.0.0.0/24",
          "192.168.0.1"
        ],
        "ip_is_private": false,
        "source_port": [
          12345
        ],
        "source_port_range": [
          "1000:2000",
          ":3000",
          "4000:"
        ],
        "port": [
          80,
          443
        ],
        "port_range": [
          "1000:2000",
          ":3000",
          "4000:"
        ],
        "process_name": [
          "curl"
        ],
        "process_path": [
          "/usr/bin/curl"
        ],
        "process_path_regex": [
          "^/usr/bin/.+"
        ],
        "package_name": [
          "com.termux"
        ],
        "user": [
          "sekai"
        ],
        "user_id": [
          1000
        ],
        "clash_mode": "direct",
        "network_type": [
          "wifi"
        ],
        "network_is_expensive": false,
        "network_is_constrained": false,
        "interface_address": {
          "en0": [
            "2000::/3"
          ]
        },
        "network_interface_address": {
          "wifi": [
            "2000::/3"
          ]
        },
        "default_interface_address": [
          "2000::/3"
        ],
        "wifi_ssid": [
          "My WIFI"
        ],
        "wifi_bssid": [
          "00:00:00:00:00:00"
        ],
        "preferred_by": [
          "tailscale",
          "wireguard"
        ],
        "rule_set": [
          "geoip-cn",
          "geosite-cn"
        ],
        // deprecated
        "rule_set_ipcidr_match_source": false,
        "rule_set_ip_cidr_match_source": false,
        "invert": false,
        "action": "route",
        "outbound": "direct"
      },
      {
        "type": "logical",
        "mode": "and",
        "rules": [],
        "invert": false,
        "action": "route",
        "outbound": "direct"
      }
    ]
  }
}

```

You can ignore the JSON Array [] tag when the content is only one item

### Default Fields

The default rule uses the following matching logic:
(domain || domain_suffix || domain_keyword || domain_regex || geosite || geoip || ip_cidr || ip_is_private) &&
(port || port_range) &&
(source_geoip || source_ip_cidr || source_ip_is_private) &&
(source_port || source_port_range) &&
other fields

`domain``domain_suffix``domain_keyword``domain_regex``geosite``geoip``ip_cidr``ip_is_private``port``port_range``source_geoip``source_ip_cidr``source_ip_is_private``source_port``source_port_range``other fields`Additionally, included rule-sets can be considered merged rather than as a single rule sub-item.

#### inbound

Tags of Inbound.

#### ip_version

4 or 6.

Not limited if empty.

#### auth_user

Username, see each inbound for details.

#### protocol

Sniffed protocol, see Protocol Sniff for details.

#### client

Since sing-box 1.10.0

Sniffed client type, see Protocol Sniff for details.

#### network

Changes in sing-box 1.13.0

Since sing-box 1.13.0, you can match ICMP echo (ping) requests via the new icmp network.

`icmp`Such traffic originates from TUN, WireGuard, and Tailscale inbounds and can be routed to Direct, WireGuard, and Tailscale outbounds.

`TUN``WireGuard``Tailscale``Direct``WireGuard``Tailscale`Match network type.

tcp, udp or icmp.

`tcp``udp``icmp`#### domain

Match full domain.

#### domain_suffix

Match domain suffix.

#### domain_keyword

Match domain using keyword.

#### domain_regex

Match domain using regular expression.

#### geosite

Deprecated in sing-box 1.8.0

Geosite is deprecated and will be removed in sing-box 1.12.0, check Migration.

Match geosite.

#### source_geoip

Deprecated in sing-box 1.8.0

GeoIP is deprecated and will be removed in sing-box 1.12.0, check Migration.

Match source geoip.

#### geoip

Deprecated in sing-box 1.8.0

GeoIP is deprecated and will be removed in sing-box 1.12.0, check Migration.

Match geoip.

#### source_ip_cidr

Match source IP CIDR.

#### ip_is_private

Since sing-box 1.8.0

Match non-public IP.

#### ip_cidr

Match IP CIDR.

#### source_ip_is_private

Since sing-box 1.8.0

Match non-public source IP.

#### source_port

Match source port.

#### source_port_range

Match source port range.

#### port

Match port.

#### port_range

Match port range.

#### process_name

Only supported on Linux, Windows, and macOS.

Match process name.

#### process_path

Only supported on Linux, Windows, and macOS.

Match process path.

#### process_path_regex

Since sing-box 1.10.0

Only supported on Linux, Windows, and macOS.

Match process path using regular expression.

#### package_name

Match android package name.

#### user

Only supported on Linux.

Match user name.

#### user_id

Only supported on Linux.

Match user id.

#### clash_mode

Match Clash mode.

#### network_type

Since sing-box 1.11.0

Only supported in graphical clients on Android and Apple platforms.

Match network type.

Available values: wifi, cellular, ethernet and other.

`wifi``cellular``ethernet``other`#### network_is_expensive

Since sing-box 1.11.0

Only supported in graphical clients on Android and Apple platforms.

Match if network is considered Metered (on Android) or considered expensive,
such as Cellular or a Personal Hotspot (on Apple platforms).

#### network_is_constrained

Since sing-box 1.11.0

Only supported in graphical clients on Apple platforms.

Match if network is in Low Data Mode.

#### interface_address

Since sing-box 1.13.0

Only supported on Linux, Windows, and macOS.

Match interface address.

#### network_interface_address

Since sing-box 1.13.0

Only supported in graphical clients on Android and Apple platforms.

Matches network interface (same values as network_type) address.

`network_type`#### default_interface_address

Since sing-box 1.13.0

Only supported on Linux, Windows, and macOS.

Match default interface address.

#### wifi_ssid

Match WiFi SSID.

See Wi-Fi State for details.

#### wifi_bssid

Match WiFi BSSID.

See Wi-Fi State for details.

#### preferred_by

Since sing-box 1.13.0

Match specified outbounds' preferred routes.

| Type | Match | 
| --- | --- |
| tailscale | Match MagicDNS domains and peers' allowed IPs | 
| wireguard | Match peers's allowed IPs | 

`tailscale``wireguard`#### rule_set

Since sing-box 1.8.0

Match rule-set.

#### rule_set_ipcidr_match_source

Since sing-box 1.8.0

Deprecated in sing-box 1.10.0

rule_set_ipcidr_match_source is renamed to rule_set_ip_cidr_match_source and will be remove in sing-box 1.11.0.

`rule_set_ipcidr_match_source``rule_set_ip_cidr_match_source`Make ip_cidr in rule-sets match the source IP.

`ip_cidr`#### rule_set_ip_cidr_match_source

Since sing-box 1.10.0

Make ip_cidr in rule-sets match the source IP.

`ip_cidr`#### invert

Invert match result.

#### action

Required

See Rule Actions for details.

#### outbound

Deprecated in sing-box 1.11.0

Moved to Rule Action.

### Logical Fields

#### type

logical

`logical`#### mode

Required

and or or

`and``or`#### rules

Required

Included rules.


---

## Rule Action

**Source URL**: <https://sing-box.sagernet.org/configuration/route/rule_action/>

# Rule Action

Changes in sing-box 1.13.0

bypass
 reject

Changes in sing-box 1.12.0

tls_fragment
 tls_fragment_fallback_delay
 tls_record_fragment
 resolve.disable_cache
 resolve.rewrite_ttl
 resolve.client_subnet

## Final actions

### route

```
{
  "action": "route", // default
  "outbound": "",

  ... // route-options Fields
}

```

You can ignore the JSON Array [] tag when the content is only one item

route inherits the classic rule behavior of routing connection to the specified outbound.

`route`#### outbound

Required

Tag of target outbound.

#### route-options Fields

See route-options fields below.

`route-options`### bypass

Since sing-box 1.13.0

Only supported on Linux with auto_redirect enabled.

`auto_redirect````
{
  "action": "bypass",
  "outbound": "",

  ... // route-options Fields
}

```

bypass bypasses sing-box at the kernel level for auto redirect connections in pre-match.

`bypass`For non-auto-redirect connections and already established connections,
if outbound is specified, the behavior is the same as route;
otherwise, the rule will be skipped.

`outbound``route`#### outbound

Tag of target outbound.

If not specified, the rule only matches in pre-match
from auto redirect, and will be skipped in other contexts.

#### route-options Fields

See route-options fields below.

`route-options`### reject

Changes in sing-box 1.13.0

Since sing-box 1.13.0, you can reject (or directly reply to) ICMP echo (ping) requests using reject action.

`reject````
{
  "action": "reject",
  "method": "default", // default
  "no_drop": false
}

```

reject reject connections

`reject`The specified method is used for reject tun connections if sniff action has not been performed yet.

`sniff`For non-tun connections and already established connections, will just be closed.

#### method

For TCP and UDP connections:

- default: Reply with TCP RST for TCP connections, and ICMP port unreachable for UDP packets.
- drop: Drop packets.

`default``drop`For ICMP echo requests:

- default: Reply with ICMP host unreachable.
- drop: Drop packets.
- reply: Reply with ICMP echo reply.

`default``drop``reply`#### no_drop

If not enabled, method will be temporarily overwritten to drop after 50 triggers in 30s.

`method``drop`Not available when method is set to drop.

`method`### hijack-dns

```
{
  "action": "hijack-dns"
}

```

hijack-dns hijack DNS requests to the sing-box DNS module.

`hijack-dns`## Non-final actions

### route-options

```
{
  "action": "route-options",
  "override_address": "",
  "override_port": 0,
  "network_strategy": "",
  "fallback_delay": "",
  "udp_disable_domain_unmapping": false,
  "udp_connect": false,
  "udp_timeout": "",
  "tls_fragment": false,
  "tls_fragment_fallback_delay": "",
  "tls_record_fragment": ""
}

```

route-options set options for routing.

`route-options`#### override_address

Override the connection destination address.

#### override_port

Override the connection destination port.

#### network_strategy

See Dial Fields for details.

Only take effect if outbound is direct without outbound.bind_interface,
outbound.inet4_bind_address and outbound.inet6_bind_address set.

`outbound.bind_interface``outbound.inet4_bind_address``outbound.inet6_bind_address`#### network_type

See Dial Fields for details.

#### fallback_network_type

See Dial Fields for details.

#### fallback_delay

See Dial Fields for details.

#### udp_disable_domain_unmapping

If enabled, for UDP proxy requests addressed to a domain,
the original packet address will be sent in the response instead of the mapped domain.

This option is used for compatibility with clients that
do not support receiving UDP packets with domain addresses, such as Surge.

#### udp_connect

If enabled, attempts to connect UDP connection to the destination instead of listen.

#### udp_timeout

Timeout for UDP connections.

Setting a larger value than the UDP timeout in inbounds will have no effect.

Default value for protocol sniffed connections:

| Timeout | Protocol | 
| --- | --- |
| 10s | dns, ntp, stun | 
| 30s | quic, dtls | 

`10s``dns``ntp``stun``30s``quic``dtls`If no protocol is sniffed, the following ports will be recognized as protocols by default:

| Port | Protocol | 
| --- | --- |
| 53 | dns | 
| 123 | ntp | 
| 443 | quic | 
| 3478 | stun | 

`dns``ntp``quic``stun`#### tls_fragment

Since sing-box 1.12.0

Fragment TLS handshakes to bypass firewalls.

This feature is intended to circumvent simple firewalls based on plaintext packet matching,
and should not be used to circumvent real censorship.

Due to poor performance, try tls_record_fragment first, and only apply to server names known to be blocked.

`tls_record_fragment`On Linux, Apple platforms, (administrator privileges required) Windows,
the wait time can be automatically detected. Otherwise, it will fall back to
waiting for a fixed time specified by tls_fragment_fallback_delay.

`tls_fragment_fallback_delay`In addition, if the actual wait time is less than 20ms, it will also fall back to waiting for a fixed time,
because the target is considered to be local or behind a transparent proxy.

#### tls_fragment_fallback_delay

Since sing-box 1.12.0

The fallback value used when TLS segmentation cannot automatically determine the wait time.

500ms is used by default.

`500ms`#### tls_record_fragment

Since sing-box 1.12.0

Fragment TLS handshake into multiple TLS records to bypass firewalls.

### sniff

```
{
  "action": "sniff",
  "sniffer": [],
  "timeout": ""
}

```

sniff performs protocol sniffing on connections.

`sniff`For deprecated inbound.sniff options, it is considered to sniff() performed before routing.

`inbound.sniff``sniff()`#### sniffer

Enabled sniffers.

All sniffers enabled by default.

Available protocol values an be found on in Protocol Sniff

#### timeout

Timeout for sniffing.

300ms is used by default.

`300ms`### resolve

```
{
  "action": "resolve",
  "server": "",
  "strategy": "",
  "disable_cache": false,
  "rewrite_ttl": null,
  "client_subnet": null
}

```

resolve resolve request destination from domain to IP addresses.

`resolve`#### server

Specifies DNS server tag to use instead of selecting through DNS routing.

#### strategy

DNS resolution strategy, available values are: prefer_ipv4, prefer_ipv6, ipv4_only, ipv6_only.

`prefer_ipv4``prefer_ipv6``ipv4_only``ipv6_only`dns.strategy will be used by default.

`dns.strategy`#### disable_cache

Since sing-box 1.12.0

Disable cache and save cache in this query.

#### rewrite_ttl

Since sing-box 1.12.0

Rewrite TTL in DNS responses.

#### client_subnet

Since sing-box 1.12.0

Append a edns0-subnet OPT extra record with the specified IP prefix to every query by default.

`edns0-subnet`If value is an IP address instead of prefix, /32 or /128 will be appended automatically.

`/32``/128`Will overrides dns.client_subnet.

`dns.client_subnet`
---

## Protocol Sniff

**Source URL**: <https://sing-box.sagernet.org/configuration/route/sniff/>

# Protocol Sniff

Changes in sing-box 1.10.0

QUIC client type detect support for QUIC
 Chromium support for QUIC
 BitTorrent support
 DTLS support
 SSH support
 RDP support

If enabled in the inbound, the protocol and domain name (if present) of by the connection can be sniffed.

#### Supported Protocols

| Network | Protocol | Domain Name | Client | 
| --- | --- | --- | --- |
| TCP | http | Host | / | 
| TCP | tls | Server Name | / | 
| UDP | quic | Server Name | QUIC Client Type | 
| UDP | stun | / | / | 
| TCP/UDP | dns | / | / | 
| TCP/UDP | bittorrent | / | / | 
| UDP | dtls | / | / | 
| TCP | ssh | / | SSH Client Name | 
| TCP | rdp | / | / | 
| UDP | ntp | / | / | 

`http``tls``quic``stun``dns``bittorrent``dtls``ssh``rdp``ntp`| QUIC Client | Type | 
| --- | --- |
| Chromium/Cronet | chromium | 
| Safari/Apple Network API | safari | 
| Firefox / uquic firefox | firefox | 
| quic-go / uquic chrome | quic-go | 

`chromium``safari``firefox``quic-go`
---

## rule-set

**Source URL**: <https://sing-box.sagernet.org/configuration/rule-set/>

Changes in sing-box 1.10.0

type: inline

`type: inline`# rule-set

Since sing-box 1.8.0

### Structure

Since sing-box 1.10.0

```
{
  "type": "inline", // optional
  "tag": "",
  "rules": []
}

```

```
{
  "type": "local",
  "tag": "",
  "format": "source", // or binary
  "path": ""
}

```

Remote rule-set will be cached if experimental.cache_file.enabled.

`experimental.cache_file.enabled````
{
  "type": "remote",
  "tag": "",
  "format": "source", // or binary
  "url": "",
  "download_detour": "", // optional
  "update_interval": "" // optional
}

```

### Fields

#### type

Required

Type of rule-set, local or remote.

`local``remote`#### tag

Required

Tag of rule-set.

### Inline Fields

Since sing-box 1.10.0

#### rules

Required

List of Headless Rule.

### Local or Remote Fields

#### format

Required

Format of rule-set file, source or binary.

`source``binary`Optional when path or url uses json or srs as extension.

`path``url``json``srs`### Local Fields

#### path

Required

Will be automatically reloaded if file modified since sing-box 1.10.0.

File path of rule-set.

### Remote Fields

#### url

Required

Download URL of rule-set.

#### download_detour

Tag of the outbound to download rule-set.

Default outbound will be used if empty.

#### update_interval

Update interval of rule-set.

1d will be used if empty.

`1d`
---

## AdGuard DNS Filer

**Source URL**: <https://sing-box.sagernet.org/configuration/rule-set/adguard/>

# AdGuard DNS Filer

Since sing-box 1.10.0

sing-box supports some rule-set formats from other projects which cannot be fully translated to sing-box,
currently only AdGuard DNS Filter.

These formats are not directly supported as source formats,
instead you need to convert them to binary rule-set.

## Convert

Use sing-box rule-set convert --type adguard [--output <file-name>.srs] <file-name>.txt to convert to binary rule-set.

`sing-box rule-set convert --type adguard [--output <file-name>.srs] <file-name>.txt`## Performance

AdGuard keeps all rules in memory and matches them sequentially,
while sing-box chooses high performance and smaller memory usage.
As a trade-off, you cannot know which rule item is matched.

## Compatibility

Almost all rules in AdGuardSDNSFilter
and rules in rule-sets listed in adguard-filter-list
are supported.

## Supported formats

### AdGuard Filter

#### Basic rule syntax

| Syntax | Supported | 
| --- | --- |
| @@ |  | 
| \|\| |  | 
| \| |  | 
| ^ |  | 
| * |  | 

`@@``\|\|``\|``^``*`#### Host syntax

| Syntax | Example | Supported | 
| --- | --- | --- |
| Scheme | https:// |  Ignored | 
| Domain Host | example.org |  | 
| IP Host | 1.1.1.1, 10.0.0. |  | 
| Regexp | /regexp/ |  | 
| Port | example.org:80 |  | 
| Path | example.org/path/ad.js |  | 

`https://``example.org``1.1.1.1``10.0.0.``/regexp/``example.org:80``example.org/path/ad.js`#### Modifier syntax

| Modifier | Supported | 
| --- | --- |
| $important |  | 
| $dnsrewrite=0.0.0.0 |  Ignored | 
| Any other modifiers |  | 

`$important``$dnsrewrite=0.0.0.0`### Hosts

Only items with 0.0.0.0 IP addresses will be accepted.

`0.0.0.0`### Simple

When all rule lines are valid domains, they are treated as simple line-by-line domain rules which,
like hosts, only match the exact same domain.


---

## Headless Rule

**Source URL**: <https://sing-box.sagernet.org/configuration/rule-set/headless-rule/>

# Headless Rule

Changes in sing-box 1.13.0

network_interface_address
 default_interface_address

Changes in sing-box 1.11.0

network_type
 network_is_expensive
 network_is_constrained

### Structure

Since sing-box 1.8.0

```
{
  "rules": [
    {
      "query_type": [
        "A",
        "HTTPS",
        32768
      ],
      "network": [
        "tcp"
      ],
      "domain": [
        "test.com"
      ],
      "domain_suffix": [
        ".cn"
      ],
      "domain_keyword": [
        "test"
      ],
      "domain_regex": [
        "^stun\\..+"
      ],
      "source_ip_cidr": [
        "10.0.0.0/24",
        "192.168.0.1"
      ],
      "ip_cidr": [
        "10.0.0.0/24",
        "192.168.0.1"
      ],
      "source_port": [
        12345
      ],
      "source_port_range": [
        "1000:2000",
        ":3000",
        "4000:"
      ],
      "port": [
        80,
        443
      ],
      "port_range": [
        "1000:2000",
        ":3000",
        "4000:"
      ],
      "process_name": [
        "curl"
      ],
      "process_path": [
        "/usr/bin/curl"
      ],
      "process_path_regex": [
        "^/usr/bin/.+"
      ],
      "package_name": [
        "com.termux"
      ],
      "network_type": [
        "wifi"
      ],
      "network_is_expensive": false,
      "network_is_constrained": false,
      "network_interface_address": {
        "wifi": [
          "2000::/3"
        ]
      },
      "default_interface_address": [
        "2000::/3"
      ],
      "wifi_ssid": [
        "My WIFI"
      ],
      "wifi_bssid": [
        "00:00:00:00:00:00"
      ],
      "invert": false
    },
    {
      "type": "logical",
      "mode": "and",
      "rules": [],
      "invert": false
    }
  ]
}

```

You can ignore the JSON Array [] tag when the content is only one item

### Default Fields

The default rule uses the following matching logic:
(domain || domain_suffix || domain_keyword || domain_regex || ip_cidr) &&
(port || port_range) &&
(source_port || source_port_range) &&
other fields

`domain``domain_suffix``domain_keyword``domain_regex``ip_cidr``port``port_range``source_port``source_port_range``other fields`#### query_type

DNS query type. Values can be integers or type name strings.

#### network

tcp or udp.

`tcp``udp`#### domain

Match full domain.

#### domain_suffix

Match domain suffix.

#### domain_keyword

Match domain using keyword.

#### domain_regex

Match domain using regular expression.

#### source_ip_cidr

Match source IP CIDR.

#### ip_cidr

ip_cidr is an alias for source_ip_cidr when rule_set_ipcidr_match_source enabled in route/DNS rules.

`ip_cidr``source_ip_cidr``rule_set_ipcidr_match_source`Match IP CIDR.

#### source_port

Match source port.

#### source_port_range

Match source port range.

#### port

Match port.

#### port_range

Match port range.

#### process_name

Only supported on Linux, Windows, and macOS.

Match process name.

#### process_path

Only supported on Linux, Windows, and macOS.

Match process path.

#### process_path_regex

Since sing-box 1.10.0

Only supported on Linux, Windows, and macOS.

Match process path using regular expression.

#### package_name

Match android package name.

#### network_type

Since sing-box 1.11.0

Only supported in graphical clients on Android and Apple platforms.

Match network type.

Available values: wifi, cellular, ethernet and other.

`wifi``cellular``ethernet``other`#### network_is_expensive

Since sing-box 1.11.0

Only supported in graphical clients on Android and Apple platforms.

Match if network is considered Metered (on Android) or considered expensive,
such as Cellular or a Personal Hotspot (on Apple platforms).

#### network_is_constrained

Since sing-box 1.11.0

Only supported in graphical clients on Apple platforms.

Match if network is in Low Data Mode.

#### network_interface_address

Since sing-box 1.13.0

Only supported in graphical clients on Android and Apple platforms.

Matches network interface (same values as network_type) address.

`network_type`#### default_interface_address

Since sing-box 1.13.0

Only supported on Linux, Windows, and macOS.

Match default interface address.

#### wifi_ssid

Only supported in graphical clients on Android and Apple platforms.

Match WiFi SSID.

#### wifi_bssid

Only supported in graphical clients on Android and Apple platforms.

Match WiFi BSSID.

#### invert

Invert match result.

### Logical Fields

#### type

logical

`logical`#### mode

Required

and or or

`and``or`#### rules

Required

Included rules.


---

## Source Format

**Source URL**: <https://sing-box.sagernet.org/configuration/rule-set/source-format/>

# Source Format

Changes in sing-box 1.13.0

version 4

`4`Changes in sing-box 1.11.0

version 3

`3`Changes in sing-box 1.10.0

version 2

`2`Since sing-box 1.8.0

### Structure

```
{
  "version": 3,
  "rules": []
}

```

### Compile

Use sing-box rule-set compile [--output <file-name>.srs] <file-name>.json to compile source to binary rule-set.

`sing-box rule-set compile [--output <file-name>.srs] <file-name>.json`### Fields

#### version

Required

Version of rule-set.

- 1: sing-box 1.8.0: Initial rule-set version.
- 2: sing-box 1.10.0: Optimized memory usages of domain_suffix rules in binary rule-sets.
- 3: sing-box 1.11.0: Added network_type, network_is_expensive and network_is_constrainted rule items.
- 4: sing-box 1.13.0: Added network_interface_address and default_interface_address rule items.

`domain_suffix``network_type``network_is_expensive``network_is_constrainted``network_interface_address``default_interface_address`#### rules

Required

List of Headless Rule.


---

## Service

**Source URL**: <https://sing-box.sagernet.org/configuration/service/>

Since sing-box 1.12.0

# Service

### Structure

```
{
  "services": [
    {
      "type": "",
      "tag": ""
    }
  ]
}

```

### Fields

| Type | Format | 
| --- | --- |
| ccm | CCM | 
| derp | DERP | 
| ocm | OCM | 
| resolved | Resolved | 
| ssm-api | SSM API | 

`ccm``derp``ocm``resolved``ssm-api`#### tag

The tag of the endpoint.


---

## CCM

**Source URL**: <https://sing-box.sagernet.org/configuration/service/ccm/>

Since sing-box 1.13.0

# CCM

CCM (Claude Code Multiplexer) service is a multiplexing service that allows you to access your local Claude Code subscription remotely through custom tokens.

It handles OAuth authentication with Claude's API on your local machine while allowing remote Claude Code to authenticate using Auth Tokens via the ANTHROPIC_AUTH_TOKEN environment variable.

`ANTHROPIC_AUTH_TOKEN`### Structure

```
{
  "type": "ccm",

  ... // Listen Fields

  "credential_path": "",
  "usages_path": "",
  "users": [],
  "headers": {},
  "detour": "",
  "tls": {}
}

```

### Listen Fields

See Listen Fields for details.

### Fields

#### credential_path

Path to the Claude Code OAuth credentials file.

If not specified, defaults to:
- $CLAUDE_CONFIG_DIR/.credentials.json if CLAUDE_CONFIG_DIR environment variable is set
- ~/.claude/.credentials.json otherwise

`$CLAUDE_CONFIG_DIR/.credentials.json``CLAUDE_CONFIG_DIR``~/.claude/.credentials.json`On macOS, credentials are read from the system keychain first, then fall back to the file if unavailable.

Refreshed tokens are automatically written back to the same location.

#### usages_path

Path to the file for storing aggregated API usage statistics.

Usage tracking is disabled if not specified.

When enabled, the service tracks and saves comprehensive statistics including:
- Request counts
- Token usage (input, output, cache read, cache creation)
- Calculated costs in USD based on Claude API pricing

Statistics are organized by model, context window (200k standard vs 1M premium), and optionally by user when authentication is enabled.

The statistics file is automatically saved every minute and upon service shutdown.

#### users

List of authorized users for token authentication.

If empty, no authentication is required.

Claude Code authenticates by setting the ANTHROPIC_AUTH_TOKEN environment variable to their token value.

`ANTHROPIC_AUTH_TOKEN`#### headers

Custom HTTP headers to send to the Claude API.

These headers will override any existing headers with the same name.

#### detour

Outbound tag for connecting to the Claude API.

#### tls

TLS configuration, see TLS.

### Example

```
{
  "services": [
    {
      "type": "ccm",
      "listen": "127.0.0.1",
      "listen_port": 8080
    }
  ]
}

```

Connect to the CCM service:

```
export ANTHROPIC_BASE_URL="http://127.0.0.1:8080"
export ANTHROPIC_AUTH_TOKEN="sk-ant-ccm-auth-token-not-required-in-this-context"

claude

```


---

## DERP

**Source URL**: <https://sing-box.sagernet.org/configuration/service/derp/>

Since sing-box 1.12.0

# DERP

DERP service is a Tailscale DERP server, similar to derper.

### Structure

```
{
  "type": "derp",

  ... // Listen Fields

  "tls": {},
  "config_path": "",
  "verify_client_endpoint": [],
  "verify_client_url": [],
  "home": "",
  "mesh_with": [],
  "mesh_psk": "",
  "mesh_psk_file": "",
  "stun": {}
}

```

### Listen Fields

See Listen Fields for details.

### Fields

#### tls

TLS configuration, see TLS.

#### config_path

Required

Derper configuration file path.

Example: derper.key

`derper.key`#### verify_client_endpoint

Tailscale endpoints tags to verify clients.

#### verify_client_url

URL to verify clients.

Object format:

```
{
  "url": "https://my-headscale.com/verify",

  ... // Dial Fields
}

```

Setting Array value to a string __URL__ is equivalent to configuring:

`__URL__````
{ "url": __URL__ }

```

#### home

What to serve at the root path. It may be left empty (the default, for a default homepage), blank for a blank page, or a URL to redirect to

`blank`#### mesh_with

Mesh with other DERP servers.

Object format:

```
{
  "server": "",
  "server_port": "",
  "host": "",
  "tls": {},

  ... // Dial Fields
}

```

Object fields:

- server: Required DERP server address.
- server_port: Required DERP server port.
- host: Custom DERP hostname.
- tls: TLS
- Dial Fields: Dial Fields

`server``server_port``host``tls``Dial Fields`#### mesh_psk

Pre-shared key for DERP mesh.

#### mesh_psk_file

Pre-shared key file for DERP mesh.

#### stun

STUN server listen options.

Object format:

```
{
  "enabled": true,

  ... // Listen Fields
}

```

Object fields:

- enabled: Required Enable STUN server.
- listen: Required STUN server listen address, default to ::.
- listen_port: Required STUN server listen port, default to 3478.
- other Listen Fields: Listen Fields

`enabled``listen``::``listen_port``3478``other Listen Fields`Setting stun value to a number __PORT__ is equivalent to configuring:

`stun``__PORT__````
{ "enabled": true, "listen_port": __PORT__ }

```


---

## OCM

**Source URL**: <https://sing-box.sagernet.org/configuration/service/ocm/>

Since sing-box 1.13.0

# OCM

OCM (OpenAI Codex Multiplexer) service is a multiplexing service that allows you to access your local OpenAI Codex subscription remotely through custom tokens.

It handles OAuth authentication with OpenAI's API on your local machine while allowing remote clients to authenticate using custom tokens.

### Structure

```
{
  "type": "ocm",

  ... // Listen Fields

  "credential_path": "",
  "usages_path": "",
  "users": [],
  "headers": {},
  "detour": "",
  "tls": {}
}

```

### Listen Fields

See Listen Fields for details.

### Fields

#### credential_path

Path to the OpenAI OAuth credentials file.

If not specified, defaults to ~/.codex/auth.json.

`~/.codex/auth.json`Refreshed tokens are automatically written back to the same location.

#### usages_path

Path to the file for storing aggregated API usage statistics.

Usage tracking is disabled if not specified.

When enabled, the service tracks and saves comprehensive statistics including:
- Request counts
- Token usage (input, output, cached)
- Calculated costs in USD based on OpenAI API pricing

Statistics are organized by model and optionally by user when authentication is enabled.

The statistics file is automatically saved every minute and upon service shutdown.

#### users

List of authorized users for token authentication.

If empty, no authentication is required.

Object format:

```
{
  "name": "",
  "token": ""
}

```

Object fields:

- name: Username identifier for tracking purposes.
- token: Bearer token for authentication. Clients authenticate by setting the Authorization: Bearer <token> header.

`name``token``Authorization: Bearer <token>`#### headers

Custom HTTP headers to send to the OpenAI API.

These headers will override any existing headers with the same name.

#### detour

Outbound tag for connecting to the OpenAI API.

#### tls

TLS configuration, see TLS.

### Example

#### Server

```
{
  "services": [
    {
      "type": "ocm",
      "listen": "127.0.0.1",
      "listen_port": 8080
    }
  ]
}

```

#### Client

Add to ~/.codex/config.toml:

`~/.codex/config.toml````
[model_providers.ocm]
name = "OCM Proxy"
base_url = "http://127.0.0.1:8080/v1"
wire_api = "responses"
requires_openai_auth = false

```

Then run:

```
codex --model-provider ocm

```

### Example with Authentication

#### Server

```
{
  "services": [
    {
      "type": "ocm",
      "listen": "0.0.0.0",
      "listen_port": 8080,
      "usages_path": "./codex-usages.json",
      "users": [
        {
          "name": "alice",
          "token": "sk-alice-secret-token"
        },
        {
          "name": "bob",
          "token": "sk-bob-secret-token"
        }
      ]
    }
  ]
}

```

#### Client

Add to ~/.codex/config.toml:

`~/.codex/config.toml````
[model_providers.ocm]
name = "OCM Proxy"
base_url = "http://127.0.0.1:8080/v1"
wire_api = "responses"
requires_openai_auth = false
experimental_bearer_token = "sk-alice-secret-token"

```

Then run:

```
codex --model-provider ocm

```


---

## Resolved

**Source URL**: <https://sing-box.sagernet.org/configuration/service/resolved/>

Since sing-box 1.12.0

# Resolved

Resolved service is a fake systemd-resolved DBUS service to receive DNS settings from other programs
(e.g. NetworkManager) and provide DNS resolution.

See also: Resolved DNS Server

### Structure

```
{
  "type": "resolved",

  ... // Listen Fields
}

```

### Listen Fields

See Listen Fields for details.

### Fields

#### listen

Required

Listen address.

127.0.0.53 will be used by default.

`127.0.0.53`#### listen_port

Required

Listen port.

53 will be used by default.

`53`
---

## SSM API

**Source URL**: <https://sing-box.sagernet.org/configuration/service/ssm-api/>

Since sing-box 1.12.0

# SSM API

SSM API service is a RESTful API server for managing Shadowsocks servers.

See https://github.com/Shadowsocks-NET/shadowsocks-specs/blob/main/2023-1-shadowsocks-server-management-api-v1.md

### Structure

```
{
  "type": "ssm-api",

  ... // Listen Fields

  "servers": {},
  "cache_path": "",
  "tls": {}
}

```

### Listen Fields

See Listen Fields for details.

### Fields

#### servers

Required

A mapping Object from HTTP endpoints to Shadowsocks Inbound tags.

Selected Shadowsocks inbounds must be configured with managed enabled.

Example:

```
{
  "servers": {
    "/": "ss-in"
  }
}

```

#### cache_path

If set, when the server is about to stop, traffic and user state will be saved to the specified JSON file
to be restored on the next startup.

#### tls

TLS configuration, see TLS.


---

## Dial Fields

**Source URL**: <https://sing-box.sagernet.org/configuration/shared/dial/>

# Dial Fields

Changes in sing-box 1.13.0

disable_tcp_keep_alive
 tcp_keep_alive
 tcp_keep_alive_interval
 bind_address_no_port

Changes in sing-box 1.12.0

domain_resolver
 domain_strategy
 netns

Changes in sing-box 1.11.0

network_strategy
 fallback_delay
 network_type
 fallback_network_type

### Structure

```
{
  "detour": "",
  "bind_interface": "",
  "inet4_bind_address": "",
  "inet6_bind_address": "",
  "bind_address_no_port": false,
  "routing_mark": 0,
  "reuse_addr": false,
  "netns": "",
  "connect_timeout": "",
  "tcp_fast_open": false,
  "tcp_multi_path": false,
  "disable_tcp_keep_alive": false,
  "tcp_keep_alive": "",
  "tcp_keep_alive_interval": "",
  "udp_fragment": false,

  "domain_resolver": "", // or {}
  "network_strategy": "",
  "network_type": [],
  "fallback_network_type": [],
  "fallback_delay": "",

  // Deprecated

  "domain_strategy": ""
}

```

You can ignore the JSON Array [] tag when the content is only one item

### Fields

#### detour

The tag of the upstream outbound.

If enabled, all other fields will be ignored.

#### bind_interface

The network interface to bind to.

#### inet4_bind_address

The IPv4 address to bind to.

#### inet6_bind_address

The IPv6 address to bind to.

#### bind_address_no_port

Since sing-box 1.13.0

Only supported on Linux.

Do not reserve a port when binding to a source address.

This allows reusing the same source port for multiple connections if the full 4-tuple (source IP, source port, destination IP, destination port) remains unique.

#### routing_mark

Only supported on Linux.

Set netfilter routing mark.

Integers (e.g. 1234) and string hexadecimals (e.g. "0x1234") are supported.

`1234``"0x1234"`#### reuse_addr

Reuse listener address.

#### netns

Since sing-box 1.12.0

Only supported on Linux.

Set network namespace, name or path.

#### connect_timeout

Connect timeout, in golang's Duration format.

A duration string is a possibly signed sequence of
decimal numbers, each with optional fraction and a unit suffix,
such as "300ms", "-1.5h" or "2h45m".
Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".

#### tcp_fast_open

Enable TCP Fast Open.

#### tcp_multi_path

Go 1.21 required.

Enable TCP Multi Path.

#### disable_tcp_keep_alive

Since sing-box 1.13.0

Disable TCP keep alive.

#### tcp_keep_alive

Since sing-box 1.13.0

Default value changed from 10m to 5m.

`10m``5m`TCP keep alive initial period.

5m will be used by default.

`5m`#### tcp_keep_alive_interval

Since sing-box 1.13.0

TCP keep alive interval.

75s will be used by default.

`75s`#### udp_fragment

Enable UDP fragmentation.

#### domain_resolver

outbound DNS rule items are deprecated and will be removed in sing-box 1.14.0, so this item will be required for outbound/endpoints using domain name in server address since sing-box 1.14.0.

`outbound`domain_resolver or route.default_domain_resolver is optional when only one DNS server is configured.

`domain_resolver``route.default_domain_resolver`Set domain resolver to use for resolving domain names.

This option uses the same format as the route DNS rule action without the action field.

`action`Setting this option directly to a string is equivalent to setting server of this options.

`server`| Outbound/Endpoints | Effected domains | 
| --- | --- |
| direct | Domain in request | 
| others | Domain in server address | 

`direct`#### network_strategy

Since sing-box 1.11.0

Only supported in graphical clients on Android and Apple platforms with auto_detect_interface enabled.

`auto_detect_interface`Strategy for selecting network interfaces.

Available values:

- default (default): Connect to default network or networks specified in network_type sequentially.
- hybrid: Connect to all networks or networks specified in network_type concurrently.
- fallback: Connect to default network or preferred networks specified in network_type concurrently, and try fallback networks when unavailable or timeout.

`default``network_type``hybrid``network_type``fallback``network_type`For fallback, when preferred interfaces fails or times out,
it will enter a 15s fast fallback state (Connect to all preferred and fallback networks concurrently),
and exit immediately if preferred networks recover.

Conflicts with bind_interface, inet4_bind_address and inet6_bind_address.

`bind_interface``inet4_bind_address``inet6_bind_address`#### network_type

Since sing-box 1.11.0

Only supported in graphical clients on Android and Apple platforms with auto_detect_interface enabled.

`auto_detect_interface`Network types to use when using default or hybrid network strategy or
preferred network types to use when using fallback network strategy.

`default``hybrid``fallback`Available values: wifi, cellular, ethernet, other.

`wifi``cellular``ethernet``other`Device's default network is used by default.

#### fallback_network_type

Since sing-box 1.11.0

Only supported in graphical clients on Android and Apple platforms with auto_detect_interface enabled.

`auto_detect_interface`Fallback network types when preferred networks are unavailable or timeout when using fallback network strategy.

`fallback`All other networks expect preferred are used by default.

#### fallback_delay

Since sing-box 1.11.0

Only supported in graphical clients on Android and Apple platforms with auto_detect_interface enabled.

`auto_detect_interface`The length of time to wait before spawning a RFC 6555 Fast Fallback connection.

For domain_strategy, is the amount of time to wait for connection to succeed before assuming
that IPv4/IPv6 is misconfigured and falling back to other type of addresses.

`domain_strategy`For network_strategy, is the amount of time to wait for connection to succeed before falling
back to other interfaces.

`network_strategy`Only take effect when domain_strategy or network_strategy is set.

`domain_strategy``network_strategy`300ms is used by default.

`300ms`#### domain_strategy

Deprecated in sing-box 1.12.0

domain_strategy is deprecated and will be removed in sing-box 1.14.0, check Migration.

`domain_strategy`Available values: prefer_ipv4, prefer_ipv6, ipv4_only, ipv6_only.

`prefer_ipv4``prefer_ipv6``ipv4_only``ipv6_only`If set, the requested domain name will be resolved to IP before connect.

| Outbound | Effected domains | Fallback Value | 
| --- | --- | --- |
| direct | Domain in request | Take inbound.domain_strategy if not set | 
| others | Domain in server address | / | 

`direct``inbound.domain_strategy`
---

## DNS01 Challenge Fields

**Source URL**: <https://sing-box.sagernet.org/configuration/shared/dns01_challenge/>

# DNS01 Challenge Fields

Changes in sing-box 1.13.0

alidns.security_token
 cloudflare.zone_token

### Structure

```
{
  "provider": "",

  ... // Provider Fields
}

```

### Provider Fields

#### Alibaba Cloud DNS

```
{
  "provider": "alidns",
  "access_key_id": "",
  "access_key_secret": "",
  "region_id": "",
  "security_token": ""
}

```

##### security_token

Since sing-box 1.13.0

The Security Token for STS temporary credentials.

#### Cloudflare

```
{
  "provider": "cloudflare",
  "api_token": "",
  "zone_token": ""
}

```

##### zone_token

Since sing-box 1.13.0

Optional API token with Zone:Read permission.

`Zone:Read`When provided, allows api_token to be scoped to a single zone.

`api_token`
---

## Listen Fields

**Source URL**: <https://sing-box.sagernet.org/configuration/shared/listen/>

# Listen Fields

Changes in sing-box 1.13.0

disable_tcp_keep_alive
 tcp_keep_alive

Changes in sing-box 1.12.0

netns
 bind_interface
 routing_mark
 reuse_addr

Changes in sing-box 1.11.0

sniff
 sniff_override_destination
 sniff_timeout
 domain_strategy
 udp_disable_domain_unmapping

### Structure

```
{
  "listen": "",
  "listen_port": 0,
  "bind_interface": "",
  "routing_mark": 0,
  "reuse_addr": false,
  "netns": "",
  "tcp_fast_open": false,
  "tcp_multi_path": false,
  "disable_tcp_keep_alive": false,
  "tcp_keep_alive": "",
  "tcp_keep_alive_interval": "",
  "udp_fragment": false,
  "udp_timeout": "",
  "detour": "",

  // Deprecated

  "sniff": false,
  "sniff_override_destination": false,
  "sniff_timeout": "",
  "domain_strategy": "",
  "udp_disable_domain_unmapping": false
}

```

### Fields

#### listen

Required

Listen address.

#### listen_port

Listen port.

#### bind_interface

Since sing-box 1.12.0

The network interface to bind to.

#### routing_mark

Since sing-box 1.12.0

Only supported on Linux.

Set netfilter routing mark.

Integers (e.g. 1234) and string hexadecimals (e.g. "0x1234") are supported.

`1234``"0x1234"`#### reuse_addr

Since sing-box 1.12.0

Reuse listener address.

#### netns

Since sing-box 1.12.0

Only supported on Linux.

Set network namespace, name or path.

#### tcp_fast_open

Enable TCP Fast Open.

#### tcp_multi_path

Go 1.21 required.

Enable TCP Multi Path.

#### disable_tcp_keep_alive

Since sing-box 1.13.0

Disable TCP keep alive.

#### tcp_keep_alive

Since sing-box 1.13.0

Default value changed from 10m to 5m.

`10m``5m`TCP keep alive initial period.

5m will be used by default.

`5m`#### tcp_keep_alive_interval

TCP keep alive interval.

75s will be used by default.

`75s`#### udp_fragment

Enable UDP fragmentation.

#### udp_timeout

UDP NAT expiration time.

5m will be used by default.

`5m`#### detour

If set, connections will be forwarded to the specified inbound.

Requires target inbound support, see Injectable.

#### sniff

Deprecated in sing-box 1.11.0

Inbound fields are deprecated and will be removed in sing-box 1.13.0, check Migration.

Enable sniffing.

See Protocol Sniff for details.

#### sniff_override_destination

Deprecated in sing-box 1.11.0

Inbound fields are deprecated and will be removed in sing-box 1.13.0.

Override the connection destination address with the sniffed domain.

If the domain name is invalid (like tor), this will not work.

#### sniff_timeout

Deprecated in sing-box 1.11.0

Inbound fields are deprecated and will be removed in sing-box 1.13.0, check Migration.

Timeout for sniffing.

300ms is used by default.

`300ms`#### domain_strategy

Deprecated in sing-box 1.11.0

Inbound fields are deprecated and will be removed in sing-box 1.13.0, check Migration.

One of prefer_ipv4 prefer_ipv6 ipv4_only ipv6_only.

`prefer_ipv4``prefer_ipv6``ipv4_only``ipv6_only`If set, the requested domain name will be resolved to IP before routing.

If sniff_override_destination is in effect, its value will be taken as a fallback.

`sniff_override_destination`#### udp_disable_domain_unmapping

Deprecated in sing-box 1.11.0

Inbound fields are deprecated and will be removed in sing-box 1.13.0, check Migration.

If enabled, for UDP proxy requests addressed to a domain, 
the original packet address will be sent in the response instead of the mapped domain.

This option is used for compatibility with clients that 
do not support receiving UDP packets with domain addresses, such as Surge.


---

## Multiplex

**Source URL**: <https://sing-box.sagernet.org/configuration/shared/multiplex/>

# Multiplex

### Inbound

```
{
  "enabled": true,
  "padding": false,
  "brutal": {}
}

```

### Outbound

```
{
  "enabled": true,
  "protocol": "smux",
  "max_connections": 4,
  "min_streams": 4,
  "max_streams": 0,
  "padding": false,
  "brutal": {}
}

```

### Inbound Fields

#### enabled

Enable multiplex support.

#### padding

If enabled, non-padded connections will be rejected.

#### brutal

See TCP Brutal for details.

### Outbound Fields

#### enabled

Enable multiplex.

#### protocol

Multiplex protocol.

| Protocol | Description | 
| --- | --- |
| smux | https://github.com/xtaci/smux | 
| yamux | https://github.com/hashicorp/yamux | 
| h2mux | https://golang.org/x/net/http2 | 

h2mux is used by default.

#### max_connections

Maximum connections.

Conflict with max_streams.

`max_streams`#### min_streams

Minimum multiplexed streams in a connection before opening a new connection.

Conflict with max_streams.

`max_streams`#### max_streams

Maximum multiplexed streams in a connection before opening a new connection.

Conflict with max_connections and min_streams.

`max_connections``min_streams`#### padding

Info

Requires sing-box server version 1.3-beta9 or later.

Enable padding.

#### brutal

See TCP Brutal for details.


---

## Pre-match

**Source URL**: <https://sing-box.sagernet.org/configuration/shared/pre-match/>

# Pre-match

Changes in sing-box 1.13.0

bypass

Pre-match is rule matching that runs before the connection is established.

### How it works

When TUN receives a connection request, the connection has not yet been established,
so no connection data can be read. In this phase, sing-box runs the routing rules in pre-match mode.

Since connection data is unavailable, only actions that do not require connection data can be executed.
When a rule matches an action that requires an established connection, pre-match stops at that rule.

### Supported actions

#### reject

Reject with TCP RST / ICMP unreachable.

See reject for details.

#### route

Route ICMP connections to the specified outbound for direct reply.

See route for details.

#### bypass

Since sing-box 1.13.0

Only supported on Linux with auto_redirect enabled.

`auto_redirect`Bypass sing-box and connect directly at kernel level.

If outbound is not specified, the rule only matches in pre-match from auto redirect,
and will be skipped in other contexts.

`outbound`For all other contexts, bypass with outbound behaves like route action.

`outbound``route`See bypass for details.


---

## TCP Brutal

**Source URL**: <https://sing-box.sagernet.org/configuration/shared/tcp-brutal/>

# TCP Brutal

### Server Requirements

- Linux
- brutal congestion control algorithm kernel module installed

`brutal`See tcp-brutal for details.

### Structure

```
{
  "enabled": true,
  "up_mbps": 100,
  "down_mbps": 100
}

```

### Fields

#### enabled

Enable TCP Brutal congestion control algorithm。

#### up_mbps, down_mbps

Required

Upload and download bandwidth, in Mbps.


---

## TLS

**Source URL**: <https://sing-box.sagernet.org/configuration/shared/tls/>

# TLS

Changes in sing-box 1.13.0

kernel_tx
 kernel_rx
 curve_preferences
 certificate_public_key_sha256
 client_certificate
 client_certificate_path
 client_key
 client_key_path
 client_authentication
 client_certificate_public_key_sha256
 ech.query_server_name

Changes in sing-box 1.12.0

fragment
 fragment_fallback_delay
 record_fragment
 ech.pq_signature_schemes_enabled
 ech.dynamic_record_sizing_disabled

Changes in sing-box 1.10.0

utls

### Inbound

```
{
  "enabled": true,
  "server_name": "",
  "alpn": [],
  "min_version": "",
  "max_version": "",
  "cipher_suites": [],
  "curve_preferences": [],
  "certificate": [],
  "certificate_path": "",
  "client_authentication": "",
  "client_certificate": [],
  "client_certificate_path": [],
  "client_certificate_public_key_sha256": [],
  "key": [],
  "key_path": "",
  "kernel_tx": false,
  "kernel_rx": false,
  "acme": {
    "domain": [],
    "data_directory": "",
    "default_server_name": "",
    "email": "",
    "provider": "",
    "disable_http_challenge": false,
    "disable_tls_alpn_challenge": false,
    "alternative_http_port": 0,
    "alternative_tls_port": 0,
    "external_account": {
      "key_id": "",
      "mac_key": ""
    },
    "dns01_challenge": {}
  },
  "ech": {
    "enabled": false,
    "key": [],
    "key_path": "",

    // Deprecated

    "pq_signature_schemes_enabled": false,
    "dynamic_record_sizing_disabled": false
  },
  "reality": {
    "enabled": false,
    "handshake": {
      "server": "google.com",
      "server_port": 443,

      ... // Dial Fields
    },
    "private_key": "UuMBgl7MXTPx9inmQp2UC7Jcnwc6XYbwDNebonM-FCc",
    "short_id": [
      "0123456789abcdef"
    ],
    "max_time_difference": "1m"
  }
}

```

### Outbound

```
{
  "enabled": true,
  "disable_sni": false,
  "server_name": "",
  "insecure": false,
  "alpn": [],
  "min_version": "",
  "max_version": "",
  "cipher_suites": [],
  "curve_preferences": [],
  "certificate": "",
  "certificate_path": "",
  "certificate_public_key_sha256": [],
  "client_certificate": [],
  "client_certificate_path": "",
  "client_key": [],
  "client_key_path": "",
  "fragment": false,
  "fragment_fallback_delay": "",
  "record_fragment": false,
  "ech": {
    "enabled": false,
    "config": [],
    "config_path": "",
    "query_server_name": "",

    // Deprecated
    "pq_signature_schemes_enabled": false,
    "dynamic_record_sizing_disabled": false
  },
  "utls": {
    "enabled": false,
    "fingerprint": ""
  },
  "reality": {
    "enabled": false,
    "public_key": "jNXHt1yRo0vDuchQlIP6Z0ZvjT3KtzVI-T4E7RoLJS0",
    "short_id": "0123456789abcdef"
  }
}

```

TLS version values:

- 1.0
- 1.1
- 1.2
- 1.3

`1.0``1.1``1.2``1.3`Cipher suite values:

- TLS_RSA_WITH_AES_128_CBC_SHA
- TLS_RSA_WITH_AES_256_CBC_SHA
- TLS_RSA_WITH_AES_128_GCM_SHA256
- TLS_RSA_WITH_AES_256_GCM_SHA384
- TLS_AES_128_GCM_SHA256
- TLS_AES_256_GCM_SHA384
- TLS_CHACHA20_POLY1305_SHA256
- TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA
- TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA
- TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA
- TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA
- TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256
- TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384
- TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256
- TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384
- TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256
- TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256

`TLS_RSA_WITH_AES_128_CBC_SHA``TLS_RSA_WITH_AES_256_CBC_SHA``TLS_RSA_WITH_AES_128_GCM_SHA256``TLS_RSA_WITH_AES_256_GCM_SHA384``TLS_AES_128_GCM_SHA256``TLS_AES_256_GCM_SHA384``TLS_CHACHA20_POLY1305_SHA256``TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA``TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA``TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA``TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA``TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256``TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384``TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256``TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384``TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256``TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256`You can ignore the JSON Array [] tag when the content is only one item

### Fields

#### enabled

Enable TLS.

#### disable_sni

Client only

Do not send server name in ClientHello.

#### server_name

Used to verify the hostname on the returned certificates unless insecure is given.

It is also included in the client's handshake to support virtual hosting unless it is an IP address.

#### insecure

Client only

Accepts any server certificate.

#### alpn

List of supported application level protocols, in order of preference.

If both peers support ALPN, the selected protocol will be one from this list, and the connection will fail if there is
no mutually supported protocol.

See Application-Layer Protocol Negotiation.

#### min_version

The minimum TLS version that is acceptable.

By default, TLS 1.2 is currently used as the minimum when acting as a
client, and TLS 1.0 when acting as a server.

#### max_version

The maximum TLS version that is acceptable.

By default, the maximum version is currently TLS 1.3.

#### cipher_suites

List of enabled TLS 1.0–1.2 cipher suites. The order of the list is ignored.
Note that TLS 1.3 cipher suites are not configurable.

If empty, a safe default list is used. The default cipher suites might change over time.

#### curve_preferences

Since sing-box 1.13.0

Set of supported key exchange mechanisms. The order of the list is ignored, and key exchange mechanisms are chosen
from this list using an internal preference order by Golang.

Available values, also the default list:

- P256
- P384
- P521
- X25519
- X25519MLKEM768

`P256``P384``P521``X25519``X25519MLKEM768`#### certificate

Server certificates chain line array, in PEM format.

#### certificate_path

Will be automatically reloaded if file modified.

The path to server certificate chain, in PEM format.

#### certificate_public_key_sha256

Since sing-box 1.13.0

Client only

List of SHA-256 hashes of server certificate public keys, in base64 format.

To generate the SHA-256 hash for a certificate's public key, use the following commands:

```
# For a certificate file
openssl x509 -in certificate.pem -pubkey -noout | openssl pkey -pubin -outform der | openssl dgst -sha256 -binary | openssl enc -base64

# For a certificate from a remote server
echo | openssl s_client -servername example.com -connect example.com:443 2>/dev/null | openssl x509 -pubkey -noout | openssl pkey -pubin -outform der | openssl dgst -sha256 -binary | openssl enc -base64

```

#### client_certificate

Since sing-box 1.13.0

Client only

Client certificate chain line array, in PEM format.

#### client_certificate_path

Since sing-box 1.13.0

Client only

The path to client certificate chain, in PEM format.

#### client_key

Since sing-box 1.13.0

Client only

Client private key line array, in PEM format.

#### client_key_path

Since sing-box 1.13.0

Client only

The path to client private key, in PEM format.

#### key

Server only

The server private key line array, in PEM format.

#### key_path

Server only

Will be automatically reloaded if file modified.

The path to the server private key, in PEM format.

#### client_authentication

Since sing-box 1.13.0

Server only

The type of client authentication to use.

Available values:

- no (default)
- request
- require-any
- verify-if-given
- require-and-verify

`no``request``require-any``verify-if-given``require-and-verify`One of client_certificate, client_certificate_path, or client_certificate_public_key_sha256 is required
if this option is set to verify-if-given, or require-and-verify.

`client_certificate``client_certificate_path``client_certificate_public_key_sha256``verify-if-given``require-and-verify`#### client_certificate

Since sing-box 1.13.0

Server only

Client certificate chain line array, in PEM format.

#### client_certificate_path

Since sing-box 1.13.0

Server only

Will be automatically reloaded if file modified.

List of path to client certificate chain, in PEM format.

#### client_certificate_public_key_sha256

Since sing-box 1.13.0

Server only

List of SHA-256 hashes of client certificate public keys, in base64 format.

To generate the SHA-256 hash for a certificate's public key, use the following commands:

```
# For a certificate file
openssl x509 -in certificate.pem -pubkey -noout | openssl pkey -pubin -outform der | openssl dgst -sha256 -binary | openssl enc -base64

# For a certificate from a remote server
echo | openssl s_client -servername example.com -connect example.com:443 2>/dev/null | openssl x509 -pubkey -noout | openssl pkey -pubin -outform der | openssl dgst -sha256 -binary | openssl enc -base64

```

#### kernel_tx

Since sing-box 1.13.0

Only supported on Linux 5.1+, use a newer kernel if possible.

Only TLS 1.3 is supported.

kTLS TX may only improve performance when splice(2) is available (both ends must be TCP or TLS without additional protocols after handshake); otherwise, it will definitely degrade performance.

`splice(2)`Enable kernel TLS transmit support.

#### kernel_rx

Since sing-box 1.13.0

Only supported on Linux 5.1+, use a newer kernel if possible.

Only TLS 1.3 is supported.

kTLS RX will definitely degrade performance even if splice(2) is in use, so enabling it is not recommended.

`splice(2)`Enable kernel TLS receive support.

## Custom TLS support

QUIC support

Only ECH is supported in QUIC.

#### utls

Client only

Not Recommended

uTLS has had repeated fingerprinting vulnerabilities discovered by researchers.

uTLS is a Go library that attempts to imitate browser TLS fingerprints by copying
ClientHello structure. However, browsers use completely different TLS stacks
(Chrome uses BoringSSL, Firefox uses NSS) with distinct implementation behaviors
that cannot be replicated by simply copying the handshake format, making detection possible.
Additionally, the library lacks active maintenance and has poor code quality,
making it unsuitable for censorship circumvention.

For TLS fingerprint resistance, use NaiveProxy instead.

uTLS is a fork of "crypto/tls", which provides ClientHello fingerprinting resistance.

Available fingerprint values:

Removed since sing-box 1.10.0

Some legacy chrome fingerprints have been removed and will fallback to chrome:

chrome_psk
 chrome_psk_shuffle
 chrome_padding_psk_shuffle
 chrome_pq
 chrome_pq_psk

- chrome
- firefox
- edge
- safari
- 360
- qq
- ios
- android
- random
- randomized

Chrome fingerprint will be used if empty.

### ECH Fields

ECH (Encrypted Client Hello) is a TLS extension that allows a client to encrypt the first part of its ClientHello
message.

The ECH key and configuration can be generated by sing-box generate ech-keypair.

`sing-box generate ech-keypair`#### pq_signature_schemes_enabled

Deprecated in sing-box 1.12.0

ECH support has been migrated to use stdlib in sing-box 1.12.0, which does not come with support for PQ signature schemes, so pq_signature_schemes_enabled has been deprecated and no longer works.

`pq_signature_schemes_enabled`Enable support for post-quantum peer certificate signature schemes.

#### dynamic_record_sizing_disabled

Deprecated in sing-box 1.12.0

dynamic_record_sizing_disabled has nothing to do with ECH, was added by mistake, has been deprecated and no longer works.

`dynamic_record_sizing_disabled`Disables adaptive sizing of TLS records.

When true, the largest possible TLS record size is always used.
When false, the size of TLS records may be adjusted in an attempt to improve latency.

#### key

Server only

ECH key line array, in PEM format.

#### key_path

Server only

Will be automatically reloaded if file modified.

The path to ECH key, in PEM format.

#### config

Client only

ECH configuration line array, in PEM format.

If empty, load from DNS will be attempted.

#### config_path

Client only

The path to ECH configuration, in PEM format.

If empty, load from DNS will be attempted.

#### query_server_name

Since sing-box 1.13.0

Client only

Overrides the domain name used for ECH HTTPS record queries.

If empty, server_name is used for queries.

`server_name`#### fragment

Since sing-box 1.12.0

Client only

Fragment TLS handshakes to bypass firewalls.

This feature is intended to circumvent simple firewalls based on plaintext packet matching,
and should not be used to circumvent real censorship.

Due to poor performance, try record_fragment first, and only apply to server names known to be blocked.

`record_fragment`On Linux, Apple platforms, (administrator privileges required) Windows,
the wait time can be automatically detected. Otherwise, it will fall back to
waiting for a fixed time specified by fragment_fallback_delay.

`fragment_fallback_delay`In addition, if the actual wait time is less than 20ms, it will also fall back to waiting for a fixed time,
because the target is considered to be local or behind a transparent proxy.

#### fragment_fallback_delay

Since sing-box 1.12.0

Client only

The fallback value used when TLS segmentation cannot automatically determine the wait time.

500ms is used by default.

`500ms`#### record_fragment

Since sing-box 1.12.0

Client only

Fragment TLS handshake into multiple TLS records to bypass firewalls.

### ACME Fields

#### domain

List of domain.

ACME will be disabled if empty.

#### data_directory

The directory to store ACME data.

$XDG_DATA_HOME/certmagic|$HOME/.local/share/certmagic will be used if empty.

`$XDG_DATA_HOME/certmagic|$HOME/.local/share/certmagic`#### default_server_name

Server name to use when choosing a certificate if the ClientHello's ServerName field is empty.

#### email

The email address to use when creating or selecting an existing ACME server account

#### provider

The ACME CA provider to use.

| Value | Provider | 
| --- | --- |
| letsencrypt (default) | Let's Encrypt | 
| zerossl | ZeroSSL | 
| https://... | Custom | 

`letsencrypt (default)``zerossl``https://...`#### disable_http_challenge

Disable all HTTP challenges.

#### disable_tls_alpn_challenge

Disable all TLS-ALPN challenges

#### alternative_http_port

The alternate port to use for the ACME HTTP challenge; if non-empty, this port will be used instead of 80 to spin up a
listener for the HTTP challenge.

#### alternative_tls_port

The alternate port to use for the ACME TLS-ALPN challenge; the system must forward 443 to this port for challenge to
succeed.

#### external_account

EAB (External Account Binding) contains information necessary to bind or map an ACME account to some other account known
by the CA.

External account bindings are "used to associate an ACME account with an existing account in a non-ACME system, such as
a CA customer database.

To enable ACME account binding, the CA operating the ACME server needs to provide the ACME client with a MAC key and a
key identifier, using some mechanism outside of ACME. §7.3.4

#### external_account.key_id

The key identifier.

#### external_account.mac_key

The MAC key.

#### dns01_challenge

ACME DNS01 challenge field. If configured, other challenge methods will be disabled.

See DNS01 Challenge Fields for details.

### Reality Fields

#### handshake

Server only

Required

Handshake server address and Dial Fields.

#### private_key

Server only

Required

Private key, generated by sing-box generate reality-keypair.

`sing-box generate reality-keypair`#### public_key

Client only

Required

Public key, generated by sing-box generate reality-keypair.

`sing-box generate reality-keypair`#### short_id

Required

A hexadecimal string with zero to eight digits.

#### max_time_difference

Server only

The maximum time difference between the server and the client.

Check disabled if empty.


---

## UDP over TCP

**Source URL**: <https://sing-box.sagernet.org/configuration/shared/udp-over-tcp/>

# UDP over TCP

It's a proprietary protocol created by SagerNet, not part of shadowsocks.

The UDP over TCP protocol is used to transmit UDP packets in TCP.

### Structure

```
{
  "enabled": true,
  "version": 2
}

```

The structure can be replaced with a boolean value when the version is not specified.

### Fields

#### enabled

Enable the UDP over TCP protocol.

#### version

The protocol version, 1 or 2.

`1``2`2 is used by default.

### Application support

| Project | UoT v1 | UoT v2 | 
| --- | --- | --- |
| sing-box | v0 (2022/08/11) | v1.2-beta9 | 
| Clash.Meta | v1.12.0 (2022/07/02) | v1.14.3 (2023/03/31) | 
| Shadowrocket | v2.2.12 (2022/08/13) | / | 

### Protocol details

#### Protocol version 1

The client requests the magic address to the upper layer proxy protocol to indicate the request: sp.udp-over-tcp.arpa

`sp.udp-over-tcp.arpa`#### Stream format

| ATYP | address | port | length | data | 
| --- | --- | --- | --- | --- |
| u8 | variable | u16be | u16be | variable | 

ATYP / address / port: Uses the SOCKS address format, but with different address types:

| ATYP | Address type | 
| --- | --- |
| 0x00 | IPv4 Address | 
| 0x01 | IPv6 Address | 
| 0x02 | Domain Name | 

`0x00``0x01``0x02`#### Protocol version 2

Protocol version 2 uses a new magic address: sp.v2.udp-over-tcp.arpa

`sp.v2.udp-over-tcp.arpa`##### Request format

| isConnect | ATYP | address | port | 
| --- | --- | --- | --- |
| u8 | u8 | variable | u16be | 

isConnect: Set to 1 to indicates that the stream uses the connect format, 0 to disable.

ATYP / address / port: Request destination, uses the SOCKS address format.

##### Connect stream format

| length | data | 
| --- | --- |
| u16be | variable | 

##### Non-connect stream format

As the same as the stream format in protocol version 1.


---

## V2Ray Transport

**Source URL**: <https://sing-box.sagernet.org/configuration/shared/v2ray-transport/>

# V2Ray Transport

V2Ray Transport is a set of private protocols invented by v2ray, and has contaminated the names of other protocols, such
as trojan-grpc in clash.

`trojan-grpc`### Structure

```
{
  "type": ""
}

```

Available transports:

- HTTP
- WebSocket
- QUIC
- gRPC
- HTTPUpgrade

Difference from v2ray-core

- No TCP transport, plain HTTP is merged into the HTTP transport.
- No mKCP transport.
- No DomainSocket transport.

You can ignore the JSON Array [] tag when the content is only one item

### HTTP

```
{
  "type": "http",
  "host": [],
  "path": "",
  "method": "",
  "headers": {},
  "idle_timeout": "15s",
  "ping_timeout": "15s"
}

```

Difference from v2ray-core

TLS is not enforced. If TLS is not configured, plain HTTP 1.1 is used.

#### host

List of host domain.

The client will choose randomly and the server will verify if not empty.

#### path

Warning

V2Ray's documentation says that the path between the server and the client must be consistent, 
but the actual code allows the client to add any suffix to the path.
sing-box uses the same behavior as V2Ray, but note that the behavior does not exist in WebSocket and HTTPUpgrade transport.

`WebSocket``HTTPUpgrade`Path of HTTP request.

The server will verify.

#### method

Method of HTTP request.

The server will verify if not empty.

#### headers

Extra headers of HTTP request.

The server will write in response if not empty.

#### idle_timeout

In HTTP2 server:

Specifies the time until idle clients should be closed with a GOAWAY frame. PING frames are not considered as activity.

In HTTP2 client:

Specifies the period of time after which a health check will be performed using a ping frame if no frames have been
received on the connection.Please note that a ping response is considered a received frame, so if there is no other
traffic on the connection, the health check will be executed every interval. If the value is zero, no health check will
be performed.

Zero is used by default.

#### ping_timeout

In HTTP2 client:

Specifies the timeout duration after sending a PING frame, within which a response must be received.
If a response to the PING frame is not received within the specified timeout duration, the connection will be closed.
The default timeout duration is 15 seconds.

### WebSocket

```
{
  "type": "ws",
  "path": "",
  "headers": {},
  "max_early_data": 0,
  "early_data_header_name": ""
}

```

#### path

Path of HTTP request.

The server will verify.

#### headers

Extra headers of HTTP request.

The server will write in response if not empty.

#### max_early_data

Allowed payload size is in the request. Enabled if not zero.

#### early_data_header_name

Early data is sent in path instead of header by default.

To be compatible with Xray-core, set this to Sec-WebSocket-Protocol.

`Sec-WebSocket-Protocol`It needs to be consistent with the server.

### QUIC

```
{
  "type": "quic"
}

```

Difference from v2ray-core

No additional encryption support:
It's basically duplicate encryption. And Xray-core is not compatible with v2ray-core in here.

### gRPC

standard gRPC has good compatibility but poor performance and is not included by default, see Installation.

```
{
  "type": "grpc",
  "service_name": "TunService",
  "idle_timeout": "15s",
  "ping_timeout": "15s",
  "permit_without_stream": false
}

```

#### service_name

Service name of gRPC.

#### idle_timeout

In standard gRPC server/client:

If the transport doesn't see any activity after a duration of this time,
it pings the client to check if the connection is still active.

In default gRPC server/client:

It has the same behavior as the corresponding setting in HTTP transport.

#### ping_timeout

In standard gRPC server/client:

The timeout that after performing a keepalive check, the client will wait for activity.
If no activity is detected, the connection will be closed.

In default gRPC server/client:

It has the same behavior as the corresponding setting in HTTP transport.

#### permit_without_stream

In standard gRPC client:

If enabled, the client transport sends keepalive pings even with no active connections.
If disabled, when there are no active connections, idle_timeout and ping_timeout will be ignored and no keepalive
pings will be sent.

`idle_timeout``ping_timeout`Disabled by default.

### HTTPUpgrade

```
{
  "type": "httpupgrade",
  "host": "",
  "path": "",
  "headers": {}
}

```

#### host

Host domain.

The server will verify if not empty.

#### path

Path of HTTP request.

The server will verify.

#### headers

Extra headers of HTTP request.

The server will write in response if not empty.


---

## Wi-Fi State

**Source URL**: <https://sing-box.sagernet.org/configuration/shared/wifi-state/>

# Wi-Fi State

Changes in sing-box 1.13.0

Linux support
 Windows support

sing-box can monitor Wi-Fi state to enable routing rules based on wifi_ssid and wifi_bssid.

`wifi_ssid``wifi_bssid`### Platform Support

| Platform | Support | Notes | 
| --- | --- | --- |
| Android |  | In graphical client | 
| Apple platforms |  | In graphical clients | 
| Linux |  | Requires supported daemon | 
| Windows |  | WLAN API | 
| Others |  |  | 

### Linux

Since sing-box 1.13.0

The following backends are supported and will be auto-detected in order of priority:

| Backend | Interface | 
| --- | --- |
| NetworkManager | D-Bus | 
| IWD | D-Bus | 
| wpa_supplicant | Unix socket | 
| ConnMan | D-Bus | 

### Windows

Since sing-box 1.13.0

Uses Windows WLAN API.


---

## Deprecated Feature List

**Source URL**: <https://sing-box.sagernet.org/deprecated/>

# Deprecated Feature List

## 1.12.0

#### Legacy DNS server formats

DNS servers are refactored,
check Migration.

Compatibility for old formats will be removed in sing-box 1.14.0.

#### outbound DNS rule item

`outbound`Legacy outbound DNS rules are deprecated
and can be replaced by dial fields,
check Migration.

`outbound`#### Legacy ECH fields

ECH support has been migrated to use stdlib in sing-box 1.12.0,
which does not come with support for PQ signature schemes,
so pq_signature_schemes_enabled has been deprecated and no longer works.

`pq_signature_schemes_enabled`Also, dynamic_record_sizing_disabled has nothing to do with ECH,
was added by mistake, has been deprecated and no longer works.

`dynamic_record_sizing_disabled`These fields will be removed in sing-box 1.13.0.

## 1.11.0

#### Legacy special outbounds

Legacy special outbounds (block / dns) are deprecated
and can be replaced by rule actions,
check Migration.

`block``dns`Old fields will be removed in sing-box 1.13.0.

#### Legacy inbound fields

Legacy inbound fields （inbound.<sniff/domain_strategy/...> are deprecated
and can be replaced by rule actions,
check Migration.

`inbound.<sniff/domain_strategy/...>`Old fields will be removed in sing-box 1.13.0.

#### Destination override fields in direct outbound

Destination override fields (override_address / override_port) in direct outbound are deprecated
and can be replaced by rule actions,
check Migration.

`override_address``override_port`#### WireGuard outbound

WireGuard outbound is deprecated and can be replaced by endpoint,
check Migration.

Old outbound will be removed in sing-box 1.13.0.

#### GSO option in TUN

GSO has no advantages for transparent proxy scenarios, is deprecated and no longer works in TUN.

Old fields will be removed in sing-box 1.13.0.

## 1.10.0

#### TUN address fields are merged

inet4_address and inet6_address are merged into address,
inet4_route_address and inet6_route_address are merged into route_address,
inet4_route_exclude_address and inet6_route_exclude_address are merged into route_exclude_address.

`inet4_address``inet6_address``address``inet4_route_address``inet6_route_address``route_address``inet4_route_exclude_address``inet6_route_exclude_address``route_exclude_address`Old fields will be removed in sing-box 1.12.0.

#### Match source rule items are renamed

rule_set_ipcidr_match_source route and DNS rule items are renamed to
rule_set_ip_cidr_match_source and will be remove in sing-box 1.11.0.

`rule_set_ipcidr_match_source``rule_set_ip_cidr_match_source`#### Drop support for go1.18 and go1.19

Due to maintenance difficulties, sing-box 1.10.0 requires at least Go 1.20 to compile.

## 1.8.0

#### Cache file and related features in Clash API

cache_file and related features in Clash API is migrated to independent cache_file options,
check Migration.

`cache_file``cache_file`#### GeoIP

GeoIP is deprecated and will be removed in sing-box 1.12.0.

The maxmind GeoIP National Database, as an IP classification database,
is not entirely suitable for traffic bypassing,
and all existing implementations suffer from high memory usage and difficult management.

sing-box 1.8.0 introduces rule-set, which can completely replace GeoIP,
check Migration.

#### Geosite

Geosite is deprecated and will be removed in sing-box 1.12.0.

Geosite, the domain-list-community project maintained by V2Ray as an early traffic bypassing solution,
suffers from a number of problems, including lack of maintenance, inaccurate rules, and difficult management.

`domain-list-community`sing-box 1.8.0 introduces rule-set, which can completely replace Geosite,
check Migration.

## 1.6.0

The following features will be marked deprecated in 1.5.0 and removed entirely in 1.6.0.

#### ShadowsocksR

ShadowsocksR support has never been enabled by default, since the most commonly used proxy sales panel in the
illegal industry stopped using this protocol, it does not make sense to continue to maintain it.

#### Proxy Protocol

Proxy Protocol is added by Pull Request, has problems, is only used by the backend of HTTP multiplexers such as nginx,
is intrusive, and is meaningless for proxy purposes.


---

## Build from source

**Source URL**: <https://sing-box.sagernet.org/installation/build-from-source/>

# Build from source

##  Requirements

### sing-box 1.11

- Go 1.23.1 - ~

### sing-box 1.10

- Go 1.20.0 - ~

### sing-box 1.9

- Go 1.18.5 - 1.22.x
- Go 1.20.0 - 1.22.x with tag with_quic, or with_utls enabled

`with_quic``with_utls`##  Simple Build

```
make

```

Or build and install binary to $GOBIN:

`$GOBIN````
make install

```

##  Custom Build

```
TAGS="tag_a tag_b" make

```

or

```
go build -tags "tag_a tag_b" ./cmd/sing-box

```

##  Build Tags

| Build Tag | Enabled by default | Description | 
| --- | --- | --- |
| with_quic |  | Build with QUIC support, see QUIC and HTTP3 DNS transports, Naive inbound, Hysteria Inbound, Hysteria Outbound and V2Ray Transport#QUIC. | 
| with_grpc | ️ | Build with standard gRPC support, see V2Ray Transport#gRPC. | 
| with_dhcp |  | Build with DHCP support, see DHCP DNS transport. | 
| with_wireguard |  | Build with WireGuard support, see WireGuard outbound. | 
| with_utls |  | Build with uTLS support for TLS outbound, see TLS. | 
| with_acme |  | Build with ACME TLS certificate issuer support, see TLS. | 
| with_clash_api |  | Build with Clash API support, see Experimental. | 
| with_v2ray_api | ️ | Build with V2Ray API support, see Experimental. | 
| with_gvisor |  | Build with gVisor support, see Tun inbound and WireGuard outbound. | 
| with_embedded_tor (CGO required) | ️ | Build with embedded Tor support, see Tor outbound. | 
| with_tailscale |  | Build with Tailscale support, see Tailscale endpoint | 
| with_naive_outbound | ️ | Build with NaiveProxy outbound support, see NaiveProxy outbound. | 

`with_quic``with_grpc``with_dhcp``with_wireguard``with_utls``with_acme``with_clash_api``with_v2ray_api``with_gvisor``with_embedded_tor``with_tailscale``with_naive_outbound`It is not recommended to change the default build tag list unless you really know what you are adding.

##  with_naive_outbound

NaiveProxy outbound requires special build configurations depending on your target platform.

### Supported Platforms

| Platform | Architectures | Mode | Requirements | 
| --- | --- | --- | --- |
| Linux | amd64, arm64 | purego | None (library included in official releases) | 
| Linux | 386, amd64, arm, arm64 | CGO | Chromium toolchain, glibc >= 2.31 at runtime | 
| Linux (musl) | 386, amd64, arm, arm64 | CGO | Chromium toolchain | 
| Windows | amd64, arm64 | purego | None (library included in official releases) | 
| Apple platforms | * | CGO | Xcode | 
| Android | * | CGO | Android NDK | 

### Windows

Use with_purego tag.

`with_purego`For official releases, libcronet.dll is included in the archive. For self-built binaries, download from cronet-go releases and place in the same directory as sing-box.exe or in a directory listed in PATH.

`libcronet.dll``sing-box.exe``PATH`### Linux (purego, amd64/arm64 only)

Use with_purego tag.

`with_purego`For official releases, libcronet.so is included in the archive. For self-built binaries, download from cronet-go releases and place in the same directory as sing-box binary or in system library path.

`libcronet.so`### Linux (CGO)

See cronet-go.

- glibc build: Requires glibc >= 2.31 at runtime
- musl build: Use with_musl tag, statically linked, no runtime requirements

`with_musl`### Apple platforms / Android

See cronet-go.


---

## Docker

**Source URL**: <https://sing-box.sagernet.org/installation/docker/>

# Docker

##  Command

```
docker run -d \
  -v /etc/sing-box:/etc/sing-box/ \
  --name=sing-box \
  --restart=always \
  ghcr.io/sagernet/sing-box \
  -D /var/lib/sing-box \
  -C /etc/sing-box/ run

```

##  Compose

```
version: "3.8"
services:
  sing-box:
    image: ghcr.io/sagernet/sing-box
    container_name: sing-box
    restart: always
    volumes:
      - /etc/sing-box:/etc/sing-box/
    command: -D /var/lib/sing-box -C /etc/sing-box/ run

```


---

## Package Manager

**Source URL**: <https://sing-box.sagernet.org/installation/package-manager/>

# Package Manager

##  Repository Installation

```
sudo mkdir -p /etc/apt/keyrings &&
   sudo curl -fsSL https://sing-box.app/gpg.key -o /etc/apt/keyrings/sagernet.asc &&
   sudo chmod a+r /etc/apt/keyrings/sagernet.asc &&
   echo '
Types: deb
URIs: https://deb.sagernet.org/
Suites: *
Components: *
Enabled: yes
Signed-By: /etc/apt/keyrings/sagernet.asc
' | sudo tee /etc/apt/sources.list.d/sagernet.sources &&
   sudo apt-get update &&
   sudo apt-get install sing-box # or sing-box-beta

```

```
sudo dnf config-manager addrepo --from-repofile=https://sing-box.app/sing-box.repo &&
sudo dnf install sing-box # or sing-box-beta

```

```
sudo dnf config-manager --add-repo https://sing-box.app/sing-box.repo &&
sudo dnf -y install dnf-plugins-core &&
sudo dnf install sing-box # or sing-box-beta

```

##  Manual Installation

The script download and install the latest package from GitHub releases
for deb or rpm based Linux distributions, ArchLinux and OpenWrt.

```
curl -fsSL https://sing-box.app/install.sh | sh

```

or latest beta:

```
curl -fsSL https://sing-box.app/install.sh | sh -s -- --beta

```

or specific version:

```
curl -fsSL https://sing-box.app/install.sh | sh -s -- --version <version>

```

##  Managed Installation

| Type | Platform | Command | Link | 
| --- | --- | --- | --- |
| AUR | Arch Linux | ? -S sing-box |  | 
| nixpkgs | NixOS | nix-env -iA nixos.sing-box |  | 
| Homebrew | macOS / Linux | brew install sing-box |  | 
| APK | Alpine | apk add sing-box |  | 
| DEB | AOSC | apt install sing-box |  | 

`? -S sing-box``nix-env -iA nixos.sing-box``brew install sing-box``apk add sing-box``apt install sing-box`| Type | Platform | Command | Link | 
| --- | --- | --- | --- |
| Homebrew | macOS | brew install sing-box |  | 

`brew install sing-box`| Type | Platform | Command | Link | 
| --- | --- | --- | --- |
| Scoop | Windows | scoop install sing-box |  | 
| Chocolatey | Windows | choco install sing-box |  | 
| winget | Windows | winget install sing-box |  | 

`scoop install sing-box``choco install sing-box``winget install sing-box`| Type | Platform | Command | Link | 
| --- | --- | --- | --- |
| Termux | Android | pkg add sing-box |  | 

`pkg add sing-box`| Type | Platform | Command | Link | 
| --- | --- | --- | --- |
| FreshPorts | FreeBSD | pkg install sing-box |  | 

`pkg install sing-box`##  Problematic Sources

| Type | Platform | Link | Promblem(s) | 
| --- | --- | --- | --- |
| DEB | AOSC | aosc-os-abbs | Problematic build tag list modification | 
| Homebrew | / | homebrew-core | Problematic build tag list modification | 
| Termux | Android | termux-packages | Problematic build tag list modification | 
| FreshPorts | FreeBSD | FreeBSD ports | Old Go  (go1.20) | 

If you are a user of them, please report issues to them:

1. Please do not modify release build tags without full understanding of the related functionality: enabling non-default
   labels may result in decreased performance; the lack of default labels may cause user confusion.
2. sing-box supports compiling with some older Go versions, but it is not recommended (especially versions that are no
   longer supported by Go).

##  Service Management

For Linux systems with systemd, usually the installation already includes a sing-box service,
you can manage the service using the following command:

| Operation | Command | 
| --- | --- |
| Enable | sudo systemctl enable sing-box | 
| Disable | sudo systemctl disable sing-box | 
| Start | sudo systemctl start sing-box | 
| Stop | sudo systemctl stop sing-box | 
| Kill | sudo systemctl kill sing-box | 
| Restart | sudo systemctl restart sing-box | 
| Logs | sudo journalctl -u sing-box --output cat -e | 
| New Logs | sudo journalctl -u sing-box --output cat -f | 

`sudo systemctl enable sing-box``sudo systemctl disable sing-box``sudo systemctl start sing-box``sudo systemctl stop sing-box``sudo systemctl kill sing-box``sudo systemctl restart sing-box``sudo journalctl -u sing-box --output cat -e``sudo journalctl -u sing-box --output cat -f`
---

## TunnelVision

**Source URL**: <https://sing-box.sagernet.org/manual/misc/tunnelvision/>

# TunnelVision

TunnelVision is an attack that uses DHCP option 121 to set higher priority routes
so that traffic does not go through the VPN.

Reference: https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2024-3661

## Status

### Android

Android does not handle DHCP option 121 and is not affected.

### Apple platforms

Update sing-box graphical client to 1.9.0-rc.16 or newer,
then enable includeAllNetworks in Settings — Packet Tunnel and you will be unaffected.

`1.9.0-rc.16``includeAllNetworks``Settings``Packet Tunnel`Note: when includeAllNetworks is enabled, the default TUN stack is changed to gvisor,
and the system and mixed stacks are not available.

`includeAllNetworks``gvisor``system``mixed`### Linux

Update sing-box to 1.9.0-rc.16 or newer, rules generated by auto-route are unaffected.

`1.9.0-rc.16``auto-route`### Windows

No solution yet.

## Workarounds

- Don't connect to untrusted networks
- Relay untrusted network through another device
- Just ignore it


---

## Hysteria 2

**Source URL**: <https://sing-box.sagernet.org/manual/proxy-protocol/hysteria2/>

# Hysteria 2

Hysteria 2 is a simple, Chinese-made protocol based on QUIC.
The selling point is Brutal, a congestion control algorithm that
tries to achieve a user-defined bandwidth despite packet loss.

Warning

Even though GFW rarely blocks UDP-based proxies, such protocols actually have far more obvious characteristics than TCP based proxies.

| Specification | Resists passive detection | Resists active probes | 
| --- | --- | --- |
| hysteria.network |  |  | 

##  Password Generator

| Generate Password | Action | 
| --- | --- |
|  | Refresh | 

````##  Difference from official Hysteria

The official program supports an authentication method called userpass,
which essentially uses a combination of <username>:<password> as the actual password,
while sing-box does not provide this alias.
To use sing-box with the official program, you need to fill in that combination as the actual password.

`<username>:<password>`##  Server Example

Replace up_mbps and down_mbps values with the actual bandwidth of your server.

`up_mbps``down_mbps````
 {
  "inbounds": [
    {
      "type": "hysteria2",
      "listen": "::",
      "listen_port": 8080,
      "up_mbps": 100,
      "down_mbps": 100,
      "users": [
        {
          "name": "sekai",
          "password": "<password>"
        }
      ],
      "tls": {
        "enabled": true,
        "server_name": "example.org",
        "key_path": "/path/to/key.pem",
        "certificate_path": "/path/to/certificate.pem"
      }
    }
  ]
}

```

```
 {
  "inbounds": [
    {
      "type": "hysteria2",
      "listen": "::",
      "listen_port": 8080,
      "up_mbps": 100,
      "down_mbps": 100,
      "users": [
        {
          "name": "sekai",
          "password": "<password>"
        }
      ],
      "tls": {
        "enabled": true,
        "server_name": "example.org",
        "acme": {
          "domain": "example.org",
          "email": "[email protected]"
        }
      }
    }
  ]
}

```

```
 {
  "inbounds": [
    {
      "type": "hysteria2",
      "listen": "::",
      "listen_port": 8080,
      "up_mbps": 100,
      "down_mbps": 100,
      "users": [
        {
          "name": "sekai",
          "password": "<password>"
        }
      ],
      "tls": {
        "enabled": true,
        "server_name": "example.org",
        "acme": {
          "domain": "example.org",
          "email": "[email protected]",
          "dns01_challenge": {
            "provider": "cloudflare",
            "api_token": "my_token"
          }
        }
      }
    }
  ]
}

```

##  Client Example

Replace up_mbps and down_mbps values with the actual bandwidth of your client.

`up_mbps``down_mbps````
{
  "outbounds": [
    {
      "type": "hysteria2",
      "server": "127.0.0.1",
      "server_port": 8080,
      "up_mbps": 100,
      "down_mbps": 100,
      "password": "<password>",
      "tls": {
        "enabled": true,
        "server_name": "example.org"
      }
    }
  ]
}

```

Tip

Use sing-box merge command to merge configuration and certificate into one file.

`sing-box merge````
{
  "outbounds": [
    {
      "type": "hysteria2",
      "server": "127.0.0.1",
      "server_port": 8080,
      "up_mbps": 100,
      "down_mbps": 100,
      "password": "<password>",
      "tls": {
        "enabled": true,
        "server_name": "example.org",
        "certificate_path": "/path/to/certificate.pem"
      }
    }
  ]
}

```

```
{
  "outbounds": [
    {
      "type": "hysteria2",
      "server": "127.0.0.1",
      "server_port": 8080,
      "up_mbps": 100,
      "down_mbps": 100,
      "password": "<password>",
      "tls": {
        "enabled": true,
        "server_name": "example.org",
        "insecure": true
      }
    }
  ]
}

```


---

## Shadowsocks

**Source URL**: <https://sing-box.sagernet.org/manual/proxy-protocol/shadowsocks/>

# Shadowsocks

Shadowsocks is the most well-known Chinese-made proxy protocol.
It exists in multiple versions, but only AEAD 2022 ciphers 
over TCP with multiplexing is recommended.

| Ciphers | Specification | Cryptographically sound | Resists passive detection | Resists active probes | 
| --- | --- | --- | --- | --- |
| Stream Ciphers | shadowsocks.org |  |  |  | 
| AEAD | shadowsocks.org |  |  |  | 
| AEAD 2022 | shadowsocks.org |  |  |  | 

(We strongly recommend using multiplexing to send UDP traffic over TCP, because
doing otherwise is vulnerable to passive detection.)

##  Password Generator

| For 2022-blake3-aes-128-gcm cipher | For other ciphers | Action | 
| --- | --- | --- |
|  |  | Refresh | 

`2022-blake3-aes-128-gcm`````````##  Server Example

```
 {
  "inbounds": [
    {
      "type": "shadowsocks",
      "listen": "::",
      "listen_port": 8080,
      "network": "tcp",
      "method": "2022-blake3-aes-128-gcm",
      "password": "<password>",
      "multiplex": {
        "enabled": true
      }
    }
  ]
}

```

```
 {
  "inbounds": [
    {
      "type": "shadowsocks",
      "listen": "::",
      "listen_port": 8080,
      "network": "tcp",
      "method": "2022-blake3-aes-128-gcm",
      "password": "<server_password>",
      "users": [
        {
          "name": "sekai",
          "password": "<user_password>"
        }
      ],
      "multiplex": {
        "enabled": true
      }
    }
  ]
}

```

##  Client Example

```
{
  "outbounds": [
    {
      "type": "shadowsocks",
      "server": "127.0.0.1",
      "server_port": 8080,
      "method": "2022-blake3-aes-128-gcm",
      "password": "<pasword>",
      "multiplex": {
        "enabled": true
      }
    }
  ]
}

```

```
{
  "outbounds": [
    {
      "type": "shadowsocks",
      "server": "127.0.0.1",
      "server_port": 8080,
      "method": "2022-blake3-aes-128-gcm",
      "password": "<server_pasword>:<user_password>",
      "multiplex": {
        "enabled": true
      }
    }
  ]
}

```


---

## Trojan

**Source URL**: <https://sing-box.sagernet.org/manual/proxy-protocol/trojan/>

# Trojan

Trojan is the most commonly used TLS proxy made in China. It can be used in various combinations.

| Protocol and implementation combination | Specification | Resists passive detection | Resists active probes | 
| --- | --- | --- | --- |
| Origin / trojan-gfw | trojan-gfw.github.io |  |  | 
| Basic Go implementation | / |  |  | 
| with privates transport by V2Ray | No formal definition |  |  | 
| with uTLS enabled | No formal definition |  |  | 

##  Password Generator

| Generate Password | Action | 
| --- | --- |
|  | Refresh | 

````##  Server Example

```
{
  "inbounds": [
    {
      "type": "trojan",
      "listen": "::",
      "listen_port": 8080,
      "users": [
        {
          "name": "example",
          "password": "password"
        }
      ],
      "tls": {
        "enabled": true,
        "server_name": "example.org",
        "key_path": "/path/to/key.pem",
        "certificate_path": "/path/to/certificate.pem"
      },
      "multiplex": {
        "enabled": true
      }
    }
  ]
}

```

```
{
  "inbounds": [
    {
      "type": "trojan",
      "listen": "::",
      "listen_port": 8080,
      "users": [
        {
          "name": "example",
          "password": "password"
        }
      ],
      "tls": {
        "enabled": true,
        "server_name": "example.org",
        "acme": {
          "domain": "example.org",
          "email": "[email protected]"
        }
      },
      "multiplex": {
        "enabled": true
      }
    }
  ]
}

```

```
{
  "inbounds": [
    {
      "type": "trojan",
      "listen": "::",
      "listen_port": 8080,
      "users": [
        {
          "name": "example",
          "password": "password"
        }
      ],
      "tls": {
        "enabled": true,
        "server_name": "example.org",
        "acme": {
          "domain": "example.org",
          "email": "[email protected]",
          "dns01_challenge": {
            "provider": "cloudflare",
            "api_token": "my_token"
          }
        }
      },
      "multiplex": {
        "enabled": true
      }
    }
  ]
}

```

##  Client Example

```
{
  "outbounds": [
    {
      "type": "trojan",
      "server": "127.0.0.1",
      "server_port": 8080,
      "password": "password",
      "tls": {
        "enabled": true,
        "server_name": "example.org"
      },
      "multiplex": {
        "enabled": true
      }
    }
  ]
}

```

Tip

Use sing-box merge command to merge configuration and certificate into one file.

`sing-box merge````
{
  "outbounds": [
    {
      "type": "trojan",
      "server": "127.0.0.1",
      "server_port": 8080,
      "password": "password",
      "tls": {
        "enabled": true,
        "server_name": "example.org",
        "certificate_path": "/path/to/certificate.pem"
      },
      "multiplex": {
        "enabled": true
      }
    }
  ]
}

```

```
{
  "outbounds": [
    {
      "type": "trojan",
      "server": "127.0.0.1",
      "server_port": 8080,
      "password": "password",
      "tls": {
        "enabled": true,
        "server_name": "example.org",
        "insecure": true
      },
      "multiplex": {
        "enabled": true
      }
    }
  ]
}

```


---

## Client

**Source URL**: <https://sing-box.sagernet.org/manual/proxy/client/>

# Client

###  Introduction

For a long time, the modern usage and principles of proxy clients
for graphical operating systems have not been clearly described.
However, we can categorize them into three types:
system proxy, firewall redirection, and virtual interface.

###  System Proxy

Almost all graphical environments support system-level proxies,
which are essentially ordinary HTTP proxies that only support TCP.

| Operating System / Desktop Environment | System Proxy | Application Support | 
| --- | --- | --- |
| Windows |  |  | 
| macOS |  |  | 
| GNOME/KDE |  |  | 
| Android | ROOT or adb (permission) is required |  | 
| Android/iOS (with sing-box graphical client) | via tun.platform.http_proxy |  | 

`tun.platform.http_proxy`As one of the most well-known proxy methods, it has many shortcomings:
many TCP clients that are not based on HTTP do not check and use the system proxy.
Moreover, UDP and ICMP traffics bypass the proxy.

```
flowchart LR
    dns[DNS query] -- Is HTTP request? --> proxy[HTTP proxy]
    dns --> leak[Leak]
    tcp[TCP connection] -- Is HTTP request? --> proxy
    tcp -- Check and use HTTP CONNECT? --> proxy
    tcp --> leak
    udp[UDP packet] --> leak
```

###  Firewall Redirection

This type of usage typically relies on the firewall or hook interface provided by the operating system,
such as Windows’ WFP, Linux’s redirect, TProxy and eBPF, and macOS’s pf.
Although it is intrusive and cumbersome to configure,
it remains popular within the community of amateur proxy open source projects like V2Ray,
due to the low technical requirements it imposes on the software.

###  Virtual Interface

All L2/L3 proxies (seriously defined VPNs, such as OpenVPN, WireGuard) are based on virtual network interfaces,
which is also the only way for all L4 proxies to work as VPNs on mobile platforms like Android, iOS.

The sing-box inherits and develops clash-premium’s TUN inbound (L3 to L4 conversion)
as the most reasonable method for performing transparent proxying.

```
flowchart TB
    packet[IP Packet]
    packet --> windows[Windows / macOS]
    packet --> linux[Linux]
    tun[TUN interface]
    windows -. route .-> tun
    linux -. iproute2 route/rule .-> tun
    tun --> gvisor[gVisor TUN stack]
    tun --> system[system TUN stack]
    assemble([L3 to L4 assemble])
    gvisor --> assemble
    system --> assemble
    assemble --> conn[TCP and UDP connections]
    conn --> router[sing-box Router]
    router --> direct[Direct outbound]
    router --> proxy[Proxy outbounds]
    router -- DNS hijack --> dns_out[DNS outbound]
    dns_out --> dns_router[DNS router]
    dns_router --> router
    direct --> adi([auto detect interface])
    proxy --> adi
    adi --> default[Default network interface in the system]
    default --> destination[Destination server]
    default --> proxy_server[Proxy server]
    proxy_server --> destination
```

##  Examples

### Basic TUN usage for Chinese users

```
{
  "dns": {
    "servers": [
      {
        "tag": "google",
        "type": "tls",
        "server": "8.8.8.8"
      },
      {
        "tag": "local",
        "type": "udp",
        "server": "223.5.5.5"
      }
    ],
    "strategy": "ipv4_only"
  },
  "inbounds": [
    {
      "type": "tun",
      "address": ["172.19.0.1/30"],
      "auto_route": true,
      // "auto_redirect": true, // On linux
      "strict_route": true
    }
  ],
  "outbounds": [
    // ...
    {
      "type": "direct",
      "tag": "direct"
    }
  ],
  "route": {
    "rules": [
      {
        "action": "sniff"
      },
      {
        "protocol": "dns",
        "action": "hijack-dns"
      },
      {
        "ip_is_private": true,
        "outbound": "direct"
      }
    ],
    "default_domain_resolver": "local",
    "auto_detect_interface": true
  }
}

```

```
{
  "dns": {
    "servers": [
      {
        "tag": "google",
        "type": "tls",
        "server": "8.8.8.8"
      },
      {
        "tag": "local",
        "type": "udp",
        "server": "223.5.5.5"
      }
    ]
  },
  "inbounds": [
    {
      "type": "tun",
      "address": ["172.19.0.1/30", "fdfe:dcba:9876::1/126"],
      "auto_route": true,
      // "auto_redirect": true, // On linux
      "strict_route": true
    }
  ],
  "outbounds": [
    // ...
    {
      "type": "direct",
      "tag": "direct"
    }
  ],
  "route": {
    "rules": [
      {
        "action": "sniff"
      },
      {
        "protocol": "dns",
        "action": "hijack-dns"
      },
      {
        "ip_is_private": true,
        "outbound": "direct"
      }
    ],
    "default_domain_resolver": "local",
    "auto_detect_interface": true
  }
}

```

```
{
  "dns": {
    "servers": [
      {
        "tag": "google",
        "type": "tls",
        "server": "8.8.8.8"
      },
      {
        "tag": "local",
        "type": "udp",
        "server": "223.5.5.5"
      },
      {
        "tag": "remote",
        "type": "fakeip",
        "inet4_range": "198.18.0.0/15",
        "inet6_range": "fc00::/18"
      }
    ],
    "rules": [
      {
        "query_type": [
          "A",
          "AAAA"
        ],
        "server": "remote"
      }
    ],
    "independent_cache": true
  },
  "inbounds": [
    {
      "type": "tun",
      "address": ["172.19.0.1/30","fdfe:dcba:9876::1/126"],
      "auto_route": true,
      // "auto_redirect": true, // On linux
      "strict_route": true
    }
  ],
  "outbounds": [
    // ...
    {
      "type": "direct",
      "tag": "direct"
    }
  ],
  "route": {
    "rules": [
      {
        "action": "sniff"
      },
      {
        "protocol": "dns",
        "action": "hijack-dns"
      },
      {
        "ip_is_private": true,
        "outbound": "direct"
      }
    ],
    "default_domain_resolver": "local",
    "auto_detect_interface": true
  }
}

```

### Traffic bypass usage for Chinese users

```
{
  "dns": {
    "servers": [
      {
        "tag": "google",
        "type": "tls",
        "server": "8.8.8.8"
      },
      {
        "tag": "local",
        "type": "https",
        "server": "223.5.5.5"
      }
    ],
    "rules": [
      {
        "rule_set": "geosite-geolocation-cn",
        "server": "local"
      },
      {
        "type": "logical",
        "mode": "and",
        "rules": [
          {
            "rule_set": "geosite-geolocation-!cn",
            "invert": true
          },
          {
            "rule_set": "geoip-cn"
          }
        ],
        "server": "local"
      }
    ]
  },
  "route": {
    "default_domain_resolver": "local",
    "rule_set": [
      {
        "type": "remote",
        "tag": "geosite-geolocation-cn",
        "format": "binary",
        "url": "https://raw.githubusercontent.com/SagerNet/sing-geosite/rule-set/geosite-geolocation-cn.srs"
      },
      {
        "type": "remote",
        "tag": "geosite-geolocation-!cn",
        "format": "binary",
        "url": "https://raw.githubusercontent.com/SagerNet/sing-geosite/rule-set/geosite-geolocation-!cn.srs"
      },
      {
        "type": "remote",
        "tag": "geoip-cn",
        "format": "binary",
        "url": "https://raw.githubusercontent.com/SagerNet/sing-geoip/rule-set/geoip-cn.srs"
      }
    ]
  },
  "experimental": {
    "cache_file": {
      "enabled": true,
      "store_rdrc": true
    },
    "clash_api": {
      "default_mode": "Enhanced"
    }
  }
}

```

```
{
  "dns": {
    "servers": [
      {
        "tag": "google",
        "type": "tls",
        "server": "8.8.8.8"
      },
      {
        "tag": "local",
        "type": "https",
        "server": "223.5.5.5"
      }
    ],
    "rules": [
      {
        "rule_set": "geosite-geolocation-cn",
        "server": "local"
      },
      {
        "type": "logical",
        "mode": "and",
        "rules": [
          {
            "rule_set": "geosite-geolocation-!cn",
            "invert": true
          },
          {
            "rule_set": "geoip-cn"
          }
        ],
        "server": "google",
        "client_subnet": "114.114.114.114/24" // Any China client IP address
      }
    ]
  },
  "route": {
    "default_domain_resolver": "local",
    "rule_set": [
      {
        "type": "remote",
        "tag": "geosite-geolocation-cn",
        "format": "binary",
        "url": "https://raw.githubusercontent.com/SagerNet/sing-geosite/rule-set/geosite-geolocation-cn.srs"
      },
      {
        "type": "remote",
        "tag": "geosite-geolocation-!cn",
        "format": "binary",
        "url": "https://raw.githubusercontent.com/SagerNet/sing-geosite/rule-set/geosite-geolocation-!cn.srs"
      },
      {
        "type": "remote",
        "tag": "geoip-cn",
        "format": "binary",
        "url": "https://raw.githubusercontent.com/SagerNet/sing-geoip/rule-set/geoip-cn.srs"
      }
    ]
  },
  "experimental": {
    "cache_file": {
      "enabled": true,
      "store_rdrc": true
    },
    "clash_api": {
      "default_mode": "Enhanced"
    }
  }
}

```

```
{
  "outbounds": [
    {
      "type": "direct",
      "tag": "direct"
    }
  ],
  "route": {
    "rules": [
      {
        "action": "sniff"
      },
      {
        "type": "logical",
        "mode": "or",
        "rules": [
          {
            "protocol": "dns"
          },
          {
            "port": 53
          }
        ],
        "action": "hijack-dns"
      },
      {
        "ip_is_private": true,
        "outbound": "direct"
      },
      {
        "type": "logical",
        "mode": "or",
        "rules": [
          {
            "port": 853
          },
          {
            "network": "udp",
            "port": 443
          },
          {
            "protocol": "stun"
          }
        ],
        "action": "reject"
      },
      {
        "rule_set": "geosite-geolocation-cn",
        "outbound": "direct"
      },
      {
        "type": "logical",
        "mode": "and",
        "rules": [
          {
            "rule_set": "geoip-cn"
          },
          {
            "rule_set": "geosite-geolocation-!cn",
            "invert": true
          }
        ],
        "outbound": "direct"
      }
    ],
    "rule_set": [
      {
        "type": "remote",
        "tag": "geoip-cn",
        "format": "binary",
        "url": "https://raw.githubusercontent.com/SagerNet/sing-geoip/rule-set/geoip-cn.srs"
      },
      {
        "type": "remote",
        "tag": "geosite-geolocation-cn",
        "format": "binary",
        "url": "https://raw.githubusercontent.com/SagerNet/sing-geosite/rule-set/geosite-geolocation-cn.srs"
      }
    ]
  }
}

```


---

## Server

**Source URL**: <https://sing-box.sagernet.org/manual/proxy/server/>

# Server

To use sing-box as a proxy protocol server, you pretty much only need to configure the inbound for that protocol.

The Proxy Protocol menu below contains descriptions and configuration examples
of recommended protocols for bypassing GFW.


---

## Migration

**Source URL**: <https://sing-box.sagernet.org/migration/>

# Migration

## 1.12.0

### Migrate to new DNS server formats

DNS servers are refactored for better performance and scalability.

References

DNS Server /
Legacy DNS Server

```
{
  "dns": {
    "servers": [
      {
        "address": "local"
      }
    ]
  }
}

```

```
{
  "dns": {
    "servers": [
      {
        "type": "local"
      }
    ]
  }
}

```

```
{
  "dns": {
    "servers": [
      {
        "address": "tcp://1.1.1.1"
      }
    ]
  }
}

```

```
{
  "dns": {
    "servers": [
      {
        "type": "tcp",
        "server": "1.1.1.1"
      }
    ]
  }
}

```

```
{
  "dns": {
    "servers": [
      {
        "address": "1.1.1.1"
      }
    ]
  }
}

```

```
{
  "dns": {
    "servers": [
      {
        "type": "udp",
        "server": "1.1.1.1"
      }
    ]
  }
}

```

```
{
  "dns": {
    "servers": [
      {
        "address": "tls://1.1.1.1"
      }
    ]
  }
}

```

```
{
  "dns": {
    "servers": [
      {
        "type": "tls",
        "server": "1.1.1.1"
      }
    ]
  }
}

```

```
{
  "dns": {
    "servers": [
      {
        "address": "https://1.1.1.1/dns-query"
      }
    ]
  }
}

```

```
{
  "dns": {
    "servers": [
      {
        "type": "https",
        "server": "1.1.1.1"
      }
    ]
  }
}

```

```
{
  "dns": {
    "servers": [
      {
        "address": "quic://1.1.1.1"
      }
    ]
  }
}

```

```
{
  "dns": {
    "servers": [
      {
        "type": "quic",
        "server": "1.1.1.1"
      }
    ]
  }
}

```

```
{
  "dns": {
    "servers": [
      {
        "address": "h3://1.1.1.1/dns-query"
      }
    ]
  }
}

```

```
{
  "dns": {
    "servers": [
      {
        "type": "h3",
        "server": "1.1.1.1"
      }
    ]
  }
}

```

```
{
  "dns": {
    "servers": [
      {
        "address": "dhcp://auto"
      },
      {
        "address": "dhcp://en0"
      }
    ]
  }
}

```

```
{
  "dns": {
    "servers": [
      {
        "type": "dhcp",
      },
      {
        "type": "dhcp",
        "interface": "en0"
      }
    ]
  }
}

```

```
{
  "dns": {
    "servers": [
      {
        "address": "1.1.1.1"
      },
      {
        "address": "fakeip",
        "tag": "fakeip"
      }
    ],
    "rules": [
      {
        "query_type": [
          "A",
          "AAAA"
        ],
        "server": "fakeip"
      }
    ],
    "fakeip": {
      "enabled": true,
      "inet4_range": "198.18.0.0/15",
      "inet6_range": "fc00::/18"
    }
  }
}

```

```
{
  "dns": {
    "servers": [
      {
        "type": "udp",
        "server": "1.1.1.1"
      },
      {
        "type": "fakeip",
        "tag": "fakeip",
        "inet4_range": "198.18.0.0/15",
        "inet6_range": "fc00::/18"
      }
    ],
    "rules": [
      {
        "query_type": [
          "A",
          "AAAA"
        ],
        "server": "fakeip"
      }
    ]
  }
}

```

```
{
  "dns": {
    "servers": [
      {
        "address": "rcode://refused"
      }
    ]
  }
}

```

```
{
  "dns": {
    "rules": [
      {
        "domain": [
          "example.com"
        ],
        // other rules

        "action": "predefined",
        "rcode": "REFUSED"
      }
    ]
  }
}

```

```
{
  "dns": {
    "servers": [
      {
        "address": "https://dns.google/dns-query",
        "address_resolver": "google"
      },
      {
        "tag": "google",
        "address": "1.1.1.1"
      }
    ]
  }
}

```

```
{
  "dns": {
    "servers": [
      {
        "type": "https",
        "server": "dns.google",
        "domain_resolver": "google"
      },
      {
        "type": "udp",
        "tag": "google",
        "server": "1.1.1.1"
      }
    ]
  }
}

```

```
{
  "dns": {
    "servers": [
      {
        "address": "1.1.1.1",
        "strategy": "ipv4_only"
      },
      {
        "tag": "google",
        "address": "8.8.8.8",
        "strategy": "prefer_ipv6"
      }
    ],
    "rules": [
      {
        "domain": "google.com",
        "server": "google"
      }
    ]
  }
}

```

```
{
  "dns": {
    "servers": [
      {
        "type": "udp",
        "server": "1.1.1.1"
      },
      {
        "type": "udp",
        "tag": "google",
        "server": "8.8.8.8"
      }
    ],
    "rules": [
      {
        "domain": "google.com",
        "server": "google",
        "strategy": "prefer_ipv6"
      }
    ],
    "strategy": "ipv4_only"
  }
}

```

```
{
  "dns": {
    "servers": [
      {
        "address": "1.1.1.1"
      },
      {
        "tag": "google",
        "address": "8.8.8.8",
        "client_subnet": "1.1.1.1"
      }
    ]
  }
}

```

```
{
  "dns": {
    "servers": [
      {
        "type": "udp",
        "server": "1.1.1.1"
      },
      {
        "type": "udp",
        "tag": "google",
        "server": "8.8.8.8"
      }
    ],
    "rules": [
      {
        "domain": "google.com",
        "server": "google",
        "client_subnet": "1.1.1.1"
      }
    ]
  }
}

```

### Migrate outbound DNS rule items to domain resolver

The legacy outbound DNS rules are deprecated and can be replaced by new domain resolver options.

References

DNS rule /
Dial Fields /
Route

```
{
  "dns": {
    "servers": [
      {
        "address": "local",
        "tag": "local"
      }
    ],
    "rules": [
      {
        "outbound": "any",
        "server": "local"
      }
    ]
  },
  "outbounds": [
    {
      "type": "socks",
      "server": "example.org",
      "server_port": 2080
    }
  ]
}

```

```
{
  "dns": {
    "servers": [
      {
        "type": "local",
        "tag": "local"
      }
    ]
  },
  "outbounds": [
    {
      "type": "socks",
      "server": "example.org",
      "server_port": 2080,
      "domain_resolver": {
        "server": "local",
        "rewrite_ttl": 60,
        "client_subnet": "1.1.1.1"
      },
      // or "domain_resolver": "local",
    }
  ],

  // or

  "route": {
    "default_domain_resolver": {
      "server": "local",
      "rewrite_ttl": 60,
      "client_subnet": "1.1.1.1"
    }
  }
}

```

### Migrate outbound domain strategy option to domain resolver

References

Dial Fields

The domain_strategy option in Dial Fields has been deprecated and can be replaced with the new domain resolver option.

`domain_strategy`Note that due to the use of Dial Fields by some of the new DNS servers introduced in sing-box 1.12,
some people mistakenly believe that domain_strategy is the same feature as in the legacy DNS servers.

`domain_strategy````
{
  "outbounds": [
    {
      "type": "socks",
      "server": "example.org",
      "server_port": 2080,
      "domain_strategy": "prefer_ipv4",
    }
  ]
}

```

```
 {
  "dns": {
    "servers": [
      {
        "type": "local",
        "tag": "local"
      }
    ]
  },
  "outbounds": [
    {
      "type": "socks",
      "server": "example.org",
      "server_port": 2080,
      "domain_resolver": {
        "server": "local",
        "strategy": "prefer_ipv4"
      }
    }
  ]
}

```

## 1.11.0

### Migrate legacy special outbounds to rule actions

Legacy special outbounds are deprecated and can be replaced by rule actions.

References

Rule Action / 
Block / 
DNS

```
{
  "outbounds": [
    {
      "type": "block",
      "tag": "block"
    }
  ],
  "route": {
    "rules": [
      {
        ...,

        "outbound": "block"
      }
    ]
  }
}

```

```
{
  "route": {
    "rules": [
      {
        ...,

        "action": "reject"
      }
    ]
  }
}

```

```
{
  "inbound": [
    {
      ...,

      "sniff": true
    }
  ],
  "outbounds": [
    {
      "tag": "dns",
      "type": "dns"
    }
  ],
  "route": {
    "rules": [
      {
        "protocol": "dns",
        "outbound": "dns"
      }
    ]
  }
}

```

```
{
  "route": {
    "rules": [
      {
        "action": "sniff"
      },
      {
        "protocol": "dns",
        "action": "hijack-dns"
      }
    ]
  }
}

```

### Migrate legacy inbound fields to rule actions

Inbound fields are deprecated and can be replaced by rule actions.

References

Listen Fields /
Rule / 
Rule Action / 
DNS Rule / 
DNS Rule Action

```
{
  "inbounds": [
    {
      "type": "mixed",
      "sniff": true,
      "sniff_timeout": "1s",
      "domain_strategy": "prefer_ipv4"
    }
  ]
}

```

```
{
  "inbounds": [
    {
      "type": "mixed",
      "tag": "in"
    }
  ],
  "route": {
    "rules": [
      {
        "inbound": "in",
        "action": "resolve",
        "strategy": "prefer_ipv4"
      },
      {
        "inbound": "in",
        "action": "sniff",
        "timeout": "1s"
      }
    ]
  }
}

```

### Migrate destination override fields to route options

Destination override fields in direct outbound are deprecated and can be replaced by route options.

References

Rule Action /
Direct

```
{
  "outbounds": [
    {
      "type": "direct",
      "override_address": "1.1.1.1",
      "override_port": 443
    }
  ]
}

```

```
{
  "route": {
    "rules": [
      {
        "action": "route-options", // or route
        "override_address": "1.1.1.1",
        "override_port": 443
      }
    ]
  }

```

### Migrate WireGuard outbound to endpoint

WireGuard outbound is deprecated and can be replaced by endpoint.

References

Endpoint /
WireGuard Endpoint /
WireGuard Outbound

```
{
  "outbounds": [
    {
      "type": "wireguard",
      "tag": "wg-out",

      "server": "127.0.0.1",
      "server_port": 10001,
      "system_interface": true,
      "gso": true,
      "interface_name": "wg0",
      "local_address": [
        "10.0.0.1/32"
      ],
      "private_key": "<private_key>",
      "peer_public_key": "<peer_public_key>",
      "pre_shared_key": "<pre_shared_key>",
      "reserved": [0, 0, 0],
      "mtu": 1408
    }
  ]
}

```

```
{
  "endpoints": [
    {
      "type": "wireguard",
      "tag": "wg-ep",
      "system": true,
      "name": "wg0",
      "mtu": 1408,
      "address": [
        "10.0.0.2/32"
      ],
      "private_key": "<private_key>",
      "listen_port": 10000,
      "peers": [
        {
          "address": "127.0.0.1",
          "port": 10001,
          "public_key": "<peer_public_key>",
          "pre_shared_key": "<pre_shared_key>",
          "allowed_ips": [
            "0.0.0.0/0"
          ],
          "persistent_keepalive_interval": 30,
          "reserved": [0, 0, 0]
        }
      ]
    }
  ]
}

```

## 1.10.0

### TUN address fields are merged

inet4_address and inet6_address are merged into address,
inet4_route_address and inet6_route_address are merged into route_address,
inet4_route_exclude_address and inet6_route_exclude_address are merged into route_exclude_address.

`inet4_address``inet6_address``address``inet4_route_address``inet6_route_address``route_address``inet4_route_exclude_address``inet6_route_exclude_address``route_exclude_address`References

TUN

```
{
  "inbounds": [
    {
      "type": "tun",
      "inet4_address": "172.19.0.1/30",
      "inet6_address": "fdfe:dcba:9876::1/126",
      "inet4_route_address": [
        "0.0.0.0/1",
        "128.0.0.0/1"
      ],
      "inet6_route_address": [
        "::/1",
        "8000::/1"
      ],
      "inet4_route_exclude_address": [
        "192.168.0.0/16"
      ],
      "inet6_route_exclude_address": [
        "fc00::/7"
      ]
    }
  ]
}

```

```
{
  "inbounds": [
    {
      "type": "tun",
      "address": [
        "172.19.0.1/30",
        "fdfe:dcba:9876::1/126"
      ],
      "route_address": [
        "0.0.0.0/1",
        "128.0.0.0/1",
        "::/1",
        "8000::/1"
      ],
      "route_exclude_address": [
        "192.168.0.0/16",
        "fc00::/7"
      ]
    }
  ]
}

```

## 1.9.5

### Bundle Identifier updates in Apple platform clients

Due to problems with our old Apple developer account,
we can only change Bundle Identifiers to re-list sing-box apps,
which means the data will not be automatically inherited.

For iOS, you need to back up your old data yourself (if you still have access to it);
for tvOS, you need to re-import profiles from your iPhone or iPad or create it manually;
for macOS, you can migrate the data folder using the following command:

```
cd ~/Library/Group\ Containers && \ 
  mv group.io.nekohasekai.sfa group.io.nekohasekai.sfavt

```

## 1.9.0

### domain_suffix behavior update

`domain_suffix`For historical reasons, sing-box's domain_suffix rule matches literal prefixes instead of the same as other projects.

`domain_suffix`sing-box 1.9.0 modifies the behavior of domain_suffix: If the rule value is prefixed with .,
the behavior is unchanged, otherwise it matches (domain|.+\.domain) instead.

`domain_suffix``.``(domain|.+\.domain)`### process_path format update on Windows

`process_path`The process_path rule of sing-box is inherited from Clash,
the original code uses the local system's path format (e.g. \Device\HarddiskVolume1\folder\program.exe),
but when the device has multiple disks, the HarddiskVolume serial number is not stable.

`process_path``\Device\HarddiskVolume1\folder\program.exe`sing-box 1.9.0 make QueryFullProcessImageNameW output a Win32 path (such as C:\folder\program.exe),
which will disrupt the existing process_path use cases in Windows.

`C:\folder\program.exe``process_path`## 1.8.0

###  Migrate cache file from Clash API to independent options

References

Clash API / 
Cache File

```
{
  "experimental": {
    "clash_api": {
      "cache_file": "cache.db", // default value
      "cahce_id": "my_profile2",
      "store_mode": true,
      "store_selected": true,
      "store_fakeip": true
    }
  }
}

```

```
{
  "experimental"  : {
    "cache_file": {
      "enabled": true,
      "path": "cache.db", // default value
      "cache_id": "my_profile2",
      "store_fakeip": true
    }
  }
}

```

###  Migrate GeoIP to rule-sets

References

GeoIP / 
Route / 
Route Rule / 
DNS Rule / 
rule-set

Tip

sing-box geoip commands can help you convert custom GeoIP into rule-sets.

`sing-box geoip````
{
  "route": {
    "rules": [
      {
        "geoip": "private",
        "outbound": "direct"
      },
      {
        "geoip": "cn",
        "outbound": "direct"
      },
      {
        "source_geoip": "cn",
        "outbound": "block"
      }
    ],
    "geoip": {
      "download_detour": "proxy"
    }
  }
}

```

```
{
  "route": {
    "rules": [
      {
        "ip_is_private": true,
        "outbound": "direct"
      },
      {
        "rule_set": "geoip-cn",
        "outbound": "direct"
      },
      {
        "rule_set": "geoip-us",
        "rule_set_ipcidr_match_source": true,
        "outbound": "block"
      }
    ],
    "rule_set": [
      {
        "tag": "geoip-cn",
        "type": "remote",
        "format": "binary",
        "url": "https://raw.githubusercontent.com/SagerNet/sing-geoip/rule-set/geoip-cn.srs",
        "download_detour": "proxy"
      },
      {
        "tag": "geoip-us",
        "type": "remote",
        "format": "binary",
        "url": "https://raw.githubusercontent.com/SagerNet/sing-geoip/rule-set/geoip-us.srs",
        "download_detour": "proxy"
      }
    ]
  },
  "experimental": {
    "cache_file": {
      "enabled": true // required to save rule-set cache
    }
  }
}

```

###  Migrate Geosite to rule-sets

References

Geosite / 
Route / 
Route Rule / 
DNS Rule / 
rule-set

Tip

sing-box geosite commands can help you convert custom Geosite into rule-sets.

`sing-box geosite````
{
  "route": {
    "rules": [
      {
        "geosite": "cn",
        "outbound": "direct"
      }
    ],
    "geosite": {
      "download_detour": "proxy"
    }
  }
}

```

```
{
  "route": {
    "rules": [
      {
        "rule_set": "geosite-cn",
        "outbound": "direct"
      }
    ],
    "rule_set": [
      {
        "tag": "geosite-cn",
        "type": "remote",
        "format": "binary",
        "url": "https://raw.githubusercontent.com/SagerNet/sing-geosite/rule-set/geosite-cn.srs",
        "download_detour": "proxy"
      }
    ]
  },
  "experimental": {
    "cache_file": {
      "enabled": true // required to save rule-set cache
    }
  }
}

```


---

## Sponsors

**Source URL**: <https://sing-box.sagernet.org/sponsors/>

# Sponsors

Do you or your friends use sing-box?

You can help keep the project bug-free and feature rich by sponsoring
the project maintainer via GitHub Sponsors.

## Commercial Sponsors

> Warp, Built for coding with multiple AI agents.

Warp, Built for coding with multiple AI agents.

## Special Sponsors

> Viral Tech, Inc.

Viral Tech, Inc.

Helping us re-list sing-box apps on the Apple Store.

> JetBrains

JetBrains

Free license for the amazing IDEs.


---

## Support

**Source URL**: <https://sing-box.sagernet.org/support/>

# Support

| Channel | Link | 
| --- | --- |
| GitHub Issues | https://github.com/SagerNet/sing-box/issues | 
| Telegram notification channel | https://t.me/yapnc | 
| Telegram user group | https://t.me/yapug | 
| Email | contact@sagernet.org | 


---


---

*End of Documentation*

For the latest updates, visit: https://sing-box.sagernet.org
