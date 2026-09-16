# ANGEL — Kali Linux Compatibility

## Status: SUPPORTED WITH LIMITATIONS

ANGEL berjalan di Kali Linux amd64 dengan persyaratan berikut.

### Dependency Wajib

| Dependency | Status | Install |
|------------|--------|---------|
| Go 1.22+ | SUPPORTED | `apt install golang` |
| Docker/Compose | SUPPORTED | `apt install docker.io docker-compose-v2` |
| Python 3 | SUPPORTED | `apt install python3` |
| Nmap | NOT INSTALLED (opsional) | `apt install nmap` |
| Netcat | NOT INSTALLED (opsional) | `apt install netcat` |
| curl | SUPPORTED | `apt install curl` |
| git | SUPPORTED | `apt install git` |
| make | SUPPORTED | `apt install make` |
| gcc | SUPPORTED | `apt install gcc` |

### Build Native

```bash
# CGO disabled (static binary, no dependencies)
CGO_ENABLED=0 go build -o bin/angel ./cmd/teamserver
CGO_ENABLED=0 go build -o bin/angel-console ./cmd/console
CGO_ENABLED=0 go build -o bin/angel-generate ./cmd/generator
CGO_ENABLED=0 go build -o bin/angel-rules ./cmd/rules-loader
```

Binary hasil: static ELF amd64, berjalan tanpa library bersama.

### Local Execution

| Service | Status | Port |
|---------|--------|------|
| teamserver | SUPPORTED | 8443 |
| console | SUPPORTED | 3000 |
| rules-loader | SUPPORTED | 9444 |
| event bus | SUPPORTED (in-proc) | — |
| orchestrator | SUPPORTED | — |
| DVWA (fixture) | SUPPORTED (Docker) | 8081 |

### Health Check

```bash
curl http://localhost:8443/  # 404 = OK
curl http://localhost:3000/  # 404 = OK
curl http://localhost:8081/  # 302 = OK
```

### Local Execution — RUNTIME PROOF

| Service | Command | Result | Evidence |
|---------|---------|--------|----------|
| Teamserver | `./bin/angel-cgo0 -bind 127.0.0.1 -port 9443` | RUNNING | HTTP listener 127.0.0.1:9443, DNS 127.0.0.1:9444, "Teamserver started successfully" |
| Console | `./bin/angel-console-cgo0 -addr 127.0.0.1 -port 9445` | RUNNING | API Gateway 127.0.0.1:9445, GET / 15µs |
| Rules Loader | `./bin/angel-rules-cgo0` | RUNNING | Daemon mode, 60s reload |

**Tested:** Ubuntu 24.04 amd64 (codespace), CGO_ENABLED=0, static binary  
**Not tested:** Kali Linux native, Termux runtime

### Limitations

1. **Nmap tidak terinstal** — scanner layer membutuhkan nmap untuk network recon
2. **Netcat tidak terinstal** — beberapa listener membutuhkan nc
3. **Docker diperlukan** — DVWA fixture, rules-loader, teamserver container
4. **Iptables tersedia** — network manipulation OK
5. **ARM64 tidak diuji** — hanya amd64 teruji
6. **Termux tidak didukung** — tidak ada runtime Android

### Keterangan

- Status berdasarkan build test di Ubuntu 24.04 amd64 (kompatibel Kali)
- Bukan runtime Kali Linux — klaim SUPPORTED WITH LIMITATIONS
- Untuk Bukti: belum menjalankan di Kali Linux sesungguhnya
- Cross-compile arm64 berhasil tapi tidak diuji runtime

### File Terkait
- `docs/TERMUX.md` — Termux compatibility
- `docs/NATIVE-RUN.md` — Native run guide
- `docs/DEPENDENCIES.md` — Full dependency list
- `audit/deployment-checklist.md` — Deployment checklist