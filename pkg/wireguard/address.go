// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package wireguard

import (
	"crypto/sha256"
	"net/netip"
)

// NetworkPrefix returns IPv6 prefix for the SideroLink.
//
// Server is using the first address in the block.
// Nodes are using random addresses from the /64 space.
func NetworkPrefix(installationID string) netip.Prefix {
	var prefixData [16]byte

	hash := sha256.Sum256([]byte(installationID))

	// Take the last 16 bytes of the clusterID's hash.
	copy(prefixData[:], hash[sha256.Size-16:])

	// Apply the ULA prefix as per RFC4193
	prefixData[0] = 0xfd

	// Apply the Talos-specific ULA Purpose suffix (SideroLink)
	// We are not importing Talos machinery package here, as Talos imports SideroLink library, and this creates an import cycle.
	prefixData[7] = 0x3

	return netip.PrefixFrom(netip.AddrFrom16(prefixData), 64).Masked()
}
