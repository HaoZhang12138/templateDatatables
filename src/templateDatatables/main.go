package main

import (
	"net/http"
	"log"
	"httpfunction"
	"os"
	"dao"
)

//程序运行时，会自动运行，判断存放上传文件的本地目录是否存在，如不存在则新建
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
	http.HandleFunc("/defaultview", httpfunction.DefaultViewHandle)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("listen and serve: ", err.Error())
	}
}
