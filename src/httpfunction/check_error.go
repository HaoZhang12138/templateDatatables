package httpfunction

import "fmt"

func check_error(err error){
	if err != nil{
		//panic(err)
		fmt.Println(err.Error())
	}

}
