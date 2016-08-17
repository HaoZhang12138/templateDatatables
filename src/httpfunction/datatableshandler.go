package httpfunction

import (
	"net/http"
	"errors"
	"log"
)

const (
	FILES_NOT_NEEDED = iota
	FILES_NEEDED_ONE
	FILES_NEEDED_ALL
)

type DataTableHandler interface {
	Upload(w http.ResponseWriter, r *http.Request, resp *Datatablesdata)(error)
	Create(w http.ResponseWriter, r *http.Request, resp *Datatablesdata)(error)
	Edit(w http.ResponseWriter, r *http.Request, resp *Datatablesdata)(error)
	Remove(w http.ResponseWriter, r *http.Request, resp *Datatablesdata)(error)
	GET(w http.ResponseWriter, r *http.Request, resp *Datatablesdata)(error)
}

func GetDatatablesHandler(tableName string)(ret DataTableHandler, err error)  {
	switch tableName {
	case "userinfo":
		ret = new(Userinfohandler)
	case "course":
		ret = new(Coursehandler)
	case "pets":
		ret = new(Petshandler)
	default:
		err = errors.New("tablename is not found")
		log.Println("failed to match tablename in func GetDatatablesHandler, err: ", err.Error())
	}
	return
}


