package db

import (
	"encoding/json"
	"fmt"
	"xorm.io/core"

	_ "github.com/go-sql-driver/mysql"
	//"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var g_dbHand []*xorm.Engine

type DBList struct {
	DBUser string `json:"db_user"`
	DBHome string `json:"db_home"`
	DBPort uint32 `json:"db_port"`
	DBName string `json:"db_name"`
	DBPass string `json:"db_pass"`
}

func UseHand(nIndex int) *xorm.Engine {
	if len(g_dbHand) >= nIndex {
		return g_dbHand[nIndex]
	}
	return nil
}

func init() {
	var dblist []DBList
	data, err := ReadConfig("./config/db.json")
	if err != nil {
		panic("读取文件错误:" + err.Error())
	}
	json.Unmarshal(data, &dblist)
	fmt.Println("MYSQL:连接参数:", dblist)
	for _, db := range dblist {
		strConn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db.DBUser,
			db.DBPass,
			db.DBHome,
			db.DBPort,
			db.DBName)
		//("mysql", "root:feidianDB*@#4@tcp(172.16.250.198:3306)/thp")
		dbHand, err := xorm.NewEngine("mysql", strConn)
		if err != nil {
			fmt.Println("Web database link failed", err)
			panic(err.Error())
		}
		if err := dbHand.Ping(); err != nil {
			fmt.Println("Test link failed", err)
			panic(err.Error())
		}
		dbHand.SetTableMapper(core.SnakeMapper{})
		dbHand.SetColumnMapper(core.SnakeMapper{})
		dbHand.ShowSQL(true)
		dbHand.SetMaxIdleConns(5)
		dbHand.SetMaxOpenConns(5)
		g_dbHand = append(g_dbHand, dbHand)
	}
	fmt.Println("Web database links are successful")
}
