package dao

import (
	"log"
	"errors"
)

func GetDataStruct(tableName string) (ret EditDataTables,err error) {
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
		log.Println("failed to match tablename in func GetDataStruct, err: ", err.Error())
	}
	return
}

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
		log.Println("failed to match tablename in func GetDataStruct, err: ", err.Error())
	}
	return
}
