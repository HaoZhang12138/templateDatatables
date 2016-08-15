package dao

type CourseDatatables struct {
	Courseid string `json:"courseid" bson:"_id"`
	Coursename string `json:"coursename" bson:"coursename"`
	Teachername string `json:"teachername" bson:"teachername"`
	Overview string        `json:"overview" bson:"overview"`
}


func (this *CourseDatatables)GetId()(string){
	return  this.Courseid
}


