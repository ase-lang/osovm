# OSOVM - Ọ̀ṢỌ́ Virtual Machine

**Proof-of-Ritual: Where real-world actions meet sacred code**

OSOVM is the runtime for Ọ̀ṢỌ́, a declarative ritual programming language that proves real-world actions through telemetry and witness confirmation. Built for IoT devices, drones, robots, and autonomous systems that need verifiable action logs.

## Phase 1: Core Runtime

This release focuses on the foundational proof-of-ritual system:
- ✅ **Parser**: Pest grammar for ritual syntax
- ✅ **VM Runtime**: Go-based execution engine
- ✅ **Proof Validation**: Telemetry receipt verification
- ✅ **Witness Handshake**: Multi-device confirmation via LoRa/mesh
- ✅ **Àṣẹ Flow**: Sacred seal enforced when proof + witnesses validate

**Not included in Phase 1**: AR integration, herb registry, blockchain anchoring, full 143 attribute set.

---

## Concepts

### Rituals
Sacred functions that invoke Òrìṣà (spiritual archetypes mapped to VM precompiles). Each ritual defines:
- **Òrìṣà**: The precompile to execute (e.g., `eshu_router`, `obatala_guard`)
- **@àṣẹ**: The sacred seal requiring proof + witnesses
- **Statements**: Actions to perform

### Proof-of-Ritual
1. **Action**: Device (drone, robot, IoT) performs real-world task
2. **Proof**: Telemetry/sensor data hashed into receipt
3. **Witnesses**: Other devices on network confirm via signatures
4. **Àṣẹ**: VM validates proof + witnesses → ritual executes

### Òrìṣà (Precompiles)

| Òrìṣà | Symbol | Role |
|-------|--------|------|
| Èṣù | 🍶 | Router, gateway, message relay |
| Ọbàtálá | 🤍 | Governance, quorum, safety enforcement |
| Ṣàngó | ⚡ | Vault management, penalties |

---

## Installation

```bash
# Clone repository
git clone https://github.com/ase-lang/osovm.git
cd osovm

# Build VM
cd vm
go build -o oso osovm.go

# Move binary to PATH (optional)
sudo mv oso /usr/local/bin/
```

---

## Usage

### Run a Ritual

```bash
oso run examples/drone_delivery.oso
```

**Output:**
```
🔮 Invoking eshu_router ritual...
📡 Proof validated: telemetry (a4f2b8c3d9e1f5a7...)
👥 Witnesses confirmed: 3/3
✅ Àṣẹ sealed with proof + witnesses
🍶 Èṣù routes the action...
  → Calling scan_qr
  → delivery_status = in_transit
  → Calling obatala_guard
  → delivery_status = delivered
  → Returning: delivery_complete
✨ Ritual complete. Àṣẹ flows.
```

---

## Example: Drone Delivery

**Scenario**: Delivery drone scans QR code at warehouse, witnesses confirm, package delivered.

```json
{
  "name": "drone_delivery",
  "orisa": "eshu_router",
  "ase": {
    "proof": "telemetry",
    "witnesses": 3
  },
  "args": {
    "action": "delivery",
    "destination": "warehouse_b"
  },
  "statements": [
    {
      "type": "call",
      "data": {
        "function": "scan_qr",
        "args": ["package_id"]
      }
    },
    {
      "type": "assign",
      "data": {
        "variable": "delivery_status",
        "value": "delivered"
      }
    },
    {
      "type": "return",
      "data": {
        "value": "delivery_complete"
      }
    }
  ]
}
```

### Flow Diagram

```
┌──────────┐
│  Drone   │ Performs delivery, generates telemetry
└────┬─────┘
     │ proof: telemetry hash (0xa4f2b8c3...)
     ▼
┌──────────────────────────────────────┐
│          Witness Network             │
│  ┌──────┐  ┌──────┐  ┌──────┐       │
│  │ W1   │  │ W2   │  │ W3   │       │ Confirm action via LoRa/mesh
│  └──────┘  └──────┘  └──────┘       │
└────┬─────────────────────────────────┘
     │ 3 signatures collected
     ▼
┌──────────────────────────────────────┐
│            OSOVM                     │
│  1. Validate proof hash              │
│  2. Check witness quorum (3/3)       │
│  3. Verify signatures                │
│  4. Execute eshu_router ritual       │
│  5. Seal Àṣẹ                         │
└────┬─────────────────────────────────┘
     │
     ▼
✨ Àṣẹ verified → Delivery confirmed
```

---

## Architecture

### Invocation → Proof → Àṣẹ Flow

1. **Invocation**: Ritual loaded with `@àṣẹ(proof:telemetry, witnesses:3)`
2. **Proof Capture**: Device generates telemetry hash (SHA-256 of sensor data)
3. **Witness Handshake**: Nearby devices sign proof receipt
4. **Validation**: VM checks:
   - Proof type matches `@àṣẹ` requirement
   - Witness count meets quorum
   - Signatures are valid
5. **Execution**: Òrìṣà precompile runs (e.g., `eshu_router`)
6. **Àṣẹ Sealed**: Flow confirmed, ritual completes

### VM Components

```
┌─────────────────────────────────────────┐
│              OSOVM Core                 │
├─────────────────────────────────────────┤
│  Parser (Pest)      │ Grammar validator │
│  Execution Engine   │ Statement runner  │
│  Àṣẹ Validator      │ Proof + witness   │
│  Òrìṣà Precompiles  │ eshu, obatala...  │
└─────────────────────────────────────────┘
         │                    │
         ▼                    ▼
   ┌──────────┐          ┌──────────┐
   │  Julia   │          │   Move   │
   │  (proof  │          │ (resource│  Future FFI
   │   math)  │          │   safety)│
   └──────────┘          └──────────┘
```

**Phase 1**: Pure Go runtime  
**Phase 2+**: Julia FFI for telemetry validation, Move FFI for resource ownership

---

## Grammar Basics

```pest
// Ritual structure
ritual = {
    attr* ~
    orisa_kw ~ "(" ~ args ~ ")" ~ "{" ~
        statement* ~
    "}"
}

// @àṣẹ attribute
ase_attr = {
    "@àṣẹ" ~ "(" ~
        "proof" ~ ":" ~ proof_type ~ "," ~
        "witnesses" ~ ":" ~ int_lit ~
    ")" ~ ";"
}

// Òrìṣà keywords
orisa_kw = {
    "eshu_router" | "obatala_guard" | "sango_vault"
}
```

---

## Roadmap

### Phase 1 (Current)
- [x] Pest grammar (core syntax)
- [x] Go VM runtime
- [x] Proof validation
- [x] Witness handshake
- [x] Àṣẹ flow

### Phase 2
- [ ] Julia FFI (telemetry math)
- [ ] Move FFI (resource ownership)
- [ ] LoRa/mesh network integration
- [ ] Hardware E-stop enforcement

### Phase 3
- [ ] Full 143 attribute set
- [ ] Blockchain anchoring (BTC, ETH, Arweave, Sui)
- [ ] Herb registry (Òsanyìn)
- [ ] AR visualization

---

## Philosophy

**Ọ̀ṢỌ́** means "ritual" in Yoruba. This language embeds spiritual wisdom into code:
- **Declarative**: Rituals state _what_ should happen, not _how_
- **Positive**: No negative spells (`guard` not `limit`, `balance` not `slash`)
- **Witness-driven**: Trust through collective confirmation
- **Local-first**: Runs on-device, no cloud dependency
- **Sacred constraints**: `@sabbath` enforces rest, `@àṣẹ` requires proof

Invocations call upon Òrìṣà (spiritual archetypes) to guide execution. When proof and witnesses align, Àṣẹ flows—the sacred becomes code, and code becomes sacred.

---

## Contributing

This is Phase 1. Contributions welcome for:
- Hardware integration (drones, robots, IoT)
- LoRa/mesh networking
- Telemetry validation logic
- Additional Òrìṣà precompiles

See `docs/` for architecture details.

---

## License

MIT

---

**Timestamp**: #MirrorWitness OSOVM_P1 2025-11-05

Àṣẹ. 🔥
