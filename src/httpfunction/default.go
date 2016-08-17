package httpfunction

import (
	"net/http"
	"log"
	"dao"
	"reflect"
)

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

func Create_default(w http.ResponseWriter, r *http.Request, resp *Datatablesdata) (err error) {
	tableName, id := Common(r)
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

func Edit_default(w http.ResponseWriter, r *http.Request, resp *Datatablesdata) (err error)  {
	tableName, id := Common(r)
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

func Remove_default(w http.ResponseWriter, r *http.Request, resp *Datatablesdata) (err error) {
	tableName, id := Common(r)
	err = Deldatatablesline(r, tableName, id)
	if err != nil {
		log.Println("failed to remove line, err: ",err.Error())
		return
	}
	return
}

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

func Common(r *http.Request)(tableName string, id []string){
	tableName = r.URL.Query().Get(URLTableName)
	id, _ = GetDataTableId(r, tableName)
	return
}

func Common_create_edit(tableName string, res []dao.DataTablesDao, rdata *Datatablesdata)(err error) {
	dataslice := make([]interface{}, 0)
	for i := range res {
		data, err := dao.GetDataStruct(tableName)
		if err != nil {
			log.Println("failed to get data struct, err: ", err.Error())
			return err
		}
		err = dao.GetOneById(tableName, res[i].GetId(), data)
		if err != nil {
			log.Println("failed to get one data by id, err: ", err.Error())
			return err
		}
		dataslice = append(dataslice, data)
	}
	rdata.Data = dataslice
	return
}

func Common_HandleFile(tableName string, res []dao.DataTablesDao, flag int, rdata *Datatablesdata)(err error){
	useToJudge, err := dao.GetDataStruct(tableName)
	if err != nil {
		log.Println("failed to get data struct, err: ", err.Error())
		return
	}
	if JudgeDataStructFileId(useToJudge) {
		err = HandleFilesData(tableName, rdata, res, flag)
		if err != nil {
			log.Println("failed to HandleFilesData, err: ", err.Error())
			return
		}
	}
	return
}
