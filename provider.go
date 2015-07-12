package gioc

import (
	"reflect"
)

type provider interface {
	apply(*injector)reflect.Value
	destroy()
}
type typeProvider struct {
	typ reflect.Type
}
func newTypeProvider(typ reflect.Type)provider{
	p:=typeProvider{}
	p.typ=typ
	return &p
}
func (this *typeProvider)apply(i *injector)reflect.Value{
	v:=reflect.New(this.typ)
	i.injectInto(v)
	return v
}
func (this *typeProvider)destroy(){
}
type valueProvider struct {
	injected bool
	value reflect.Value
}
func newValueProvider(value reflect.Value)provider{
	p:=valueProvider{}
	p.value=value
	return &p
}
func (this *valueProvider)apply(i *injector)reflect.Value{
	if !this.injected{
		i.injectInto(this.value)
		this.injected=true
	}
	return this.value
}
func (this *valueProvider)destroy(){

}
type singletonProvider struct {
	typ reflect.Type
	value reflect.Value
}

func newSingletonProvider (typ reflect.Type)provider{
	p:=singletonProvider{}
	p.typ=typ
	return &p
}
func (this *singletonProvider)apply(i *injector)reflect.Value{
	if !this.value.IsValid(){
		this.value=reflect.New(this.typ)
		i.injectInto(this.value)
	}
	return this.value
}
func (this *singletonProvider)destroy(){

}
