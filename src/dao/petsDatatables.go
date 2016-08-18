package dao

//上传文件的例子
type PetsDatatables struct {
	Petid string `json:"petid" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Category string `json:"category" bson:"category"`
	Color string `json:"color" bson:"color"`
	FileId string `json:"fileid" json:"fileid"`
}

func (this *PetsDatatables)GetId()(interface{}){
	return  this.Petid
}
