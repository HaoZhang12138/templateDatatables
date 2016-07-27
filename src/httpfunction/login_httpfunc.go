package httpfunction

import (
	"net/http"
	"dao"
	"io"
	"handlehtml"
)

func LoginHandle(w http.ResponseWriter, r *http.Request){

	if r.Method == "GET" {
		err := handlehtml.ReadHtml(w, "login", nil)
		check_error(err)
	}

	if r.Method == "POST" {

		var userinfo dao.UserLoginInfo
		loadUserInfo(r, &userinfo)

		flag, err := userinfo.Exist()
		check_error(err)

		if !flag {
			io.WriteString(w, "账号或者密码错误")
			return
		}

		flag, err = userinfo.Check()
		check_error(err)

		if !flag {
			io.WriteString(w, "账号或者密码错误")
			return
		}

		http.Redirect(w, r, "/main",
			http.StatusFound)
	}
}