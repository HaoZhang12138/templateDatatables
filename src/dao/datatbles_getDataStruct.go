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
		log.Println(err.Error(), " failed to match tablename in func GetDataStruct")
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
		log.Println(err.Error(), " failed to match tablename in func GetDataStruct")
	}
	return
}
