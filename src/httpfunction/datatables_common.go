package httpfunction

import (
	"reflect"
	"net/http"
	"errors"
	"strings"
	"log"
	"dao"
	"os"
	"io"
	"strconv"
	"github.com/bitly/go-simplejson"
	"fmt"
)

func GetUploadTableName(tableName string) string {
	return  tableName + "_uploadfile"
}

func JudgeDataStructFileId(data interface{}) (flag bool) {
	flag = false
	v := reflect.ValueOf(data).Elem()
	for i := 0; i < v.NumField(); i++ {
		name := v.Type().Field(i).Name
		if name == "FileId" {
			flag = true
			break
		}
	}
	return
}

func GetDataTableId(r *http.Request, tableName string)(id []string, err error ){

	action := r.FormValue("action")
	id = make([]string, 0)
	if action == "create" {
		id = append(id, "0")
	}else {
		var idInJson string
		idInJson, err = dao.GetTableIdInJson(tableName)
		if err != nil {
			log.Println("failed to GetTableIdInJson, err: ", err.Error())
			return
		}
		id = make([]string, 0)
		err = errors.New("failed to parse id")
		for k,v := range r.Form{
			if strings.Contains(k,"[" + idInJson + "]"){
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

func Getpostfile(r *http.Request, tableName string) (uploadid string, err error){
	var fileTmp dao.Uploadfile
	file,handler, err := r.FormFile("upload")
	if err != nil {
		log.Println("failed to get formfile, err: ", err.Error())
		return
	}
	defer file.Close()

	fileTmp.Filename = handler.Filename
	t, err := os.Create(dao.UPLOAD_DIR + "/" + fileTmp.Filename)
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

	fileTmp.Systempath = dao.UPLOAD_DIR + "/" + fileTmp.Filename
	fileTmp.Webpath = dao.PreWebPath + fileTmp.Filename

	fileInfo, err := os.Stat(fileTmp.Systempath)
	if err != nil {
		log.Println("failed to get the file station, err: ", err.Error())
		return
	}
	fileTmp.Filesize = strconv.Itoa(int(fileInfo.Size()))
	fileTmp.Id = strconv.Itoa(dao.GetNextId())
	fileTmp.Insert(GetUploadTableName(tableName))
	uploadid = fileTmp.Id
	return
}

func Createdatatablesline(r *http.Request, tableName string, id []string) (res []dao.EditDataTables, err error) {

	//mutexlock.Lock()
	//defer mutexlock.Unlock()
	res = make([]dao.EditDataTables, len(id))
	for i := range id {
		res[i], err = dao.GetDataStruct(tableName)
		if err != nil {
			log.Println("failed to GetDataStruct in func Createdatatablesline, err: ", err.Error())
			return
		}
		dao.CommonLoadFromPostForm(r, tableName, id[i], res[i])
		err = dao.Insert(tableName, res[i])
		if err != nil {
			log.Println("failed to create datatables row, err: ", err.Error())
			return
		}
	}
	return
}

func Editdatatablesline_HandleFile (r *http.Request, tableName string, id string, res dao.EditDataTables) (err error){

	var fileTmp dao.Uploadfile
	fileTmp.Id, err = dao.GetFileId(tableName, res.GetId())
	newfileid := r.FormValue("data[" + id + "][fileid]")
	if err != nil {
		log.Println("failed to get fileId in func Editdatatablesline, err: ", err.Error())
		return
	}
	if fileTmp.Id != "" && fileTmp.Id != newfileid {
		err = fileTmp.Remove(GetUploadTableName(tableName))
		if err != nil {
			log.Println("failed to remove uploadfile in func deldatatablesline, err: ", err.Error())
			return
		}
	}
	return
}

func Editdatatablesline(r *http.Request, tableName string, id []string) (res []dao.EditDataTables,err error) {

	//mutexlock.Lock()
	//defer mutexlock.Unlock()
	res = make([]dao.EditDataTables, len(id))
	for i := range id {
		res[i], err = dao.GetDataStruct(tableName)
		if err != nil {
			log.Println("failed to GetDataStruct in func Createdatatablesline, err: ", err.Error())
			return
		}
		dao.CommonLoadFromPostForm(r, tableName, id[i], res[i])
		if JudgeDataStructFileId(res[i]) {
			err = Editdatatablesline_HandleFile(r, tableName, id[i], res[i])
			if err != nil {
				log.Println("failed to handlefile in func Editdatatablesline, err: ", err.Error())
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

func Deldatatablesline_HandleFile (tableName string, res dao.EditDataTables) (err error){

	var fileTmp dao.Uploadfile
	fileTmp.Id, err = dao.GetFileId(tableName, res.GetId())
	if err != nil {
		log.Println("failed to get fileId in func Deldatatablesline, err: ", err.Error())
		return
	}
	if fileTmp.Id != ""{

		err = fileTmp.Remove(GetUploadTableName(tableName))
		if err != nil {
			log.Println("failed to remove uploadfile in func deldatatablesline, err: ", err.Error())
			return
		}
	}
	return
}

func Deldatatablesline(r *http.Request, tableName string, id []string) (err error) {

	//mutexlock.Lock()
	//defer mutexlock.Unlock()
	res := make([]dao.EditDataTables,len(id))
	for i := range id {
		res[i], err = dao.GetDataStruct(tableName)
		if err != nil {
			log.Println("failed to GetDataStruct in func Createdatatablesline, err: ", err.Error())
			return
		}
		dao.CommonLoadFromPostForm(r, tableName, id[i], res[i])
		if JudgeDataStructFileId(res[i]) {
			err = Deldatatablesline_HandleFile(tableName, res[i])
			if err != nil {
				log.Println("failed to handlefile in func Deldatatablesline, err: ", err.Error())
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
func HandleFilesData(tableName string, returndata *Datatablesdata, res []dao.EditDataTables, flag int) (err error) {
	fmt.Println("using func HandleFilesData")
	if flag == -1 {
		return
	}else if flag == 0 {
		returndata.Files.Files = simplejson.New()
		filearray := make([]dao.Uploadfile, 0)
		filearray, err = dao.GetAllUploadfile(GetUploadTableName(tableName))
		if err != nil {
			log.Println("failed to GetAllUploadfile in func ViewHandle, err: ", err.Error())
			return
		}
		for i := range filearray{
			returndata.Files.Files.Set(filearray[i].Id, filearray[i])
		}
	}else if flag == 1 {
		var fileone dao.Uploadfile
		returndata.Files.Files = simplejson.New()
		for i := range res {
			fileone.Id, err = dao.GetFileId(tableName, res[i].GetId())
			if err != nil {
				log.Println("failed to get fileId in func Deldatatablesline, err: ", err.Error())
				return
			}
			if fileone.Id != "" {
				err = fileone.GetOneUploadfile(GetUploadTableName(tableName))
				if err != nil {
					log.Println("failed to get one uploadfile, err: ", err.Error())
					return
				}
				returndata.Files.Files.Set(fileone.Id, fileone)
			}
		}
	}else {
		err = errors.New("flag is not found")
		return
	}
	return
}

