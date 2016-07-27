package httpfunction

import (
	"net/http"
	"dao"
	"handlehtml"
	"io"
	"sync"
)

var lock sync.Mutex

func loadUserInfo(r *http.Request, info *dao.UserLoginInfo){
	info.User = r.FormValue("user")
	info.Pass = r.FormValue("pass")
	info.Name = r.FormValue("name")
	info.Age = r.FormValue("age")
	info.Sex = r.FormValue("sex")
	info.Tel = r.FormValue("tel")
}

func AddHandle(w http.ResponseWriter, r *http.Request){

	if r.Method == "GET" {

		err := handlehtml.ReadHtml(w, "add", nil)
		check_error(err)
	}

	if r.Method == "POST" {

		lock.Lock()
		defer lock.Unlock()

		var userinfo dao.UserLoginInfo
		loadUserInfo(r, &userinfo)
		flag,err := userinfo.Exist()
		check_error(err)

		//time.Sleep(15 * time.Second)

		if flag {
			io.WriteString(w, "账号 " + userinfo.User + " 已存在")
			return
		}else {
			err := userinfo.Insert()
			check_error(err)
			handlehtml.ReadHtml(w,"addsuccess", nil)
		}

	}
}
