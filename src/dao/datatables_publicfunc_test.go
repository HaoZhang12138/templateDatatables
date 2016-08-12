package dao

import (
	"testing"
	"log"
	"encoding/json"
	"reflect"
)

func TestGetOneById(t *testing.T) {
	tmp := UserInfoDatatables{}
	tmp.Id = "123456"
	tmp.Name = "test"

	err := Insert("userinfo", tmp)
	if err != nil {
		t.Error("failed to insert %v", err.Error())
		return
	}
	tt , err := GetDataStruct("userinfo")
	if err != nil {
		t.Error("failed to get datastruct %v", err.Error())
		//return
	}

	err = GetOneById("userinfo", tmp.Id, tt)
	j, _ := json.Marshal(tt)
	log.Println(string(j))

	err = Remove("userinfo",tmp.Id)
	if err != nil {
		t.Error("failed to remove %v", err.Error())
		return
	}
}


func TestCommonLoadFromPostForm(t *testing.T) {
	tmp, _ := GetDataStruct("userinfo")
	flag := false
	v := reflect.ValueOf(tmp).Elem()
	for i := 0; i < v.NumField(); i++ {

		name := v.Type().Field(i).Name
		if name == "FileId" {
			flag = true
			break
		}
	}
	log.Println(flag)

	flag = false
	tmp , _ = GetDataStruct("test")
	v = reflect.ValueOf(tmp).Elem()
	for i := 0; i < v.NumField(); i++ {

		name := v.Type().Field(i).Name
		if name == "FileId" {
			flag = true
			break
		}
	}
	log.Println(flag)

}