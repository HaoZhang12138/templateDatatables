package handlehtml

import "net/http"

func ReadHtml(w http.ResponseWriter, name string, locals map[string]interface{}) error{
	t := Templates[name + ".html"]
	err := t.Execute(w, locals)
	return err
}
