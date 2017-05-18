package main

import (

	"net/http"
	//"strings"
	"log"
	//"strings"
	//"time"
	//"crypto/md5"

	//"strconv"
	//"go/types"

)


func main() {
	//http.HandleFunc("/", sayhelloName) //设置访问的路由
	http.HandleFunc("/",index)
	http.HandleFunc("/api/",index2)
	err := http.ListenAndServe(":9999", nil) //设置监听的端口请访问http://localhost:9999/
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}



}







