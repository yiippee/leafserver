package msg

import (
	//"github.com/name5566/leaf/network/json"
	"github.com/name5566/leaf/network/protobuf"
)

//var Processor network.Processor

//var Processor = json.NewProcessor()
var Processor = protobuf.NewProcessor()

// 一个结构体定义了一个 JSON 消息的格式
// 消息名为 Hello
//type Hello struct {
//	Name string
//}

func init() {
	// 这里我们注册了一个 JSON 消息 Hello
	Processor.Register(&Hello{})
	Processor.Register(&Hello2{})
}
