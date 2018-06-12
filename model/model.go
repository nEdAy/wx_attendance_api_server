package model

import (
	"fmt"
	"time"

	"github.com/google/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nEdAy/wx_attendance_api_server/config"
	"log"
)

// 数据库操作对象
var DB *gorm.DB

type Model struct {
	Id         int   `gorm:"column:id;primary_key" json:"id"`
	CreateTime int64 `gorm:"column:create_time" json:"create_time"`
	//CreatedOn  int `json:"created_on"`
	//ModifiedOn int `json:"modified_on"`
	//DeletedOn  int `json:"deleted_on"`
}

// 初始化数据库
func Setup() {
	logger.Infoln("正在与数据库建立连接...")
	// 连接数据库
	var err error
	DB, err = gorm.Open(config.Database.Type,
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			config.Database.User,
			config.Database.Password,
			config.Database.Host,
			config.Database.Port,
			config.Database.Name))
	if err != nil {
		log.Fatalln(err)
	}
	// 设置表名前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.Database.TablePrefix + defaultTableName
	}
	// 禁用表名多元化
	DB.SingularTable(true)
	// 是否开启debug模式
	if config.Database.Debug {
		// debug 模式
		DB = DB.Debug()
	}
	/*	DB.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
		DB.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
		DB.Callback().Delete().Replace("gorm:delete", deleteCallback)*/
	// 连接池最大连接数
	DB.DB().SetMaxIdleConns(config.Database.MaxIdleConns)
	// 默认打开连接数
	DB.DB().SetMaxOpenConns(config.Database.MaxOpenConns)

	// 开启协程ping MySQL数据库查看连接状态
	go func() {
		for {
			// ping
			err = DB.DB().Ping()
			if err != nil {
				logger.Infoln(err)
			}
			// 间隔30s ping一次
			time.Sleep(config.Database.PingInterval)
		}
	}()
	logger.Infoln("与数据库建立连接成功!")
}

// 关闭连接
func Close() {
	if DB != nil {
		defer DB.Close()
	}
}

/*
// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
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
*/
