package main

import (
"../controller"
"fmt"
)

func main() {
	fmt.Printf("please input type(1-sync.Map; 2-Mysql):")
	fmt.Scan(&controller.Method)
	switch controller.Method {
	case 1:
		controller.InitMapstore()
	case 2:
		controller.ConnectToMysql()
	}
	controller.StartWeb()
}

