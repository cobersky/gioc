package gioc

import (
	"reflect"
)

type typeDescribe struct {
	fields []reflect.StructField
	initMethod reflect.Method
	hasInitMethod bool
}
var typeDescribes = map[string]*typeDescribe{}
func getTypeDescribe(typ reflect.Type)*typeDescribe{
	key:=typ.String()
	if td,ok:=typeDescribes[key];ok{
		return td
	}
	return createTypeDescribe(typ)
}
func createTypeDescribe(typ reflect.Type)*typeDescribe{
	td:=&typeDescribe{}
	td.initMethod,td.hasInitMethod=typ.MethodByName("Init")
	for typ.Kind()==reflect.Ptr{
		typ=typ.Elem()
	}
	for i:=0;i<typ.NumField();i++{
		f:=typ.Field(i)
		tag:=f.Tag.Get("inject")
		if tag!=""{
			td.fields=append(td.fields,f)
		}
	}
	typeDescribes[typ.String()]=td
	return td
}
