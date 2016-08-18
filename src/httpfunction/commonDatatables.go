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

//公共的从postform读取数据的方法
func CommonLoadFromPostForm(req *http.Request,tableName string, id string,ptrdata interface{})(err error)  {
	// parse form
	if err = req.ParseForm();err !=nil{
		log.Println("failed to parseform, err: ",err.Error())
		return
	}

	v := reflect.ValueOf(ptrdata).Elem()
	fields := make(map[string]reflect.Value)
	for i := 0; i < v.NumField(); i++ {
		prefix := "data["+id+"]["
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		jsonName := tag.Get("json")
		if jsonName == ""{
			jsonName = strings.ToLower(fieldInfo.Name)
		}
		prefix = prefix+jsonName+"]"
		fields[prefix] = v.Field(i)
	}

	// take from form
	for name, values := range req.Form {
		f,found := fields[name]
		if !found{
			continue
		}
		for _,value := range values {
			err = populate(f,value)
			if err != nil {
				log.Println("failed to populate, err: ",err.Error())
				return
			}
		}
	}
	idInJson, err := dao.GetTableIdInJson(tableName)
	if err != nil {
		log.Println("failed to GetTableIdInJson, err: ", err.Error())
		return
	}
	createPrefix := "data[0][" + idInJson + "]"
	if createId, exist := fields[createPrefix]; exist{
		autoId := strconv.Itoa(dao.GetNextId())
		populate(createId,autoId)
	}
	return
}

//根据反射的类型设置值
func populate(v reflect.Value, value string) error  {
	switch v.Kind() {
	case reflect.String:
		if v.CanSet(){
			v.SetString(value)
		}else {
			return errors.New("string field can't be set")
		}

	case reflect.Int:
		i,err := strconv.ParseInt(value,10,64)
		if err != nil {
			log.Println("failed to pasrse to int, err: ",err.Error())
			return err
		}
		if v.CanSet(){
			v.SetInt(i)
		}else {
			return errors.New("int field can't be set")
		}

	case reflect.Int64:
		i,err := strconv.ParseInt(value,10,64)
		if err != nil {
			log.Println("failed to pasrse to int, err :",err.Error())
			return err
		}
		if v.CanSet(){
			v.SetInt(i)
		}else {
			return errors.New("int64 field can't be set")
		}

	case reflect.Bool:
		b,err := strconv.ParseBool(value)
		if err != nil {
			log.Println("failed to pasrse to bool, err: ",err.Error())
			return err
		}
		if v.CanSet(){
			v.SetBool(b)
		}else {
			return errors.New("bool field can't be set")
		}
	case reflect.Float32:
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Println("failed to parse to float32, err: ", err.Error())
			return err
		}
		if v.CanSet() {
			v.SetFloat(f)
		}else {
			return  errors.New("float32 field can't be set")
		}
	case reflect.Float64:
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Println("failed to parse to float64, err: ", err.Error())
			return err
		}
		if v.CanSet() {
			v.SetFloat(f)
		}else {
			return  errors.New("float64 field can't be set")
		}

	default:
		return fmt.Errorf("unsupported kind %s",v.Type())
	}
	return nil
}

//得到存放上传文件信息的数据库表的名字
func GetUploadTableName(tableName string) string {
	return  tableName + "_uploadfile"
}

//判断struct有无上传文件字段
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

//从postform中得到数据的主键字段值
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

//接受上传的文件，存在本地，如对上传文件的存放有特殊要求，请改写
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

//处理创建请求
func Createdatatablesline(r *http.Request, tableName string, id []string) (res []dao.DataTablesDao, err error) {

	res = make([]dao.DataTablesDao, len(id))
	for i := range id {
		res[i], err = dao.GetDataStruct(tableName)
		if err != nil {
			log.Println("failed to GetDataStruct in func Createdatatablesline, err: ", err.Error())
			return
		}
		err = CommonLoadFromPostForm(r, tableName, id[i], res[i])
		if err != nil {
			log.Println("failed to load data from post form, err: ", err.Error())
			return
		}
		err = dao.Insert(tableName, res[i])
		if err != nil {
			log.Println("failed to create datatables row, err: ", err.Error())
			return
		}
	}
	return
}

//如有文件，处理编辑时对文件的请求
func Editdatatablesline_HandleFile (r *http.Request, tableName string, id string, res dao.DataTablesDao) (err error){

	var fileTmp dao.Uploadfile
	fileTmp.Id, err = dao.GetFileId(tableName, res.GetId())
	newFileId := r.FormValue("data[" + id + "][fileid]")
	if err != nil {
		log.Println("failed to get fileId in func Editdatatablesline, err: ", err.Error())
		return
	}
	if fileTmp.Id != "" && fileTmp.Id != newFileId {
		err = fileTmp.Remove(GetUploadTableName(tableName))
		if err != nil {
			log.Println("failed to remove uploadfile in func deldatatablesline, err: ", err.Error())
			return
		}
	}
	return
}

//处理编辑请求
func Editdatatablesline(r *http.Request, tableName string, id []string) (res []dao.DataTablesDao,err error) {

	res = make([]dao.DataTablesDao, len(id))
	for i := range id {
		res[i], err = dao.GetDataStruct(tableName)
		if err != nil {
			log.Println("failed to GetDataStruct in func Createdatatablesline, err: ", err.Error())
			return
		}
		err = CommonLoadFromPostForm(r, tableName, id[i], res[i])
		if err != nil {
			log.Println("failed to load data from post form, err: ", err.Error())
			return
		}
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

//如有文件，删除时对其进行处理
func Deldatatablesline_HandleFile (tableName string, res dao.DataTablesDao) (err error){

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

//处理删除请求
func Deldatatablesline(r *http.Request, tableName string, id []string) (err error) {

	res := make([]dao.DataTablesDao,len(id))
	for i := range id {
		res[i], err = dao.GetDataStruct(tableName)
		if err != nil {
			log.Println("failed to GetDataStruct in func Createdatatablesline, err: ", err.Error())
			return
		}
		err = CommonLoadFromPostForm(r, tableName, id[i], res[i])
		if err != nil {
			log.Println("failed to load data from post form, err: ", err.Error())
			return
		}
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

//回复给前端的数据中， 如有文件， 加上相应的文件信息
func HandleFilesData(tableName string, returndata *Datatablesdata, res []dao.DataTablesDao, flag int) (err error) {

	if flag == FILES_NOT_NEEDED{
		return
	}else if flag == FILES_NEEDED_ALL {
		returndata.Files.Files = simplejson.New()
		fileArray := make([]dao.Uploadfile, 0)
		fileArray, err = dao.GetAllUploadfile(GetUploadTableName(tableName))
		if err != nil {
			log.Println("failed to GetAllUploadfile in func ViewHandle, err: ", err.Error())
			return
		}
		for i := range fileArray{
			returndata.Files.Files.Set(fileArray[i].Id, fileArray[i])
		}
	}else if flag == FILES_NEEDED_ONE {
		var fileOne dao.Uploadfile
		returndata.Files.Files = simplejson.New()
		for i := range res {
			fileOne.Id, err = dao.GetFileId(tableName, res[i].GetId())
			if err != nil {
				log.Println("failed to get fileId in func Deldatatablesline, err: ", err.Error())
				return
			}
			if fileOne.Id != "" {
				err = fileOne.GetOneUploadfile(GetUploadTableName(tableName))
				if err != nil {
					log.Println("failed to get one uploadfile, err: ", err.Error())
					return
				}
				returndata.Files.Files.Set(fileOne.Id, fileOne)
			}
		}
	}else {
		err = errors.New("flag is not found")
		return
	}
	return
}