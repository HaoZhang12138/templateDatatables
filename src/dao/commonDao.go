package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func GetAll(tableName string, data interface{})(err error){
	f := func(c *mgo.Collection) (interface{}, error) {
		return nil, c.Find(bson.M{}).All(data)
	}
	_, err = doCllection(tableName, f)
	if err != nil {
		log.Println("failed to get all data, err: ", err.Error())
		return
	}
	return
}

func Insert(tableName string, data interface{})(err error) {

	f := func(c *mgo.Collection) (interface{}, error) {
		return nil, c.Insert(data)
	}
	_, err = doCllection(tableName, f)
	return
}

func Remove(tableName string, id interface{})(err error){
	f := func(c *mgo.Collection) (interface{}, error) {
		return nil, c.Remove(bson.M{"_id": id})
	}
	_, err = doCllection(tableName, f)
	return
}

func Update(tableName string, id interface{}, data interface{})(err error) {
	f := func(c *mgo.Collection) (interface{}, error) {
		return nil, c.UpdateId(id, data)
	}
	_, err = doCllection(tableName, f)
	return
}

func GetFileId(tableName string, id interface{})(ret string, err error){

	ans := make(map[string]string)
	f := func(c *mgo.Collection) (interface{}, error) {
		return nil, c.Find(bson.M{"_id": id}).Select(bson.M{"fileid": 1}).One(&ans)
	}
	_, err = doCllection(tableName, f)
	if err != nil {
		log.Println("failed to get fileId in function GetfileId, err: ", err.Error())
		return
	}
	ret = ans["fileid"]

	return

}

func GetOneById(tableName string, id interface{}, data interface{})(err error){

	f := func(c *mgo.Collection) (interface{}, error) {
		return nil, c.Find(bson.M{"_id": id}).One(data)
	}
	_, err = doCllection(tableName, f)

	if err != nil {
		log.Println("failed to get one userinfo, err: ", err.Error())
		return nil
	}
	return
}
