package models

import (
	"rs/utils"
	"rs/utils/datebase"
	"time"
)

func DBInit() {
	datebase.CreateDBEngine()
	if utils.ConfValue.DBType == "mysql" {
		datebase.DB.Set("gorm:table_options", "DEFAULT CHARSET=utf8 AUTO_INCREMENT=1").AutoMigrate(&Conf{}, &Role{}, &User{}, &Api{})
	} else if utils.ConfValue.DBType == "pgsql" {
		datebase.DB.AutoMigrate(&Conf{}, &Role{}, &User{}, &Api{})
	}

	datebase.DB.Model(&User{}).AddForeignKey("role_id", "roles(role_id)", "RESTRICT", "RESTRICT")
}

type OrmModel struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"delete_at"`
}

type Conf struct {
	OrmModel
	ConfID    int64  `json:"conf_id" gorm:"primary_key;unique_index"`
	ConfGroup string `json:"conf_group"`
	ConfName  string `json:"conf_name" gorm:"index:idx_name_code;unique_index"`
	ConfValue string `json:"conf_value"`
}
