package dao

type UserinfoDatatables struct {
	User string `json:"user"`
	Pass string `json:"pass"`
	Name string `json:"name"`
	Sex string  `json:"sex"`
	Age string  `json:"age"`
	Tel string  `json:"tel"`
	Id string `json:"id" bson:"_id"` // if you change json:"id", you should have change TableIdInJson in same

	FileId string `json:"fileid"` // if you need upload file, add this and not change it
}
const TableIdInJson = "id"

func (this *UserinfoDatatables)GetId()(string){
	return  this.Id //change Id to you primary key name
}
