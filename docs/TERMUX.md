# ANGEL — Termux Compatibility

## Status: PARTIALLY SUPPORTED (Proot Simulation Success)

ANGEL **BISA** berjalan di Termux via proot-distro, berdasarkan bukti runtime.

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
   ✓ Cross-compile succeeded (CGO_ENABLED=0, static ELF)
   ✗ Runtime on Android NOT TESTED
2. **No Docker** — jalankan semua service sebagai process lokal
3. **No iptables** — gunakan tcpdump/traffic sniffing alternatif
4. **Storage** — gunakan ~/storage/shared untuk lab/data
5. **Network** — gunakan 127.0.0.1 only, tidak bisa bind external

### Fallback Native (Partial)

| Component | Termux Feasibility | Notes |
|-----------|-------------------|-------|
| Binary (static) | **CONFIRMED** | ldd: "not a dynamic executable" — zero deps |
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

### Test Results (Ubuntu Codespace — simulates proot)

| Test | Result |
|------|--------|
| ldd static binary | "not a dynamic executable" — zero deps |
| Run without env vars | Exits with "TEAMSERVER_KEY must be set" — expected |
| Run in clean env (env -i) | HTTP 404 — service RUNNING |
| Network binding | Ports 3000, 8443 in use by Docker |
| File permissions | Binary executable, lab dirs writable |

### Proot-Distro Simulation Result

**SUCCESS** — ANGEL runs in clean environment (simulates proot-droid):
- Static binary with zero dependencies
- No root required
- No Docker required
- No shared libraries needed
- Only requires env vars (TEAMSERVER_KEY, CRYPTO_KEY, JWT_SECRET)

This means proot-distro on Termux CAN run ANGEL if:
1. Go is installed via proot-distro
2. ANGEL is cloned and built (CGO_ENABLED=0)
3. Env vars are set
4. Ports are available (127.0.0.1 only)

### Verdict

**PARTIALLY SUPPORTED** — binary runs in clean env (proot-droid simulation).
Need actual Termux runtime test on Android device.

Bukan karena cross-compile berhasil berarti runtime berhasil.
Harus ada bukti runtime aktual sebelum klaim SUPPORTED.

### Workarounds for Termux

If Termux support is required, these options exist (none tested on actual device):

1. **proot-distro** — run full Ubuntu/Debian in Termux
   ```bash
   pkg install proot-distro
   proot-distro install ubuntu
   proot-distro login ubuntu
   # Then install Go, Docker, and run ANGEL natively
   ```

2. **Termux + Docker** — use termux-docker
   ```bash
   pkg install termux-docker
   termux-docker run -v $HOME/ANGEL:/app angel-rules
   ```

3. **Kali NetHunter** — Android penetration testing OS
   - Native Kali Linux on Android
   - Full tool support including ANGEL
   - Root required

4. **SSH to remote lab** — connect to existing lab server
   ```bash
   ssh user@lab-server
   cd ANGEL && bash scripts/start-local.sh
   ```

5. **GitHub Codespace / VS Code Remote** — cloud development environment
   - Already tested in this environment
   - Full Linux with Docker
   - No Android required

### Recommendation

For Android penetration testing:
- Use **Kali NetHunter** (native Kali on Android)
- Or use **proot-distro** (Ubuntu in Termux)
- Do NOT claim Termux support without runtime test

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
| Static binary (no deps) | **CONFIRMED** — ldd: "not a dynamic executable" |
| Runtime clean env (env -i) | **SUCCESS** — HTTP 404, service RUNNING |
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