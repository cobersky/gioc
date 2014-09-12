package gioc

import (
	"reflect"
	"fmt"
	"sync"
)

type injector struct {
	parent Injector
	providers map[string]provider
	sync.Mutex
}
func NewInjector() Injector {
	i := injector{}
	i.providers = make(map[string]provider,64)
	return &i
}

func (this *injector) Map(ptr interface{}, name string) Mapping {
	typ := reflect.TypeOf(ptr)
	return this.mapType(typ, name)
}

func (this *injector) CreateChild() Injector {
	i := NewInjector()
	i.SetParent(this)
	return i
}
func (this *injector) UnMap(ptr interface{}, name string) {
	typ := reflect.TypeOf(ptr)
	this.UnMap(typ, name)
}

func (this *injector) InjectInto(p interface{}) {
	val := reflect.ValueOf(p)
	this.injectInto(val)
}

func (this *injector) GetInstance(ptr interface{}, name string) interface{} {
	typ := reflect.TypeOf(ptr)
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return this.getInstance(typ, name)
}

func (this *injector) NewInstance(ptr interface{}, name string) interface{} {
	return this.newInstance(reflect.TypeOf(ptr), name).Interface()
}
func (this*injector) GetParent() Injector {
	return this.parent
}
func (this*injector) SetParent(i Injector) {
	this.parent = i
}

func (this*injector) HasMapping(ptr interface{}, name string, deeply bool) (b bool) {
	return this.hasMapping(reflect.TypeOf(ptr), name, deeply)
}

/*************************************************/
/**private methods*/
func (this *injector) createMapping(typ reflect.Type, name, key string) Mapping {
	if _,ok:=this.providers[key];ok{
		panic("this mapping is already exists!")
	}
	mp := newMapping(this, typ, name, key)
	return mp
}

func (this *injector) getProvider(typ reflect.Type, name string) provider {
	i := this
	for i != nil {
		if p, ok := i.providers[generateUid(typ, name)]; ok {
			return p
		}
		if this.parent == nil {
			i = nil
		}else {
			i = this.parent.(*injector)
		}
	}
	return nil
}

func (this *injector) newInstance(typ reflect.Type, name string) reflect.Value {
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	p := this.getProvider(typ, name)
	if p != nil {
		return p.apply(this)
	}else {
		v := reflect.New(typ)
		this.injectInto(v)
		return v
	}
}

func (this *injector) getInstance(typ reflect.Type, name string) reflect.Value {
	p := this.getProvider(typ, name)
	if p != nil {
		return p.apply(this)
	}
	panic("can't find mapping!")
}

func (this *injector) injectInto(val reflect.Value) {
	for val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		tField := typ.Field(i)
		name := tField.Tag.Get("inject")
		if len(name) > 0 {
			vField := val.Field(i)
			if vField.CanSet() {
				vField.Set(this.getInstance(tField.Type, name))
			}
		}
	}
}

func (this *injector) unMap(typ reflect.Type, name string) {
	if this.hasMapping(typ, name, false) {
		this.Lock()
		delete(this.providers, generateUid(typ, name))
		this.Unlock()
	}
}

func (this*injector) hasMapping(typ reflect.Type, name string, deeply bool) (b bool) {
	i := this
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	b = false
	for i != nil {
		_, b = i.providers[generateUid(typ, name)]
		if deeply {
			break
		}
		if this.parent == nil {
			i = nil
		}else {
			i = this.parent.(*injector)
		}
	}
	return
}

func (this *injector) mapType(typ reflect.Type, name string) Mapping {
	fmt.Println(typ)
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	mappingKey := generateUid(typ, name)
	return this.createMapping(typ, name, mappingKey)
}
/****************************************************/
func generateUid(typ reflect.Type, name string) string {
	if len(name) > 0 {
		return typ.String() + ":" + name
	}else {
		return typ.String() + ":-"
	}
}

