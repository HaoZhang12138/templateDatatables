package dao

import (
	"log"
	"errors"
)

type test struct {
	Name string `json:"name"`
	Age string `json:"age"`
}

func (this *test)GetId()(string){
	return  this.Name
}

func GetDataStruct(tableName string) (ret EditDataTables,err error) {
	switch tableName {
	case "userinfo":
		ret = new(UserinfoDatatables)
	case "test":
		ret = new(test)

	default:
		err = errors.New("tablename is not found")
		log.Println("failed to match tablename in func GetDataStruct, err: ", err.Error())
	}
	return
}

func GetDataStructSilce(tableName string) (ret interface{},err error) {
	switch tableName {
	case "userinfo":
		ret = new([]UserinfoDatatables)
	default:
		err = errors.New("tablename is not found")
		log.Println("failed to match tablename in func GetDataStruct, err: ", err.Error())
	}
	return
}
