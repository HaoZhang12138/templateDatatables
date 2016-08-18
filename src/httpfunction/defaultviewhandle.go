package httpfunction

import (
	"net/http"
	"log"
	"encoding/json"
)

//如果不想自定义handler， 可使用此默认handler进行处理，可通过设置ajax的url属性来调用
func DefaultViewHandle(w http.ResponseWriter, r *http.Request){
	var resp Datatablesdata
	var err error
	if r.Method == "POST"{
		action := r.FormValue("action")
		if action == "upload" {
			err = Upload_default(w, r, &resp)
		} else if action == "create" {
			err = Create_default(w, r, &resp)
		}else if action == "edit" {
			err = Edit_default(w, r, &resp)
		} else if action == "remove"{
			err = Remove_default(w, r, &resp)
		}
		if err != nil {
			log.Println("failed to action: ", action, " err: ", err.Error())
			return
		}
	} else if r.Method == "GET" {
		err = GET_default(w, r, &resp)
		if err != nil {
			log.Println("failed to Excute GET, err :", err.Error())
			return
		}
	}

	response, err := json.Marshal(resp)
	if err != nil {
		log.Println("failed to marshal to json, err: ", err.Error())
		return
	}
	w.Write(response)
}


