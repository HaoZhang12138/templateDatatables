package dao

import (
	"testing"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func TestUserLoginInfo_Check(t *testing.T) {
	t1 := UserLoginInfo{User: "1", Pass: "2"}
	ok, err := t1.Check()
	if err != nil {
		t.Error(err.Error())
	}
	if ok != false {
		t.Error("Check error")
	}

}

func TestUserLoginInfo_Check2(t *testing.T) {
	t1 := UserLoginInfo{User: "1", Pass: "1"}
	ok, err := t1.Check()
	if err != nil {
		t.Error(err.Error())
	}

	if ok != true {
		t.Error("Check2 error")
	}
}

func TestUserLoginInfo_Exist(t *testing.T) {
	t1 := UserLoginInfo{User: "10"}

	ok, err := t1.Exist()
	if err != nil {
		t.Error(err.Error())
	}

	if ok != false {
		t.Error("Exist function error")
	}
}

func TestUserLoginInfo_Exist2(t *testing.T) {
	t1 := UserLoginInfo{User:"a"}
	ok, err := t1.Exist()
	if err != nil {
		t.Error(err.Error())
	}

	if ok != true {
		t.Error("Exist function2 error")
	}
}


func TestUserLoginInfo_Insert(t *testing.T) {
	t1 := UserLoginInfo{User:"test", Pass:"test"}

	err := t1.Insert()

	if err != nil {
		t.Error(err.Error())
	}
}

func TestListkey(t *testing.T) {
	f := func(c *mgo.Collection) (interface{}, error){
		info,err := c.RemoveAll(bson.M{})

		return info,err
	}
	doCllection("userinfo", f)

	_, err := Listkey()
	if err != nil {
		t.Error(err.Error())
	}
}
