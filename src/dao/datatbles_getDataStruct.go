package dao

import (
	"log"
	"errors"
)

func GetDataStruct(tableName string) (ret EditDataTables,err error) {
	switch tableName {
	case "userinfo":
		ret = new(UserInfoDatatables)
		break
	default:
		err = errors.New("tablename is not found")
		log.Println("failed to match tablename in func GetDataStruct, err: ", err.Error())
	}
	return
}

func GetDataStructSilce(tableName string) (ret interface{},err error) {
	switch tableName {
	case "userinfo":
		ret = new([]UserInfoDatatables)
		break
	default:
		err = errors.New("tablename is not found")
		log.Println("failed to match tablename in func GetDataStruct, err: ", err.Error())
	}
	return
}
