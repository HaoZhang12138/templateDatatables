package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func (this *UserLoginInfo)GetAll()(ret []interface{}, err error){

	ret = make([]interface{}, 0)
	f := func(c *mgo.Collection) (interface{}, error) {
		return nil, c.Find(bson.M{}).All(&ret)
	}
	_, err = doCllection(collections, f)
	if err != nil {
		log.Println("failed to get all data")
		return
	}
	return
}

func (this *UserLoginInfo) GetOneById()(interface{}) {

	var result interface{}
	f := func(c *mgo.Collection) (interface{}, error) {
		return nil, c.Find(bson.M{"_id": this.Id}).One(&result)
	}
	_, err := doCllection(collections, f)
	if err != nil {
		log.Println("failed to get one userinfo")
		return nil
	}

	return result
}

func (this *UserLoginInfo) Insert()(err error) {

	f := func(c *mgo.Collection) (interface{}, error) {
		return nil, c.Insert(this)
	}
	_, err = doCllection(collections, f)
	return
}

func (this *UserLoginInfo) Remove()(err error){
	f := func(c *mgo.Collection) (interface{}, error) {
		return nil, c.Remove(bson.M{"_id": this.Id})
	}
	_, err = doCllection(collections, f)
	return

}

func (this *UserLoginInfo) Update()(err error) {
	f := func(c *mgo.Collection) (interface{}, error) {
		return nil, c.UpdateId(this.Id, this)
	}
	_, err = doCllection(collections, f)
	return
}

func (this *UserLoginInfo) GetfileId()(ret string, err error){
	ans := make(map[string]string)
	f := func(c *mgo.Collection) (interface{}, error) {
		return nil, c.Find(bson.M{"_id": this.Id}).Select(bson.M{"fileid": 1}).One(&ans)
	}
	_, err = doCllection(collections, f)
	if err != nil {
		log.Println("failed to get fileId in function GetfileId")
		return
	}
	ret = ans["fileid"]

	return

}
