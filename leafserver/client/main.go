package main

import (
	"./proto"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3564")
	if err != nil {
		panic(err)
	}

	// Hello 消息（JSON 格式）
	// 对应游戏服务器 Hello 消息结构体
	//data := []byte(`{
	//	"Hello": {
	//		"Name": "leaf"
	//	}
	//}`)
	//
	//// len + data
	//m := make([]byte, 2+len(data))
	//
	//// 默认使用大端序
	//binary.BigEndian.PutUint16(m, uint16(len(data)))
	//
	//copy(m[2:], data)
	//
	//// 发送消息
	//conn.Write(m)

	// protobuf
	// 创建一个消息 Test
	test := &msg.Hello{
		// 使用辅助函数设置域的值
		Name: *proto.String("Leaf"),
	}

	//playId := make([]byte, 2)
	//binary.BigEndian.PutUint16(play)

	// 进行编码
	data, err := proto.Marshal(test)
	if err != nil {
		fmt.Println("marshaling error: ", err)
	}
	// len + data
	m := make([]byte, 4+len(data))

	binary.BigEndian.PutUint16(m, uint16(len(data) + 2))
	//binary.BigEndian.PutUint16(m, uint16(0))
	//// 默认使用大端序
	//binary.BigEndian.PutUint16(m, uint16(len(data)))

	copy(m[4:], data)

	// 发送消息
	conn.Write(m)
}
