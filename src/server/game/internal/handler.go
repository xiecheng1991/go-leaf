package internal

import(
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/gate"
	"reflect"
	"server/msg"
)

func init() {
	handler( &msg.Hello{}, handleHello )
}

func handler( m interface{}, h interface{} ) {
	skeleton.RegisterChanRPC( reflect.TypeOf(m), h )
}

func handleHello( args []interface{} ) {
	m := args[0].( *msg.Hello ) //断言
	a := args[1].( gate.Agent )
	log.Debug( "Hello %v", m.Name )
	//给发送者一个回应消息
	a.WriteMsg( &msg.Hello{
		Name : "Client",
	} )
}
