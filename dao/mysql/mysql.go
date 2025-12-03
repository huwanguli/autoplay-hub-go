package mysql

import (
	"autoplay-hub/settings"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// db小写不对外开放，保证对mysql数据库的操作只在mysql中进行
var db *sqlx.DB

func Init(mc *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
		mc.User,
		mc.Password,
		mc.Host,
		mc.Port,
		mc.DBName,
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect mysql failed", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(mc.MaxOpenConns)
	db.SetMaxIdleConns(mc.MaxIdleConns)
	return
}

func Close() {
	_ = db.Close()
}
