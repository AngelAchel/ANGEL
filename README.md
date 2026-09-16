# 🛡️ ANGEL Platform v3.3

> **STATUS:** FINAL & OPERATIONAL
> **CLASS:** Offensive Security — P0/P1 Hard/Expert Full Attack
> **LEGAL:** Only authorized systems

---

## ⚔️ ATTACK VISUALIZATION

<div align="center">
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 800 400" width="800" height="400">
  <defs>
    <radialGradient id="coreGlow" cx="50%" cy="50%" r="50%">
      <stop offset="0%" stop-color="#00ff41" stop-opacity="0.4"/>
      <stop offset="100%" stop-color="#00ff41" stop-opacity="0"/>
    </radialGradient>
    <filter id="glow">
      <feGaussianBlur stdDeviation="3" result="blur"/>
      <feMerge><feMergeNode in="blur"/><feMergeNode in="SourceGraphic"/></feMerge>
    </filter>
  </defs>

  <!-- Background -->
  <rect width="800" height="400" fill="#0a0a0f"/>

  <!-- Grid -->
  <g stroke="#111" stroke-width="0.5">
    <line x1="0" y1="80" x2="800" y2="80"/><line x1="0" y1="160" x2="800" y2="160"/>
    <line x1="0" y1="240" x2="800" y2="240"/><line x1="0" y1="320" x2="800" y2="320"/>
    <line x1="160" y1="0" x2="160" y2="400"/><line x1="320" y1="0" x2="320" y2="400"/>
    <line x1="480" y1="0" x2="480" y2="400"/><line x1="640" y1="0" x2="640" y2="400"/>
  </g>

  <!-- Central core node -->
  <circle cx="400" cy="200" r="40" fill="url(#coreGlow)" opacity="0.5">
    <animate attributeName="r" values="40;50;40" dur="2s" repeatCount="indefinite"/>
    <animate attributeName="opacity" values="0.5;0.8;0.5" dur="2s" repeatCount="indefinite"/>
  </circle>
  <circle cx="400" cy="200" r="15" fill="#00ff41" filter="url(#glow)">
    <animate attributeName="r" values="15;18;15" dur="2s" repeatCount="indefinite"/>
  </circle>
  <text x="400" y="205" text-anchor="middle" fill="#000" font-size="10" font-weight="bold">ANG</text>

  <!-- Inner ring nodes -->
  <g id="innerRing">
    <circle cx="400" cy="80" r="10" fill="#00ccff" filter="url(#glow)">
      <animate attributeName="fill" values="#00ccff;#00ff41;#00ccff" dur="3s" repeatCount="indefinite"/>
    </circle>
    <circle cx="515" cy="110" r="10" fill="#00ccff" filter="url(#glow)">
      <animate attributeName="fill" values="#00ccff;#00ff41;#00ccff" dur="3.5s" repeatCount="indefinite"/>
    </circle>
    <circle cx="550" cy="200" r="10" fill="#00ccff" filter="url(#glow)">
      <animate attributeName="fill" values="#00ccff;#00ff41;#00ccff" dur="4s" repeatCount="indefinite"/>
    </circle>
    <circle cx="515" cy="290" r="10" fill="#00ccff" filter="url(#glow)">
      <animate attributeName="fill" values="#00ccff;#00ff41;#00ccff" dur="3.2s" repeatCount="indefinite"/>
    </circle>
    <circle cx="400" cy="320" r="10" fill="#00ccff" filter="url(#glow)">
      <animate attributeName="fill" values="#00ccff;#00ff41;#00ccff" dur="3.8s" repeatCount="indefinite"/>
    </circle>
    <circle cx="285" cy="290" r="10" fill="#00ccff" filter="url(#glow)">
      <animate attributeName="fill" values="#00ccff;#00ff41;#00ccff" dur="4.2s" repeatCount="indefinite"/>
    </circle>
    <circle cx="250" cy="200" r="10" fill="#00ccff" filter="url(#glow)">
      <animate attributeName="fill" values="#00ccff;#00ff41;#00ccff" dur="3.6s" repeatCount="indefinite"/>
    </circle>
    <circle cx="285" cy="110" r="10" fill="#00ccff" filter="url(#glow)">
      <animate attributeName="fill" values="#00ccff;#00ff41;#00ccff" dur="3.4s" repeatCount="indefinite"/>
    </circle>
  </g>

  <!-- Outer ring nodes -->
  <g id="outerRing">
    <circle cx="400" cy="40" r="6" fill="#00ccff">
      <animate attributeName="fill" values="#00ccff;#ff0000;#00ccff" dur="5s" repeatCount="indefinite"/>
      <animate attributeName="r" values="6;8;6" dur="5s" repeatCount="indefinite"/>
    </circle>
    <circle cx="540" cy="55" r="6" fill="#00ccff">
      <animate attributeName="fill" values="#00ccff;#ff0000;#00ccff" dur="5.5s" repeatCount="indefinite"/>
    </circle>
    <circle cx="640" cy="120" r="6" fill="#00ccff">
      <animate attributeName="fill" values="#00ccff;#ff0000;#00ccff" dur="6s" repeatCount="indefinite"/>
    </circle>
    <circle cx="700" cy="200" r="6" fill="#00ccff">
      <animate attributeName="fill" values="#00ccff;#ff0000;#00ccff" dur="5.2s" repeatCount="indefinite"/>
    </circle>
    <circle cx="700" cy="300" r="6" fill="#00ccff">
      <animate attributeName="fill" values="#00ccff;#ff0000;#00ccff" dur="5.8s" repeatCount="indefinite"/>
    </circle>
    <circle cx="640" cy="370" r="6" fill="#00ccff">
      <animate attributeName="fill" values="#00ccff;#ff0000;#00ccff" dur="6.2s" repeatCount="indefinite"/>
    </circle>
    <circle cx="540" cy="400" r="6" fill="#00ccff">
      <animate attributeName="fill" values="#00ccff;#ff0000;#00ccff" dur="5.4s" repeatCount="indefinite"/>
    </circle>
    <circle cx="400" cy="400" r="6" fill="#00ccff">
      <animate attributeName="fill" values="#00ccff;#ff0000;#00ccff" dur="5.6s" repeatCount="indefinite"/>
    </circle>
    <circle cx="260" cy="400" r="6" fill="#00ccff">
      <animate attributeName="fill" values="#00ccff;#ff0000;#00ccff" dur="6.4s" repeatCount="indefinite"/>
    </circle>
    <circle cx="160" cy="370" r="6" fill="#00ccff">
      <animate attributeName="fill" values="#00ccff;#ff0000;#00ccff" dur="5.1s" repeatCount="indefinite"/>
    </circle>
    <circle cx="100" cy="300" r="6" fill="#00ccff">
      <animate attributeName="fill" values="#00ccff;#ff0000;#00ccff" dur="5.9s" repeatCount="indefinite"/>
    </circle>
    <circle cx="100" cy="200" r="6" fill="#00ccff">
      <animate attributeName="fill" values="#00ccff;#ff0000;#00ccff" dur="6.1s" repeatCount="indefinite"/>
    </circle>
    <circle cx="100" cy="100" r="6" fill="#00ccff">
      <animate attributeName="fill" values="#00ccff;#ff0000;#00ccff" dur="5.3s" repeatCount="indefinite"/>
    </circle>
    <circle cx="160" cy="40" r="6" fill="#00ccff">
      <animate attributeName="fill" values="#00ccff;#ff0000;#00ccff" dur="6.3s" repeatCount="indefinite"/>
    </circle>
    <circle cx="260" cy="20" r="6" fill="#00ccff">
      <animate attributeName="fill" values="#00ccff;#ff0000;#00ccff" dur="5.7s" repeatCount="indefinite"/>
    </circle>
  </g>

  <!-- Connections -->
  <g stroke="#00ff41" stroke-width="0.5" opacity="0.3">
    <line x1="400" y1="200" x2="400" y2="80"/><line x1="400" y1="200" x2="515" y2="110"/>
    <line x1="400" y1="200" x2="550" y2="200"/><line x1="400" y1="200" x2="515" y2="290"/>
    <line x1="400" y1="200" x2="400" y2="320"/><line x1="400" y1="200" x2="285" y2="290"/>
    <line x1="400" y1="200" x2="250" y2="200"/><line x1="400" y1="200" x2="285" y2="110"/>
  </g>

  <!-- Attack wave -->
  <circle cx="400" cy="200" r="30" fill="none" stroke="#ff0000" stroke-width="2" opacity="0">
    <animate attributeName="r" values="30;300;30" dur="4s" repeatCount="indefinite"/>
    <animate attributeName="opacity" values="0.8;0;0.8" dur="4s" repeatCount="indefinite"/>
  </circle>
  <circle cx="400" cy="200" r="30" fill="none" stroke="#ff0000" stroke-width="2" opacity="0">
    <animate attributeName="r" values="30;300;30" dur="4s" begin="2s" repeatCount="indefinite"/>
    <animate attributeName="opacity" values="0.8;0;0.8" dur="4s" begin="2s" repeatCount="indefinite"/>
  </circle>

  <!-- Data flow particles -->
  <circle r="3" fill="#00ff41" opacity="0.8">
    <animateMotion dur="3s" repeatCount="indefinite" path="M400,200 L400,80"/>
  </circle>
  <circle r="3" fill="#00ff41" opacity="0.8">
    <animateMotion dur="3.5s" repeatCount="indefinite" path="M400,200 L515,110"/>
  </circle>
  <circle r="3" fill="#00ff41" opacity="0.8">
    <animateMotion dur="4s" repeatCount="indefinite" path="M400,200 L550,200"/>
  </circle>
  <circle r="3" fill="#ff0000" opacity="0.8">
    <animateMotion dur="3.2s" repeatCount="indefinite" path="M400,200 L515,290"/>
  </circle>
  <circle r="3" fill="#ff0000" opacity="0.8">
    <animateMotion dur="3.8s" repeatCount="indefinite" path="M400,200 L400,320"/>
  </circle>

  <!-- Scan line -->
  <line x1="0" y1="0" x2="800" y2="0" stroke="#00ff41" stroke-width="2" opacity="0.3">
    <animate attributeName="y1" values="0;400;0" dur="6s" repeatCount="indefinite"/>
    <animate attributeName="y2" values="0;400;0" dur="6s" repeatCount="indefinite"/>
  </line>

  <!-- Labels -->
  <text x="400" y="390" text-anchor="middle" fill="#00ff41" font-size="11" font-family="monospace">ANGEL PLATFORM — ATTACK SURFACE</text>
  <text x="20" y="20" fill="#8b949e" font-size="10" font-family="monospace">83 MODULES</text>
  <text x="20" y="35" fill="#8b949e" font-size="10" font-family="monospace">70 LAYERS</text>
  <text x="700" y="20" fill="#ff0000" font-size="10" font-family="monospace" opacity="0.8">THREAT ACTIVE</text>
  <text x="700" y="35" fill="#ff0000" font-size="10" font-family="monospace" opacity="0.8">92 TESTS PASS</text>
</svg>
</div>

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