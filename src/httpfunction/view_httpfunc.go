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

type uploadid struct {
	Id string `json:"id"`
}

type ret struct {
	Data []dao.UserLoginInfo  `json:"data"`
	Files uploadfile_tmp `json:"files,omitempty"`
	Upload uploadid `json:"upload,omitempty"`
}

type uploadfile_tmp struct {
	Files *simplejson.Json `json:"files"`
}

const UPLOAD_DIR = "/home/zh/GoPro/firstRedis/html/uploads"
const PreWebPath = "/uploads/"

func loadUserInfoForDatatables(r *http.Request,id string)(ret dao.EditDataTables){

	var info *dao.UserLoginInfo
	info = new(dao.UserLoginInfo)
	info.User = r.FormValue("data[" + id + "][user]")
	info.Pass = r.FormValue("data[" + id + "][pass]")
	info.Name = r.FormValue("data[" + id + "][name]")
	info.Age  = r.FormValue("data[" + id + "][age]")
	info.Tel  = r.FormValue("data[" + id + "][tel]")
	info.Sex  = r.FormValue("data[" + id + "][sex]")
	info.Id   = r.FormValue("data[" + id + "][id]")
	info.FileId = r.FormValue("data[" + id + "][file]")

	if info.Id == "" {
		info.Id = strconv.Itoa(dao.GetNextId())
	}
	ret = info
	return
}


func GetDataTableId(r *http.Request)(id string, err error ){

	action := r.FormValue("action")
	if action == "create" {
		id = "0"
	}else {
		err = errors.New("failed to parse id")
		for k,v := range r.Form{
			if strings.Contains(k,"id"){
				id = v[0]
				err = nil
				break
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


func Createdatatablesline(w http.ResponseWriter,r *http.Request, id string) (res dao.EditDataTables, err error) {
	res = loadUserInfoForDatatables(r, id)
	res.Insert()
	if err != nil {
		log.Println("failed to create datatables row")
	}

	return
}

func Editdatatablesline(w http.ResponseWriter,r *http.Request, id string) (res dao.EditDataTables,err error) {

	err = Deldatatablesline(w,r,id)
	if err != nil {
		log.Println("failed to delete datatables row")
		return
	}

	res = loadUserInfoForDatatables(r,id)
	err = res.Insert()
	if err != nil {
		log.Println("failed to insert datatables row")
		return
	}

	return
}


func Deldatatablesline(w http.ResponseWriter,r *http.Request, id string) (err error) {

	res := loadUserInfoForDatatables(r, id)

	var filetmp dao.Uploadfile
	filetmp.Id, err = res.GetfileId()
	if err != nil {
		log.Println("failed to get fileId from userlogininfo")
		return
	}
	if filetmp.Id != "" {
		err = filetmp.LoadUploadfile()
		check_error(err)
		err = filetmp.Remove()
		check_error(err)
	}

	err = res.Remove()
	if err != nil {
		log.Println("failed to delete datatables line")
	}
	return
}



func ViewHandle(w http.ResponseWriter, r *http.Request){


	var tmp ret
	var err error
	if r.Method == "POST"{
		if r.ContentLength != 0 {
			action := r.FormValue("action")
			if action == "upload" {
				tmp.Upload.Id, err = Getpostfile(w, r)
				check_error(err)

			} else {
				id, _ := GetDataTableId(r)
				if action == "create"  || action == "edit"{
					var res dao.EditDataTables
					if action == "create" {
						res, _ = Createdatatablesline(w,r,id)
					}else {
						res, _ = Editdatatablesline(w,r,id)
					}

					data, ok := res.GetOneById().(dao.UserLoginInfo)
					if ok {
						tmp.Data = append(tmp.Data, data)
					} else {
						log.Println("failed to type assertion")
						return
					}

				}else if action == "remove"{
					Deldatatablesline(w,r,id)
				}
			}
		}
	} else if r.Method == "GET" {
		tmp.Data, err = dao.GetAllUserinfo()
		check_error(err)
	}

	tmp.Files.Files = simplejson.New()
	res, err := dao.GetAllUploadfile()
	check_error(err)
	len := len(res)
	for i := 0; i < len; i++ {
		tmp.Files.Files.Set(res[i].Id, res[i])
	}

	t, err := json.Marshal(tmp)
	check_error(err)
	w.Write(t)
}


