package dao

import (
	"net/http"
)

type EditDataTables interface {
	LoadDataFromPostForm(*http.Request,string)
	GetId()(string)
}
