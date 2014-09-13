package gioc

import (
	"reflect"
	"fmt"
)

type mapping struct {
	*injector
	typ  reflect.Type
	name string
	uid  string
}

func newMapping(injector *injector, typ reflect.Type, name, uid string) *mapping {
	mp := mapping{}
	mp.typ = typ
	mp.uid = uid
	mp.name = name
	mp.injector = injector
	return &mp
}

func (this *mapping) ToType(ptr interface{}) {
	typ := reflect.TypeOf(ptr)
	if typ.Implements(this.typ) {
		this.toType(typ)
	}else {
		panic(fmt.Sprintf("type:%s didn't implements %s", typ.String(), this.typ.String()))
	}
}
func (this *mapping) toType(typ reflect.Type) {
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	this.injector.addProvider(this.uid, newTypeProvider(typ))
}
func (this *mapping) ToValue(ptr interface{}) {
	typ := reflect.TypeOf(ptr)
	if typ.Implements(this.typ) {
		this.injector.addProvider(this.uid, newValueProvider(reflect.ValueOf(ptr)))
	}else {
		panic(fmt.Sprintf("type:%s didn't implements %s", typ.String(), this.typ.String()))
	}
}
func (this *mapping) toSingleton(typ reflect.Type) {
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() == reflect.Struct {
		this.injector.addProvider(this.uid, newSingletonProvider(typ))
	}else {
		panic("only struct type can be mapped as singleton!")
	}
}
func (this *mapping) ToSingleton(ptr interface{}) {
	typ := reflect.TypeOf(ptr)
	if typ.Implements(this.typ) {
		this.toSingleton(typ)
	}else {
		panic(fmt.Sprintf("type:%s didn't implements %s", typ.String(), this.typ.String()))
	}
}
func (this *mapping) AsSingleton() {
	if this.typ.Kind() == reflect.Struct {
		this.toSingleton(this.typ)
	}else {
		panic("only struct type could be mapped as singleton!")
	}
}
