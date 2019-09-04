package main

import (
	. "fmt"
	v1 "rs/api/v1"
	"rs/models"
	orm "rs/utils/datebase"
)

func main() {
	Println("rs is running...")
	defer orm.Eloquent.Close()
	models.DBInit()
	router := v1.InitRouter()
	_ = router.Run(":80")
}
