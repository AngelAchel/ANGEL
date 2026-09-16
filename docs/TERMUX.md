# ANGEL — Termux Compatibility

## Status: NOT SUPPORTED (Belum Diuji)

ANGEL **TIDAK didukung** di Termux saat ini. Berikut analisis kompatibilitas:

### Dependency Check

| Dependency | Termux Status | ANGEL Requirement | Verdict |
|------------|---------------|-------------------|---------|
| Go 1.22+ | NOT AVAILABLE | Required | BLOCKER |
| Docker | NOT AVAILABLE | Required (services) | BLOCKER |
| Docker Compose | NOT AVAILABLE | Required | BLOCKER |
| Python 3 | AVAILABLE (pkg install python) | Optional | OK |
| Nmap | NOT AVAILABLE | Optional (layer 06-10) | WARNING |
| Netcat | NOT AVAILABLE | Optional | WARNING |
| curl | AVAILABLE | Required (health check) | OK |
| git | AVAILABLE | Required | OK |
| make | NOT AVAILABLE | Required (build) | BLOCKER |
| gcc | NOT AVAILABLE | Required (CGO) | BLOCKER |
| iptables | NOT AVAILABLE | Required (network) | BLOCKER |
| Storage | LIMITED | Required (lab/data) | WARNING |
| Network | LIMITED | Required (bind 0.0.0.0) | WARNING |

### Root Cause NOT SUPPORTED

1. **Go tidak tersedia** — Termux tidak punya Go package resmi
2. **Docker tidak berjalan** — Termux tidak dukung Docker daemon
3. **Make tidak tersedia** — Tidak bisa build via Makefile
4. **gcc tidak tersedia** — Tidak bisa compile CGO
5. **iptables tidak tersedia** — Tidak bisa network manipulation
6. **Storage terbatas** — Termux storage terbatas ~/storage
7. **Network terbatas** — Termux bound ke localhost, tidak bisa bind 0.0.0.0

### What Would Be Needed for Termux Support

1. **Go binary** — cross-compile static binary untuk Android arm64
2. **No Docker** — jalankan semua service sebagai process lokal
3. **No iptables** — gunakan tcpdump/traffic sniffing alternatif
4. **Storage** — gunakan ~/storage/shared untuk lab/data
5. **Network** — gunakan 127.0.0.1 only, tidak bisa bind external

### Fallback Native (Partial)

| Component | Termux Feasibility | Notes |
|-----------|-------------------|-------|
| Binary (static) | POSSIBLE | CGO_ENABLED=0 cross-compile |
| Event bus | POSSIBLE | in-process, no Docker needed |
| Orchestrator | POSSIBLE | in-process, no Docker needed |
| C2 server | POSSIBLE | bind 127.0.0.1 only |
| Console | POSSIBLE | bind 127.0.0.1 only |
| Rules loader | POSSIBLE | local rules only |
| DVWA fixture | NOT POSSIBLE | requires Docker |
| Teamserver | POSSIBLE | bind 127.0.0.1 only |
| Network scanning | NOT POSSIBLE | no nmap, no root |
| Packet injection | NOT POSSIBLE | no root, no iptables |
| Docker services | NOT POSSIBLE | no Docker daemon |

### Verdict

**NOT SUPPORTED** — belum ada bukti runtime di Termux.

Bukan karena cross-compile berhasil berarti runtime berhasil.
Harus ada bukti runtime aktual sebelum klaim SUPPORTED.

### Testing Required

Untuk mendukung Termux, diperlukan:
1. Install Go di Termux (manual compile atau termux-packages)
2. Cross-compile static binary untuk arm64 Android
3. Test runtime di perangkat Android aktual
4. Test network binding di Android
5. Test storage access di Android
6. Test signal handling di Android
7. Test process lifecycle di Android

### Status Tiap Komponen

| Component | Status |
|-----------|--------|
| Build (amd64) | SUPPORTED |
| Build (arm64 cross-compile) | SUPPORTED (not tested runtime) |
| Runtime Termux | NOT TESTED |
| Docker | NOT SUPPORTED |
| Network binding | NOT TESTED |
| Storage | NOT TESTED |
| Signal handling | NOT TESTED |
| Process lifecycle | NOT TESTED |
| Event bus | NOT TESTED |
| C2 server | NOT TESTED |
| Console | NOT TESTED |
| Rules loader | NOT TESTED |
| Orchestrator | NOT TESTED |
| DVWA fixture | NOT SUPPORTED |
| Network scanning | NOT SUPPORTED |
| Packet injection | NOT SUPPORTED |

### Peringatan

JANGAN klaim Termux support sebelum:
- [ ] Go terinstal di Termux
- [ ] Binary berhasil di-cross-compile untuk arm64 Android
- [ ] Binary berhasil di-run di perangkat Android aktual
- [ ] Service berhasil binding ke port
- [ ] Event bus berhasil berjalan
- [ ] Storage berhasil akses
- [ ] Signal handling berhasil (Ctrl+C, SIGTERM)

### File Terkait
- `docs/KALI.md` — Kali Linux compatibility
- `docs/NATIVE-RUN.md` — Native run guide
- `docs/DEPENDENCIES.md` — Full dependency list
- `audit/deployment-checklist.md` — Deployment checklist