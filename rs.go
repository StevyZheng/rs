package main

import (
	. "fmt"
	"rs/api"
	"rs/libs/net"
	"rs/models"
	"rs/utils/datebase/mysql"
)

func main() {
	var n net.Net
	n.InitNet()
	it := n.Interfaces.Iterator()
	for it.Next() {
		println(it.Value().(net.Interface).Name)
	}
	Println("rs is running...")
	defer mysql.Eloquent.Close()
	models.DBInit()
	router := api.InitRouter()
	_ = router.Run(":80")
}
