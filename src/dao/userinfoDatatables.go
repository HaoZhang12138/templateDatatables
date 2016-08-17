package dao

type UserinfoDatatables struct {
	User string `json:"user"`
	Pass string `json:"pass"`
	Name string `json:"name"`
	Sex string  `json:"sex"`
	Age string  `json:"age"`
	Tel string  `json:"tel"`
	Id string `json:"id" bson:"_id"`
	FileId string `json:"fileid"`
}

func (this *UserinfoDatatables)GetId()(interface{}){
	return  this.Id
}

