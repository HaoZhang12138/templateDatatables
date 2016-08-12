package main

import (
	"net/http"
	"log"
	"httpfunction"
	"os"
)
const UPLOAD_DIR = "/home/zh/GoPro/templateDatatables/html/uploads"
func init() {
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
	http.HandleFunc("/view", httpfunction.ViewHandle)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("listen and serve: ", err.Error())
	}
}
