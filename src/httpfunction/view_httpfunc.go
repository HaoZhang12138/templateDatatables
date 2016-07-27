package httpfunction

import (
	"net/http"
	"dao"
	"io"
	"handlehtml"
)

func ViewHandle(w http.ResponseWriter, r *http.Request){

	var userinfo dao.UserLoginInfo
	loadUserInfo(r, &userinfo)

	flag, err := userinfo.Exist()
	check_error(err)

	if !flag {
		io.WriteString(w, "用戶名为 " + userinfo.User + " 的用戶不存在")
		return
	}

	info,err := userinfo.GetUserinfo()
	check_error(err)

	locals := make(map[string]interface{})
	locals["info"] = info
	handlehtml.ReadHtml(w,"view", locals)
}
