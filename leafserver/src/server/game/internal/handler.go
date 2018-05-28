package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"reflect"
	"server/msg"
	"time"
	"github.com/golang/protobuf/proto"
)

func init() {
	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	handler(&msg.Hello{}, handleHello)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleHello(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.Hello)
	// 消息的发送者
	a := args[1].(gate.Agent)

	time.Sleep(time.Second*5)
	// 输出收到的消息的内容
	log.Debug("hello %v", m.GetName())

	// 给发送者回应一个 Hello 消息
	a.WriteMsg(&msg.Hello{
		Name: *proto.String("client"),
	})
}
