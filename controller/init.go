package controller

import (
	"../dataStore/mapstore"
	"../dataStore/Mysql"
	"github.com/gin-gonic/gin"
	"log"
)

//连接到数据库
func ConnectToMysql() {
	var err error
	Mysql.Db, err = Mysql.ConnectToMysql()
	if err != nil {
		log.Fatal("cannot connect to mysql")
	}
	log.Println("connect to mysql success")
	count = Mysql.CountNum()
}

//加载文件
func InitMapstore() {
	mapstore.LoadFromFile()
}

//运行web端
func StartWeb() {
	r := gin.Default()
	r.LoadHTMLGlob("../view/*")
	r.GET("/", HandleGet)
	r.POST("/",HandlePost)
	r.GET("/f:scz",JumpToLongURL)
	v := r.Group("/admin")
	{
		v.GET("/getall",GetAll)
		v.GET("/deleteall", DeleteAll)
	}
	r.Run()
}