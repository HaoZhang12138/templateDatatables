package main

import (
	"net/http"
	"log"
	"httpfunction"
	"handlehtml"
//	"dao"
)

func init() {
	handlehtml.Inithtml()
//	dao.InitRedis()
}

func main() {
	
	http.HandleFunc("/main", httpfunction.MainHandle)
	http.HandleFunc("/view", httpfunction.ViewHandle)
	http.HandleFunc("/add" , httpfunction.AddHandle)
	http.HandleFunc("/list", httpfunction.ListHandle)
	http.HandleFunc("/", httpfunction.LoginHandle)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("listen and serve: ", err.Error())
	}
}
