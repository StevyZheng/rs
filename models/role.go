package models

import (
	orm "rs/utils/datebase"
)

type Role struct {
	OrmModel
	RoleID      int64  `json:"role_id" gorm:"primary_key"  sql:"AUTO_INCREMENT"`
	RoleName    string `json:"role_name" gorm:"type:varchar(32);unique_index"`
	RoleDetails string `json:"role_details"`
}

func (r *Role) RoleGetFromName() (role Role, err error) {
	if err = orm.Eloquent.Where("role_name = ?", r.RoleName).First(&role).Error; err != nil {
		return
	}
	return
}

func (r *Role) RoleList() (roles []Role, err error) {
	if err = orm.Eloquent.Find(&roles).Error; err != nil {
		return
	}
	return
}

func (r Role) GetRoleNameFromRoleID(roleId int64) (roleName string) {
	orm.Eloquent.Where("role_id = ?", roleId).First(&r)
	return r.RoleName
}

func (r *Role) GetRoleIDFromRoleName(roleName string) (roleId int64) {
	orm.Eloquent.Where("role_name = ?", roleName).First(r)
	return r.RoleID
}

func (r Role) RoleInsert() (roleId int64, err error) {
	//添加数据
	result := orm.Eloquent.Create(&r)
	roleId = r.RoleID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//修改
func (r *Role) RoleUpdate(roleId int64) (updateRole Role, err error) {

	if err = orm.Eloquent.Select([]string{"role_id", "role_name"}).First(&updateRole, roleId).Error; err != nil {
		return
	}
	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Model(&updateRole).Updates(&r).Error; err != nil {
		return
	}
	return
}

//删除数据
func (r *Role) RoleDestroy(roleId int64) (Result Role, err error) {

	if err = orm.Eloquent.Select([]string{"role_id"}).First(&r, roleId).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Delete(&r).Error; err != nil {
		return
	}
	Result = *r
	return
}
