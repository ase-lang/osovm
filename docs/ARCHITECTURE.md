# OSOVM Architecture

## Overview

OSOVM is a virtual machine for executing proof-of-ritual programs written in the Ọ̀ṢỌ́ language. The system validates real-world actions through cryptographic proofs and witness consensus.

## Core Components

### 1. Parser (Pest Grammar)

**File**: `grammar/ORISA.pest`

Defines the declarative syntax for rituals. Key elements:
- **Attributes**: `@àṣẹ`, `@whitegate`, `@sabbath`, etc.
- **Òrìṣà Keywords**: `eshu_router`, `obatala_guard`, `sango_vault`
- **Statements**: Assignments, calls, returns

Phase 1 uses JSON for simplicity; Phase 2 will use full Pest parser in Go.

### 2. VM Runtime (Go)

**File**: `vm/osovm.go`

Execution engine with:
- **Ritual Loader**: Parses .oso files
- **Àṣẹ Validator**: Verifies proof + witnesses
- **Òrìṣà Precompiles**: Executes spiritual archetypes
- **Statement Executor**: Runs ritual logic

### 3. Proof System

**Proof Types**:
- `telemetry`: Sensor data (GPS, QR codes, camera)
- `cognition`: AI decision logs
- `hardware`: Robot state confirmations

**Proof Structure**:
```json
{
  "type": "telemetry",
  "receipt": "sha256_hash_of_data",
  "timestamp": 1730851200,
  "device_id": "drone_001"
}
```

### 4. Witness Network

Devices on LoRa/mesh network confirm proofs:
1. **Broadcast**: Prover sends proof to network
2. **Verify**: Witnesses validate telemetry/action
3. **Sign**: Witnesses sign proof receipt
4. **Collect**: Prover gathers signatures
5. **Submit**: Prover submits to OSOVM

**Witness Structure**:
```json
{
  "device_id": "witness_001",
  "signature": "sha256(device_id + proof_receipt)",
  "timestamp": 1730851200
}
```

### 5. Òrìṣà Precompiles

Spiritual archetypes mapped to VM functions:

| Òrìṣà | Function | Phase 1 Behavior |
|-------|----------|------------------|
| Èṣù | `eshu_router()` | Logs routing; future: tithe enforcement |
| Ọbàtálá | `obatala_guard()` | Validates quorum; future: governance |
| Ṣàngó | `sango_vault()` | Logs vault ops; future: treasury splits |

Phase 2+ will add FFI to Julia (proof math) and Move (resource safety).

## Execution Flow

```
┌─────────────────────────────────────────┐
│ 1. Load Ritual                          │
│    - Parse .oso file                    │
│    - Extract @àṣẹ requirements          │
└──────────┬──────────────────────────────┘
           │
           ▼
┌─────────────────────────────────────────┐
│ 2. Validate Àṣẹ                         │
│    - Check proof type matches           │
│    - Verify hash format                 │
│    - Count witnesses (quorum)           │
│    - Validate signatures                │
└──────────┬──────────────────────────────┘
           │
           ▼
┌─────────────────────────────────────────┐
│ 3. Execute Òrìṣà Precompile             │
│    - eshu_router / obatala_guard / etc. │
└──────────┬──────────────────────────────┘
           │
           ▼
┌─────────────────────────────────────────┐
│ 4. Run Statements                       │
│    - Assignments, calls, returns        │
└──────────┬──────────────────────────────┘
           │
           ▼
┌─────────────────────────────────────────┐
│ 5. Seal Àṣẹ                             │
│    - Log: "Ritual complete. Àṣẹ flows." │
└─────────────────────────────────────────┘
```

## Future Architecture (Phase 2+)

### FFI Integration

```
┌─────────────────────────────────────────┐
│              OSOVM (Go)                 │
├─────────────────────────────────────────┤
│  ┌─────────────┐   ┌─────────────┐     │
│  │   Parser    │   │  Execution  │     │
│  │   (Pest)    │   │   Engine    │     │
│  └─────────────┘   └──────┬──────┘     │
└──────────────────────────┼─────────────┘
                           │
          ┌────────────────┴────────────────┐
          ▼                                 ▼
    ┌──────────┐                      ┌──────────┐
    │  Julia   │                      │   Move   │
    │          │                      │          │
    │ - Proof  │                      │ - Linear │
    │   math   │                      │   types  │
    │ - Crypto │                      │ - Asset  │
    │   verify │                      │   safety │
    └──────────┘                      └──────────┘
```

**Julia FFI**: Validate telemetry math, ZK proofs, statistical analysis  
**Move FFI**: Enforce resource ownership, prevent double-spends

### Network Layer

```
┌─────────────────────────────────────────┐
│          Device Network                 │
│  ┌──────┐  ┌──────┐  ┌──────┐          │
│  │Drone │  │Robot │  │ IoT  │          │
│  └──┬───┘  └──┬───┘  └──┬───┘          │
└─────┼─────────┼─────────┼──────────────┘
      │         │         │
      └─────────┴─────────┘
              │
         LoRa/Mesh
              │
      ┌───────▼────────┐
      │ Witness Relay  │
      │  (Gossip)      │
      └───────┬────────┘
              │
      ┌───────▼────────┐
      │    OSOVM       │
      │  (On-device)   │
      └────────────────┘
```

## Security Considerations

### Proof Integrity
- All proofs must be SHA-256 hashed
- Timestamps prevent replay attacks
- Device IDs prevent spoofing

### Witness Trust
- Minimum quorum enforced (e.g., 3 witnesses)
- Signatures verified against device public keys
- Geographic distribution checks (future)

### Sabbath Enforcement
- `@sabbath` blocks execution on specified days
- Prevents maintenance during sacred periods
- VM returns error if invoked on Sabbath

## Testing Strategy

### Unit Tests
- Proof validation logic
- Witness signature verification
- Òrìṣà precompile execution

### Integration Tests
- Full ritual execution
- Multi-witness scenarios
- Error handling (insufficient witnesses, invalid proofs)

### Hardware Tests
- Drone QR scanning
- Robot telemetry streaming
- LoRa mesh communication

## Performance Targets

- **Parse time**: <100ms for 1000-line ritual
- **Validation**: <50ms for 10 witnesses
- **Execution**: <500ms for typical ritual
- **Memory**: <50MB per VM instance

## Extensibility

### Adding New Òrìṣà
1. Define keyword in `grammar/ORISA.pest`
2. Add precompile function in `vm/osovm.go`
3. Document in `README.md` Òrìṣà table

### Adding New Attributes
1. Define syntax in `grammar/ORISA.pest`
2. Add validation in `vm/osovm.go`
3. Update schema in `docs/SCHEMA.md` (future)

### Adding FFI
1. Define C bindings for Julia/Move
2. Add FFI calls in precompiles
3. Test with mock data, then real computation

---

**Last Updated**: 2025-11-05  
**Phase**: 1 (Core Runtime)
