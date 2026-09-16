# 🛡️ ANGEL Platform

> **STATUS:** OPERATIONAL Red Team
> **CLASS:** Offensive Security — P0/P1 Hard/Expert Full Attack
> **LEGAL:** Only authorized systems

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

## >> LAYER MAP

```
layer01-05  ==================  16 modul
layer06-10  ======  5 modul
layer11-15  ======  5 modul
layer16-21  =======  6 modul
layer22-25  =====  4 modul
layer26-40  ============  15 modul
layer41-60  ==================  18 modul
layer61-70  ==========  10 modul
```

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

**ANGEL -SOFTWARE ENGGINER CYBER SECURITY**
