# ANGEL Platform v3.3

**Offensive Security Platform untuk Engagement Resmi**

> **Status:** FINAL & OPERATIONAL
> **Tujuan:** P0/P1, Hard/Expert, Full Attack, No Demo, No Placeholder
> **Legalitas:** Hanya digunakan pada sistem yang telah diizinkan

---

## QUICK START

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

## LAYANAN

| Service | Port | Keterangan |
|---------|------|------------|
| Teamserver | 8443 | C2 server (HTTP/HTTPS/DNS/SMB) |
| Console | 3000 | API Gateway |
| DVWA | 8081 | Target latihan |

---

## FITUR UTAMA

| Kategori | Detail |
|----------|--------|
| **70 Layer** | 70 rentang layer, 83 paket Go |
| **376+ file Go** | Terstruktur per layer |
| **92 test automated** | Semua passing |
| **3 binary** | `angel`, `angel-console`, `angel-generate` |
| **Event Bus** | Semua komunikasi lewat central event bus |
| **700+ teknik** | Fallback chains, anti-analysis, opsec |

---

## CARA PAKAI

### Setup
```bash
make deps      # Install Go deps
make build     # Build 3 binary
make test      # Jalankan 92 package test
```

### Docker
```bash
docker compose up -d          # Mulai semua service
docker compose ps             # Cek status
docker compose logs -f angel-teamserver  # Lihat logs
docker compose down           # Hentikan
```

### Generate Implant
```bash
./bin/angel-generate -os linux -arch amd64 -server http://localhost:8443 -out lab/implants
./bin/angel-generate -os windows -arch amd64 -server http://localhost:8443 -out lab/implants
```

### Test Lab
```bash
./lab/test_lab.sh         # Basic (10 endpoint)
./lab/test_lab_full.sh    # Comprehensive (25 check)
./lab/test_lab_layers.sh  # Per-layer (83 package)
```

---

## ARSITEKTUR

```
Tier 1: Infrastructure   (Terraform/Ansible, 4 VPC nodes)
Tier 2: C2 Framework     (Implant, Teamserver, Listeners)
Tier 3: Orchestrator     (LangGraph + Brain + Fireteam)
Tier 4: API Gateway      (RBAC, Rate limiting, Auth)
Tier 5: Frontend         (Angular dashboard, agent console)
```

Setiap layer terpisah fungsional, komunikasi lewat **Event Bus Protocol**.

---

## LEGALITAS

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

## REPOSITORY

- **GitHub:** https://github.com/SealAngel7/ANGEL
- **Blueprint:** STRUKTUR_ANGEL.md
- **Test Scenarios:** tests/TEST_SCENARIOS.md
- **Dokumentasi:** docs/ directory
- **Usage Guide:** USAGE_GUIDE.md

---

## PERINGATAN PENTING

⚠️ **Hanya gunakan pada sistem yang telah diizinkan.**
⚠️ **Pastikan memiliki kontrak, izin tertulis, dan persetujuan founder.**
⚠️ **Platform dirancang untuk engagement offensive security resmi.**
⚠️ **Semua aktivitas dilacak melalui evidence ledger (CHAIN_CUSTODY).**

---

**ANGEL Platform - Offensive Security Framework untuk Engagement Resmi**