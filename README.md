# 🛡️ ANGEL Platform v3.3

**Offensive Security Platform untuk Engagement Resmi**

> **Status:** FINAL & OPERATIONAL
> **Tujuan:** P0/P1, Hard/Expert, Full Attack, No Demo, No Placeholder
> **Legalitas:** Hanya digunakan pada sistem yang telah diizinkan

---

## 🚀 QUICK START

```bash
# 1. Clone repository
git clone <REPO_URL>
cd ANGEL

# 2. Set env secrets (WAJIB)
export TEAMSERVER_KEY=<strong-random-key>
export CRYPTO_KEY=<strong-random-key>
export JWT_SECRET=<strong-random-secret>

# 3. Build binary
make build

# 4. Jalankan lab (Docker)
docker compose up -d

# 5. Test end-to-end
./lab/test_lab_full.sh

# 6. Mulai engagement
make engage SCOPE=target.txt
```

---

## 📡 LAYANAN

| Service | Port | Keterangan |
|---------|------|------------|
| 🎯 Teamserver | 8443 | C2 server (HTTP/HTTPS/DNS/SMB) |
| 🌐 Console | 3000 | API Gateway |
| 🎮 DVWA | 8081 | Target latihan |

---

## ⚔️ FITUR UTAMA

| Kategori | Detail |
|----------|--------|
| **70 Layer** | 70 rentang layer, 83 paket Go |
| **376+ file Go** | Terstruktur per layer |
| **92 test automated** | Semua passing ✅ |
| **3 binary** | `angel`, `angel-console`, `angel-generate` |
| **Event Bus** | Semua komunikasi lewat central event bus |
| **700+ teknik** | Fallback chains, anti-analysis, opsec |

### Layer Distribution

```
layer01-05  ████████████████  16 modul (C2, Implant, Listener, DB, Sqli, etc.)
layer06-10  █████  5 modul (Evasion, Kerberos, Lateral, Persistence, Rootkit)
layer11-15  █████  5 modul (Brain, Collector, Credential, Destruction, Orchestrator)
layer16-21  ██████  6 modul (Cleanup, Evidence, Exploit, Infra, OSINT, Report)
layer22-25  ████  4 modul (Auth Bypass, Destruction Chain, Implant Gen, Net Evasion)
layer26-40  ████████████  15 modul (AI, API, Cloud, Container, IR, Malware, Mobile, etc.)
layer41-60  ██████████████  18 modul (Cache Smuggle, Cert Forgery, CSRF, DNSSEC, IoT, IPv6, LDAP, etc.)
layer61-70  ██████████  10 modul (ARP/DHCP, Biz Logic, Crypto, Deser, GraphQL, gRPC, etc.)
```

---

## 📖 CARA PAKAI LENGKAP

### 1. Setup Awal

```bash
# Clone
git clone <REPO_URL>
cd ANGEL

# Set environment secrets (WAJIB - tanpa ini panic)
export TEAMSERVER_KEY="ganti-dengan-key-amanan"
export CRYPTO_KEY="ganti-dengan-key-amanan"
export JWT_SECRET="ganti-dengan-secret"

# Install deps + build
make deps
make build
```

### 2. Jalankan Lab

```bash
# Mulai semua service
docker compose up -d

# Cek status
docker compose ps

# Lihat logs
docker compose logs -f angel-teamserver
```

### 3. Test

```bash
# Test cepat
make test

# Lab basic (10 endpoint)
./lab/test_lab.sh

# Lab comprehensive (25 check)
./lab/test_lab_full.sh

# Lab per-layer (83 package)
./lab/test_lab_layers.sh
```

### 4. Generate Implant

```bash
# Linux x64
./bin/angel-generate -os linux -arch amd64 -server http://localhost:8443 -out lab/implants

# Windows x64
./bin/angel-generate -os windows -arch amd64 -server http://localhost:8443 -out lab/implants
```

### 5. Mulai Engagement

```bash
# Define target
echo "192.168.1.1" > target.txt
echo "192.168.1.2" >> target.txt

# Start engagement
make engage SCOPE=target.txt
```

### 6. Monitor

```bash
# Dashboard
http://localhost:3000

# Cek status
docker compose ps
docker compose logs -f angel-teamserver
```

### 7. Post-Engagement

```bash
# Hentikan
docker compose down

# Cleanup
make cleanup

# Generate laporan
make report
```

---

## 🏗️ ARSITEKTUR

```
┌─────────────────────────────────────────────────┐
│              ANGEL PLATFORM v3.3                 │
├─────────────────────────────────────────────────┤
│                                                 │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐      │
│  │ Frontend │→ │ Gateway  │→ │ Teamserver│     │
│  │ (Angular)│  │ (port 3000)│  │ (port 8443)│    │
│  └──────────┘  └──────────┘  └────┬─────┘     │
│                                   │            │
│  ┌────────────────────────────────┼────────┐  │
│  │           Event Bus            │        │  │
│  └──┬─────┬─────┬─────┬─────┬────┼──┬─────┘  │
│     │     │     │     │     │    │          │
│  ┌──┴──┐┌┴──┐┌┴──┐┌──┴┐┌──┴┐┌─┴─┐┌─┴─┐     │
│  │Brain││Impl││Auth││Mob││AI ││Web││IoT│...  │
│  │     ││Gen ││Byp ││key││   ││3  ││   │     │
│  └─────┘└───┘└────┘└───┘└───┘└───┘└───┘     │
│                                                 │
│  ┌───────────────────────────────────────────┐ │
│  │  Target: DVWA (port 8081)                 │ │
│  └───────────────────────────────────────────┘ │
└─────────────────────────────────────────────────┘
```

Setiap layer terpisah fungsional, komunikasi lewat **Event Bus Protocol**.

---

## 📋 MAKEFILE TARGETS

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

## 🔒 LEGALITAS

- Seluruh aktivitas hanya pada sistem yang telah diizinkan
- Harus memiliki kontrak, izin tertulis, dan persetujuan founder
- Hanya digunakan untuk engagement resmi offensive security

### Prinsip Dasar
1. **"No copy-paste"** — tiap baris ditulis sendiri
2. **"Signature-free"** — defender gak kenal
3. **"Modular"** — tiap modul jalan sendiri, tapi orchestrated
4. **"Evidentiary"** — tiap action ada bukti
5. **"Clean"** — post-engagement, semua hilang

---

## 📁 STRUKTUR REPOSITORI

```
ANGEL/
├── README.md              # File ini
├── USAGE_GUIDE.md         # Panduan lengkap
├── STRUKTUR_ANGEL.md      # Blueprint 70 layer
├── TEST_SCENARIOS.md      # Test scenarios
├── Makefile               # Build/test/release pipeline
├── .env.example           # Template environment
├── docker-compose.yml     # Docker orchestration
├── go.mod                 # Go module
├── bin/                   # Compiled binaries (gitignored)
├── lab/                   # Lab environment & test scripts
│   ├── test_lab.sh
│   ├── test_lab_full.sh
│   └── test_lab_layers.sh
├── cmd/                   # Entry points
│   ├── teamserver/
│   ├── console/
│   └── generator/
├── c2/                    # C2 core
├── gateway/               # API Gateway
├── orchestrator/          # LangGraph + Brain
├── frontend/              # Angular dashboard
├── infra/                 # Terraform + Ansible + WireGuard
├── modules/               # 70 layer modules
│   ├── layer01-05/        # 16 modul
│   ├── layer06-10/        # 5 modul
│   ├── layer11-15/        # 5 modul
│   ├── layer16-21/        # 6 modul
│   ├── layer22-25/        # 4 modul
│   ├── layer26-40/        # 15 modul
│   ├── layer41-60/        # 18 modul
│   └── layer61-70/        # 10 modul
├── pkg/                   # Shared packages
├── scripts/               # Build/lint/release pipeline
├── tests/                 # Test scenarios
└── docs/                  # Dokumentasi
```

---

## 🔧 TROUBLESHOOTING

| Masalah | Solusi |
|---------|--------|
| Teamserver tidak start | Cek `TEAMSERVER_KEY` sudah di-set |
| Implant tidak register | Cek network ke teamserver: `curl http://localhost:8443/` |
| Dashboard tidak akses | Cek `JWT_SECRET` sudah di-set |
| Panic/exit | Pastikan env var TEAMSERVER_KEY & JWT_SECRET di-set |
| Docker build gagal | Pastikan `cmd/teamserver/` & `cmd/console/` ada |

---

## 📊 STATUS

| Komponen | Status |
|----------|--------|
| Build | ✅ OK |
| Test | ✅ 92 packages PASS |
| Layer | ✅ 83/83 functional |
| Lab | ✅ 25 PASS / 0 FAIL |
| Docker | ✅ 3 containers running |
| Secret | ✅ Env var enforced |
| Placeholder | ✅ Bersih |
| Deploy | ⏳ Butuh env var asli + server |

---

## 📌 REFERENCE

- **GitHub:** https://github.com/SealAngel7/ANGEL
- **Blueprint:** STRUKTUR_ANGEL.md
- **Test Scenarios:** tests/TEST_SCENARIOS.md
- **Dokumentasi:** docs/ directory
- **Usage Guide:** USAGE_GUIDE.md

---

## ⚠️ PERINGATAN

⚠️ **Hanya gunakan pada sistem yang telah diizinkan.**
⚠️ **Pastikan memiliki kontrak, izin tertulis, dan persetujuan founder.**
⚠️ **Platform dirancang untuk engagement offensive security resmi.**
⚠️ **Semua aktivitas dilacak melalui evidence ledger (CHAIN_CUSTODY).**

---

**ANGEL Platform - Offensive Security Framework untuk Engagement Resmi**