package httpfunction

import "net/http"

type Userinfohandler struct {

}

func (Userinfohandler) Upload(w http.ResponseWriter, r *http.Request, resp *Datatablesdata)(error) {
	return  Upload_default(w, r, resp)
}

func (Userinfohandler) Create(w http.ResponseWriter, r *http.Request, resp *Datatablesdata)(error) {
	return Create_default(w, r, resp)
}

func (Userinfohandler) Edit(w http.ResponseWriter, r *http.Request, resp *Datatablesdata)(error) {
	return Edit_default(w, r, resp)
}

func (Userinfohandler) Remove(w http.ResponseWriter, r *http.Request, resp *Datatablesdata)(error) {
	return  Remove_default(w, r, resp)
}

func (Userinfohandler) GET(w http.ResponseWriter, r *http.Request, resp *Datatablesdata)(error) {
	return GET_default(w, r, resp)
}

