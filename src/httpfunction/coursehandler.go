package httpfunction

import "net/http"

type Coursehandler struct {

}

func (Coursehandler) Upload(w http.ResponseWriter, r *http.Request, resp *Datatablesdata)(error) {
	return  Upload_default(w, r, resp)
}

func (Coursehandler) Create(w http.ResponseWriter, r *http.Request, resp *Datatablesdata)(error) {
	return Create_default(w, r, resp)
}

func (Coursehandler) Edit(w http.ResponseWriter, r *http.Request, resp *Datatablesdata)(error) {
	return Edit_default(w, r, resp)
}

func (Coursehandler) Remove(w http.ResponseWriter, r *http.Request, resp *Datatablesdata)(error) {
	return  Remove_default(w, r, resp)
}

func (Coursehandler) GET(w http.ResponseWriter, r *http.Request, resp *Datatablesdata)(error) {
	return GET_default(w, r, resp)
}