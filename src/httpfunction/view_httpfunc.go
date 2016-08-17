package httpfunction

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/bitly/go-simplejson"
)

type Datatablesdata struct {
	Data interface{}  `json:"data"`
	Files uploadfile_tmp `json:"files,omitempty"`
	Upload uploadid `json:"upload,omitempty"`
}

type uploadid struct {
	Id string `json:"id"`
}

type uploadfile_tmp struct {
	Files *simplejson.Json `json:"files"`
}

const URLTableName = "tableName"

func ViewHandle(w http.ResponseWriter, r *http.Request){

	var resp Datatablesdata
	var err error
	tableName := r.URL.Query().Get(URLTableName)
	handler, err := GetDatatablesHandler(tableName)
	if err != nil {
		log.Println("failed to get handler, err: ", err.Error())
		return
	}
	if r.Method == "POST"{
		action := r.FormValue("action")
		if action == "upload" {
			err = handler.Upload(w, r, &resp)
		} else if action == "create" {
			err = handler.Create(w, r, &resp)
		}else if action == "edit" {
			err = handler.Edit(w, r, &resp)
		} else if action == "remove"{
			err = handler.Remove(w, r, &resp)
		}
		if err != nil {
			log.Println("failed to action: ", action, " err: ", err.Error())
			return
		}
	} else if r.Method == "GET" {
		err = handler.GET(w, r, &resp)
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





