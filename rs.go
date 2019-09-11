package main

import (
	. "fmt"
	"rs/api"
	"rs/models"
	"rs/utils/datebase/mysql"
)

func main() {
	Println("rs is running...")
	defer mysql.Eloquent.Close()
	models.DBInit()
	router := api.InitRouter()
	_ = router.Run(":80")
}
