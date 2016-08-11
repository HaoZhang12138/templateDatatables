package dao

import (
	"net/http"
	"strconv"
)


type UserInfoDatatables struct {
	User string `json:"user"`
	Pass string `json:"pass"`
	Name string `json:"name"`
	Sex string  `json:"sex"`
	Age string  `json:"age"`
	Tel string  `json:"tel"`
	Id string `json:"id" bson:"_id"`
	FileId string `json:"fileid"`
}


func (this *UserInfoDatatables)LoadDataFromPostForm(r *http.Request,id string){

	this.User = r.FormValue("data[" + id + "][user]")
	this.Pass = r.FormValue("data[" + id + "][pass]")
	this.Name = r.FormValue("data[" + id + "][name]")
	this.Age  = r.FormValue("data[" + id + "][age]")
	this.Tel  = r.FormValue("data[" + id + "][tel]")
	this.Sex  = r.FormValue("data[" + id + "][sex]")
	this.Id   = r.FormValue("data[" + id + "][id]")
	this.FileId = r.FormValue("data[" + id + "][fileid]")

	if this.Id == "" {
		this.Id = strconv.Itoa(GetNextId())
	}
	return
}

func (this *UserInfoDatatables)GetId()(string){
	return  this.Id
}
