package component

import (
	structFactory "../../../pkg/struct_factory"
	"fmt"
	"reflect"
)

type Factory struct {
	registry map[string]reflect.Type
}

func NewFactory() *Factory {
	return &Factory{
		registry: make(map[string]reflect.Type),
	}
}

func (factory *Factory) Register(name string, reflectType reflect.Type) {
	factory.registry[name] = reflectType
}

func (factory *Factory) CreateMessageTransport(transportType string) (TransportInterface, error) {
	fmt.Printf("Factory registry %v\n", factory.registry)
	fmt.Printf("Factory create %v\n", transportType)

	t := factory.registry[transportType]
	v := reflect.New(t)
	structFactory.InitializeStruct(t, v.Elem())
	c := v.Interface().(TransportInterface)

	return c, nil
}
