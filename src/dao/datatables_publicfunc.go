package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func  GetAll(tableName string)(ret []interface{}, err error){

	ret = make([]interface{}, 0)
	f := func(c *mgo.Collection) (interface{}, error) {
		return nil, c.Find(bson.M{}).All(&ret)
	}
	_, err = doCllection(tableName, f)
	if err != nil {
		log.Println("failed to get all data")
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

func Remove(tableName string, Id string)(err error){
	f := func(c *mgo.Collection) (interface{}, error) {
		return nil, c.Remove(bson.M{"_id": Id})
	}
	_, err = doCllection(tableName, f)
	return
}


func Update(tableName string, Id string, data interface{})(err error) {
	f := func(c *mgo.Collection) (interface{}, error) {
		return nil, c.UpdateId(Id, data)
	}
	_, err = doCllection(tableName, f)
	return
}

func GetFileId(tableName string, Id string)(ret string, err error){

	ans := make(map[string]string)
	f := func(c *mgo.Collection) (interface{}, error) {
		return nil, c.Find(bson.M{"_id": Id}).Select(bson.M{"fileid": 1}).One(&ans)
	}
	_, err = doCllection(tableName, f)
	if err != nil {
		log.Println("failed to get fileId in function GetfileId")
		return
	}
	ret = ans["fileid"]

	return

}

func GetOneById(tableName string, Id string)(interface{}) {

	var result interface{}
	f := func(c *mgo.Collection) (interface{}, error) {
		return nil, c.Find(bson.M{"_id": Id}).One(&result)
	}
	_, err := doCllection(tableName, f)
	if err != nil {
		log.Println("failed to get one userinfo")
		return nil
	}

	return result
}
