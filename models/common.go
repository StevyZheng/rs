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
	orm.Eloquent.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Role{}, &User{})
}

type OrmModel struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeleteAt  time.Time `json:"delete_at"`
}

/*type OrmModel struct {
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
	DeleteAt  int64 `json:"delete_at"`
}*/
