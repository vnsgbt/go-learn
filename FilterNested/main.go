package main

import (
	"reflect"
	"strings"
	"encoding/json"
	"bytes"
	"fmt"
)

type InnerStruct struct {
	ID      string `json:"id"`
	Private string `json:"private" private:"true"`
}
type MyStruct struct {
	ID      string `json:"id"`
	Age     int `json:"age"`
	Inner   InnerStruct `json:"inner"`
	Private string `json:"private" private:"true"`
}

type MyTypes map[string]MyFields

type MyFields struct {
	T       reflect.Type
	Private bool
}

func main() {
	innerStruct := InnerStruct{
		ID: "inner",
		Private: "private",
	}
	myStruct := MyStruct{
		ID: "vms",
		Age: 5,
		Inner: innerStruct,
		Private: "private",
	}
	rs, err := FilterNestedStruct(myStruct)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rs)
}

func filter(typeinf reflect.Type) map[string]interface{} {
	fmt.Printf("\n[Type to be filtered] %v\n", typeinf)
	pmap := make(map[string]interface{})
	for i := 0; i < typeinf.NumField(); i++ {
		field := typeinf.Field(i)
		if val := field.Tag.Get("private"); val == "true" {
			name := field.Tag.Get("json")
			name = (strings.Split(name, ","))[0]
			pmap[name] = nil
		}
	}
	return pmap
}

func parseFilter(nestInf interface{}, typeinf reflect.Type) map[string]interface{} {
	nestMap := nestInf.(map[string]interface{})
	fmt.Printf("[Incoming Type]: %v", typeinf)
	filterMap := filter(typeinf)
	fmt.Printf("[Filtered Map]: %v", filterMap)
	for key := range filterMap {
		delete(nestMap, key)
	}
	fmt.Printf("\n[After Delete] %v\n", nestMap)
	for key := range nestMap {
		if reflect.TypeOf(nestMap[key]).Kind() == reflect.Map {
			fmt.Printf("\n[Each Key] %v \n", nestMap[key])
			innerInf := func() (reflect.Type, bool) {
				rType := reflect.TypeOf(nestMap[key])
				fmt.Printf("[rType] %v", rType)
				for i := 0; i < rType.NumField(); i++ {
					field := rType.Field(i).Tag.Get("json")
					if key == field {
						return rType, true
					}
				}
				return nil, false
			}
			fmt.Printf("[InnerMap] %v\n", nestMap[key])
			if inf, ok := innerInf(); ok {
				nestMap[key] = parseFilter(nestMap[key], inf)
			}
		}
	}
	return nestMap
}

func FilterNestedStruct(source interface{}) (map[string]interface{}, error) {
	// Decode source to interface map
	omap := make(map[string]interface{})
	if b, err := json.Marshal(source); err == nil {
		if err = json.NewDecoder(bytes.NewReader(b)).Decode(&omap); err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}
	rs := parseFilter(omap, reflect.TypeOf(source))
	fmt.Println(rs)
	return rs, nil
}
