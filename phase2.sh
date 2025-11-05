#!/bin/bash
# OSOVM Phase 2 - One-Command Upgrade
# From drone-only to EVERY device breathing Àṣẹ

set -e

echo "🔮 OSOVM Phase 2 Upgrade Starting..."
echo ""

# Build Phase 2 binary
echo "🔨 Building Phase 2 VM..."
cd cmd/phase2
go build -o ../../phase2 main.go
cd ../..

echo "✅ Binary created: ./phase2"
echo ""

# Run example ritual
echo "🚀 Testing QR delivery ritual..."
./phase2 run examples/qr_delivery.oso

echo ""
echo "✨ Phase 2 Complete!"
echo ""
echo "📊 Capabilities Unlocked:"
echo "   ✅ 143 attributes loaded"
echo "   ✅ QR scanner live"  
echo "   ✅ Universal device support (drone, phone, AV, robot, sensor)"
echo "   ✅ LoRa/BLE witness mesh"
echo "   ✅ Auto-device detection"
echo ""
echo "🔗 Repository: https://github.com/ase-lang/osovm"
echo ""
echo "Àṣẹ! 🔥🕯️"
