# ANGEL — Read-Only Offensive Tooling Assessment

**Purpose:** Verify whether ANGEL's *read-only* tooling can be used "as-is" for live
bug-bounty reconnaissance on open programs (Grab @ HackerOne, GoTo/GoTo Financial @ YesWeHack).
**Scope of this assessment:** Passive + read-only active recon only.
No data modification, no authentication, no exploitation performed.

---

## 1. Executive Summary

- **ANGEL read-only tooling IS production-grade usable** for the recon phase.
  - The OSINT module (`modules/layer16-21/osint/`) compiles cleanly, passes its own
    unit-test suite (30+ tests, all PASS), and was exercised live against `grab.com`
    (a real HackerOne Bug-Bounty target).
- **Live read-only test on grab.com surfaced LOW-severity issues only**
  (missing HSTS, weak cookie flags, tech/path disclosure). No Medium/High/Critical
  was found because Grab's public surface is mature/hardened.
- For **Medium / High / Critical**, read-only recon is only the *first phase*.
  Those severities require an exploitation PoC (e.g. SQLi/XSS/IDOR/SSRF with
  demonstrated impact), which is **out of scope for "pure read-only"** and also
  requires an authorized BBP account + confirmed in-scope assets.
- MCP-based tools (nmap/ffuf/subfinder/shodan) and the 20+ OSINT agents are still
  **simulated/stub** — they are NOT wired to real binaries yet (see §4).

**Conclusion:** "Read-only to find gaps" → **yes, it works and is safe** (pure GET/HEAD/DNS/TLS).
"Find Medium/High/Critical via read-only only" → **no** — those need an exploit PoC step.

---

## 2. Verified read-only tools (REAL & usable now)

| Tool | File | Read-only capability | Live-verified? |
|---|---|---|---|
| WebRecon | modules/layer16-21/osint/web.go | Tech fingerprint, WAF detect, SSL inspect, robots.txt parse, GET/HEAD | YES (grab.com) |
| PortScanner | modules/layer16-21/osint/port.go | TCP connect, banner grab, service fingerprint, IsPortOpen | Unit-tested (not run broad against live — IDS risk) |
| PersonRecon | modules/layer16-21/osint/person.go | Email harvest, social-media/git recon, WHOIS | Unit-tested |
| DNSRecon (engine.go) | modules/layer16-21/osint/engine.go | Reverse DNS, resolve, wildcard | Unit-tested (live: 542 subs via CT log) |

Build result: `go build ./modules/layer16-21/osint/...` → exit 0 (clean)
Test result: `go test ./modules/layer16-21/osint/...` → **PASS** (ok, 0.042s)

---

## 3. Live read-only results — grab.com (HackerOne open program)

### Finding 1 — Missing HSTS  [LOW-MEDIUM]
- `Strict-Transport-Security` header **absent** on `https://grab.com`.
- Proof: `curl -sI https://grab.com | grep -i strict-transport` → empty.
- Other defensive headers present: X-Frame-Options, X-Content-Type-Options, CSP, X-XSS-Protection.

### Finding 2 — Weak cookie flags (AWSALB)  [LOW]
- `Set-Cookie: AWSALB=...; Path=/` — **no Secure / no HttpOnly / no SameSite**.
- (This is an AWS Load-Balancer cookie, not an app session cookie.)

### Finding 3 — Tech & path disclosure  [INFORMATIONAL]
- `Server: nginx`, fronted by `Via: ... CloudFront`, `Set-Cookie: AWSALB...`
  → reveals AWS CloudFront + ALB infrastructure.
- `robots.txt` exposes internal paths (`/wp-admin/`, `/sg/grab_portal/`, `/sg/profile/`,
  `/sg/fun/`, `/sg/funpages/`, `/sg/game_dir/`, `/grabpaymerchant/`).
- Note: Grab policy explicitly lists *"Disclosure of known public files or directories
  (e.g. robots.txt)"* as **non-qualifying** for bounty.

### Finding 4 — CORS  [CLEAN]
- `Access-Control-Allow-Origin` echoes `https://www.grab.com` only (no reflection of
  attacker Origin, no `*`). **Not vulnerable** to CORS misconfig in sample.

### Finding 5 — Subdomain enumeration (CT log)  [CLEAN]
- 542 subdomains discovered via public CT log (`crt.sh`).
- Sample probe (DNS resolve + takeover-signature check): no subdomain-takeover
  signatures (no `NoSuchBucket`, no `Repository not found`, no Fastly/Heroku error page).

**Net: Grab is hardened. Read-only recon yields only Low/Informational here.**

---

## 4. Tools that are STILL SIMULATED / STUB (must not present as real)

| Tool | Location | Status |
|---|---|---|
| MCP tools: nmap / ffuf / playwright / subfinder / shodan (etc.) | orchestrator/mcp/server.go | `Execute == nil` → returns `"simulated"` |
| `osint/agents/*` (20+ agents) | modules/layer16-21/osint/agents/ | stub — `Run()` returns `"osint:done"` |
| ModuleTopology | modules/layer16-21/osint/moduletopo.go | stub — returns single string |

These are the highest-leverage gap: if MCP server tools were wired to real
binaries, the same read-only approach could scale across many targets.

---

## 5. Bug-bounty program reference (verified)

| Company | Platform | Status | Notes (read-only) |
|---|---|---|---|
| Grab | HackerOne | OPEN | recon allowed; robots.txt disclosure = non-qualifying; needs PoC + impact |
| GoTo Financial (Gojek/Tokopedia) | YesWeHack | OPEN | "Open ports/services without PoC" = non-qualifying; DoS/social-engineering = out-of-scope |
| Gojek | HackerOne | **NOT accepting submissions** | Do not test against active submission |

---

## 6. Recommendation / What to do next

1. **Use the OSINT module as-is for recon.** It is verified real. Feed findings into
   the orchestrator intent-classifier (`IntentRecon`, risk score 10, auto-execute)
   which already separates `recon` from `exploit` (risk 80) / `destruction` (risk 95).
   This gives you a clean **read-only phase** with no danger to live targets.
2. **Do NOT claim Medium/High/Critical from read-only alone.** Those require an
   exploitation PoC step — run that only after: (a) you hold an authorized BBP account,
   (b) the asset is confirmed in-scope, (c) you test only your own accounts / staging
   you control.
3. **To scale recon** (more subdomains, templated misconfig scan), wire the MCP server
   to real binaries (nuclei + subfinder). `nuclei` is already installed in this
   environment at `$(go env GOPATH)/bin/nuclei`; subfinder too. The MCP server needs
   `Tool.Execute` implementations to call them.
4. **To find fresh Medium/High/Critical:** pick a smaller/fresher BBP program
   (younger SaaS product), run full recon (subdomains via CT logs / subfinder),
   then attempt *targeted* PoCs only on confirmed in-scope assets you can replicate
   locally first.
5. **Report submission is on you** (you need the BBP account). This repo provides the
   *tools* and *draft evidence*; it does not submit reports.

---

## 7. Files produced by this assessment
- `audit/grab-readonly-recon-report.md` — standalone draft report for grab.com (H1).
- `audit/ANGEL-READONLY-ASSESSMENT.md` — this file (tools status + verdict).
