package util

func PutFixed64(dst string, value uint64) string {
	return dst + string(EncodeFixed64(value))
}	

func EncodeFixed64(value uint64) []byte {
	var buffer []byte
	buffer = append(buffer, byte(value))
	buffer = append(buffer, byte(value >> 8))
	buffer = append(buffer, byte(value >> 16))
	buffer = append(buffer, byte(value >> 24))
	buffer = append(buffer, byte(value >> 32))
	buffer = append(buffer, byte(value >> 40))
	buffer = append(buffer, byte(value >> 48))
	buffer = append(buffer, byte(value >> 56))
	return buffer 
}