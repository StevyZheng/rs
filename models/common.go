package models

import (
	orm "rs/utils/datebase"
	"time"
)

func DBInit() {
	AutoMigrate()
	orm.Eloquent.Model(&User{}).AddForeignKey("role_id", "roles(role_id)", "RESTRICT", "RESTRICT")
}

func AutoMigrate() {
	orm.Eloquent.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Conf{}, &Role{}, &User{}, &Api{})
}

type OrmModel struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"delete_at"`
}

/*type OrmModel struct {
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
	DeleteAt  int64 `json:"delete_at"`
}*/

type Conf struct {
	OrmModel
	ConfID    int64  `json:"conf_id" gorm:"primary_key;unique_index"`
	ConfGroup string `json:"conf_group"`
	ConfName  string `json:"conf_name" gorm:"index:idx_name_code;unique_index"`
	ConfValue string `json:"conf_value"`
}
