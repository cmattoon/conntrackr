package typeconv

import "strconv"

// Converts a 32-bit hex string to a uint32
func Hex2uint32(hex string) uint32 {
	n, _ := strconv.ParseUint(hex, 16, 32)
	return uint32(n)
}
