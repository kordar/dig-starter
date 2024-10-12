package dig_starter

import (
	goframeworkdig "github.com/kordar/goframework-dig"
	logger "github.com/kordar/gologger"
	"go.uber.org/dig"
)

func GetDig(namespace string) *dig.Container {
	if !goframeworkdig.HasDigInstance(namespace) {
		logger.Fatalf("dig instance %s not exist.", namespace)
	}
	return goframeworkdig.GetDigInstance(namespace)
}

func ProvideE(constructor interface{}, opts ...dig.ProvideOption) {
	ProvideEByNamespace(defaultNamespace, constructor, opts...)
}

func Provide(constructor interface{}, opts ...dig.ProvideOption) error {
	return ProvideByNamespace(defaultNamespace, constructor, opts...)
}

func ProvideByNamespace(namespace string, constructor interface{}, opts ...dig.ProvideOption) error {
	instance := GetDig(namespace)
	return instance.Provide(constructor, opts...)
}

func ProvideEByNamespace(namespace string, constructor interface{}, opts ...dig.ProvideOption) {
	err := ProvideByNamespace(namespace, constructor, opts...)
	if err != nil {
		logger.Fatalf("[provide %s] %v", namespace, err)
	}
}

func Invoke(function interface{}, opts ...dig.InvokeOption) error {
	return InvokeByNamespace(defaultNamespace, function, opts...)
}

func InvokeE(function interface{}, opts ...dig.InvokeOption) {
	InvokeEByNamespace(defaultNamespace, function, opts...)
}

func InvokeByNamespace(namespace string, function interface{}, opts ...dig.InvokeOption) error {
	instance := GetDig(namespace)
	return instance.Invoke(function, opts...)
}

func InvokeEByNamespace(namespace string, function interface{}, opts ...dig.InvokeOption) {
	err := InvokeByNamespace(namespace, function, opts...)
	if err != nil {
		logger.Fatalf("[invoke %s] %v", namespace, err)
	}
}

func DecorateByNamespace(namespace string, decorator interface{}, opts ...dig.DecorateOption) error {
	instance := GetDig(namespace)
	return instance.Decorate(decorator, opts...)
}

func DecorateEByNamespace(namespace string, decorator interface{}, opts ...dig.DecorateOption) {
	err := DecorateByNamespace(namespace, decorator, opts...)
	if err != nil {
		logger.Fatalf("[decorate %s] %v", namespace, err)
	}
}

func Decorate(decorator interface{}, opts ...dig.DecorateOption) error {
	return DecorateByNamespace(defaultNamespace, decorator, opts...)
}

func DecorateE(decorator interface{}, opts ...dig.DecorateOption) {
	DecorateEByNamespace(defaultNamespace, decorator, opts...)
}

func ScopeByNamespace(namespace string, name string, opts ...dig.ScopeOption) *dig.Scope {
	instance := GetDig(namespace)
	return instance.Scope(name, opts...)
}

func Scope(name string, opts ...dig.ScopeOption) *dig.Scope {
	return ScopeByNamespace(defaultNamespace, name, opts...)
}
