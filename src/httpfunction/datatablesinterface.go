package httpfunction

import (
	"net/http"
	"github.com/bitly/go-simplejson"
)

//url中用于传表名的字段
const URLTableName = "tableName"

const (
	FILES_NOT_NEEDED = iota
	FILES_NEEDED_ONE
	FILES_NEEDED_ALL
)
//http回复的数据
type Datatablesdata struct {
	Data interface{}  `json:"data"`
	Files uploadfile_tmp `json:"files,omitempty"`
	Upload uploadid `json:"upload,omitempty"`
}

type uploadid struct {
	Id string `json:"id"`
}

type uploadfile_tmp struct {
	Files *simplejson.Json `json:"files"`
}

//handler接口
type DataTableHandler interface {
	Upload(w http.ResponseWriter, r *http.Request, resp *Datatablesdata)(error)
	Create(w http.ResponseWriter, r *http.Request, resp *Datatablesdata)(error)
	Edit(w http.ResponseWriter, r *http.Request, resp *Datatablesdata)(error)
	Remove(w http.ResponseWriter, r *http.Request, resp *Datatablesdata)(error)
	GET(w http.ResponseWriter, r *http.Request, resp *Datatablesdata)(error)
}



