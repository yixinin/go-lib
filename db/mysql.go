package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

type MysqlConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
	Prefix   string `yaml:"prefix"`
}

var Mysql *xorm.Engine

func InitMysql(conf *MysqlConfig) {
	var err error
	Mysql, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", conf.User, conf.Password, conf.Host, conf.Port, conf.DB))
	if err != nil {
		panic(err)
	}
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, conf.Prefix)
	Mysql.SetTableMapper(tbMapper)
	// SyncTable()

}

func SyncTable(s ...interface{}) {
	Mysql.Sync2(s...)
}
