package httpfunction

import "github.com/bitly/go-simplejson"

import (
	"net/http"
	"dao"
	"encoding/json"
	"errors"
	"strings"
	"strconv"
	"log"
	"os"
	"io"
	"reflect"
	"fmt"
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

const UPLOAD_DIR = "/home/zh/GoPro/firstRedis/html/uploads"
const PreWebPath = "/uploads/"
const URLTableName = "tableName"
//var mutexlock sync.Mutex

func GetDataTableId(r *http.Request)(id []string, err error ){

	action := r.FormValue("action")
	id = make([]string, 0)
	if action == "create" {
		id = append(id, "0")
	}else {
		id = make([]string, 0)
		err = errors.New("failed to parse id")
		for k,v := range r.Form{
			if strings.Contains(k,"[id]"){
				id = append(id, v[0])
				err = nil
			}
		}
		if err != nil {
			log.Println("failed to get datatables id, err: ", err.Error())
		}
	}
	return
}

func Getpostfile(w http.ResponseWriter, r *http.Request) (UploadId string, err error){
	var fileTmp dao.Uploadfile
	file,handler, err := r.FormFile("upload")
	if err != nil {
		log.Println("failed to get formfile, err: ", err.Error())
		return
	}
	defer file.Close()

	fileTmp.Filename = handler.Filename
	t, err := os.Create(UPLOAD_DIR + "/" + fileTmp.Filename)
	defer t.Close()
	if err != nil {
		log.Println("fail to create file, err: ", err.Error())
		return
	}
	_, err = io.Copy(t, file)
	if err != nil {
		log.Println("failed to copy file, err: ", err.Error())
		return
	}

	fileTmp.Systempath = UPLOAD_DIR + "/" + fileTmp.Filename
	fileTmp.Webpath = PreWebPath + fileTmp.Filename

	fileInfo, err := os.Stat(fileTmp.Systempath)
	if err != nil {
		log.Println("failed to get the file station, err: ", err.Error())
		return
	}
	fileTmp.Filesize = strconv.Itoa(int(fileInfo.Size()))
	fileTmp.Id = strconv.Itoa(dao.GetNextId())
	fileTmp.Insert()
	UploadId = fileTmp.Id
	return
}

func Createdatatablesline(w http.ResponseWriter,r *http.Request, id []string, tableName string) (res []dao.EditDataTables, err error) {

	//mutexlock.Lock()
	//defer mutexlock.Unlock()
	res = make([]dao.EditDataTables, len(id))
	for i := range id {
		res[i], err = dao.GetDataStruct(tableName)
		if err != nil {
			log.Println("failed to GetDataStruct in func Createdatatablesline, err: ", err.Error())
			return
		}
		res[i].LoadDataFromPostForm(r, id[i])
		err = dao.Insert(tableName, res[i])
		if err != nil {
			log.Println("failed to create datatables row, err: ", err.Error())
			return
		}
	}

	return
}

func Editdatatablesline(w http.ResponseWriter,r *http.Request, id []string, tableName string) (res []dao.EditDataTables,err error) {

	//mutexlock.Lock()
	//defer mutexlock.Unlock()
	res = make([]dao.EditDataTables, len(id))
	for i := range id {
		res[i], err = dao.GetDataStruct(tableName)
		if err != nil {
			log.Println(err.Error(), " failed to GetDataStruct in func Createdatatablesline")
			return
		}
		res[i].LoadDataFromPostForm(r, id[i])

		var filetmp dao.Uploadfile
		filetmp.Id, err = dao.GetFileId(tableName, res[i].GetId())
		NowFileId := r.FormValue("data[" + id[i] + "][fileid]")
		if err != nil {
			log.Println("failed to get fileId in func Editdatatablesline, err: ", err.Error())
			return
		}
		if filetmp.Id != "" && filetmp.Id != NowFileId{
			//if need delete local file you should have this
			//err = filetmp.GetOneUploadfile()
			//if err != nil {
			//	log.Println("failed to load uploadfile in func deldatatablesline")
			//	return
			// }
			err = filetmp.Remove()
			if err != nil {
				log.Println("failed to remove uploadfile in func deldatatablesline, err: ", err.Error())
				return
			}
		}

		err = dao.Update(tableName, res[i].GetId(), res[i])
		if err != nil {
			log.Println("failed to update datatables line, err: ", err.Error())
			return
		}
	}
	return
}

func Deldatatablesline(w http.ResponseWriter,r *http.Request, id []string, tableName string) (err error) {

	//mutexlock.Lock()
	//defer mutexlock.Unlock()
	res := make([]dao.EditDataTables,len(id))
	for i := range id {
		res[i], err = dao.GetDataStruct(tableName)
		if err != nil {
			log.Println("failed to GetDataStruct in func Createdatatablesline, err: ", err.Error())
			return
		}
		res[i].LoadDataFromPostForm(r, id[i])

		var filetmp dao.Uploadfile
		filetmp.Id, err = dao.GetFileId(tableName, res[i].GetId())
		if err != nil {
			log.Println("failed to get fileId in func Deldatatablesline, err: ", err.Error())
			return
		}
		if filetmp.Id != ""{
			//if need delete local file you should have this
			//err = filetmp.GetOneUploadfile()
			//if err != nil {
			//	log.Println("failed to load uploadfile in func deldatatablesline")
			//	return
			//}
			err = filetmp.Remove()
			if err != nil {
				log.Println("failed to remove uploadfile in func deldatatablesline, err: ", err.Error())
				return
			}
		}

		err = dao.Remove(tableName, res[i].GetId())
		if err != nil {
			log.Println("failed to delete datatables line, err: ", err.Error())
			return
		}

	}

	return
}

//flag is -1 to handle remove, flag is 0 to handle GET and upload, flag is 1 to handle edit and create
func HandleFilesData(tableName string, ReturnData *Datatablesdata, res []dao.EditDataTables, flag int) (err error) {
	if flag == -1 {
		return
	}else if flag == 0 {
		ReturnData.Files.Files = simplejson.New()
		FileArray := make([]dao.Uploadfile, 0)
		FileArray, err = dao.GetAllUploadfile()
		if err != nil {
			log.Println("failed to GetAllUploadfile in func ViewHandle, err: ", err.Error())
			return
		}
		for i := range FileArray{
			ReturnData.Files.Files.Set(FileArray[i].Id, FileArray[i])
		}
	}else if flag == 1 {
		var FileOne dao.Uploadfile
		ReturnData.Files.Files = simplejson.New()
		for i := range res {
			FileOne.Id, err = dao.GetFileId(tableName, res[i].GetId())
			if err != nil {
				log.Println("failed to get fileId in func Deldatatablesline, err: ", err.Error())
				return
			}
			if FileOne.Id != "" {
				err = FileOne.GetOneUploadfile()
				if err != nil {
					log.Println("failed to get one uploadfile, err: ", err.Error())
					return
				}
				ReturnData.Files.Files.Set(FileOne.Id, FileOne)
			}
		}
	}else {
		err = errors.New("flag is not found")
		return
	}
	return
}

func ViewHandle(w http.ResponseWriter, r *http.Request){

	var ReturnData Datatablesdata
	var res []dao.EditDataTables
	var err error
	var flag int
	tableName := r.URL.Query().Get(URLTableName)

	if r.Method == "POST"{
		if r.ContentLength != 0 {
			action := r.FormValue("action")
			if action == "upload" {
				ReturnData.Upload.Id, err = Getpostfile(w, r)
				if err != nil {
					log.Println("failed to action upload in func ViewHandle, err: ", err.Error())
					return
				}
				flag = 0
			} else {
				id, _ := GetDataTableId(r)
				if action == "create"  || action == "edit"{
					if action == "create" {
						res, _ = Createdatatablesline(w, r, id, tableName)
					}else {
						res, _ = Editdatatablesline(w, r, id, tableName)
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
					ReturnData.Data = dataslice
					flag = 1

				}else if action == "remove"{
					Deldatatablesline(w, r, id, tableName)
					flag = -1
				}
			}
		}
	} else if r.Method == "GET" {

		ReturnData.Data, err = dao.GetDataStructSilce(tableName)
		if err != nil {
			log.Println("falied to get datastruct silce, err: ", err.Error())
			return
		}
		err = dao.GetAll(tableName, ReturnData.Data)
		if err != nil {
			log.Println("failed to get all data, err: ",err.Error())
			return
		}

		if reflect.ValueOf(ReturnData.Data).Elem().Len() == 0 {
			fmt.Println("ReturnData.Data is empty")
			ReturnData.Data = make([]interface{}, 0)
		}
		flag = 0
	}

	HandleFilesData(tableName, &ReturnData, res, flag)

	ReturnDataJson, err := json.Marshal(ReturnData)
	if err != nil {
		log.Println("failed to marshal to json in func ViewHandle, err: ", err.Error())
		return
	}
	w.Write(ReturnDataJson)
}




