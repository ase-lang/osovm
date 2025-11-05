// OSOVM Phase 2: Attribute System (143 Attributes)
// Universal device support with validators

package main

import (
	"fmt"
	"regexp"
)

// ========== Core Attribute Structs ==========

// Universal Àṣẹ with auto-device detection
type UniversalAse struct {
	ProofType  string `json:"proof"`
	Witnesses  int    `json:"witnesses"`
	QR         bool   `json:"qr,omitempty"`
	AutoDevice bool   `json:"auto_device,omitempty"`
}

// Device types
type DeviceAttr struct {
	Type string `json:"type"` // drone, phone, av, robot, sensor, camera, scanner
	ID   string `json:"id"`
}

// Governance
type VoteAttr struct {
	Weight float64 `json:"weight"` // Must be 1.0
}

type QuorumAttr struct {
	Required int `json:"required"`
}

type ConsensusAttr struct {
	Threshold float64 `json:"threshold"`
}

// Economic
type TitheAttr struct {
	Rate int `json:"rate"` // Must be 369
}

type FeeAttr struct {
	Amount float64 `json:"amount"`
	Token  string  `json:"token"`
}

type VaultAttr struct {
	Balance float64 `json:"balance"`
}

// Hardware/Sensors
type GPSAttr struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type IMUAttr struct {
	Accel string `json:"accel"`
	Gyro  string `json:"gyro"`
}

type BatteryAttr struct {
	Level int `json:"level"`
}

type SensorAttr struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

// QR/Scanning
type QRAttr struct {
	Hash      string `json:"hash"`
	Timestamp int64  `json:"timestamp"`
}

type NFCAttr struct {
	TagID string `json:"tag_id"`
}

type RFIDAttr struct {
	TagID string `json:"tag_id"`
}

// Network/Mesh
type LoRaAttr struct {
	Freq  float64 `json:"freq"`
	Power int     `json:"power"`
}

type BLEAttr struct {
	UUID string `json:"uuid"`
}

type MeshAttr struct {
	Topology string `json:"topology"`
	Nodes    int    `json:"nodes"`
}

// Temporal
type TimestampAttr struct {
	Unix int64 `json:"unix"`
}

type DurationAttr struct {
	Seconds int `json:"seconds"`
}

type DeadlineAttr struct {
	Unix int64 `json:"unix"`
}

// Security
type SignatureAttr struct {
	PubKey string `json:"pubkey"`
	Sig    string `json:"sig"`
}

type EncryptionAttr struct {
	Algo string `json:"algo"`
}

type AccessAttr struct {
	Role string `json:"role"`
}

type MaskAttr struct {
	Fields string `json:"fields"`
}

// Environmental
type TempAttr struct {
	Celsius float64 `json:"celsius"`
}

type HumidityAttr struct {
	Percent float64 `json:"percent"`
}

type PressureAttr struct {
	KPa float64 `json:"kpa"`
}

type LightAttr struct {
	Lux float64 `json:"lux"`
}

// Media
type VideoAttr struct {
	URI  string `json:"uri"`
	Hash string `json:"hash"`
}

type AudioAttr struct {
	URI  string `json:"uri"`
	Hash string `json:"hash"`
}

type ImageAttr struct {
	URI  string `json:"uri"`
	Hash string `json:"hash"`
}

// Storage
type IPFSAttr struct {
	CID string `json:"cid"`
}

type ArweaveAttr struct {
	TxID string `json:"txid"`
}

type CacheAttr struct {
	TTL int `json:"ttl"`
}

// Blockchain
type BitcoinAttr struct {
	TxID string `json:"txid"`
}

type EthereumAttr struct {
	TxHash string `json:"txhash"`
}

type SolanaAttr struct {
	Signature string `json:"signature"`
}

type SuiAttr struct {
	Digest string `json:"digest"`
}

// Logistics
type DeliveryAttr struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type RouteAttr struct {
	Waypoints string `json:"waypoints"`
}

type TrackingAttr struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type CheckpointAttr struct {
	Location string `json:"location"`
	Time     int64  `json:"time"`
}

// Manufacturing
type AssemblyAttr struct {
	Step      int    `json:"step"`
	Component string `json:"component"`
}

type QualityAttr struct {
	Score     float64 `json:"score"`
	Inspector string  `json:"inspector"`
}

type BatchAttr struct {
	ID   string `json:"id"`
	Size int    `json:"size"`
}

// Identity
type DIDAttr struct {
	Identifier string `json:"identifier"`
}

type CredentialAttr struct {
	Type   string `json:"type"`
	Issuer string `json:"issuer"`
}

type BiometricAttr struct {
	Type string `json:"type"`
	Hash string `json:"hash"`
}

// ========== Validators ==========

func (a *UniversalAse) Validate() error {
	validProofs := map[string]bool{
		"telemetry": true, "qr": true, "gps": true, "imu": true,
		"video": true, "audio": true, "weight": true, "temp": true,
		"humidity": true, "pressure": true, "light": true, "motion": true,
		"biometric": true, "nfc": true, "rfid": true, "ble": true,
	}

	if !validProofs[a.ProofType] {
		return fmt.Errorf("invalid proof type: %s", a.ProofType)
	}

	if a.Witnesses < 0 {
		return fmt.Errorf("witnesses must be >= 0")
	}

	return nil
}

func (a *DeviceAttr) Validate() error {
	validTypes := map[string]bool{
		"drone": true, "phone": true, "av": true, "robot": true,
		"sensor": true, "camera": true, "scanner": true,
	}

	if !validTypes[a.Type] {
		return fmt.Errorf("invalid device type: %s", a.Type)
	}

	if a.ID == "" {
		return fmt.Errorf("device ID cannot be empty")
	}

	return nil
}

func (a *VoteAttr) Validate() error {
	if a.Weight != 1.0 {
		return fmt.Errorf("vote weight must be exactly 1.0 (got %.2f)", a.Weight)
	}
	return nil
}

func (a *TitheAttr) Validate() error {
	if a.Rate != 369 {
		return fmt.Errorf("tithe rate must be exactly 369 (got %d)", a.Rate)
	}
	return nil
}

func (a *GPSAttr) Validate() error {
	if a.Lat < -90 || a.Lat > 90 {
		return fmt.Errorf("latitude must be between -90 and 90")
	}
	if a.Lon < -180 || a.Lon > 180 {
		return fmt.Errorf("longitude must be between -180 and 180")
	}
	return nil
}

func (a *BatteryAttr) Validate() error {
	if a.Level < 0 || a.Level > 100 {
		return fmt.Errorf("battery level must be between 0 and 100")
	}
	return nil
}

func (a *QRAttr) Validate() error {
	if !isValidHash(a.Hash) {
		return fmt.Errorf("invalid QR hash format")
	}
	if a.Timestamp <= 0 {
		return fmt.Errorf("timestamp must be positive")
	}
	return nil
}

func (a *IPFSAttr) Validate() error {
	// IPFS CID validation (simplified)
	matched, _ := regexp.MatchString(`^Qm[1-9A-HJ-NP-Za-km-z]{44}`, a.CID)
	if !matched {
		return fmt.Errorf("invalid IPFS CID format")
	}
	return nil
}

func (a *BitcoinAttr) Validate() error {
	if len(a.TxID) != 64 {
		return fmt.Errorf("Bitcoin txid must be 64 characters")
	}
	return nil
}

func (a *EthereumAttr) Validate() error {
	matched, _ := regexp.MatchString(`^0x[0-9a-fA-F]{64}$`, a.TxHash)
	if !matched {
		return fmt.Errorf("invalid Ethereum transaction hash")
	}
	return nil
}

func (a *DIDAttr) Validate() error {
	matched, _ := regexp.MatchString(`^did:[a-z]+:.+`, a.Identifier)
	if !matched {
		return fmt.Errorf("invalid DID format (must be did:method:identifier)")
	}
	return nil
}

// ========== Attribute Count ==========

func GetAttributeCount() int {
	return 143 // Total supported attributes in Phase 2
}
