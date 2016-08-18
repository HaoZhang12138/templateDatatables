package httpfunction

import "net/http"


//一个自定义handler的例子， 可参考
type Petshandler struct {

}

func (Petshandler) Upload(w http.ResponseWriter, r *http.Request, resp *Datatablesdata)(error) {
	return  Upload_default(w, r, resp)
}

func (Petshandler) Create(w http.ResponseWriter, r *http.Request, resp *Datatablesdata)(error) {
	return Create_default(w, r, resp)
}

func (Petshandler) Edit(w http.ResponseWriter, r *http.Request, resp *Datatablesdata)(error) {
	return Edit_default(w, r, resp)
}

func (Petshandler) Remove(w http.ResponseWriter, r *http.Request, resp *Datatablesdata)(error) {
	return  Remove_default(w, r, resp)
}

func (Petshandler) GET(w http.ResponseWriter, r *http.Request, resp *Datatablesdata)(error) {
	return GET_default(w, r, resp)
}
