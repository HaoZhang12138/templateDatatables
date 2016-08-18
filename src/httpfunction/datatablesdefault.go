package httpfunction

import (
	"net/http"
	"log"
	"dao"
	"reflect"
)
//默认的上传方法
func Upload_default(w http.ResponseWriter, r *http.Request, resp *Datatablesdata)(err error) {
	tableName := r.URL.Query().Get(URLTableName)
	resp.Upload.Id, err = Getpostfile(r,tableName)
	if err != nil {
		log.Println("failed to action upload in func ViewHandle, err: ", err.Error())
		return
	}
	err = Common_HandleFile(tableName, nil, FILES_NEEDED_ALL, resp)
	if err != nil {
		log.Println("failed to handle file in upload, err: ", err.Error())
		return
	}
	return
}

//默认的创建方法
func Create_default(w http.ResponseWriter, r *http.Request, resp *Datatablesdata) (err error) {
	tableName, id, err := Get_tableName_id(r)
	res, err := Createdatatablesline(r, tableName, id)
	if err != nil {
		log.Println("failed to create a new line, err: ", err.Error())
		return
	}
	err = Common_create_edit(tableName, res, resp)
	if err != nil {
		log.Println("failed to add data to Datatablesdata, err: ", err.Error())
		return
	}
	err = Common_HandleFile(tableName, res, FILES_NEEDED_ONE, resp)
	if err != nil {
		log.Println("failed to handle file in create, err: ", err.Error())
		return
	}
	return
}

//默认的编辑方法
func Edit_default(w http.ResponseWriter, r *http.Request, resp *Datatablesdata) (err error)  {
	tableName, id ,err := Get_tableName_id(r)
	res, err := Editdatatablesline(r, tableName, id)
	if err != nil {
		log.Println("failed to edit line, err: ", err.Error())
		return
	}
	err = Common_create_edit(tableName, res, resp)
	if err != nil {
		log.Println("failed to add data to Datatablesdata, err: ", err.Error())
		return
	}
	err = Common_HandleFile(tableName, res, FILES_NEEDED_ONE, resp)
	if err != nil {
		log.Println("failed to handle file in edit, err: ", err.Error())
		return
	}
	return
}

//默认的删除方法
func Remove_default(w http.ResponseWriter, r *http.Request, resp *Datatablesdata) (err error) {
	tableName, id ,err := Get_tableName_id(r)
	err = Deldatatablesline(r, tableName, id)
	if err != nil {
		log.Println("failed to remove line, err: ",err.Error())
		return
	}
	resp.Data, err = dao.GetDataStructSilce(tableName)
	if err != nil {
		log.Println("falied to get datastruct silce, err: ", err.Error())
		return
	}

	return
}

//默认的GET方法
func GET_default(w http.ResponseWriter, r *http.Request, resp *Datatablesdata) (err error){
	tableName := r.URL.Query().Get(URLTableName)
	resp.Data, err = dao.GetDataStructSilce(tableName)
	if err != nil {
		log.Println("falied to get datastruct silce, err: ", err.Error())
		return
	}
	err = dao.GetAll(tableName, resp.Data)
	if err != nil {
		log.Println("failed to get all data, err: ",err.Error())
		return
	}
	if reflect.ValueOf(resp.Data).Elem().Len() == 0 {
		log.Println("response is empty")
		resp.Data = make([]interface{}, 0)
	}
	err = Common_HandleFile(tableName, nil, FILES_NEEDED_ALL, resp)
	if err != nil {
		log.Println("failed to handle file in GET, err: ", err.Error())
		return
	}

	return
}

//用于获取请求的表名和主键
func Get_tableName_id(r *http.Request)(tableName string, id []string, err error){
	tableName = r.URL.Query().Get(URLTableName)
	id, err = GetDataTableId(r, tableName)
	if err != nil {
		log.Println("failed to get post data Id, err: ", err.Error())
		return
	}
	return
}

//用于创建和编辑之后，加载要回应的数据
func Common_create_edit(tableName string, res []dao.DataTablesDao, resp *Datatablesdata)(err error) {

	resp.Data, err = dao.GetDataStructSilce(tableName)
	if err != nil{
		log.Println("failed to get data struct slice, err: ", err.Error())
		return
	}
	var ids []interface{}
	for i := range res {
		ids = append(ids, res[i].GetId())
	}
	err = dao.GetDataByIdSlice(tableName, ids, resp.Data)
	if err != nil{
		log.Println("failed to get data by id array, err: ", err.Error())
		return
	}
	return
}

//用于除删除外的所有请求之后， 判断是否要加载文件信息，是则在函数内加载
func Common_HandleFile(tableName string, res []dao.DataTablesDao, flag int, resp *Datatablesdata)(err error){
	useToJudge, err := dao.GetDataStruct(tableName)
	if err != nil {
		log.Println("failed to get data struct, err: ", err.Error())
		return
	}
	if JudgeDataStructFileId(useToJudge) {
		err = HandleFilesData(tableName, resp, res, flag)
		if err != nil {
			log.Println("failed to HandleFilesData, err: ", err.Error())
			return
		}
	}
	return
}
