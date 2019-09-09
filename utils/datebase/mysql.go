package datebase

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
	"sync"
	"time"
)

type MysqlConnectPool struct {
}

var instance *MysqlConnectPool
var once sync.Once

var Eloquent *gorm.DB

func GetInstance() *MysqlConnectPool {
	once.Do(func() {
		instance = &MysqlConnectPool{}
	})
	return instance
}

func init() {
	var err error
	Eloquent, err = gorm.Open("mysql", "root:000000@tcp(127.0.0.1:3306)/rs?charset=utf8&parseTime=True&loc=Local&timeout=50ms")

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if Eloquent.Error != nil {
		fmt.Printf("database error %v", Eloquent.Error)
	}

	Eloquent.DB().SetMaxIdleConns(1000)
	Eloquent.DB().SetMaxOpenConns(10240)
	Eloquent.DB().SetConnMaxLifetime(2 * time.Hour)
	Eloquent.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	Eloquent.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	Eloquent.Callback().Delete().Replace("gorm:delete", deleteCallback)
}

func (m *MysqlConnectPool) GetMysqlDB() (dbCon *gorm.DB) {
	return dbCon
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := gorm.NowFunc()
		if createTimeField, ok := scope.FieldByName("CreatedAt"); ok {
			if createTimeField.IsBlank {
				_ = createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("UpdatedAt"); ok {
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(nowTime)
			}
		}
	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("UpdatedAt", gorm.NowFunc())
	}
}

// 注册删除钩子在删除之前
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		//deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedAt")
		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedTempAt")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(gorm.NowFunc()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
