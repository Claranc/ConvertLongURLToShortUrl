package mapstore

import (
	"../../model"
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)


//CREATE
func Create(key,value string, db *sync.Map) {
	db.Store(key,value)
}

//READ
func Read(key string, db *sync.Map) (string, error) {
	s,ok := db.Load(key)
	if ok == false {
		return "",errors.New("not exist")
	}
	return s.(string),nil
}

//DELETE
func Delete(key string, db *sync.Map) {
	db.Delete(key)
}


//运行时加载落盘的KV
func LoadFromFile() {
	s2l, err := os.OpenFile(s2lName,  os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	if err != nil {
		log.Println("cannot open the file of shortToLong")
	}
	scanner := bufio.NewScanner(s2l)
	count := 0
	for scanner.Scan() {
		temp := string(scanner.Text())
		if len(temp) == 0 {
			break
		}
		count++
		if str:= strings.Split(temp, " "); len(str) != 2 && len(str) != 0 {
			panic("context of s2l error")
		} else {
			Create(str[0], str[1], &model.ShortToLong)
			Create(str[1], str[0], &model.LongToShort)
		}
	}
	model.Num += count
	defer func() {
		s2l.Close()
	}()
}

//写入文件里备份
func WriteToFile(shorturl, longurl string) {
	s2l, err := os.OpenFile("S2L.dat", os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	if err != nil {
		log.Println("connot open the file of shortToLong")
	}
	writer1 := bufio.NewWriter(s2l)
	str1 := fmt.Sprintf("%s %s\n", shorturl, longurl)
	writer1.WriteString(str1)
	writer1.Flush()
	defer func() {
		s2l.Close()
	}()
}

//删除文件
func DeleteFile() {
	s2l, err := os.OpenFile("S2L.dat", os.O_RDWR|os.O_CREATE|os.O_TRUNC,0644)
	if err != nil {
		log.Println("connot open the file of shortToLong")
	}
	defer s2l.Close()
}
