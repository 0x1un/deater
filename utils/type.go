package utils

import (
	"bytes"
	"encoding/binary"
)

func BytesToInt(b []byte) int {
	bytesBuf := bytes.NewBuffer(b)
	var x int
	err := binary.Read(bytesBuf, binary.BigEndian, &x)
	if err != nil {
		return -1
	}
	return x
}

func IntToBytes(n int) []byte {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, uint32(n))
	return bs
}
