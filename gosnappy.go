package gosnappy

// #cgo LDFLAGS: -lsnappy
// #include <snappy-c.h>
import "C"

import (
	"fmt"
	"unsafe"
)

func Encode(input []byte) ([]byte, error) {
	inputLen := C.size_t(len(input))
	outputLen := C.snappy_max_compressed_length(inputLen)
	output := make([]byte, outputLen)
	res := C.snappy_compress((*_Ctype_char)(unsafe.Pointer(&input[0])), inputLen, (*_Ctype_char)(unsafe.Pointer(&output[0])), &outputLen)
	if res != C.SNAPPY_OK {
		return nil, fmt.Errorf("snappy error: %d", res)
	}
	return output[0:outputLen], nil
}

func Decode(input []byte) ([]byte, error) {
	inputLen := C.size_t(len(input))
	var outputLen C.size_t
	C.snappy_uncompressed_length((*_Ctype_char)(unsafe.Pointer(&input[0])), inputLen, &outputLen)
	output := make([]byte, outputLen)
	res := C.snappy_uncompress((*_Ctype_char)(unsafe.Pointer(&input[0])), inputLen, (*_Ctype_char)(unsafe.Pointer(&output[0])), &outputLen)
	if res != C.SNAPPY_OK {
		return nil, fmt.Errorf("snappy error: %d", res)
	}
	return output[0:outputLen], nil
}
