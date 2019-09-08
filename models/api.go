package models

import (
	orm "rs/utils/datebase"
)

type Api struct {
	OrmModel
	ApiID  int64  `json:"api_id"`
	Url    string `json:"url"`
	Method string `json:"method"`
	Param  string `json:"param"`
	Return string `json:"return"`
}

func (a *Api) ApiList() (apis []Api, err error) {
	if err = orm.Eloquent.Find(&apis).Error; err != nil {
		return
	}
	return
}

func (a Api) ApiInsert() (apiId int64, err error) {
	//添加数据
	result := orm.Eloquent.Create(&a)
	apiId = a.ApiID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

func (a *Api) ApiUpdate(apiId int64) (updateApi Api, err error) {

	if err = orm.Eloquent.Select([]string{"api_id", "url"}).First(&updateApi, apiId).Error; err != nil {
		return
	}
	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Model(&updateApi).Updates(&a).Error; err != nil {
		return
	}
	return
}

//删除数据
func (a *Api) ApiDestroy(apiId int64) (Result Api, err error) {

	if err = orm.Eloquent.Select([]string{"role_id"}).First(&a, apiId).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Delete(&a).Error; err != nil {
		return
	}
	Result = *a
	return
}
