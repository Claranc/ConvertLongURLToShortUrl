package controller

import (
	"../dataStore/mapstore"
	"../model"
	"log"
	"strings"
	"sync"
)


//将长网址处理成短网址的函数
func convertLongToShort(longurl string) string {
	if res,err := mapstore.Read(longurl,&model.LongToShort); err == nil {
		return res
	}
	mu.Lock()
	output := convertTenToOtherJinzhi(num+count, jinzhi)
	num++
	mu.Unlock()
	return output
}

//十进制数转n进制数的函数
func convertTenToOtherJinzhi(num int, jinzhi int) string {
	if num < 0 {
		log.Fatal("wrong input")
	}
	var res string
	if num == 0 {
		return string(code62[0])
	}
	for(num > 0) {
		res = string(code62[num%jinzhi]) + res
		num /= jinzhi
	}
	return res
}

func readShortUrl(shorturl string, db *sync.Map) (string, error) {
	longurl,err := mapstore.Read(shorturl, db)
	return  longurl, err
}


//验证输入的长网址是否合法
func CheckValidOfLongUrl(longurl string) bool {
	var flag1 bool = strings.HasPrefix(longurl, validhead[0]) || strings.HasPrefix(longurl,validhead[1])
	if flag1 == true {
		for _,v := range longurl {
			if (v >= '0' && v <= '9') || (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || v == '-' ||
				v == '%' || v == '&' || v == '?' || v == '_' || v == '.' || v == ':' || v == '/'{

			} else {
				flag1 = false
				break
			}
		}
	}
	return flag1
}
