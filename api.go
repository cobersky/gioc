package gioc

import ."reflect"

type Injector interface {
	Map(ptr interface{}, name string) Mapping
	MapByType(typ Type, name string)Mapping
	CreateChild() Injector
	UnMap(ptr interface{}, name string)
	UnMapByType(typ Type, name string)
	InjectInto(p interface{})
	GetParent() Injector
	SetParent(Injector)
	GetInstance(ptr interface{}, name string) interface{}
	GetInstanceByType(typ Type, name string) Value
	InstantiationUnMapped(typ Type) Value
	HasMapping(ptr interface{}, name string, deeply bool) bool
	HasMappingOfType(typ Type, name string, deeply bool) bool
}
type Mapping interface {
	ToType(ptr interface{})
	ToValue(ptr interface{})
	ToSingleton(typ interface{})
	AsSingleton()
}
