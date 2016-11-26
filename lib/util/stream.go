package util

import (
	"bytes"
)

func WriteUint32(value uint32, out *bytes.Buffer) {
	out.Write([]byte{byte(value >> 24 & 0xff)})
	out.Write([]byte{byte(value >> 16 & 0xff)})
	out.Write([]byte{byte(value >> 8 & 0xff)})
	out.Write([]byte{byte(value >> 0 & 0xff)})
}
