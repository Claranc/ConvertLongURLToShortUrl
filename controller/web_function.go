package controller

import (
	"../dataStore/mapstore"
	"../dataStore/Mysql"
	"../model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//网页入口
func HandleGet(c *gin.Context) {
	c.HTML(http.StatusOK, "run.html", gin.H{
	})
}

//返回存储的列表
func GetAll(c *gin.Context) {
	var resLToS string
	var resSToL string
	model.LongToShort.Range(func(k, v interface{}) bool {
		resLToS += fmt.Sprintf("longurl: %s,   shorturl: %s \n", k, v)
		return true
	})
	model.ShortToLong.Range(func(k, v interface{}) bool {
		resSToL += fmt.Sprintf("shorturl: %s,   longurl: %s \n", k, v)
		return true
	})
	c.String(200, "%s \n\n\n%s", resLToS, resSToL)
}

//清空数据
func DeleteAll(c *gin.Context) {
	if Method == 1 {
		model.LongToShort.Range(func(k, v interface{}) bool {
			mapstore.Delete(k.(string),&model.LongToShort)
			return true
		})
		model.ShortToLong.Range(func(k, v interface{}) bool {
			mapstore.Delete(k.(string), &model.ShortToLong)
			return true
		})
		mapstore.DeleteFile()
		c.String(200, "OK")
	} else if Method == 2 {
		Mysql.DeleteAll()
		count = 0
		num = 19960117
		c.String(200, "OK")
	}
}

//POST
func HandlePost(c *gin.Context) {
	longurl := c.PostForm("longurl")
	if Method == 1 {
		if len(longurl) > 0 {
			valid := CheckValidOfLongUrl(longurl)
			if valid {
				if len(longurl) > 15 {
					shorturl, exist := mapstore.Read(longurl, &model.LongToShort);
					if exist != nil {
						str := convertLongToShort(string(longurl))
						//shorturl = fmt.Sprintf("http://www.%s.com", str)
						shorturl = fmt.Sprintf("http://fengxinjie.club:8080/f" + str)
						mapstore.Create(longurl, shorturl, &model.LongToShort)
						mapstore.Create(shorturl, longurl, &model.ShortToLong)
						mapstore.WriteToFile(shorturl,longurl)
					}
					c.HTML(http.StatusOK, "run.html", gin.H{
						"output": shorturl,
						"longurl": longurl,
					})
				} else {
					c.HTML(http.StatusOK, "shortLen.html", gin.H{
						"output": "你的网址已经足够短了，来个长点的吧",
					})
				}
			} else {
				c.HTML(http.StatusOK, "shortLen.html", gin.H{
					"output": "你的网址不合法，请重新输入",
				})
			}
		} else {
			shorturl := c.PostForm("shorturl")
			longurl, err := readShortUrl(string(shorturl), &model.ShortToLong)
			if err != nil {
				c.HTML(http.StatusOK, "Longnotexist.html", gin.H{
					"output": "查询不到该短网址的信息呢",
				})
			} else {
				c.HTML(http.StatusOK, "run.html", gin.H{
					"longurl":  longurl,
					"output2": longurl,
				})
			}
		}
	} else if Method == 2 {
		if len(longurl) > 0 {
			valid := CheckValidOfLongUrl(longurl)
			if valid {
				if len(longurl) > 15 {
					Mysql.ConnectToMysql()
					shorturl,ok := Mysql.FindShortUrl(longurl)
					if ok == false {
						str := convertLongToShort(string(longurl))
						shorturl = fmt.Sprintf("http://fengxinjie.club:8080/f" + str)
						Mysql.InsertValue(shorturl,longurl)
					}
					c.HTML(http.StatusOK, "run.html", gin.H{
						"output": shorturl,
						"longurl": longurl,
					})
				} else {
					c.HTML(http.StatusOK, "shortLen.html", gin.H{
						"output": "你的网址已经足够短了，来个长点的吧",
					})
				}

			} else {
				c.HTML(http.StatusOK, "shortLen.html", gin.H{
					"output": "你的网址不合法，请重新输入",
				})
			}
		} else {
			shorturl := c.PostForm("shorturl")
			longurl, ok := Mysql.FindLongUrl(shorturl)
			if ok == false {
				c.HTML(http.StatusOK, "Longnotexist.html", gin.H{
					"output": "查询不到该短网址的信息呢",
				})
			} else {
				c.HTML(http.StatusOK, "run.html", gin.H{
					"longurl": longurl,
					"output2": longurl,
				})
			}
		}
	}
}

func Print(c *gin.Context) {
	str := c.Param("scz")
	shorturl := "http://fengxinjie.club:8080/f"+str
	switch Method {
	case 1:
		longurl,err := mapstore.Read(shorturl, &model.ShortToLong)
		if err != nil {
			c.String(200, "长网址不存在")
		}
		c.Redirect(302, longurl)
	case 2:
		Mysql.ConnectToMysql()
		longurl,OK := Mysql.FindLongUrl(shorturl)
		if OK == false {
			c.String(200, "长网址不存在")
		}
		c.Redirect(302, longurl)
	}


}