/* Conversion functions from Go types to their bytecode representation
 *
 */
package BytecodeGenerator

// Convert int64 to a byte slice; big endian
func I64tobyteslice(x int64) []byte {
	b := make([]byte, 8)
	b[7] = byte(x)
	b[6] = byte(x >> 8)
	b[5] = byte(x >> 16)
	b[4] = byte(x >> 24)
	b[3] = byte(x >> 32)
	b[2] = byte(x >> 40)
	b[1] = byte(x >> 48)
	b[0] = byte(x >> 56)
	return b
}

// Convert uint64 to a byte slice; big endian
func Ui64tobyteslice(x uint64) []byte {
	b := make([]byte, 8)
	b[7] = byte(x)
	b[6] = byte(x >> 8)
	b[5] = byte(x >> 16)
	b[4] = byte(x >> 24)
	b[3] = byte(x >> 32)
	b[2] = byte(x >> 40)
	b[1] = byte(x >> 48)
	b[0] = byte(x >> 56)
	return b
}

// Convert uint32 to a byte slice; big endian
func Ui32tobyteslice(x uint32) []byte {
	b := make([]byte, 4)
	b[3] = byte(x)
	b[2] = byte(x >> 8)
	b[1] = byte(x >> 16)
	b[0] = byte(x >> 24)
	return b
}
