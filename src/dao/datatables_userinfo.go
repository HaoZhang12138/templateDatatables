package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"errors"
)

func GetAllUserinfo()(ret []UserLoginInfo, err error){

	result := make([]UserLoginInfo, 0)
	f := func(c *mgo.Collection) (interface{}, error) {
		return nil, c.Find(bson.M{}).All(&result)
	}

	_, err = doCllection(collections, f)
	if err != nil {
		return
	}
	ret = result
	return
}

func (this *UserLoginInfo) GetOneById()(interface{}) {

	var result UserLoginInfo
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
	f := func(c *mgo.Collection) (interface{}, error) {
		var ret UserLoginInfo
		err := c.Find(bson.M{"_id": this.Id}).One(&ret)
		return ret, err
	}
	ans, err := doCllection(collections, f)
	if err != nil {
		log.Println("failed to get fileId")
		return
	}
	t, ok := ans.(UserLoginInfo)
	if !ok {
		err = errors.New("fail to type assertion")
	}
	ret = t.FileId
	return

}
