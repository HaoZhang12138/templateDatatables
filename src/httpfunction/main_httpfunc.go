package httpfunction

import (
	"net/http"
	"handlehtml"
)


func MainHandle(w http.ResponseWriter, r *http.Request){
	err := handlehtml.ReadHtml(w, "main", nil)
	check_error(err)
}