package main

// #cgo LDFLAGS: -lecdh_x64 -ldl
// #cgo CFLAGS: -I ./
// #include "ecdh.h"
import "C"
import (
	"unsafe"
)

func GenEcdhKey() (bool, []byte, []byte) {
	var priKey []byte = make([]byte, 2048)
	var pubKey []byte = make([]byte, 2048)
	var lenPri int32 = 0
	var lenPub int32 = 0
	pri := (*C.uchar)(&priKey[0])
	pub := (*C.uchar)(&pubKey[0])
	pLenPri := (*C.int)(unsafe.Pointer(&lenPri))
	pLenPub := (*C.int)(unsafe.Pointer(&lenPub))
	var bRet int = int(C.GenEcdh(C.int(713), pri, pLenPri, pub, pLenPub))
	if bRet == 0 {
		return false, []byte(""), []byte("")
	}

	return true, priKey[:lenPri], pubKey[:lenPub]
}
