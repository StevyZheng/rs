package main

import (
	. "fmt"
	"rs/api"
	"rs/models"
	orm "rs/utils/datebase"
)

func main() {
	Println("rs is running...")
	defer orm.Eloquent.Close()
	models.DBInit()
	router := api.InitRouter()
	_ = router.Run(":80")
}
