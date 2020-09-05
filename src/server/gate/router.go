package gate

import(
	"server/game"
	"server/msg"
	"server/login"
)

func init() {
	//空结构体Hello{} 程序中只会存在一个
	//game.ChanRPC
	msg.Processor.SetRouter( &msg.Hello{}, game.ChanRPC )
	msg.Processor.SetRouter( &msg.Register{}, login.ChanRPC )
	msg.Processor.SetRouter( &msg.CreateRole{}, log.ChanRPC )
}
