package internal

import (
	"github.com/name5566/leaf/module"
	"server/base"
)

var (
	//每个模块都创建了一个包级私有的skeleton
	skeleton = base.NewSkeleton()
	//得到一个ChanRPC 通过channel进行rpc通信
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
}

func (m *Module) OnDestroy() {

}
