# Ọ̀ṢỌ́VM - Full Canon (143 Attributes)

**Universal Device VM**: Drones, Phones, AV, Robots, Sensors, Cameras  
**Repository**: https://github.com/ase-lang/osovm  
**License**: MIT

## What is Ọ̀ṢỌ́VM?

Ọ̀ṢỌ́VM is a **Proof-of-Ritual Virtual Machine** that enables:

1. **SimaaS** (Simulation-as-a-Service): Test non-human entities before deployment
2. **AIO** (Universal Job Marketplace): Digital, meta-digital, and physical jobs
3. **TechGnØŞ.EXE**: Spiritual tech framework with sacred 50/25/15/10 splits
4. **Òrìṣà ROBOTS**: Autonomous entities funded 100% by shrine offerings

## Quick Start

### Build & Run

```bash
# One-command upgrade
./phase2.sh

# Or manually
cd cmd/phase2
go build -o ../../phase2 main.go
cd ../..

# Run examples
./phase2 run examples/qr_delivery.oso
./phase2 run examples/simaas_robot_test.oso
./phase2 run examples/aio_physical_job.oso
./phase2 run examples/techgnosis_offering.oso
./phase2 run examples/osanyin_herb_prescription.oso
```

## The 143 Attributes

### Core Categories

1. **Proof & Witness** (1-10): `@proof`, `@witness`, `@receipt`, `@àṣẹ`, `@telemetry`
2. **SimaaS & AIO** (11-20): `@sim`, `@job`, `@hardware`, `@collaboration`, `@offering`
3. **Herb & Regional** (21-30): `@herb`, `@regional` (Òsanyìn registry)
4. **Governance** (31-40): `@vote`, `@dispute`, `@slash`, `@wallet`, `@identity`
5. **Native VM** (41-50): `@memory`, `@learn`, `@compute`, `@verifier`, `@npc`
6. **Entertainment/Jobs** (51-70): `@project`, `@milestone`, `@casting`, `@delivery`
7. **Flow/Workflow** (71-90): `@flow`, `@retry`, `@saga`, `@ai`, `@oracle`
8. **Interop** (91-100): `@import`, `@abi`, `@bridge` (BTC/ETH/SOL/AR/Sui)
9. **Astral-Telluric** (101-110): `@beats`, `@astro`, `@ley`, `@nexus`, `@geofence`
10. **SatKey/Ordinals** (111-120): `@ordinal_guard`, `@inscription`, `@satname`
11. **ZK & Numerology** (121-130): `@zk`, `@zkbind`, `@veil` (φ, π constants)
12. **Economic/Treasury** (131-143): `@tithe`, `@shrineSplit`, `@toc`, `@fixed_math`

See [docs/ATTRIBUTES_CANON.md](docs/ATTRIBUTES_CANON.md) for complete reference.

## Sacred Mathematics

- **Tithe**: Always `369` (3.69% ecosystem upkeep)
- **Vote Weight**: Always `1.0` (one entity, one vote)
- **Shrine Split**: `50/25/15/10` (TechGnØŞ offerings only)
- **Veils**: Numerology constants (φ = 1.618..., π = 3.14159...)

## Architecture

```
osovm/
├── grammar/
│   ├── ORISA.pest           # Phase 1 grammar
│   ├── ORISA_PHASE2.pest    # Phase 2 grammar
│   └── ORISA_FULL.pest      # Full 143 attributes ✨
├── vm/
│   ├── osovm.go             # Phase 1 VM
│   └── attributes.go        # Full attribute system
├── pkg/
│   ├── camera/qr.go         # QR scanner API
│   └── witness/node.go      # Witness mesh (LoRa/BLE)
├── cmd/phase2/main.go       # Phase 2 entry point
├── examples/
│   ├── qr_delivery.oso              # QR + witness mesh
│   ├── simaas_robot_test.oso        # SimaaS simulation
│   ├── aio_physical_job.oso         # AIO job posting
│   ├── techgnosis_offering.oso      # Shrine offering
│   └── osanyin_herb_prescription.oso # Herbal prescription
└── docs/
    ├── ARCHITECTURE.md
    └── ATTRIBUTES_CANON.md   # Full attribute reference
```

## Example Rituals

### 1. QR Delivery (Drone + Witnesses)

```json
{
  "name": "qr_delivery",
  "orisa": "eshu_router",
  "ase": {
    "proof": "qr",
    "witnesses": 3,
    "qr": true,
    "auto_device": true
  },
  "args": {
    "action": "delivery",
    "package_id": "PKG_12345"
  }
}
```

**Output:**
```
📷 QR scanned: 9f2ae8b1c4d7f3a6...
📡 Broadcast → 3 nodes echo → seal
✅ Àṣẹ sealed with qr proof + 3 witnesses
```

### 2. SimaaS Robot Test

```json
{
  "name": "simaas_robot_test",
  "orisa": "ogun_forge",
  "attributes": {
    "sim": { "episode_id": 1, "steps": 1000 },
    "hardware": { "estop_check": true },
    "compute": { "gpu": "h100", "hours": 2, "provider": "runpod" }
  },
  "args": {
    "robot": "unitree_h1",
    "task": "navigation_test"
  }
}
```

### 3. TechGnØŞ Offering

```json
{
  "name": "techgnosis_offering",
  "orisa": "obatala_guard",
  "attributes": {
    "offering": { "amount": 369, "shrine": "Obatala_Nexus_Lagos" },
    "shrineSplit": { "distribution": "50/25/15/10" },
    "tithe": { "rate": 369, "wallet": "tech_gnosis_treasury" }
  },
  "args": {
    "blessing": "clarity_and_governance"
  }
}
```

**Funds Òrìṣà ROBOTS 100%** 🤖

### 4. Òsanyìn Herb Prescription

```json
{
  "name": "osanyin_herb_prescription",
  "orisa": "osanyin_herb",
  "attributes": {
    "herb": { "key": "cerasee", "class": "ONCO" },
    "regional": { "locale": "DO-Santiago" },
    "veil": { "constant": "φ" }
  },
  "args": {
    "prescription": "detox_protocol",
    "duration_days": 21
  }
}
```

## Òrìṣà Precompiles

| Òrìṣà | Role | Purpose |
|-------|------|---------|
| **Èṣù** (`eshu_router`) | Router & Gateway | Enforces 3.69% tithe, routes actions |
| **Ọbàtálá** (`obatala_guard`) | Governance | Quorum validation, whitegate protection |
| **Ṣàngó** (`sango_vault`) | Vault & Penalties | Manages treasury, slashing |
| **Ọ̀yá** (`oya_witness`) | Witness Network | Coordinates witness mesh |
| **Yemọja** (`yemoja_cache`) | Memory & Cache | Persists ritual state |
| **Ògún** (`ogun_forge`) | Hardware & Manufacturing | Enforces safety, builds simulations |
| **Ọ̀ṣun** (`oshun_river`) | Flow & Abundance | Manages data streams |
| **Òsanyìn** (`osanyin_herb`) | Herbs & Healing | Prescription validation |
| **Ọrúnmìlà** (`orunmila_oracle`) | Oracle & Divination | External data binding |

## Device Support

- ✅ **Drones**: Telemetry, GPS, QR scanning
- ✅ **Phones**: Camera, BLE, WiFi mesh
- ✅ **AVs**: LiDAR, radar, vision systems
- ✅ **Robots**: E-stop safety, IMU, force sensors
- ✅ **Sensors**: Temp, humidity, pressure, light
- ✅ **Cameras**: QR/barcode scanning, CV

## Network Types

- **LoRa**: Long-range mesh (1-10km)
- **BLE**: Short-range mesh (10-100m)
- **WiFi**: Local mesh
- **Mesh**: Hybrid topologies

## Blockchain Anchoring

Supports cross-chain imports:
- **Bitcoin** (`@import(chain:'btc')`)
- **Ethereum** (`@import(chain:'eth')`)
- **Solana** (`@import(chain:'sol')`)
- **Arweave** (`@import(chain:'ar')`)
- **Sui** (`@import(chain:'sui')`)

## Roadmap

### Phase 1 ✅
- Core VM runtime
- Proof-of-Ritual flow
- Telemetry & witness validation
- Òrìṣà precompiles

### Phase 2 ✅
- 143 attributes
- QR scanner API
- Universal device support
- Witness mesh (LoRa/BLE)

### Phase 3 (Next)
- [ ] Integrate Pest parser (native Ọ̀ṢỌ́ syntax)
- [ ] Julia FFI (proof validation math)
- [ ] Move FFI (resource safety)
- [ ] Production LoRa/BLE mesh
- [ ] Cross-chain anchoring (BTC/ETH/SOL/AR/Sui)
- [ ] Òsanyìn herb registry
- [ ] SimaaS API (GPU simulation)
- [ ] AIO job marketplace
- [ ] TechGnØŞ shrine network

## Philosophy

> *"Every device breathes Àṣẹ"*

Ọ̀ṢỶ́VM is **local-first**, **proof-driven**, and **spiritually aligned**:

1. **No cloud required**: Runs on Termux, Raspberry Pi, drones
2. **Proof-of-Ritual**: Real-world actions verified by witnesses
3. **Sacred economics**: 3.69% tithe, 50/25/15/10 shrine splits
4. **Universal**: Works on any device with a CPU

## Community

- **GitHub**: https://github.com/ase-lang/osovm
- **Documentation**: [docs/](docs/)
- **Examples**: [examples/](examples/)

## License

MIT - See [LICENSE](LICENSE)

---

**Àṣẹ!** 🔥🕯️

*Èmi ni Bínò ÈL Guà — ọmọ Kọ́dà, aṣáájú ọ̀nà tuntun-tuntun.*

*Kí ìmọ́lẹ̀ Ọbàtálá máa tàn lọ́nà wa.*
