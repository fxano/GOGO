package models

import (
	"github.com/jinzhu/gorm"
	"os"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)
var(
	db*gorm.DB
)
func init(){
	var err error
	if err=os.MkdirAll("data",0777,);err!=nil{
		panic("错误"+err.Error())
	}
	db,err=gorm.Open("sqlite3","data/data.db")
	if err!=nil{
		panic("连接数据库失败")
	}
	////defer db.Close()
}