package conversion

import (
	"fmt"
	"github.com/goinggo/mapstructure"
)

type Person struct {
	Name string `json:"name111"` // json 命名不影响
	Age int// 主要是字段名，要与 MAP中名字一样,todo 首字母不能小写
	Ext string// 没有map的字段，为默认值
}



func main() {
	mapInstance := make(map[string]interface{})
	mapInstance["Name"] = "liang637210"
	mapInstance["Age"] = "28"
	//var person Person
	person :=Person{}
	MapToStruct(mapInstance,&person)

	//var person Person
	////将 map 转换为指定的结构体
	//if err := mapstructure.Decode(mapInstance, &person); err != nil {
	//	log.Println(err)
	//}
	fmt.Println(person)
	//fmt.Println(person.Ext)
}



func MapToStructStr(m map[string]string,structData interface{}) (err error) {
	//fmt.Println(reflect.TypeOf(structData).Kind())
	//fmt.Println(reflect.ValueOf(structData))

	if err = mapstructure.Decode(m, structData); err != nil {
		return
	}
	return nil
	//fmt.Println(person)
	//fmt.Println(person.Ext)
}


func MapToStruct(m map[string]interface{},structData interface{}) (err error) {
	//fmt.Println(reflect.TypeOf(structData).Kind())
	//fmt.Println(reflect.ValueOf(structData))

	if err = mapstructure.Decode(m, structData); err != nil {
		return
	}
	return nil
	//fmt.Println(person)
	//fmt.Println(person.Ext)
}