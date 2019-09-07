package models

import (
	orm "rs/utils/datebase"
)

type Role struct {
	OrmModel
	RoleID   int64  `json:"id" gorm:"primary_key;unique_index"`
	RoleName string `json:"role_name" gorm:"type:varchar(32);unique_index;index:idx_name_code"`
}

func (r *Role) RoleList() (roles []Role, err error) {
	if err = orm.Eloquent.Find(&roles).Error; err != nil {
		return
	}
	return
}

func (r Role) GetRoleNameFromRoleID(role_id int64) (role_name string) {
	orm.Eloquent.Where("role_id = ?", role_id).First(&r)
	return r.RoleName
}

func (r Role) GetRoleIDFromRoleName(roleName string) (role_id int64) {
	orm.Eloquent.Where("role_name = ?", roleName).First(&r)
	return r.RoleID
}

func (r Role) RoleInsert() (role_id int64, err error) {
	//添加数据
	result := orm.Eloquent.Create(&r)
	role_id = r.RoleID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//修改
func (r *Role) RoleUpdate(role_id int64) (updateRole Role, err error) {

	if err = orm.Eloquent.Select([]string{"role_id", "role_name"}).First(&updateRole, role_id).Error; err != nil {
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
func (r *Role) RoleDestroy(role_id int64) (Result Role, err error) {

	if err = orm.Eloquent.Select([]string{"role_id"}).First(&r, role_id).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Delete(&r).Error; err != nil {
		return
	}
	Result = *r
	return
}
