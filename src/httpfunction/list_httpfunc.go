package httpfunction

import (
	"net/http"
	"dao"
	"handlehtml"
)

func ListHandle(w http.ResponseWriter, r *http.Request){

	info,err := dao.Listkey()
	check_error(err)

	locals := make(map[string]interface{})
	locals["info"] = info
	handlehtml.ReadHtml(w,"list", locals)
}

