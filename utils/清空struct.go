package utils

import (
	"fmt"
	"reflect"
)

type A struct {
	d map[string]interface{}
	dd []string
}
func main()  {
	a:=A{}
	a.d=make(map[string]interface{},0)
	a.d["aa"]=1
	a.dd=make([]string,0)
	fmt.Printf("%v\n",a)
	Clear(&a)
	fmt.Printf("%v\n",a)
	a.d["aa"]=1
}
func Clear(v interface{}) { // 必须传入指针
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}