package main

import (
	"fmt"
	// 辅助库
	//"github.com/golang/protobuf/proto"
)

func main() { // 创建一个消息 Test
	a := &LoginInfo{}
	fmt.Println(a)

	accountRequest := ManualAuthAccountRequest_AesKey{
		Len: 16,
	}
	fmt.Println(accountRequest)
}
