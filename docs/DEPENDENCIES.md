# ANGEL — Dependencies

## Runtime Dependencies

### Mandatory

| Dependency | Version | Purpose | Install |
|------------|---------|---------|---------|
| Go | 1.22+ | Build & run | `apt install golang` |
| curl | 7.68+ | Health check | `apt install curl` |
| git | 2.34+ | Clone repo | `apt install git` |
| make | 4.3+ | Build automation | `apt install make` |

### Optional (Native Mode)

| Dependency | Version | Purpose | Install |
|------------|---------|---------|---------|
| nmap | 7.9+ | Network scanning | `apt install nmap` |
| netcat | 1.219+ | Network tool | `apt install netcat` |
| python3 | 3.10+ | Scripts | `apt install python3` |
| iptables | 1.8+ | Network rules | `apt install iptables` |

### Optional (Docker Mode)

| Dependency | Version | Purpose | Install |
|------------|---------|---------|---------|
| Docker | 24.0+ | Container runtime | `apt install docker.io` |
| Docker Compose | 2.0+ | Orchestration | `apt install docker-compose-v2` |

## Build Dependencies

| Dependency | Version | Purpose | Required |
|------------|---------|---------|----------|
| Go | 1.22+ | Compiler | YES |
| gcc | 13+ | CGO compiler | Only if CGO enabled |
| make | 4.3+ | Build tool | YES |
| git | 2.34+ | Source control | YES |

## Go Module Dependencies

```
github.com/angel-platform/angel
├── pkg/eventbus        (in-process event bus)
├── pkg/types           (shared types)
├── pkg/crypto          (encryption)
├── pkg/logger          (logging)
├── pkg/purchase        (license)
├── pkg/supabase        (cloud rules)
├── c2/generate         (implant generator)
├── c2/implant          (implant runtime)
├── c2/profiles         (malleable profiles)
├── gateway/            (HTTP API)
│   ├── auth            (JWT + RBAC)
│   ├── config          (configuration)
│   └── middleware      (HTTP middleware)
├── modules/layer01-05/ (C2, listener, sqli, nosql, resilience)
├── modules/layer06-10/ (evasion, kerberos, persistence, rootkit)
├── modules/layer11-15/ (collector, credential, destruction, orchestrator)
├── modules/layer16-21/ (cleanup, evidence, exploit, infra, osint)
├── modules/layer22-25/ (authbypass, implantgen, netevasion)
├── modules/layer26-40/ (AI, container, mobile, web3, wireless)
├── modules/layer41-60/ (ldap, iot, csrf, dnssec, methodology)
├── modules/layer61-70/ (crypto, vlan, deser, graphql, grpc)
├── cmd/teamserver      (C2 server entrypoint)
├── cmd/console         (console entrypoint)
├── cmd/generator       (implant generator entrypoint)
└── cmd/rules-loader    (rules loader entrypoint)
```

## External Services (Optional)

| Service | Purpose | Required |
|---------|---------|----------|
| Supabase | Cloud rules storage | No (local fallback) |
| SQLite | Database | No (in-memory for tests) |
| DVWA | Test fixture | No (Docker only) |

## Architecture Support

| Architecture | Build | Runtime | Status |
|--------------|-------|---------|--------|
| amd64 (x86_64) | YES | YES | SUPPORTED |
| arm64 (aarch64) | YES | NOT TESTED | NOT TESTED |
| 386 (x86) | YES | NOT TESTED | NOT TESTED |
| Android (Termux) | CROSS-COMPILE | NOT TESTED | NOT SUPPORTED |
| Windows | YES (cross) | NOT TESTED | NOT TESTED |
| macOS | YES (cross) | NOT TESTED | NOT TESTED |

## Network Requirements

| Port | Service | Protocol | Direction |
|------|---------|----------|-----------|
| 8443 | Teamserver | TCP | Listen |
| 3000 | Console | TCP | Listen |
| 9444 | Rules Loader | TCP | Listen |
| 8080 | Teamserver HTTP | TCP | Listen |
| 443 | Teamserver HTTPS | TCP | Listen |
| 5353 | DNS Listener | UDP | Listen |
| 4455 | SMB Beacon | TCP | Listen |
| 8081 | DVWA | TCP | Connect |
| 9443 | WebSocket | TCP | Listen |

## Storage Requirements

| Path | Purpose | Size |
|------|---------|------|
| ./lab/data | Database, rules | 100MB+ |
| ./lab/logs | Log files | 50MB+ |
| ./lab/implants | Generated implants | 10MB+ |
| ./audit | Audit reports | 10MB |
| ./docs | Documentation | 5MB |

## Permission Requirements

| Permission | Purpose | Linux | Termux |
|------------|---------|-------|--------|
| Network bind | Listen ports | OK | localhost only |
| File read/write | Lab data | OK | ~/storage/shared |
| Signal handling | Graceful shutdown | OK | Limited |
| Process fork | Subprocess | OK | Limited |
| iptables | Network rules | OK | NOT AVAILABLE |
| Raw socket | Packet sniff | root | NOT AVAILABLE |

## Minimum System Requirements

| Resource | Minimum | Recommended |
|----------|---------|-------------|
| RAM | 1GB | 4GB |
| Disk | 1GB | 5GB |
| CPU | 1 core | 2+ cores |
| OS | Linux | Linux (Kali/Ubuntu) |
| Architecture | amd64 | amd64/arm64 |

## Not Supported

- Windows (no native build tested)
- macOS (no native build tested)
- Termux (no runtime test)
- Android (no runtime test)
- iOS (not applicable)
- ARM 32-bit (not tested)
- RISC-V (not tested)
- MIPS (not tested)