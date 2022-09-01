package byte64

import (
	"strings"
)

const encodeChars = "-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"

// Encode function
func Encode(raw []byte) string {
	sb := strings.Builder{}
	b := make([]byte, 3)
	size := 0
	for i := 0; i < len(raw); i++ {
		b[size] = raw[i]
		size++
		if size == 3 {
			sb.WriteByte(encodeChars[b[0]&0x3f])
			sb.WriteByte(encodeChars[(b[1]&0x0f)<<2+b[0]>>6])
			sb.WriteByte(encodeChars[(b[2]&0x03)<<4+b[1]>>4])
			sb.WriteByte(encodeChars[b[2]>>2])

			size = 0
		}
	}

	if size == 1 {
		// output
		sb.WriteByte(encodeChars[b[0]&0x3f])
		sb.WriteByte(encodeChars[b[0]>>6])
		size = 0
	}

	if size == 2 {
		// output
		sb.WriteByte(encodeChars[b[0]&0x3f])
		sb.WriteByte(encodeChars[(b[1]&0x0f)<<2+b[0]>>6])
		sb.WriteByte(encodeChars[b[1]>>4])
		size = 0
	}

	return sb.String()
}
