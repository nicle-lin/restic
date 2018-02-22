// init database
package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
)

type DBJsonConf struct {
	AliasName string
	DBHost    string
	DBUser    string
	DBPwd     string
	DBName    string
}

func InitMysqlDB() error {
	data, err := ioutil.ReadFile("conf/dbconf.json")
	if err != nil {
		return err
	}
	var confs []DBJsonConf
	err = json.Unmarshal(data, &confs)
	if err != nil {
		return err
	}
	for _, v := range confs {
		err = registerMysqlDB(v.AliasName, v.DBUser, v.DBPwd, v.DBHost, v.DBName)
		if err != nil {
			return err
		}
	}
	return nil
}

func registerMysqlDB(aliasName, user, password, host string, dbName string) error {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	return orm.RegisterDataBase(aliasName, "mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&loc=Local", user, password, host, dbName))
}

func init() {
	err := InitMysqlDB()
	if err != nil {
		panic(err)
	}
	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
	}

}