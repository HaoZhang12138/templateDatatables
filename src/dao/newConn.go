package dao

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"log"
	"gopkg.in/mgo.v2/bson"
)

/*const password = "root"
const host = "127.0.0.1:6379"
const redisDB = 0

var Pool *redis.Pool
func InitRedis() {
	Pool = &redis.Pool{
		MaxIdle: 20,
		IdleTimeout: 60 * time.Second,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_,err := c.Do("ping")

			return err
		},
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}
			_, err = c.Do("auth", password)
			if err != nil {
				return nil, err
			}

			_, err = c.Do("select", redisDB)
			return c, err
		},
	}
}*/


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

/*func GetNextId() func() int {
	session := getSession()
	counter := session.DB(dataBase).C(countertable)
	cid :="counterid"
	return func() int {
		defer session.Close()
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
		return doc.Seq
	}
}*/

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

