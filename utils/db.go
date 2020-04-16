package utils

import (
	"fmt"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
	"xorm.io/xorm/names"
)

var Engine *xorm.Engine

func DbInit() error {
	connectStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		Config.Database.Host,
		Config.Database.Port,
		Config.Database.User,
		Config.Database.DbName,
		Config.Database.Pwd,
	)
	var err error
	Engine, err = xorm.NewEngine("postgres", connectStr)
	if err != nil {
		return err
	}

	//SetMaxIdleConns用于设置闲置的连接数
	Engine.DB().SetMaxIdleConns(5)
	//SetMaxOpenConns用于设置最大打开的连接数
	//Engine.DB().SetMaxOpenConns(100)

	// 启用Logger，显示详细日志
	if Config.Mode == "debug" {
		Engine.ShowSQL(true)
		Engine.Logger().SetLevel(log.LOG_DEBUG)
	}

	Engine.SetMapper(names.GonicMapper{})

	Engine.Sync2()
	return nil
}
