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
)

type Datatablesdata struct {
	Data []interface{}  `json:"data"`
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

func GetDataTableId(r *http.Request)(id []string, err error ){

	action := r.FormValue("action")
	id = make([]string, 0)
	if action == "create" {
		id = append(id, "0")
	}else {
		id = make([]string, 0)
		err = errors.New("failed to parse id")
		for k,v := range r.Form{
			if strings.Contains(k,"_id"){
				id = append(id, v[0])
				err = nil
			}
		}
		if err != nil {
			log.Println("failed to get datatables id")
		}
	}
	return
}

func Getpostfile(w http.ResponseWriter, r *http.Request) (uploadid string, err error){
	var filetmp dao.Uploadfile
	file,handler, err := r.FormFile("upload")
	if err != nil {
		log.Println("failed to get formfile")
		return
	}
	defer file.Close()

	filetmp.Filename = handler.Filename
	t, err := os.Create(UPLOAD_DIR + "/" + filetmp.Filename)
	defer t.Close()
	if err != nil {
		log.Println("fail to create file")
		return
	}
	_, err = io.Copy(t, file)
	if err != nil {
		log.Println("failed to copy file")
		return
	}

	filetmp.Systempath = UPLOAD_DIR + "/" + filetmp.Filename
	filetmp.Webpath = PreWebPath + filetmp.Filename

	fileinfo, err := os.Stat(filetmp.Systempath)
	if err != nil {
		log.Println("failed to get the file station")
		return
	}
	filetmp.Filesize = strconv.Itoa(int(fileinfo.Size()))
	filetmp.Id = strconv.Itoa(dao.GetNextId())

	//database
	filetmp.Insert()

	uploadid = filetmp.Id

	return
}

func Createdatatablesline(w http.ResponseWriter,r *http.Request, id []string, tableName string) (res []dao.EditDataTables, err error) {

	res = make([]dao.EditDataTables, len(id))
	for i := range id {
		res[i], err = dao.GetDataStruct(tableName)
		if err != nil {
			log.Println(err.Error(), " failed to GetDataStruct in func Createdatatablesline")
			return
		}
		res[i].LoadDataFromPostForm(r, id[i])
		err = dao.Insert(tableName, res[i])
		if err != nil {
			log.Println("failed to create datatables row")
			return
		}
	}

	return
}

func Editdatatablesline(w http.ResponseWriter,r *http.Request, id []string, tableName string) (res []dao.EditDataTables,err error) {

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
			log.Println("failed to get fileId in func Editdatatablesline")
			return
		}
		if filetmp.Id != "" && filetmp.Id != NowFileId{
			//if need delete local file you should have this
			//err = filetmp.LoadUploadfile()
			//if err != nil {
			//	log.Println("failed to load uploadfile in func deldatatablesline")
			//	return
			// }
			err = filetmp.Remove()
			if err != nil {
				log.Println("failed to remove uploadfile in func deldatatablesline")
				return
			}
		}

		err = dao.Update(tableName, res[i].GetId(), res[i])
		if err != nil {
			log.Println("failed to update datatables line")
			return
		}
	}
	return
}

func Deldatatablesline(w http.ResponseWriter,r *http.Request, id []string, tableName string) (err error) {

	res := make([]dao.EditDataTables,len(id))
	for i := range id {
		res[i], err = dao.GetDataStruct(tableName)
		if err != nil {
			log.Println(err.Error(), " failed to GetDataStruct in func Createdatatablesline")
			return
		}
		res[i].LoadDataFromPostForm(r, id[i])

		var filetmp dao.Uploadfile
		filetmp.Id, err = dao.GetFileId(tableName, res[i].GetId())
		if err != nil {
			log.Println("failed to get fileId in func Deldatatablesline")
			return
		}
		if filetmp.Id != ""{
			//if need delete local file you should have this
			//err = filetmp.LoadUploadfile()
			//if err != nil {
			//	log.Println("failed to load uploadfile in func deldatatablesline")
			//	return
			//}
			err = filetmp.Remove()
			if err != nil {
				log.Println("failed to remove uploadfile in func deldatatablesline")
				return
			}
		}

		err = dao.Remove(tableName, res[i].GetId())
		if err != nil {
			log.Println("failed to delete datatables line")
			return
		}
	}

	return
}


func ViewHandle(w http.ResponseWriter, r *http.Request){

	var tmp Datatablesdata
	var err error
	flag := false
	tableName := r.URL.Query().Get(URLTableName)

	if r.Method == "POST"{
		if r.ContentLength != 0 {
			action := r.FormValue("action")
			if action == "upload" {
				tmp.Upload.Id, err = Getpostfile(w, r)
				if err != nil {
					log.Println(err.Error(), " failed to action upload in func ViewHandle")
					return
				}
				flag = true
			} else {
				id, _ := GetDataTableId(r)
				if action == "create"  || action == "edit"{
					var res []dao.EditDataTables
					if action == "create" {
						res, _ = Createdatatablesline(w, r, id, tableName)
					}else {
						res, _ = Editdatatablesline(w, r, id, tableName)
					}

					for i := range res {
						//data, _:= dao.GetDataStruct(tableName)
						//dao.GetOneById(tableName, res[i].GetId(), &data)
						//fmt.Println(data)
						tmp.Data = append(tmp.Data, dao.GetOneById(tableName, res[i].GetId()))
					}
					flag = true

				}else if action == "remove"{
					Deldatatablesline(w, r, id, tableName)
				}
			}
		}
	} else if r.Method == "GET" {
		tmp.Data, err = dao.GetAll(tableName)
		if err != nil {
			log.Println(err.Error(), " failed to GET in func ViewHandle")
			return
		}
		flag = true
	}
	if flag {
		tmp.Files.Files = simplejson.New()
		res, err := dao.GetAllUploadfile()
		if err != nil {
			log.Println(err.Error(), " failed to GetAllUploadfile in func ViewHandle")
			return
		}
		for i := range res{
			tmp.Files.Files.Set(res[i].Id, res[i])
		}
	}

	t, err := json.Marshal(tmp)
	if err != nil {
		log.Println(err.Error(), " failed to marshal to json in func ViewHandle")
		return
	}
	w.Write(t)
}


