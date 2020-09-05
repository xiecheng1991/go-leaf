package internal

import (
	"fmt"
	"reflect"
	"github.com/name5566/leaf/db/mysql"
	"server/msg"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

var db *mysql.Mysql

func init() {
	//打开一个db链接
	var err error
	db, err = mysql.Open()
	if err != nil {
		panic( err )
	}
	handleMsg( &msg.Register{}, handler_register )
}

func handler_register( args []interface{} ) {
	m := args[0].( *msg.Register )
//	a := args[1].(gate.Agent)
	log.Debug( "recieve register msg, account:%s, password:%s", m.Account, m.PassWord )
	sql := fmt.Sprintf( "insert into account values('%s', '%s')", m.Account, m.PassWord )
	err := db.Exec( sql )
	if err != nil {
		fmt.Println( err )
		return
	}
	//返回一个创建角色成功的消息
	log.Debug( "register ok" )
}

func handler_create( args []interface{} ) {
	m := args[0].(*msg.CreateRole)
	a := args[1].(gate.Agent)
}

func handler_login( args []interface{} ) {
	m := args[0].( *msg.Login )
	a := args[1].( gate.Agent )

	sql := fmt.Sprintf( "select * from account where account = '%s' and password = '%s'", m.Account, m.PassWord )

	rows, err := db.Query( "select * from account where account = '%s', password = '%s'", m.Account, m.PassWord )
	defer rows.Close()
	if err != nil {
		fmt.Println( err )
		return
	}

	if !row.Next() {
		log.Debug( "invaliad account and passworld" )
		return
	}

	//返回login success

}
