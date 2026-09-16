# 🛡️ ANGEL Platform v3.3

> **STATUS:** FINAL & OPERATIONAL
> **CLASS:** Offensive Security — P0/P1 Hard/Expert Full Attack
> **LEGAL:** Only authorized systems

---

## ██ INITIALIZE

```bash
git clone <REPO_URL> && cd ANGEL
export TEAMSERVER_KEY=<strong-random-key>
export CRYPTO_KEY=<strong-random-key>
export JWT_SECRET=<strong-random-secret>
make build && docker compose up -d
./lab/test_lab_full.sh
```

---

## ██ DEPLOY

```bash
docker compose up -d
docker compose ps
docker compose logs -f angel-teamserver
```

---

## ██ GENERATE IMPLANT

```bash
./bin/angel-generate -os linux -arch amd64 -server http://localhost:8443 -out lab/implants
./bin/angel-generate -os windows -arch amd64 -server http://localhost:8443 -out lab/implants
```

---

## ██ ENGAGE

```bash
echo "192.168.1.1" > target.txt
make engage SCOPE=target.txt
```

---

## ██ MONITOR

```bash
http://localhost:3000
docker compose ps
docker compose logs -f angel-teamserver
```

---

## ██ LAYER MAP

```
layer01-05  ████████████████  16 modul
layer06-10  █████  5 modul
layer11-15  █████  5 modul
layer16-21  ██████  6 modul
layer22-25  ████  4 modul
layer26-40  ████████████  15 modul
layer41-60  ██████████████  18 modul
layer61-70  ██████████  10 modul
```

---

## ██ STATUS

| Build | Test | Layer | Lab | Docker | Secret |
|:-----:|:----:|:-----:|:---:|:------:|:------:|
| ✅ OK | ✅ 92 PASS | ✅ 83/83 | ✅ 25/0 | ✅ 3/3 | ✅ Env |

---

## ██ LEGAL

⚠️ Only authorized systems. Contract + written permission required.
⚠️ For official offensive security engagement only.
⚠️ All activity logged via evidence ledger.

---

**ANGEL Platform — Offensive Security Framework**