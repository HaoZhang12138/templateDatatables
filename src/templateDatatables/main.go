package main

import (
	"net/http"
	"log"
	"httpfunction"
	"os"
	"dao"
)

func init() {
	_, err := os.Stat(dao.UPLOAD_DIR)
	if err != nil {
		err = os.Mkdir(dao.UPLOAD_DIR, os.ModePerm)
		if err != nil {
			log.Println("failed to create dir")
		}
	}
}

func main() {

	http.Handle("/", http.FileServer(http.Dir("./html")))
	http.HandleFunc("/view", httpfunction.ViewHandle)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("listen and serve: ", err.Error())
	}
}
