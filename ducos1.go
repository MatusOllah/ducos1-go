// Package ducos1 implements the DUCO-S1 mining algorithm used in Duino-Coin mining.
package ducos1

import (
	"crypto/sha1"
	"encoding/hex"
	"strconv"
)

// DUCOS1 computes a DUCO-S1 hash.
func DUCOS1(lastHash string, expectedHash string, difficulty int) string {
	hasher := sha1.New()
	for i := 0; i <= difficulty*100; i++ {
		hasher.Write([]byte(lastHash + strconv.Itoa(i)))
		hash := hex.EncodeToString(hasher.Sum(nil)[:])
		if hash == expectedHash {
			return hash
		}
		hasher.Reset()
	}
	return ""
}
