// OSOVM Phase 2: QR Scanner API
// Works on any device with camera (phone, drone, AV, robot)

package camera

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// QRScan represents a QR code scan result
type QRScan struct {
	RawData   string `json:"raw_data"`
	Hash      string `json:"hash"`
	Timestamp int64  `json:"timestamp"`
	DeviceID  string `json:"device_id"`
}

// Scanner interface for different camera backends
type Scanner interface {
	Scan() (*QRScan, error)
	GetDeviceID() string
}

// MockScanner for Phase 2 demo (simulates scan)
type MockScanner struct {
	DeviceID string
}

func NewMockScanner(deviceID string) *MockScanner {
	return &MockScanner{DeviceID: deviceID}
}

func (s *MockScanner) GetDeviceID() string {
	return s.DeviceID
}

// Scan simulates QR code scanning
// Production: integrate with platform-specific camera APIs
// - Android: CameraX + ML Kit Barcode
// - iOS: AVFoundation
// - Linux: V4L2 + ZBar
func (s *MockScanner) Scan() (*QRScan, error) {
	// Simulate QR code data (in production, this comes from camera)
	mockQRData := fmt.Sprintf("DELIVERY_%d_CHECKPOINT", time.Now().Unix())

	// Hash the QR code for proof
	hash := sha256.Sum256([]byte(mockQRData))
	hashStr := hex.EncodeToString(hash[:])

	return &QRScan{
		RawData:   mockQRData,
		Hash:      hashStr,
		Timestamp: time.Now().Unix(),
		DeviceID:  s.DeviceID,
	}, nil
}

// ScanAndBroadcast scans QR and broadcasts to witness mesh
func ScanAndBroadcast(scanner Scanner, witnessCount int) (*QRScan, error) {
	// 1. Scan QR code
	scan, err := scanner.Scan()
	if err != nil {
		return nil, fmt.Errorf("QR scan failed: %w", err)
	}

	fmt.Printf("📷 QR scanned: %s\n", scan.Hash[:16]+"...")
	fmt.Printf("📡 Broadcasting to %d witnesses...\n", witnessCount)

	// 2. Broadcast to witness mesh (LoRa/BLE)
	// Production: actual mesh broadcast
	// Phase 2: simulated
	time.Sleep(100 * time.Millisecond) // Simulate network delay

	fmt.Printf("✅ Witnesses confirmed scan\n")

	return scan, nil
}

// ========== Platform Detection ==========

// DetectCameraDevice auto-detects available camera
func DetectCameraDevice() string {
	// Production: check /dev/video*, AVFoundation, CameraX
	// Phase 2: return mock device
	return "camera_auto_detected"
}

// ========== Camera API Entry Points ==========

// ScanQRCode is the main API for ritual @qr attribute
func ScanQRCode(deviceID string, witnessCount int) (*QRScan, error) {
	scanner := NewMockScanner(deviceID)
	return ScanAndBroadcast(scanner, witnessCount)
}
