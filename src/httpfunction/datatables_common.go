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
)

func GetUploadTableName(tableName string) string {
	return "uploadfor" + tableName
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

func GetDataTableId(r *http.Request)(id []string, err error ){

	action := r.FormValue("action")
	id = make([]string, 0)
	if action == "create" {
		id = append(id, "0")
	}else {
		id = make([]string, 0)
		err = errors.New("failed to parse id")
		for k,v := range r.Form{
			if strings.Contains(k,"[" + dao.TableIdInJson + "]"){
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

func Getpostfile(w http.ResponseWriter, r *http.Request, tableName string) (UploadId string, err error){
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
	fileTmp.Insert(GetUploadTableName(tableName))
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
		dao.CommonLoadFromPostForm(r, id[i], res[i])
		err = dao.Insert(tableName, res[i])
		if err != nil {
			log.Println("failed to create datatables row, err: ", err.Error())
			return
		}
	}
	return
}

func Editdatatablesline_HandleFile (r *http.Request, id string, tableName string, res dao.EditDataTables) (err error){

	var fileTmp dao.Uploadfile
	fileTmp.Id, err = dao.GetFileId(tableName, res.GetId())
	NowFileId := r.FormValue("data[" + id + "][fileid]")
	if err != nil {
		log.Println("failed to get fileId in func Editdatatablesline, err: ", err.Error())
		return
	}
	if fileTmp.Id != "" && fileTmp.Id != NowFileId {
		err = fileTmp.Remove(GetUploadTableName(tableName))
		if err != nil {
			log.Println("failed to remove uploadfile in func deldatatablesline, err: ", err.Error())
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
			log.Println("failed to GetDataStruct in func Createdatatablesline, err: ", err.Error())
			return
		}
		dao.CommonLoadFromPostForm(r, id[i], res[i])
		if JudgeDataStructFileId(res[i]) {
			err = Editdatatablesline_HandleFile(r, id[i], tableName, res[i])
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
		dao.CommonLoadFromPostForm(r, id[i], res[i])
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
func HandleFilesData(tableName string, ReturnData *Datatablesdata, res []dao.EditDataTables, flag int) (err error) {
	if flag == -1 {
		return
	}else if flag == 0 {
		ReturnData.Files.Files = simplejson.New()
		FileArray := make([]dao.Uploadfile, 0)
		FileArray, err = dao.GetAllUploadfile(GetUploadTableName(tableName))
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
				err = FileOne.GetOneUploadfile(GetUploadTableName(tableName))
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

