package utils

import (
	"encoding/binary"
)

func BytesToInt64(buf []byte) int64 {
	return int64(binary.LittleEndian.Uint64(buf))
}
func BytesToInt32(buf []byte) int32 {
	return int32(binary.LittleEndian.Uint64(buf))
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToUint64(buf []byte) uint64 {
	return binary.LittleEndian.Uint64(buf)
}

func BytesToUint16(buf []byte) uint16 {
	return binary.LittleEndian.Uint16(buf)
}
func BytesToUint32(buf []byte) uint32 {
	return binary.LittleEndian.Uint32(buf)
}
