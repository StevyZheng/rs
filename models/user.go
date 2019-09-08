package models

import (
	orm "rs/utils/datebase"
)

type User struct {
	OrmModel
	UserID   int64  `json:"id" gorm:"primary_key" sql:"AUTO_INCREMENT"`
	UserName string `json:"user_name" gorm:"type:varchar(64);unique_index" `
	Password string `json:"password" gorm:"type:varchar(256)"`
	Email    string `json:"email" gorm:"type:varchar(128)"`
	Role     Role   `json:"role" gorm:"auto_preload;foreignkey:RoleID"`
	RoleID   int64  `json:"role_id"`
}

//列表
func (user *User) UserList() (users []User, err error) {
	//orm.Eloquent.Model(&user).Related(&user.Role).Find(&user.Role)
	if err = orm.Eloquent.Find(&users).Error; err != nil {
		return
	}
	var r Role
	for i, value := range users {
		users[i].Role.RoleID = value.RoleID
		orm.Eloquent.Where("role_id = ?", users[i].RoleID).First(&r)
		users[i].Role.RoleName = r.RoleName
		users[i].Role.CreatedAt = r.CreatedAt
		users[i].Role.UpdatedAt = r.UpdatedAt
		users[i].Role.DeleteAt = r.DeleteAt
	}
	return
}

//添加
func (user User) UserInsert() (userId int64, err error) {

	//添加数据
	if 0 == user.RoleID {
		user.RoleID = user.Role.GetRoleIDFromRoleName(user.Role.RoleName)
	}
	result := orm.Eloquent.Create(&user)
	userId = user.UserID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//修改
func (user *User) UserUpdate(userId int64) (updateUser User, err error) {

	if err = orm.Eloquent.Select([]string{"user_id", "user_name"}).First(&updateUser, userId).Error; err != nil {
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
func (user *User) UserDestroy(userId int64) (Result User, err error) {

	if err = orm.Eloquent.Select([]string{"id"}).First(&user, userId).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Delete(&user).Error; err != nil {
		return
	}
	Result = *user
	return
}
