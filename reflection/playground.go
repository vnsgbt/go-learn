package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello, playground")

	tst2 := 10
	tst3 := 1.2
	mym := map[string]interface{}{}
	mym["id"] = "vns"

	fmt.Println(reflect.TypeOf(tst2))
	fmt.Println(reflect.TypeOf(tst3))

	mymtp := reflect.TypeOf(mym).Elem()

	fmt.Printf("\n%t\n", mymtp)

	maptp := reflect.TypeOf((map[string]interface{})(nil))
	fmt.Println(maptp.Implements(mymtp))

	var a interface{} = make(map[string]interface{})
	fmt.Println(reflect.TypeOf(a).Kind() == reflect.Map)

	type MyStruct struct {
		ID  string `json:"id"`
		Age int `json:"age"`
	}

	myStruct := &MyStruct{
		ID: "vms",
		Age: 5,
	}

	st := reflect.TypeOf(myStruct).Elem()
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i).Tag.Get("json")
		fmt.Printf("\n%v\n", field)
	}

	ss := reflect.ValueOf(myStruct).Elem()
	for i := 0; i < ss.NumField(); i++ {
		field := ss.Field(i).Interface()
		fmt.Printf("\n%v, %v \n", field, reflect.TypeOf(field).Kind() == reflect.Int)
	}
}
