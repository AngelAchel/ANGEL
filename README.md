# 🛡️ ANGEL TOOLKIT

> **STATUS:** OPERATIONAL Red Team
> 
> **CLASS:** Offensive Security — P0/P1 Hard/Expert Full Attack
> 
> **LEGAL:** Only authorized systems
> 
> **Built:** Cyber Security Software Engineer
---

## ⚔️ ATTACK

![Attack Visualization](attack_viz.gif)

---

## >> INITIALIZE

```bash
git clone <REPO_URL> && cd ANGEL
export TEAMSERVER_KEY=<strong-random-key>
export CRYPTO_KEY=<strong-random-key>
export JWT_SECRET=<strong-random-secret>
make build && docker compose up -d
./lab/test_lab_full.sh
```

---

## >> DEPLOY

```bash
docker compose up -d
docker compose ps
docker compose logs -f angel-teamserver
```

---

## >> GENERATE IMPLANT

```bash
./bin/angel-generate -os linux -arch amd64 -server http://localhost:8443 -out lab/implants
./bin/angel-generate -os windows -arch amd64 -server http://localhost:8443 -out lab/implants
```

---

## >> ENGAGE

```bash
echo "192.168.1.1" > target.txt
make engage SCOPE=target.txt
```

---

## >> MONITOR

```bash
http://localhost:3000
docker compose ps
docker compose logs -f angel-teamserver
```

---

## >> DAMPAK PERUSAHAAN

| Skenario | Dampak |
|----------|--------|
| 💀 Credential Dump | Semua password bocor, root access |
| 🔥 Data Exfiltration | Data permanen keluar, tak terlacak |
| 🧠 Ransomware | Semua file terenkripsi, backup hancur |
| 🕵️ Espionage | Intel kompetitor dicuri |
| 🔌 Persistence | Backdoor permanen, tidak hilang |
| 🧹 Anti-Forensik | Bukti dihapus, tidak ada jejak |
| 💀 Privilege Escalation | User biasa jadi root |
| 🌐 Lateral Movement | Semua server terinfeksi |
| 🔥 Destruction | Database, log, backup — semua hilang |
| 💀 Complete Takeover | Full kontrol infrastruktur |


---

## >> IMPACT

| Capability | Detail |
|------------|--------|
| 🎯 70 Layer Attack Surface | Full offensive stack |
| ⚔️ 700+ Techniques | Fallback chains, anti-analysis |
| 🧬 83 Modules | C2, Evasion, Credential, Exploit, Malware |
| 🔌 Event Bus | Real-time cross-layer orchestration |
| 🛡️ 92 Tests | All passing — zero regressions |
| 💀 Signature-Free | Defender gak kenal |
| 🔥 Privilege Escalation | Auto escalate |
| 🕵️ OSINT + Recon | Passive & active intel |
| 💀 Exfiltration | Data out, no trace |
| 🧹 Cleanup | Leave zero evidence |


---

## >> LEGAL

⚠️ Only authorized systems. Contract + written permission required.
⚠️ For official offensive security engagement only.
⚠️ All activity logged via evidence ledger.

---

**ANGEL Platform — Offensive Security Framework**
GitHub: https://github.com/AngelAchel/ANGEL
