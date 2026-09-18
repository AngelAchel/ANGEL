# ANGEL Platform

[![CI](https://github.com/AngelAchel/ANGEL/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/AngelAchel/ANGEL/actions/workflows/ci.yml)
[![Go Report](https://golangci-lint.run/badge/github.com/AngelAchel/ANGEL)](https://golangci-lint.run/)
[![License](https://img.shields.io/badge/License-Custom-blue)](LICENSE)

> **ANGEL** is an offensive security framework for authorized penetration testing and red team engagements.
> 70-layer attack surface with real-time cross-module event bus orchestration.

## Architecture

ANGEL is organized into **13 module groups** spanning **70 security layers**:

| Layer Range | Domain |
|-------------|--------|
| 01–05 | Brain, C2 Implant, Listener, DBPost, Decoy, Resilience |
| 06–10 | Evasion, Kerberos, Lateral Movement, Persistence, Rootkit |
| 11–15 | Brain, Collector, Credential, Destruction, Orchestrator |
| 16–21 | Cleanup, Evidence, Exploit, Infra, OSINT, Report |
| 22–25 | Auth Bypass, Crypto, Destruction Chain, Implant Gen, Net Evasion |
| 26–40 | AI, API, C2, Cloud, Container, IR, Malware, Mobile, Physical, Social Engineering |
| 41–60 | Cache Smuggle, Cert Forgery, CSRF, DNSSEC, IoT, IPv6, LDAP, MDNS, SCADA, SQL Inject |
| 61–70 | ARP/DHCP, Biz Logic, Crypto, Deserialization, GraphQL, gRPC, Memory, Race Condition, VLAN |

All modules communicate via the **event bus** (`modules/eventbus/`) for real-time cross-layer orchestration.

## Quick Start

```bash
git clone https://github.com/AngelAchel/ANGEL.git
cd ANGEL
cp .env.example .env.local
# Edit .env.local — set TEAMSERVER_KEY, CRYPTO_KEY, JWT_SECRET
make build && make test && make lint
make setup
```

## Deployment

### Docker Compose (Recommended)
```bash
make setup
docker compose up -d
docker compose ps
```

Services:
- **angel-teamserver** — C2 server on port `8443`
- **angel-console** — API Gateway / Dashboard on port `3000`
- **angel-rules** — Rules engine (one-shot)
- **dvwa** — Target practice on port `8081`

### Native (Kali Linux)
```bash
CGO_ENABLED=0 make build
bash scripts/start-local.sh
bash scripts/health-check.sh
```

## Usage

```bash
# Generate implant
./bin/angel-generate -os linux -arch amd64 -server http://localhost:8443 -out lab/implants

# Define target scope
echo "192.168.1.1" > target.txt

# Start engagement
make engage SCOPE=target.txt

# Monitor dashboard
open http://localhost:3000

# Cleanup
make cleanup
```

## Make Targets

| Target | Description |
|--------|-------------|
| `make build` | Build all binaries |
| `make test` | Run test suite (93 packages) |
| `make lint` | Run golangci-lint |
| `make fmt` | Format code |
| `make setup` | Install deps + build + deploy |
| `make engage` | Start penetration engagement |
| `make docker` | Build Docker image |
| `make clean` | Remove build artifacts |

## Verification

```bash
make build    # All binaries compile
make test     # 93 packages, 0 failures
make lint     # 0 issues
./lab/test_lab_full.sh  # End-to-end lab tests
```

## Project Structure

```
ANGEL/
├── cmd/               # Entry points (teamserver, console, generator, rules-loader)
├── c2/                # C2 core (implant, profiles)
├── gateway/           # API Gateway (auth, RBAC, rate limit)
├── orchestrator/      # Orchestration engine
├── frontend/          # Angular dashboard
├── modules/           # 70-layer security modules
├── pkg/               # Shared packages (crypto, eventbus, logger, types)
├── scripts/           # Build, lint, health check scripts
├── tests/             # E2E and integration tests
├── docs/              # Documentation
├── lab/               # Lab environment (configs, implants, logs)
├── Dockerfile.*       # Container builds
└── docker-compose.yml # Docker orchestration
```

## Testing

```bash
make test                              # Full test suite
./lab/test_lab.sh                      # Layer-by-layer test
./lab/test_lab_full.sh                 # Full 70-layer test
```

## Legal

⚠️ **ANGEL is for authorized offensive security testing only.**

- Written permission required before any engagement
- Only target systems within scope
- All activity is logged via the evidence ledger
- Post-engagement cleanup ensures no artifacts remain

## Documentation

- [USAGE_GUIDE.md](USAGE_GUIDE.md) — Full usage guide
- [STRUKTUR_ANGEL.md](STRUKTUR_ANGEL.md) — 70-layer blueprint
- [docs/NATIVE-RUN.md](docs/NATIVE-RUN.md) — Native run (no Docker)
- [docs/KALI.md](docs/KALI.md) — Kali Linux compatibility
- [docs/TERMUX.md](docs/TERMUX.md) — Termux compatibility
- [docs/PRODUCTION.md](docs/PRODUCTION.md) — Production deployment

## License

Custom license — see LICENSE file.

---

**ANGEL Platform** — Offensive Security Framework
GitHub: https://github.com/AngelAchel/ANGEL
