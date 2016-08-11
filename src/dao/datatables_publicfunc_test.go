package dao

import (
	"testing"
	"log"
	"encoding/json"
)

func TestGetOneById(t *testing.T) {
	tmp := UserInfoDatatables{}
	tmp.Id = "123456"
	tmp.Name = "test"

	err := Insert("userinfo", tmp)
	if err != nil {
		log.Println("failed to insert %v", err.Error())
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
		log.Println("failed to remove %v", err.Error())
		return
	}
}
