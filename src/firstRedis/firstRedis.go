package main

import (
	"net/http"
	"log"
//	"httpfunction"
	"handlehtml"
	"httpfunction"
	"os"
)
const UPLOAD_DIR = "/home/zh/GoPro/firstRedis/html/uploads"
func init() {
	handlehtml.Inithtml()

	_, err := os.Stat(UPLOAD_DIR)
	if err != nil {
		err = os.Mkdir(UPLOAD_DIR, os.ModePerm)
		if err != nil {
			log.Println("failed to create dir")
		}
	}
}

func main() {


	http.Handle("/", http.FileServer(http.Dir("./html")))
	//http.Handle("/", http.FileServer(http.Dir("./uploads")))

	http.HandleFunc("/main", httpfunction.MainHandle)
	http.HandleFunc("/view", httpfunction.ViewHandle)
	http.HandleFunc("/add" , httpfunction.AddHandle)
	http.HandleFunc("/list", httpfunction.ListHandle)
	http.HandleFunc("/login", httpfunction.LoginHandle)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("listen and serve: ", err.Error())
	}
}
