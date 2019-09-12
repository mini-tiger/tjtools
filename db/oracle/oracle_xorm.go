package oracle

import (
	_ "github.com/mattn/go-oci8"
	"github.com/go-xorm/xorm"
	"os"
)

func NewEngine(dsn string,sqllog string) (engine *xorm.Engine, err error) {
	engine, err = xorm.NewEngine("oci8", dsn)
	if err != nil {
		return
	}
	engine.SetMaxOpenConns(20)
	engine.SetMaxIdleConns(10)
	err = engine.Ping()

	f, err := os.Create(sqllog)
	if err != nil {
		//println(err.Error())
		return
	}
	xorm.NewSimpleLogger(f) //todo 将SQL语句写到日志
	return
}
