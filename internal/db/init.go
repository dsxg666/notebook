package db

import (
	"database/sql"
	"fmt"
	"github.com/dsxg666/notebook/global"
	"github.com/dsxg666/notebook/pkg/setting"
	_ "github.com/go-sql-driver/mysql"
)

func initMysql(dbSetting *setting.DatabaseSettingS) *sql.DB {
	dbConn, err := sql.Open(dbSetting.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t",
		dbSetting.UserName,
		dbSetting.Password,
		dbSetting.Host,
		dbSetting.DBName,
		dbSetting.Charset,
		dbSetting.ParseTime,
	))
	if err != nil {
		global.SugarLogger.Error("sql.Open err:", err)
		return nil
	}
	// 最大连接数
	dbConn.SetMaxOpenConns(global.DatabaseSetting.MaxOpenConns)
	// 闲置连接数
	dbConn.SetMaxIdleConns(global.DatabaseSetting.MaxIdleConns)
	return dbConn
}

func InitDB(dbSetting *setting.DatabaseSettingS) *sql.DB {
	return initMysql(dbSetting)
}