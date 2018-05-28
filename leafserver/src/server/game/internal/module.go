package internal

import (
	"github.com/name5566/leaf/module"
	"server/base"
)

var (
	// 在程序进入main.go的main()函数之前,该模块的skeleton会通过base.NewSkeleton()被实例化
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton // 如果一个Module是基于Skeleton实现的,则Skeleton就为这个Module提供了ChanRPC的功能.
}

func (m *Module) OnInit() {
	// 在game模块中,OnInit()逻辑其实就是让自己的Module指向我们之前实例化的Skeleton
	// 所以当调用game.Module.Run()的时候,其实调用的逻辑是Skeleton的Run()函数
	m.Skeleton = skeleton
}

func (m *Module) OnDestroy() {

}
