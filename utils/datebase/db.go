package datebase

import (
	"github.com/jinzhu/gorm"
	"rs/utils"
	"rs/utils/datebase/mysql"
	"rs/utils/datebase/pgsql"
)

var DB *gorm.DB

func CreateDBEngine() {
	dbType := utils.ConfValue.DBType
	if dbType == "pgsql" {
		DB = pgsql.Eloquent
	} else if dbType == "mysql" {
		DB = mysql.Eloquent
	}
}
