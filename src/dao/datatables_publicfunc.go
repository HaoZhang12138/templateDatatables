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

func GetOneById(tableName string, Id string, data interface{})(err error){
	f := func(c *mgo.Collection) (interface{}, error) {
		return nil, c.Find(bson.M{"_id": Id}).One(data)
	}
	_, err = doCllection(tableName, f)

	if err != nil {
		log.Println("failed to get one userinfo")
		return nil
	}
	return
}

/*
func CommonLoadFromPostForm(req * http.Request,id string,ptrdata interface{})(err error)  {
	// parse form
	if err = req.ParseForm();err !=nil{
		log.Error("failed to parseform:%v",err)
		return
	}

	v := reflect.ValueOf(ptrdata).Elem()
	fields := make(map[string]reflect.Value)
	for i := 0; i < v.NumField(); i++ {
		prefix := "data["+id+"]["
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		jsonname := tag.Get("json")
		if jsonname == ""{
			jsonname = strings.ToLower(fieldInfo.Name)
		}
		prefix = prefix+jsonname+"]"
		fields[prefix] = v.Field(i)
	}

	// take from form
	for name, values := range req.Form{
		f,found := fields[name]
		if !found{
			continue
		}
		for _,value :=range values{
			err = populate(f,value)
			if err != nil {
				log.Error("failed to populate err:%v",err)
				return
			}
		}
	}
	createprefix := "data[0][_id]"
	if createid, exist := fields[createprefix];exist{
		autoid,err := dao.GetDataTableAutoIncId()
		if err != nil {
			log.Error("failed allocate id for create :%v",err)
			return err
		}
		populate(createid,autoid)
	}
	return
}

func populate(v reflect.Value, value string) error  {
	switch v.Kind() {
	case reflect.String:
		if v.CanSet(){
			v.SetString(value)
		}else {
			return errors.New("string field can't be set")
		}

	case reflect.Int:
		i,err := strconv.ParseInt(value,10,64)
		if err != nil {
			log.Error("failed to pasrse %v to int",value)
			return err
		}
		if v.CanSet(){
			v.SetInt(i)
		}else {
			return errors.New("int field can't be set")
		}

	case reflect.Int64:
		i,err := strconv.ParseInt(value,10,64)
		if err != nil {
			log.Error("failed to pasrse %v to int",value)
			return err
		}
		if v.CanSet(){
			v.SetInt(i)
		}else {
			return errors.New("int64 field can't be set")
		}

	case reflect.Bool:
		b,err := strconv.ParseBool(value)
		if err != nil {
			log.Error("failed to pasrse %v to int",value)
			return err
		}
		if v.CanSet(){
			v.SetBool(b)
		}else {
			return errors.New("bool field can't be set")
		}

	default:
		return fmt.Errorf("unsupported kind %s",v.Type())
	}
	return nil
}*/

