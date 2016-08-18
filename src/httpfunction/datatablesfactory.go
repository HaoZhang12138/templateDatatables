package httpfunction

import (
	"errors"
	"log"
)

//如果自定义了handler， 在此处添加case
func GetDatatablesHandler(tableName string)(ret DataTableHandler, err error)  {
	switch tableName {
	case "pets":
		ret = new(Petshandler)
	default:
		err = errors.New("tablename is not found")
		log.Println("failed to match tablename in func GetDatatablesHandler, err: ", err.Error())
	}
	return
}

