# Deployment Instructions

## Push to GitHub

### 1. Create Repository on GitHub
1. Go to https://github.com/new
2. Repository name: `osovm`
3. Owner: `ase-lang` (or create organization first)
4. Description: "Ọ̀ṢỌ́ Virtual Machine - Proof-of-Ritual Execution Engine"
5. Public repository
6. Do NOT initialize with README (already exists)

### 2. Add Remote and Push

```bash
cd ~/osovm
git remote add origin https://github.com/ase-lang/osovm.git
git branch -M main
git push -u origin main
```

### 3. Repository URL

Once pushed: **https://github.com/ase-lang/osovm**

---

## Test OSOVM Locally

### Build

```bash
cd vm
go build -o oso osovm.go
```

### Run Example

```bash
./oso run ../examples/drone_delivery.oso
```

---

## Phase 1 Deliverables Checklist

- [x] Pest grammar: `grammar/ORISA.pest`
- [x] Go VM runtime: `vm/osovm.go`
- [x] Example rituals:
  - [x] `examples/drone_delivery.oso`
  - [x] `examples/factory_handoff.oso`
  - [x] `examples/qr_checkpoint.oso`
- [x] README with Invocation → Proof → Àṣẹ flow
- [x] Git repo initialized
- [ ] Push to GitHub (manual step)

---

**#MirrorWitness OSOVM_P1 2025-11-05**

Àṣẹ! 🔥
