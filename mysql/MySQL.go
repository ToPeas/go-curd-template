package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github/ToPeas/go-curd-templatepkg/setting"
	"os"
	"time"
	"xorm.io/core"

	"log"
)

var engine *xorm.Engine

func Setup() {
	var err error
	engine, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s", setting.Config.Database.Username, setting.Config.Database.Password, setting.Config.Database.Host, setting.Config.Database.Port, setting.Config.Database.DBName, setting.Config.Database.Charset))
	if err != nil {
		panic(fmt.Errorf("xorm: %w", err))
	}

	engine.TZLocation = setting.Config.App.TZLocation

	engine.ShowSQL(true)
	if setting.Config.App.Debug {
		engine.Logger().SetLevel(core.LOG_DEBUG)
	} else {
		_ = os.Mkdir("logs", 0777)
		f, err := os.Create("logs/xorm.log")
		if err != nil {
			log.Fatalln(err.Error())
		}
		engine.SetLogger(xorm.NewSimpleLogger(f))
	}

	engine.SetTableMapper(core.GonicMapper{})
	engine.SetColumnMapper(core.GonicMapper{})

	engine.SetMaxIdleConns(10)
	engine.SetMaxOpenConns(100)
	engine.SetConnMaxLifetime(60 * time.Second)
	//engine.Sync2(User{}, App{})
}
