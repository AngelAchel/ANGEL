# ANGEL

<p align="center">

<img src="https://readme-typing-svg.demolab.com?font=Fira+Code&size=28&duration=3000&color=FF4444&center=true&vCenter=true&width=600&lines=ANGEL+Platform;Offensive+Security+Framework;70+Layers+of+Attack+Surface;Red+Team+Operations+Engine" alt="ANGEL Platform" />

</p>

<p align="center">

[![CI](https://github.com/AngelAchel/ANGEL/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/AngelAchel/ANGEL/actions/workflows/ci.yml)
[![Go Report](https://golangci-lint.run/badge/github.com/AngelAchel/ANGEL)](https://golangci-lint.run)
[![License](https://img.shields.io/badge/License-Custom-red)](LICENSE)
[![Docker](https://img.shields.io/badge/Docker-Deploy-blue)](https://github.com/AngelAchel/ANGEL)

</p>

> **ANGEL** — Offensive security framework for authorized penetration testing and red team engagements.
> 70-layer attack surface with real-time cross-module event bus orchestration.

---

## Capabilities

ANGEL delivers a complete red team toolkit spanning the full offensive lifecycle:

| Phase | Capabilities |
|-------|-------------|
| **Reconnaissance** | OSINT, network enumeration, service discovery, vulnerability mapping |
| **Initial Access** | C2 implant delivery, phishing, supply chain, credential stuffing |
| **Execution** | Command execution, script injection, memory-resident payloads |
| **Persistence** | Scheduled tasks, registry hooks, service installation, rootkit |
| **Privilege Escalation** | Kernel exploits, token manipulation, ACL abuse, credential dumping |
| **Defense Evasion** | Anti-analysis, anti-debug, sleep masking, syscall obfuscation |
| **Credential Access** | LSASS, browser wallets, VPN tokens, MFA bypass, biometric |
| **Discovery** | Domain recon, user enumeration, share scanning, subnet mapping |
| **Lateral Movement** | SMB, RDP, WinRM, SSH, WMI, Pass-the-Hash, Pass-the-Ticket |
| **Collection** | Filesystem, email, browser history, keylogger, screen capture |
| **Exfiltration** | Encrypted channels, DNS tunneling, HTTP(S) covert, cloud storage |
| **Impact** | Data destruction, ransomware simulation, service disruption |

---

## Architecture

ANGEL operates as a distributed, event-driven framework. Every module communicates
through a real-time event bus enabling parallel, coordinated operations across all layers.

```
+-------------------------------------------------------------+
|                        OPERATIONS CENTER                          |
|                  ANGEL Console - Web Dashboard                    |
|              REST API . WebSocket . Real-time Monitoring        |
+------------------------------------------------------------+
|                       |
|            +----------+---------+
|            |      EVENT BUS       |
|            |  (Real-time Orchestration)|
|            |  96 modules connected |
|            |  Publish / Subscribe |
|            +----------+---------+
|                       |
+-----------------------+------------------------+
|                       |                        |
+----+-----------------+ +-------------+ +------+-----------------+
| C2 TEAMSERVER     | | IMPLANT GEN  | | RULES ENGINE      |
| Port 8443         | | 7 OS/Arch     | | JSON-based        |
| Encrypted         | | AES-256      | | Auto-execution    |
| Heartbeats        | | Domain Front  | |                    |
+-------------------+ +-------------+ +----------------------+
```

**70 Layers - 13 Domains:**

```
Layer 01-05   ████████████████████  Brain . C2 Implant . Listener . DBPost . Resilience
Layer 06-10   ████████████████████  Evasion . Kerberos . Lateral . Persistence . Rootkit
Layer 11-15   ████████████████████  Brain . Collector . Credential . Destruction . Orchestrator
Layer 16-21   ████████████████████  Cleanup . Evidence . Exploit . Infra . OSINT . Report
Layer 22-25   ████████████████████  Auth Bypass . Crypto . Destruction Chain . Implant Gen . Net Evasion
Layer 26-40   ████████████████████  AI . Cloud . Mobile . Supply Chain . Wireless . Zero Trust
Layer 41-60   ████████████████████  SQLi . XSS . CSRF . DNSSEC . IoT . SCADA . Compliance
Layer 61-70   ████████████████████  Memory . GraphQL . gRPC . Race Condition . VLAN . Deserialization
```

---

## Quick Start

### Prerequisites

- Go 1.22+
- Node.js 20+
- Docker & Docker Compose
- Git

### Clone and Deploy

```bash
git clone https://github.com/AngelAchel/ANGEL.git
cd ANGEL
cp .env.example .env.local
make build && make test && make lint
make setup
```

### Docker Deployment (Recommended)

```bash
make engage SCOPE=target.txt
docker compose up -d
```

Services launch automatically:

| Service | Port | Purpose |
|---------|------|---------|
| **angel-teamserver** | `8443` | C2 server with encrypted channels |
| **angel-console** | `3000` | Dashboard & API gateway |
| **angel-rules** | -- | Rules engine (one-shot) |
| **dvwa** | `8081` | Target practice environment |

### Native Deployment (Kali Linux)

```bash
CGO_ENABLED=0 make build
bash scripts/start-local.sh
bash scripts/health-check.sh
```

---

## Operational Flow

```
+----------+     +----------+     +----------+     +----------+
|  TARGET  |---->|  IMPLANT |---->|  C2      |---->|  EXFIL   |
|  SCAN    |     |  DELIVER |     |  MANAGE  |     |  COLLECT |
+----------+     +----------+     +----------+     +----------+
      |                |                |                |
      v                v                v                v
 70 layers       AES-256 encrypted    Event bus        Real-time
 parallel         domain fronting     orchestration    command & ctrl
```

### Generate Implant

```bash
./bin/angel-generate -os linux -arch amd64 -server http://localhost:8443 -out lab/implants
./bin/angel-generate -os windows -arch amd64 -server https://target.com -evasion domain-fronting
```

### Define Scope and Engage

```bash
echo "192.168.1.0/24" > target.txt
make engage SCOPE=target.txt
```

### Monitor Dashboard

```
http://localhost:3000
```

Real-time monitoring of all agents, tasks, results, and activity logs.

### Cleanup

```bash
make cleanup
```

---

## Make Targets

| Target | Description |
|--------|-------------|
| `make build` | Build all 5 binaries (teamserver, console, rules, generate, doctor) |
| `make test` | Run all test suites across 93 packages |
| `make lint` | golangci-lint, staticcheck, go vet, errcheck |
| `make report` | Generate HTML engagement report |
| `make setup` | Initialize environment, copy .env.local |
| `make engage` | Start Docker services with target scope |
| `make cleanup` | Stop and remove all containers |
| `make doctor` | Verify system compatibility (Go, Docker, Node, ports) |

---

## Module Coverage

| Domain | Layers | Key Tools |
|--------|--------|-----------|
| **C2 Framework** | 01-05 | Cobalt Strike compatible, Havoc, Sliver protocols |
| **Evasion & Stealth** | 06-10 | 11 sleep handlers, 7 syscall methods, 15 anti-analysis |
| **Credential Theft** | 11-15 | LSASS, browser wallets, gaming, cloud, MFA, biometric |
| **Lateral Movement** | 06-10 | SMB beacon, WinRM, SSH, WMI, Pass-the-Hash |
| **Web Exploitation** | 41-60 | SQLi (8 DBMS), XSS, CSRF, cache smuggling, SSRF |
| **Network Attacks** | 41-60 | DNSSEC, IPv6, LDAP, SCADA, IoT, MDNS |
| **AI & Cloud** | 26-40 | LLM injection, cloud takeover, container escape |
| **Deserialization** | 61-70 | Java, Python, PHP, .NET, Ruby gadget chains |
| **Forensics** | 61-70 | Memory analysis, timeline reconstruction, artifact recovery |

---

## Deployment Options

### Docker Compose

Production-ready with TLS encryption, health checks, and automated recovery.

```yaml
angel-teamserver: 8443/tcp, 443/tcp, 5353/udp
angel-console:     3000/tcp
dvwa:              8081/tcp
```

### Native Linux/Kali

Direct binary execution with system-level access for kernel exploitation and hardware attacks.

### Cross-Platform

Compiles natively for Linux, Windows, and macOS (amd64/arm64).

---

## Security & Compliance

- **Authorized use only** — ANGEL is designed for penetration testing with explicit written authorization.
- **No hardcoded credentials** — All secrets loaded from environment configuration.
- **End-to-end encryption** — AES-256 on all C2 channels, ECDH key exchange.
- **Event bus isolation** — Module communication via encrypted event bus, no direct imports.
- **Audit logging** — Full chain-of-custody logging for compliance reports.

---

## Documentation

- **[USAGE_GUIDE.md](USAGE_GUIDE.md)** — Complete operational guide
- **[STRUKTUR_ANGEL.md](STRUKTUR_ANGEL.md)** — Full architecture blueprint (70 layers)
- **[docs/API.md](docs/API.md)** — API reference
- **[docs/PRODUCTION.md](docs/PRODUCTION.md)** — Production deployment guide
- **[docs/NATIVE-RUN.md](docs/NATIVE-RUN.md)** — Native Linux setup

---

## License

Custom License — See [LICENSE](LICENSE) for details.

**ANGEL** is intended for authorized security professionals conducting penetration tests
and red team operations. Unauthorized use is illegal.

---

<p align="center">

<img src="https://img.shields.io/badge/ANGEL-Platform-red?style=for-the-badge" alt="ANGEL Platform">

**70 Layers . 96 Modules . Event-Driven . Production-Ready**

</p>
