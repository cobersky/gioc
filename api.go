package gioc

type Injector interface {
	Map(ptr interface {},name string)Mapping
	CreateChild()Injector
	UnMap(ptr interface {},name string)
	InjectInto(p interface {})
	GetParent()Injector
	SetParent(Injector)
	GetInstance(ptr interface {},name string)interface {}
	NewInstance(ptr interface {},name string)interface {}
	HasMapping(ptr interface {},name string,deeply bool)bool
}
type Mapping interface {
	ToType(ptr interface {})
	ToValue(ptr interface {})
	ToSingleton(typ interface {})
	AsSingleton()
}
