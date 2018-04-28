package main

import (
	"bytes"
	"compress/zlib"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"fmt"
	"math/big"
	mrand "math/rand"
	"strings"
	"time"
)

//RandomStr 随机生成字符串
func RandomStr(length int) []byte {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := mrand.New(mrand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return result
}

func GetMd5(src string) string {
	data := []byte(strings.TrimSpace(src))
	has := md5.Sum(data)
	ret := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return ret
}

// 进行zlib压缩
func DoZlibCompress(src []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	w.Write(src)
	w.Close()
	return in.Bytes()
}

func fromBase10(base10 string) *big.Int {
	i, ok := new(big.Int).SetString(base10, 10)
	if !ok {
		panic("bad number: " + base10)
	}
	return i
}

/*
func PKCS1v15(raw string, k *rsa.PrivateKey) {
	// 加密数据
	encData, err := rsa.EncryptPKCS1v15(rand.Reader, &k.PublicKey, []byte(raw))
	CheckErr(err)

	// 将加密信息转换为16进制
	fmt.Println(hex.EncodeToString(encData))

	// 解密数据
	decData, err := rsa.DecryptPKCS1v15(rand.Reader, k, encData)
	CheckErr(err)

	fmt.Println(string(decData))
}
*/
// Rsa加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	pub := &rsa.PublicKey{
		N: fromBase10("28451871049931367000280397980315941493900129515342596978911559687990314360389032587440776677027204713391568456885285049251487633608731647183467169168881911527826624481487591327384831906488048909401577611922812327263514418984933031922276030058409673698944286410157636442703854841054883577764295855609055424607908529646258978732803548772153882771376598661378357620270911570592259824592983240228765987019924029891220246156951679001803386278265765263294008064317769795655401414404284566271952991617207133906501324250043672867665318381453808219063463146255586300194092972814576468544100433701118961141427623372047206165351"),
		E: 65537,
	}
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData) //RSA算法加密
}

func main() {
	encData, _ := RsaEncrypt([]byte("0"))
	fmt.Println(encData)

	// 将加密信息转换为16进制
	fmt.Println(hex.EncodeToString(encData))
	return
}
