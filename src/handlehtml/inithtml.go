package handlehtml

import (
	"io/ioutil"
	"path"
	"log"
	"html/template"
)

const HTML_DIR = "/home/zh/GoPro/firstRedis/html"

var Templates = make(map[string]*template.Template)
func Inithtml() {
	fileInfoArr, err := ioutil.ReadDir(HTML_DIR)
	if err != nil {
		panic(err)
		return
	}
	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		templatePath = HTML_DIR + "/" + templateName
		log.Println("Loading template:", templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		Templates[templateName] = t
	}
}
