package httpfunction

import (
	"net/http"
	"dao"
	"encoding/json"
	"log"
	"reflect"
	"fmt"
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
//var mutexlock sync.Mutex

func ViewHandle(w http.ResponseWriter, r *http.Request){

	var returndata Datatablesdata
	var res []dao.EditDataTables
	var err error
	var flag int
	tableName := r.URL.Query().Get(URLTableName)

	if r.Method == "POST"{
		if r.ContentLength != 0 {
			action := r.FormValue("action")
			if action == "upload" {
				returndata.Upload.Id, err = Getpostfile(r,tableName)
				if err != nil {
					log.Println("failed to action upload in func ViewHandle, err: ", err.Error())
					return
				}
				flag = 0
			} else {
				id, _ := GetDataTableId(r, tableName)
				if action == "create"  || action == "edit"{
					if action == "create" {
						res, _ = Createdatatablesline(r, tableName, id)
					}else {
						res, _ = Editdatatablesline(r, tableName, id)
					}
					dataslice := make([]interface{}, 0)
					for i := range res {
						data, err := dao.GetDataStruct(tableName)
						if err != nil {
							log.Println("failed to get datastruct, err: ", err.Error())
							return
						}
						err = dao.GetOneById(tableName, res[i].GetId(), data)
						if err != nil {
							log.Println("failed to get one data by id, err: ", err.Error())
							return
						}
						dataslice = append(dataslice, data)
					}
					returndata.Data = dataslice
					flag = 1

				}else if action == "remove"{
					Deldatatablesline(r, tableName, id)
					flag = -1
				}

			}
		}
	} else if r.Method == "GET" {

		returndata.Data, err = dao.GetDataStructSilce(tableName)
		if err != nil {
			log.Println("falied to get datastruct silce, err: ", err.Error())
			return
		}
		err = dao.GetAll(tableName, returndata.Data)
		if err != nil {
			log.Println("failed to get all data, err: ",err.Error())
			return
		}

		if reflect.ValueOf(returndata.Data).Elem().Len() == 0 {
			fmt.Println("ReturnData.Data is empty")
			returndata.Data = make([]interface{}, 0)
		}
		flag = 0
	}

	useToJudge, err := dao.GetDataStruct(tableName)
	if err != nil {
		log.Println("failed to get data struct, err: ", err.Error())
		return
	}
	if JudgeDataStructFileId(useToJudge) {
		HandleFilesData(tableName, &returndata, res, flag)
	}

	response, err := json.Marshal(returndata)
	if err != nil {
		log.Println("failed to marshal to json, err: ", err.Error())
		return
	}
	w.Write(response)
}




