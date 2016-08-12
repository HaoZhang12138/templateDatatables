package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Uploadfile struct {
	Id string `json:"id" bson:"_id"`
	Filename string `json:"filename"`
	Filesize string `json:"filesize"`
	Webpath string `json:"webpath"`
	Systempath string  `json:"systempath"`
}


func (this *Uploadfile) GetOneUploadfile(uploadtablename string)(err error) {
	f := func(c *mgo.Collection) (interface{}, error) {
		return nil, c.Find(bson.M{"_id": this.Id}).One(this)
	}
	_, err = doCllection(uploadtablename, f)
	return
}


func (this *Uploadfile) Insert(uploadtablename string)(err error) {
	f := func(c *mgo.Collection) (interface{}, error) {
		return nil, c.Insert(this)
	}
	_, err = doCllection(uploadtablename, f)
	return
}


func (this *Uploadfile) Remove(uploadtablename string)(err error){
	f := func(c *mgo.Collection) (interface{}, error) {

		/*err := os.Remove(this.Systempath)
		if err != nil {
			log.Println("failed to delete file")
			return nil,err
		}*/ // delete the local file

		return nil, c.Remove(bson.M{"_id": this.Id})
	}
	_, err = doCllection(uploadtablename, f)
	return
}

func GetAllUploadfile(uploadtablename string) (ret []Uploadfile, err error) {
	result := make([]Uploadfile, 0)
	f := func(c *mgo.Collection) (interface{}, error) {
		return nil, c.Find(bson.M{}).All(&result)
	}

	_, err = doCllection(uploadtablename, f)
	if err != nil {
		return
	}
	ret = result
	return
}