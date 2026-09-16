# ANGEL Platform v3.2

**Offensive Security Platform untuk Engagement Resmi**

> **Status:** FINAL & EXECUTABLE
> **Tujuan:** P0/P1, Hard/Expert, Full Attack, No Demo, No Placeholder
> **Legalitas:** Hanya digunakan pada sistem yang telah diizinkan

---

## QUICK START

```bash
# 1. Clone repository
git clone <REPO_URL>
cd ANGEL

# 2. Build binary
make build

# 3. Jalankan lab (Docker)
make up

# 4. Test end-to-end
make test

# 5. Mulai engagement
make engage SCOPE=target.txt
```

## LAYANAN

| Service | Port | Keterangan |
|---------|------|------------|
| Teamserver | 8443 | C2 server (HTTP/HTTPS) |
| Console | 3000 | Agent management UI |
| Rules | 9444 | Rules engine |
| DVWA | 8081 | Target latihan |

---

## FITUR UTAMA

| kategori | detail |
|----------|---------|
| **70 Layer** | C2 Framework, Evasion, AD Attack, Persistence, Rootkit |
| **427 file Go** | Terstruktur per layer 1-70 |
| **92 test automated** | Semua passing (1,346 manual scenarios di TEST_SCENARIOS.md) |
| **3 binary** | `angel`, `angel-console`, `angel-rules` |
| **Event Bus** | Semua komunikasi lewat central event bus |
| **700+ teknik** | Fallback chains, anti-analysis, opsec procedures |

---

## CARA PAKAI

### Setup Lengkap
```bash
make setup     # Install deps + buat .env.local
make build     # Build 3 binary
make test      # Jalankan 1,346 test case
make engage SCOPE=target.txt  # Start engagement
```

### Manajemen C2
```bash
docker compose up -d angel-teamserver   # Start teamserver
docker compose up -d angel-console      # Start console
docker compose logs angel-teamserver    # Lihat logs
```

### Generate Implant
```bash
./bin/angel-generate -os linux -arch amd64 -server http://teamserver:8443 -out lab/implants
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
- Harus memiliki kontrak, izin polisi, dan persetujuan founder
- Hanya digunakan untuk engagement resmi offensive security

### Prinsip Dasar
1. **"No copy-paste"** — tiap baris ditulis sendiri
2. **"Signature-free"** — defender gak kenal
3. **"Modular"** — tiap modul jalan sendiri, tapi orchestrated
4. **"Evidentiary"** — tiap action ada bukti
5. **"Clean"** — post-engagement, semua hilang

---

## GITHUB & RESOURCES

- **Repository:** https://github.com/angel-framework/angel
- **Blueprint:** STRUKTUR_ANGEL.md v3.2
- **Test Scenarios:** tests/TEST_SCENARIOS.md (1,346 test)
- **Dokumentasi:** docs/ directory

---

## PERINGATAN PENTING

⚠️ **Hanya gunakan pada sistem yang telah diizinkan.**  
⚠️ **Pastikan memiliki kontrak, izin tertulis, dan persetujuan founder.**  
⚠️ **Platform dirancang untuk engagement offensive security resmi.**  
⚠️ **Semua aktivitas dilacak melalui evidence ledger (CHAIN_CUSTODY).**

---

**ANGEL Platform - Offensive Security Framework untuk Engagement Resmi**