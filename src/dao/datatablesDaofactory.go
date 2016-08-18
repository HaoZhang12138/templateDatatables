package dao

import (
	"log"
	"errors"
)

//自定义结构体， 在此处添加关于结构体的case
func GetDataStruct(tableName string) (ret DataTablesDao,err error) {
	switch tableName {
	case "userinfo":
		ret = new(UserinfoDatatables)
	case "course":
		ret = new(CourseDatatables)
	case "pets":
		ret = new(PetsDatatables)
	default:
		err = errors.New("tablename is not found")
		log.Println("failed to match tablename in func GetDataStruct, err: ", err.Error())
	}
	return
}

//自定义结构体， 在此处添加关于结构体数组的case
func GetDataStructSilce(tableName string) (ret interface{},err error) {
	switch tableName {
	case "userinfo":
		ret = new([]UserinfoDatatables)
	case "course":
		ret = new([]CourseDatatables)
	case "pets":
		ret = new([]PetsDatatables)
	default:
		err = errors.New("tablename is not found")
		log.Println("failed to match tablename in func GetDataStructSilce, err: ", err.Error())
	}
	return
}

//自定义结构体， 在此处添加关于结构体主键的case
func GetTableIdInJson(tableName string)(idInJson string, err error) {
	switch tableName {
	case "userinfo":
		idInJson = "id"
	case "course":
		idInJson = "courseid"
	case "pets":
		idInJson = "petid"
	default:
		err = errors.New("tablename is not found")
		log.Println("failed to match tablename in func GetTableIdInJson, err: ", err.Error())
	}
	return
}
