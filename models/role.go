package models

import (
	"rs/utils/datebase"
)

type Role struct {
	OrmModel
	RoleID      int64  `json:"role_id" gorm:"primary_key;unique_index;AUTO_INCREMENT"`
	RoleName    string `json:"role_name" gorm:"type:varchar(32);unique_index"`
	RoleDetails string `json:"role_details"`
}

func (r *Role) RoleGetFromName() (role Role, err error) {
	if err = datebase.DB.Where("role_name = ?", r.RoleName).First(&role).Error; err != nil {
		return
	}
	return
}

func (r *Role) RoleList() (roles []Role, err error) {
	if err = datebase.DB.Find(&roles).Error; err != nil {
		return
	}
	return
}

func (r Role) GetRoleNameFromRoleID(roleId int64) (roleName string, err error) {
	if err = datebase.DB.Where("role_id = ?", roleId).First(&r).Error; err != nil {
		return
	}
	roleName = r.RoleName
	return
}

func (r *Role) GetRoleIDFromRoleName(roleName string) (roleId int64, err error) {
	if err = datebase.DB.Where("role_name = ?", roleName).First(&r).Error; err != nil {
		return
	}
	roleId = r.RoleID
	return
}

func (r Role) RoleInsert() (roleId int64, err error) {
	//添加数据
	if err = datebase.DB.Create(&r).Error; err != nil {
		return
	}
	roleId = r.RoleID
	return
}

//修改成r
func (r *Role) RoleUpdate(roleId int64) (updateRole Role, err error) {
	if err = datebase.DB.Where("role_id = ?", roleId).First(&updateRole).Error; err != nil {
		return
	}
	if err = datebase.DB.Model(&updateRole).Updates(&r).Error; err != nil {
		return
	}
	updateRole = *r
	return
}

//删除数据
func (r *Role) RoleDestroy(roleId int64) (Result Role, err error) {
	r.RoleID = roleId
	if err = datebase.DB.Delete(&r).Error; err != nil {
		return
	}
	return
}
