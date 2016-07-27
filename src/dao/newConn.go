package dao

import (
	"gopkg.in/mgo.v2"
	"fmt"
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

var mgoSession *mgo.Session
func getSession() *mgo.Session {

	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(host)
		if err != nil {
			fmt.Println(err.Error())
			//panic(err)
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
