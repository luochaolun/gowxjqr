package main

/*
#include <stdio.h>
*/
import "C"
import (
	"fmt"
	"math/rand"
	"syscall"
	"time"
	"unsafe"
	// 辅助库
	//"github.com/golang/protobuf/proto"
)

//RandomStr 随机生成字符串
func RandomStr(length int) []byte {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return result
}

func GenEcdhKey() (bool, []byte, []byte) {
	h, err := syscall.LoadLibrary("./dll/ecdh_x64.dll")
	if err != nil {
		return false, []byte{}, []byte{}
	}
	defer syscall.FreeLibrary(h)

	proc, err := syscall.GetProcAddress(h, "GenEcdh")
	if err != nil {
		fmt.Println("GetProcAddress error!")
		return false, []byte{}, []byte{}
	}

	priKey := make([]byte, 2048)
	pubKey := make([]byte, 2048)
	var lenPri int32 = 0
	var lenPub int32 = 0
	pri := uintptr(unsafe.Pointer(&priKey))
	pub := uintptr(unsafe.Pointer(&pubKey))

	pLenPri := uintptr(unsafe.Pointer(&lenPri))
	pLenPub := uintptr(unsafe.Pointer(&lenPub))

	r, k, err := syscall.Syscall6(uintptr(proc),
		5,
		uintptr(713),
		pri,
		pLenPri,
		pub,
		pLenPub,
		0)
	if err != nil {
		return false, []byte{}, []byte{}
	}
	fmt.Printf("%+v\n", priKey)
	fmt.Printf("%+v\n", pubKey)
	fmt.Println(lenPri)
	fmt.Println(lenPub)
	fmt.Println("-----------")
	fmt.Println(r)
	fmt.Println(k)
	fmt.Println("-----------")

	return false, []byte{}, []byte{}
}

func main() {
	b, EcdhPubKey, _ := GenEcdhKey()
	if !b {
		fmt.Println("初始化出错!")
		return
	}

	var ecdhLen int32
	ecdhLen = int32(len(EcdhPubKey))
	a := &LoginInfo{}
	fmt.Println(a)

	var l int32
	l = 16
	var nid int32
	nid = 713
	aes := ManualAuthAccountRequest_AesKey{
		Len: &l,
		Key: RandomStr(16),
	}
	ecdh := ManualAuthAccountRequest_Ecdh{
		Nid: &nid,
		EcdhKey: &ManualAuthAccountRequest_Ecdh_EcdhKey{
			Len: &ecdhLen,
			Key: EcdhPubKey,
		},
	}
	fmt.Println(aes)
	fmt.Println(ecdh)
}
