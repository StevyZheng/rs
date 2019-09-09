package models

import (
	orm "rs/utils/datebase"
)

type User struct {
	OrmModel
	UserID   int64  `json:"user_id" gorm:"primary_key" sql:"AUTO_INCREMENT"`
	UserName string `json:"user_name" gorm:"type:varchar(64);unique_index" `
	Password string `json:"password" gorm:"type:varchar(256)"`
	Email    string `json:"email" gorm:"type:varchar(128)"`
	Role     Role   `json:"role" gorm:"auto_preload;foreignkey:RoleID"`
	RoleID   int64  `json:"role_id"`
}

func (u *User) UserGetFromName() (user User, err error) {
	if err = orm.Eloquent.Where("user_name = ?", u.UserName).First(&user).Error; err != nil {
		return
	}
	user.Role.RoleID = user.RoleID
	var r Role
	orm.Eloquent.Where("role_id = ?", user.RoleID).First(&r)
	user.Role.RoleName = r.RoleName
	user.Role.RoleDetails = r.RoleDetails
	user.Role.CreatedAt = r.CreatedAt
	user.Role.UpdatedAt = r.UpdatedAt
	user.Role.DeletedAt = r.DeletedAt
	return
}

//列表
func (u *User) UserList() (users []User, err error) {
	//orm.Eloquent.Model(&user).Related(&user.Role).Find(&user.Role)
	if err = orm.Eloquent.Find(&users).Error; err != nil {
		return
	}
	var r Role
	for i, value := range users {
		users[i].Role.RoleID = value.RoleID
		orm.Eloquent.Where("role_id = ?", users[i].RoleID).First(&r)
		users[i].Role.RoleName = r.RoleName
		users[i].Role.RoleDetails = r.RoleDetails
		users[i].Role.CreatedAt = r.CreatedAt
		users[i].Role.UpdatedAt = r.UpdatedAt
		users[i].Role.DeletedAt = r.DeletedAt
	}
	return
}

//添加
func (u User) UserInsert() (userId int64, err error) {
	//添加数据
	if 0 == u.RoleID {
		u.RoleID = u.Role.GetRoleIDFromRoleName(u.Role.RoleName)
		u.Role = Role{}
	}
	result := orm.Eloquent.Create(&u)
	userId = u.UserID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//修改
func (u *User) UserUpdate(userId int64) (updateUser User, err error) {
	if err = orm.Eloquent.Select([]string{"user_id", "user_name"}).First(&updateUser, userId).Error; err != nil {
		return
	}
	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Model(&updateUser).Updates(&u).Error; err != nil {
		return
	}
	return
}

//删除数据
func (u *User) UserDestroyFromID(userId int64) (Result User, err error) {
	if err = orm.Eloquent.Select([]string{"user_id"}).First(&u, userId).Error; err != nil {
		return
	}
	if err = orm.Eloquent.Delete(&u).Error; err != nil {
		return
	}
	Result = *u
	return
}

func (u *User) UserDestroyFromName(userName string) (Result User, err error) {
	if err = orm.Eloquent.Where("user_name = ?", userName).First(&u).Error; err != nil {
		return
	}
	if err = orm.Eloquent.Delete(&u).Error; err != nil {
		return
	}
	Result = *u
	return
}
