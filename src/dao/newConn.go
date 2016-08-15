package dao

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"log"
	"gopkg.in/mgo.v2/bson"
)

const host = "127.0.0.1:27017"
const dataBase = "myuser"
const countertable = "counter"

var mgoSession *mgo.Session
func getSession() *mgo.Session {

	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(host)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	return mgoSession.Clone()
}

func doCllection(col string, f func(*mgo.Collection) (interface{}, error) ) (interface{}, error) {
	session := getSession()
	defer session.Close()
	c :=  session.DB(dataBase).C(col)
	return f(c)
}

func GetNextId() (id int) {
	f := func(counter *mgo.Collection) (interface{}, error) {
		cid :="counterid"
		change := mgo.Change{
			Update: bson.M{"$inc": bson.M{"seq": 1}},
			Upsert: true,
			ReturnNew: true,
		}
		doc := struct{ Seq int }{}
		if _, err := counter.Find(bson.M{"_id": cid}).Apply(change, &doc); err != nil {
			panic(fmt.Errorf("get counter failed:", err.Error()))
		}
		log.Println("seq:", doc)
		return doc.Seq, nil
	}
	n, _ := doCllection(countertable, f)
	id,_ = n.(int)
	return id
}

