package controller

import (
	"../dataStore/mapstore"
	"LongShortUrlConvert/dataStore/Mysql"
	"github.com/gin-gonic/gin"
	"log"
)

func ConnectToMysql() {
	var err error
	Mysql.Db, err = Mysql.ConnectToMysql()
	if err != nil {
		log.Fatal("cannot connect to mysql")
	}
	log.Println("connect to mysql success")
	count = Mysql.CountNum()
}


func InitMapstore() {
	mapstore.LoadFromFile()
}

func StartWeb() {
	r := gin.Default()
	r.LoadHTMLGlob("view/*")
	r.GET("/", HandleGet)
	r.POST("/",HandlePost)
	r.GET("/f:scz",Print)
	v := r.Group("/admin")
	{
		v.GET("/getall",GetAll)
		v.GET("/deleteall", DeleteAll)
	}
	r.Run()
}