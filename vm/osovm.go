// OSOVM - Ọ̀ṢỌ́ Virtual Machine
// Phase 1: Proof-of-Ritual Core Runtime
// Handles ritual execution, witness validation, and Àṣẹ flow

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// ========== Core Types ==========

type ProofType string

const (
	ProofTelemetry ProofType = "telemetry"
	ProofCognition ProofType = "cognition"
	ProofHardware  ProofType = "hardware"
)

// Àṣẹ Attribute - The sacred seal
type AseAttr struct {
	ProofType ProofType `json:"proof"`
	Witnesses int       `json:"witnesses"`
}

// Ritual - Sacred function invoking Òrìṣà
type Ritual struct {
	Name       string            `json:"name"`
	Orisa      string            `json:"orisa"`
	Ase        *AseAttr          `json:"ase,omitempty"`
	Args       map[string]string `json:"args"`
	Statements []Statement       `json:"statements"`
}

// Statement - Ritual instructions
type Statement struct {
	Type  string                 `json:"type"` // "call", "assign", "return"
	Data  map[string]interface{} `json:"data"`
}

// Proof - Real-world action verification
type Proof struct {
	Type      ProofType `json:"type"`
	Receipt   string    `json:"receipt"`   // Hash of telemetry/action
	Timestamp int64     `json:"timestamp"`
	DeviceID  string    `json:"device_id"`
}

// Witness - Network confirmation
type Witness struct {
	DeviceID  string `json:"device_id"`
	Signature string `json:"signature"`
	Timestamp int64  `json:"timestamp"`
}

// Execution Context
type Context struct {
	Ritual    *Ritual
	Proof     *Proof
	Witnesses []Witness
	Variables map[string]interface{}
}

// ========== VM State ==========

type VM struct {
	Rituals map[string]*Ritual
	Context *Context
}

func NewVM() *VM {
	return &VM{
		Rituals: make(map[string]*Ritual),
	}
}

// ========== Parser (Simplified for Phase 1) ==========

func (vm *VM) LoadRitual(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read ritual: %w", err)
	}

	// Simple JSON-based ritual format for Phase 1
	// Production would use Pest parser
	var ritual Ritual
	if err := json.Unmarshal(data, &ritual); err != nil {
		return fmt.Errorf("failed to parse ritual: %w", err)
	}

	vm.Rituals[ritual.Name] = &ritual
	return nil
}

// ========== Execution Engine ==========

func (vm *VM) Execute(ritualName string, proof *Proof, witnesses []Witness) error {
	ritual, exists := vm.Rituals[ritualName]
	if !exists {
		return fmt.Errorf("ritual '%s' not found", ritualName)
	}

	// Initialize context
	vm.Context = &Context{
		Ritual:    ritual,
		Proof:     proof,
		Witnesses: witnesses,
		Variables: make(map[string]interface{}),
	}

	fmt.Printf("🔮 Invoking %s ritual...\n", ritual.Orisa)

	// 1. Validate Àṣẹ requirements
	if ritual.Ase != nil {
		if err := vm.validateAse(); err != nil {
			return fmt.Errorf("❌ Àṣẹ validation failed: %w", err)
		}
		fmt.Println("✅ Àṣẹ sealed with proof + witnesses")
	}

	// 2. Execute Òrìṣà precompile
	if err := vm.executeOrisa(); err != nil {
		return fmt.Errorf("❌ Òrìṣà execution failed: %w", err)
	}

	// 3. Execute statements
	for _, stmt := range ritual.Statements {
		if err := vm.executeStatement(&stmt); err != nil {
			return fmt.Errorf("❌ Statement execution failed: %w", err)
		}
	}

	fmt.Println("✨ Ritual complete. Àṣẹ flows.")
	return nil
}

// ========== Àṣẹ Validation ==========

func (vm *VM) validateAse() error {
	ase := vm.Context.Ritual.Ase
	proof := vm.Context.Proof
	witnesses := vm.Context.Witnesses

	// 1. Verify proof type matches
	if proof == nil {
		return fmt.Errorf("proof required but not provided")
	}
	if proof.Type != ase.ProofType {
		return fmt.Errorf("proof type mismatch: expected %s, got %s", ase.ProofType, proof.Type)
	}

	// 2. Validate proof receipt (check hash format)
	if !isValidHash(proof.Receipt) {
		return fmt.Errorf("invalid proof receipt hash")
	}

	// 3. Check witness quorum
	if len(witnesses) < ase.Witnesses {
		return fmt.Errorf("insufficient witnesses: need %d, got %d", ase.Witnesses, len(witnesses))
	}

	// 4. Validate witness signatures (simplified for Phase 1)
	for i, w := range witnesses {
		if !vm.validateWitness(&w, proof) {
			return fmt.Errorf("witness %d signature invalid", i)
		}
	}

	fmt.Printf("📡 Proof validated: %s (%s)\n", proof.Type, proof.Receipt[:16]+"...")
	fmt.Printf("👥 Witnesses confirmed: %d/%d\n", len(witnesses), ase.Witnesses)

	return nil
}

func (vm *VM) validateWitness(w *Witness, proof *Proof) bool {
	// Phase 1: Simple hash-based signature
	// Production: Use Ed25519 signatures over LoRa/mesh
	expected := computeWitnessSignature(w.DeviceID, proof.Receipt)
	return w.Signature == expected
}

func computeWitnessSignature(deviceID, proofReceipt string) string {
	data := deviceID + proofReceipt
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func isValidHash(hash string) bool {
	if len(hash) != 64 {
		return false
	}
	_, err := hex.DecodeString(hash)
	return err == nil
}

// ========== Òrìṣà Precompiles ==========

func (vm *VM) executeOrisa() error {
	orisa := vm.Context.Ritual.Orisa

	switch orisa {
	case "eshu_router":
		return vm.eshuRouter()
	case "obatala_guard":
		return vm.obatalaGuard()
	case "sango_vault":
		return vm.sangoVault()
	default:
		return fmt.Errorf("unknown Òrìṣà: %s", orisa)
	}
}

// Èṣù - Router & Gateway
func (vm *VM) eshuRouter() error {
	fmt.Println("🍶 Èṣù routes the action...")
	// Phase 1: Log routing logic
	// Production: Enforce 3.69% tithe, conditional triggers
	return nil
}

// Ọbàtálá - Governance & Safety
func (vm *VM) obatalaGuard() error {
	fmt.Println("🤍 Ọbàtálá enforces quorum...")
	// Phase 1: Validate quorum from witnesses
	// Production: Enforce consensus, mask policies
	return nil
}

// Ṣàngó - Vault & Penalties
func (vm *VM) sangoVault() error {
	fmt.Println("⚡ Ṣàngó manages the vault...")
	// Phase 1: Log vault operations
	// Production: Handle TechGnØŞ splits, penalties
	return nil
}

// ========== Statement Execution ==========

func (vm *VM) executeStatement(stmt *Statement) error {
	switch stmt.Type {
	case "call":
		function := stmt.Data["function"].(string)
		fmt.Printf("  → Calling %s\n", function)
	case "assign":
		variable := stmt.Data["variable"].(string)
		value := stmt.Data["value"]
		vm.Context.Variables[variable] = value
		fmt.Printf("  → %s = %v\n", variable, value)
	case "return":
		value := stmt.Data["value"]
		fmt.Printf("  → Returning: %v\n", value)
	}
	return nil
}

// ========== CLI Entry Point ==========

func main() {
	// Handle Termux-specific behavior where Args[1] is the full binary path
	argsOffset := 1
	if len(os.Args) > 1 && os.Args[1][0] == '/' {
		argsOffset = 2
	}
	
	if len(os.Args) < argsOffset+2 {
		fmt.Println("Usage: oso run <ritual.oso>")
		os.Exit(1)
	}

	command := os.Args[argsOffset]
	ritualPath := os.Args[argsOffset+1]

	if command != "run" {
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}

	vm := NewVM()

	// Load ritual
	if err := vm.LoadRitual(ritualPath); err != nil {
		fmt.Printf("Error loading ritual: %v\n", err)
		os.Exit(1)
	}

	// Extract ritual name from path
	parts := strings.Split(ritualPath, "/")
	filename := parts[len(parts)-1]
	ritualName := strings.TrimSuffix(filename, ".oso")

	// Mock proof + witnesses for Phase 1 demo
	proof := &Proof{
		Type:      ProofTelemetry,
		Receipt:   "a4f2b8c3d9e1f5a7b2c8d4e9f1a3b5c7d2e4f6a8b1c3d5e7f9a2b4c6d8e1f3a5",
		Timestamp: time.Now().Unix(),
		DeviceID:  "drone_001",
	}

	witnesses := []Witness{
		{
			DeviceID:  "witness_001",
			Signature: computeWitnessSignature("witness_001", proof.Receipt),
			Timestamp: time.Now().Unix(),
		},
		{
			DeviceID:  "witness_002",
			Signature: computeWitnessSignature("witness_002", proof.Receipt),
			Timestamp: time.Now().Unix(),
		},
		{
			DeviceID:  "witness_003",
			Signature: computeWitnessSignature("witness_003", proof.Receipt),
			Timestamp: time.Now().Unix(),
		},
	}

	// Execute ritual
	if err := vm.Execute(ritualName, proof, witnesses); err != nil {
		fmt.Printf("Execution failed: %v\n", err)
		os.Exit(1)
	}
}
