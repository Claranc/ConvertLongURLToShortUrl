package mysql

import (
	"LongShortUrlConvert/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db = &gorm.DB{}
var Count = 0

//连接到数据库
func ConnectToMysql() (*gorm.DB, error) {
	var err error
	var str string = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8",name, passwd, mysqlAddr, databaseName)
	Db,err = gorm.Open("mysql", str)
	return Db,err
}

//读出当前数据表里有多少条数据
func CountNum() int {
	Db.Find(&model.Shorttolong{}).Count(&Count)
	return Count
}

//创建数据表
func initTable() {
	Db.Set("gorm:table_options","collate=utf8_bin").CreateTable(&model.Shorttolong{})
}

//插入一条长短网址记录
func InsertValue(shorturl, longurl string) {
	//ConnectToMysql()
	cb := model.Shorttolong{ Shorturl:shorturl, Longurl:longurl}
	Db.Model(&model.Shorttolong{}).Create(&cb)
}

//删除该条短网址记录
func DeleteValue(shorturl string) {
	//ConnectToMysql()
	Db.Where("shorturl = ?", shorturl).Delete(&model.Shorttolong{})
}

//从长网址查询短网址
func FindShortUrl(longurl string) (string, bool) {
	//ConnectToMysql()
	res := &model.Shorttolong{}
	Db.Where("longurl = ?", longurl).First(res)
	if len(res.Shorturl) == 0 {
		return "", false
	}
	return res.Shorturl, true
}

//从短网址查询长网址
func FindLongUrl(shorturl string) (string, bool) {
	//ConnectToMysql()
	res := &model.Shorttolong{}
	Db.Where("shorturl = ?", shorturl).First(res)
	if len(res.Longurl) == 0 {
		return "", false
	}
	return res.Longurl,true
}

//清空数据表
func DeleteAll() {
	ConnectToMysql()
	Db.DropTable(&model.Shorttolong{})
	initTable()
	Count = 0
}