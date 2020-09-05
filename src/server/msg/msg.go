package msg

import (
	//"github.com/name5566/leaf/network"
	"github.com/name5566/leaf/network/json"
)

//var Processor network.Processor
var Processor = json.NewProcessor()

func init() {
	Processor.Register( &Hello{} )
	Processor.Register( &Register{} )
	Processor.Register( &Login )
}

//注册一个Hello消息,至于发送到那个模块
//就需要看leafserver/gate/router.go
type Hello struct {
	Name string
}

type Register struct {
	Account string
	PassWord string
}

//账号登陆
type Login struct {
	Account string
	PassWord string
}

//产生一个角色id 所属账号
type CreateRole struct {
	Name string
}

type EnterGame struct {
	Rid int
}


