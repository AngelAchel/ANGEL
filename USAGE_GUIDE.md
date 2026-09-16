# ANGEL Platform v3.2 - USAGE GUIDE

> **Status:** FINAL & OPERATIONAL
> **Tujuan:** Offensive security engagement untuk testing P0/P1
> **Legalitas:** Hanya digunakan pada sistem yang telah diizinkan
> **Prinsip:** "No copy-paste" — tiap baris ditulis sendiri

---

## PRASYARAT

- Go 1.27+
- Docker + Docker Compose
- Make

---

## SECTION 1: INSTALASI & BUILD

### Langkah 1: Clone & Enter
```bash
git clone <repo>
cd angel
```

### Langkah 2: Set Environment Secrets (WAJIB)
```bash
export TEAMSERVER_KEY="<strong-random-key-min-32-char>"
export CRYPTO_KEY="<strong-random-key-min-32-char>"
export JWT_SECRET="<strong-random-secret>"
```

> **Catatan:** Tanpa env var ini, aplikasi akan panic di startup.
> Untuk persisten, tambahkan ke `.env.example` atau `~/.bashrc`.

### Langkah 3: Build Binary
```bash
make build
```

**Output:**
- `bin/angel` — Teamserver/C2 server
- `bin/angel-console` — API Gateway / Console
- `bin/angel-generate` — Implant generator

### Langkah 4: Test Suite
```bash
make test
```

**Output:** 92 packages, semestinya ALL PASS.

### Langkah 5: Lint
```bash
make lint
```

---

## SECTION 2: DEPLOYMENT LAB (DOCKER COMPOSE)

### Jalankan Semua Service
```bash
docker compose up -d
```

**Service:**
- `angel-teamserver` — C2 di port 8443, 8080, 443, 5353/udp, 4455
- `angel-console` — API Gateway di port 3000
- `angel-rules` — Rules engine (one-shot, bukan server)
- `dvwa` — Target latihan DVWA di port 8081

### Cek Status
```bash
docker compose ps
docker logs angel-teamserver
docker logs angel-rules
```

### Test Lab
```bash
./lab/test_lab.sh
```

### Hentikan
```bash
docker compose down
```

---

## SECTION 3: OPERASIONAL ENGAGEMENT

### Generate Implant
```bash
# Linux x64
./bin/angel-generate -os linux -arch amd64 -server http://localhost:8443 -out lab/implants

# Windows x64
./bin/angel-generate -os windows -arch amd64 -server http://localhost:8443 -out lab/implants
```

### Define Target Scope
```bash
echo "192.168.1.1" > target.txt
echo "192.168.1.2" >> target.txt
```

### Mulai Engagement
```bash
make engage SCOPE=target.txt
```

Atau via Docker:
```bash
docker compose up -d angel-teamserver angel-console
```

### During Engagement
- **Dashboard:** `http://localhost:3000`
- **Teamserver:** `http://localhost:8443`
- **Monitor logs:** `docker compose logs -f angel-teamserver`
- **Cek rules:** `docker logs angel-rules`
- **Status:** `docker compose ps`

---

## SECTION 4: POST-ENGAGEMENT CLEANUP

```bash
docker compose down
make cleanup
```

### Generate Laporan
```bash
make report
# Output: report.html
```

---

## SECTION 5: STRUKTUR REPOSITORI

```
ANGEL/
├── Makefile           # Entry point: make build / test / release
├── .env.example       # Template environment
├── go.mod             # Go module
├── cmd/               # Entry points (teamserver, console, generator, rules-loader)
├── c2/                # C2 core (implant, profiles, generator)
├── gateway/           # API Gateway (auth, RBAC, rate limit)
├── orchestrator/      # LangGraph orchestration + Brain
├── frontend/          # Angular dashboard
├── infra/             # Terraform + Ansible + WireGuard
├── modules/           # 70 layer modules (layer01-05 s/d layer61-70)
├── pkg/               # Shared packages (crypto, eventbus, logger, types)
├── scripts/           # Build / lint / release pipeline
├── tests/             # Test scenarios TC-001..TC-1346
├── docs/              # Dokumentasi + report template
└── lab/               # Lab environment (configs, implants, logs)
```

---

## SECTION 6: MAKEFILE TARGETS

| Target | Deskripsi | Contoh |
|--------|-----------|--------|
| `all` | Build + test + lint | `make all` |
| `build` | Build 3 binary | `make build` |
| `deps` | Install Go deps | `make deps` |
| `setup` | Deps + build | `make setup` |
| `test` | Run all tests | `make test` |
| `lint` | golangci-lint | `make lint` |
| `fmt` | Format kode | `make fmt` |
| `tidy` | Go mod tidy | `make tidy` |
| `clean` | Bersihkan artifacts | `make clean` |
| `engage` | Mulai engagement | `make engage SCOPE=target.txt` |
| `implant-generate` | Generate implant | `make implant-generate OS=linux TARGET=amd64` |
| `infra-deploy` | Deploy infra | `make infra-deploy ENV=production` |
| `c2-deploy` | Deploy C2 | `make c2-deploy` |
| `report` | Generate laporan | `make report` |
| `help` | Daftar target | `make help` |

---

## SECTION 7: LEGALITAS & PRINSIP

### Legalitas
- Hanya untuk sistem yang telah diizinkan
- Harus punya kontrak + izin tertulis

### Prinsip
1. "No copy-paste" — tiap baris ditulis sendiri
2. "If I can't explain every line, it doesn't go in"
3. "Signature-free" — defender gak kenal
4. "Modular" — tiap modul jalan sendiri, di-orchestrate
5. "Evidentiary" — tiap aksi ada bukti
6. "Clean" — post-engagement, semua hilang
7. "Resilient" — setiap kegagalan ada fallback
8. "Adaptive" — beradaptasi dengan environment
9. "Autonomous" — keputusan tanpa operator jika perlu
10. "Observable" — tiap aksi log dan terukur

---

## SECTION 8: TROUBLESHOOTING

**Teamserver tidak start:**
- Cek `TEAMSERVER_KEY` sudah di-set: `echo $TEAMSERVER_KEY`
- Cek port 8444 tidak bentrok: `ss -tlnp | grep 8443`

**Implant tidak register:**
- Cek network ke teamserver: `curl http://localhost:8443/`
- Cek listener: `docker compose ps angel-teamserver`

**Dashboard tidak akses:**
- Cek `JWT_SECRET` sudah di-set
- Cek port 3000: `curl http://localhost:3000/api/v1/health`

**Semua panic/exit:**
- Pastikan env var TEAMSERVER_KEY dan JWT_SECRET sudah di-set sebelum build

---

## SECTION 9: REFERENCE

- `STRUKTUR_ANGEL.md` — Blueprint 70 layer
- `TEST_SCENARIOS.md` — 1.346 test case (TC-001..TC-1346)
- `docs/report_template.md` — Template laporan