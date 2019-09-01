package models

import (
	"github.com/jinzhu/gorm"
	orm "rs/utils/datebase"
)

type Role struct {
	gorm.Model
	RoleName string `json:"role_name" gorm:"type:varchar(32);index:idx_name_code"`
}

type User struct {
	gorm.Model
	UserName string `json:"user_name" gorm:"type:varchar(64)"`
	Password string `json:"password" gorm:"type:varchar(256)"`
	Email    string `json:"email" gorm:"type:varchar(128)"`
	Role     Role   `json:"role"`
}

//列表
func (user *User) Users() (users []User, err error) {
	if err = orm.Eloquent.Find(&users).Error; err != nil {
		return
	}
	return
}

//添加
func (user User) Insert() (id uint, err error) {

	//添加数据
	result := orm.Eloquent.Create(&user)
	id =user.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//修改
func (user *User) Update(id uint) (updateUser User, err error) {

	if err = orm.Eloquent.Select([]string{"id", "username"}).First(&updateUser, id).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Model(&updateUser).Updates(&user).Error; err != nil {
		return
	}
	return
}

//删除数据
func (user *User) Destroy(id uint) (Result User, err error) {

	if err = orm.Eloquent.Select([]string{"id"}).First(&user, id).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Delete(&user).Error; err != nil {
		return
	}
	Result = *user
	return
}