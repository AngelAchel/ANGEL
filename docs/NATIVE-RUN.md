# ANGEL — Native Run Guide

## Quick Start (No Docker)

### Prerequisites
- Go 1.22+
- curl
- git
- make
- gcc (for CGO builds)

### Build

```bash
# Clone repo
git clone https://github.com/AngelAchel/ANGEL.git
cd ANGEL

# Build all binaries (CGO disabled - static, no dependencies)
CGO_ENABLED=0 make build

# Or build individually
CGO_ENABLED=0 go build -o bin/angel ./cmd/teamserver
CGO_ENABLED=0 go build -o bin/angel-console ./cmd/console
CGO_ENABLED=0 go build -o bin/angel-generate ./cmd/generator
CGO_ENABLED=0 go build -o bin/angel-rules ./cmd/rules-loader
```

### Set Environment Variables

```bash
export TEAMSERVER_KEY="your-key-here"
export CRYPTO_KEY="your-crypto-key-here"
export JWT_SECRET="your-jwt-secret-here"
```

### Start Services (Local)

```bash
# Terminal 1: Teamserver
./bin/angel -bind 127.0.0.1 -port 8443

# Terminal 2: Console
./bin/angel-console -addr 127.0.0.1 -port 3000

# Terminal 3: Rules Loader
./bin/angel-rules
```

### Health Check

```bash
# Check teamserver
curl http://127.0.0.1:8443/  # Expected: 404 (service running)

# Check console
curl http://127.0.0.1:3000/  # Expected: 404 (service running)

# Check rules loader
curl http://127.0.0.1:9444/  # Expected: 404 (service running)
```

### Run Tests

```bash
# Full test suite
go test -count=1 ./...

# Lab test
bash lab/test_lab_full.sh

# With race detection
go test -race -count=1 ./...

# With coverage
go test -count=1 -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```

### Stop Services

```bash
# Kill all angel processes
pkill -f "angel"

# Or specific
pkill -f "angel-cgo0"
pkill -f "angel-console-cgo0"
pkill -f "angel-rules-cgo0"
```

### Using Scripts

```bash
# Start all local services
bash scripts/start-local.sh

# Stop all local services
bash scripts/stop-local.sh

# Health check all services
bash scripts/health-check.sh

# Doctor check
./bin/angel-doctor doctor
```

### Runtime Proof (Actual Execution)

Date: 2026-09-16  
Environment: Ubuntu 24.04 amd64, CGO_ENABLED=0  
Binary: static ELF, no shared library dependency

| Service | Command | Output | Status |
|---------|---------|--------|--------|
| Teamserver | `./bin/angel-cgo0 -bind 127.0.0.1 -port 9443` | "Teamserver started successfully", HTTP listener 127.0.0.1:9443, DNS listener 127.0.0.1:9444 | RUNNING |
| Console | `./bin/angel-console-cgo0 -addr 127.0.0.1 -port 9445` | "API Gateway starting on 127.0.0.1:9445", GET / 15µs | RUNNING |
| Rules Loader | `./bin/angel-rules-cgo0` | Daemon mode, 60s reload ticker | RUNNING |

All services started and responded to requests before SIGTERM.

### Native vs Docker

| Feature | Native | Docker |
|---------|--------|--------|
| Build | CGO_ENABLED=0 | Docker build |
| Dependencies | None (static) | Docker engine |
| Ports | 127.0.0.1 only | 0.0.0.0 |
| Storage | ./lab/data | ./lab/data (volume) |
| Network | localhost only | docker network |
| DVWA | NOT AVAILABLE | vulnerables/web-dvwa |
| Rules | Local file only | + Supabase |
| Speed | Faster | Slower (container overhead) |
| Compatibility | Linux only | Linux + Docker |

### Native Limitations

1. **No Docker** — DVWA fixture unavailable
2. **No Docker network** — services must use different ports
3. **127.0.0.1 only** — no external access
4. **No Supabase** — rules limited to local file
5. **No container restart** — process must be managed manually

### Troubleshooting

| Issue | Solution |
|-------|----------|
| "TEAMSERVER_KEY not set" | export TEAMSERVER_KEY=... |
| "address already in use" | pkill angel, restart |
| "permission denied" | chmod +x bin/angel* |
| "command not found" | export PATH=$PATH:$(pwd)/bin |
| "go: command not found" | Install Go 1.22+ |
| "curl: command not found" | apt install curl |

### Architecture

```
[Console:3000] ←→ [Teamserver:8443] ←→ [DVWA:8081]
      ↑                  ↑
  HTTP API           C2 Protocol
      ↑                  ↑
  [EventBus] ←→ [Orchestrator]
      ↑
  [RulesLoader:9444]
```

All services communicate via in-process EventBus.
No external message broker required for native mode.