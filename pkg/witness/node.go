// OSOVM Phase 2: Witness Mesh Node
// Any device can join as witness (phone, drone, sensor, AV)

package witness

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Node represents a witness node in the mesh network
type Node struct {
	DeviceID   string `json:"device_id"`
	DeviceType string `json:"device_type"` // phone, drone, av, robot, sensor
	Network    string `json:"network"`     // lora, ble, wifi, mesh
	PublicKey  string `json:"public_key"`
	Online     bool   `json:"online"`
}

// WitnessSignature is proof that a node witnessed an action
type WitnessSignature struct {
	DeviceID  string `json:"device_id"`
	Signature string `json:"signature"`
	Timestamp int64  `json:"timestamp"`
	Network   string `json:"network"`
}

// CreateNode initializes a new witness node
func CreateNode(deviceID, deviceType, network string) *Node {
	// Generate mock public key (production: real Ed25519 keypair)
	pubKey := generateMockPublicKey(deviceID)

	return &Node{
		DeviceID:   deviceID,
		DeviceType: deviceType,
		Network:    network,
		PublicKey:  pubKey,
		Online:     true,
	}
}

// WitnessAction signs a proof hash
func (n *Node) WitnessAction(proofHash string) (*WitnessSignature, error) {
	if !n.Online {
		return nil, fmt.Errorf("node %s is offline", n.DeviceID)
	}

	// Production: Ed25519 signature over LoRa/BLE
	// Phase 2: SHA256(deviceID + proofHash)
	signature := computeSignature(n.DeviceID, proofHash)

	fmt.Printf("👁️  Witness %s (%s) signed via %s\n", n.DeviceID, n.DeviceType, n.Network)

	return &WitnessSignature{
		DeviceID:  n.DeviceID,
		Signature: signature,
		Timestamp: time.Now().Unix(),
		Network:   n.Network,
	}, nil
}

// BroadcastProof sends proof to witness mesh
func BroadcastProof(proofHash string, requiredWitnesses int, network string) ([]*WitnessSignature, error) {
	fmt.Printf("📡 Broadcasting proof to mesh network (%s)...\n", network)

	// Auto-discover nearby witness nodes
	nodes := discoverWitnessNodes(network, requiredWitnesses)

	if len(nodes) < requiredWitnesses {
		return nil, fmt.Errorf("insufficient witness nodes: found %d, need %d", len(nodes), requiredWitnesses)
	}

	// Collect witness signatures
	signatures := make([]*WitnessSignature, 0, requiredWitnesses)

	for i := 0; i < requiredWitnesses; i++ {
		sig, err := nodes[i].WitnessAction(proofHash)
		if err != nil {
			return nil, fmt.Errorf("witness %s failed: %w", nodes[i].DeviceID, err)
		}
		signatures = append(signatures, sig)
	}

	fmt.Printf("✅ Collected %d/%d witness signatures\n", len(signatures), requiredWitnesses)
	return signatures, nil
}

// ========== Helper Functions ==========

func discoverWitnessNodes(network string, count int) []*Node {
	// Production: actual LoRa/BLE discovery
	// Phase 2: create mock nodes
	nodes := make([]*Node, count)

	deviceTypes := []string{"phone", "drone", "sensor", "av"}

	for i := 0; i < count; i++ {
		deviceType := deviceTypes[i%len(deviceTypes)]
		nodes[i] = CreateNode(
			fmt.Sprintf("witness_%s_%d", network, i+1),
			deviceType,
			network,
		)
	}

	return nodes
}

func computeSignature(deviceID, proofHash string) string {
	data := deviceID + proofHash
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func generateMockPublicKey(deviceID string) string {
	hash := sha256.Sum256([]byte(deviceID + "_pubkey"))
	return hex.EncodeToString(hash[:])
}

// ========== Network Types ==========

const (
	NetworkLoRa = "lora"
	NetworkBLE  = "ble"
	NetworkWiFi = "wifi"
	NetworkMesh = "mesh"
)

// ValidateNetwork checks if network type is supported
func ValidateNetwork(network string) bool {
	valid := map[string]bool{
		NetworkLoRa: true,
		NetworkBLE:  true,
		NetworkWiFi: true,
		NetworkMesh: true,
	}
	return valid[network]
}
