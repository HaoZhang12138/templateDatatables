package dao

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

type UserLoginInfo struct {
	User string `json:"user"`
	Pass string `json:"pass"`
	Name string `json:"name"`
	Sex string  `json:"sex"`
	Age string  `json:"age"`
	Tel string  `json:"tel"`
	Id string `json:"id" bson:"_id"`
	FileId string `json:"fileid"`
}

/*func (this *UserLoginInfo) Insert()(err error) {
	c := Pool.Get()
	defer c.Close()

	_, err = c.Do("hmset", this.User, "pass", this.Pass, "name", this.Name,
	"sex", this.Sex, "age", this.Age, "tel", this.Tel)

	return
}


func (this *UserLoginInfo) Exist()(exist bool,err error) {
	user := this.User
	c := Pool.Get()
	defer c.Close()
	exist, err = redis.Bool(c.Do("exists", user))
	return
}

func (this *UserLoginInfo) Check()(ok bool, err error){
	c := Pool.Get()
	defer c.Close()
	tmp, err := redis.String(c.Do("HGET", this.User, "pass"))
	if err != nil {
		return
	}
	if tmp != this.Pass {
		ok = false
	}else{
		ok = true
	}
	return
}

func (this *UserLoginInfo) Getall()(info map[string]string, err error){
	c := Pool.Get()
	defer c.Close()
	info, err = redis.StringMap(c.Do("hgetall", this.User))
	if err != nil {
		return
	}
	delete(info, "pass")
	delete(info, "user")

	return
}

func Listkey() (info interface{}, err error) {
	c := Pool.Get()
	defer c.Close()
	info,err = redis.Strings(c.Do("keys", "*"))
	return
}*/

const collections = "userinfo"

func (this *UserLoginInfo) Exist()(exist bool,err error) {


	f := func(c *mgo.Collection) (interface{}, error){
		return c.Find(bson.M{"user": this.User}).Count()
	}

	n, err := doCllection(collections, f)

	exist = true
	if n == 0 {
		exist = false
	}

	return
}

func (this *UserLoginInfo) Check()(ok bool, err error){

	result := UserLoginInfo{}
	f := func(c *mgo.Collection) (interface{}, error){
		err := c.Find(bson.M{"user": this.User}).One(&result)
		return nil, err
	}

	_, err = doCllection(collections, f)
	if err != nil {
		return
	}

	ok = true
	if result.Pass != this.Pass {
		ok = false
	}

	return
}

func Listkey() (info []string, err error) {

	result := []UserLoginInfo{}

	f := func(c *mgo.Collection) (interface{}, error){
		return nil, c.Find(bson.M{}).All(&result)
	}

	_, err = doCllection(collections, f)

	if err != nil {
		return
	}

	lens := len(result)
	info = make([]string, lens)
	for i := 0; i < lens; i++ {
		info = append(info, result[i].User)
	}
	return
}









