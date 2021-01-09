package builders

import (
	. "github.com/liujh2010/genesis/configuration"
	. "github.com/liujh2010/genesis/service"
)

type ServiceProviderContructors func() ServiceProvider

type IServiceProviderBuilder interface {
	ProvideService(constructors ServiceProviderContructors) IServiceProviderBuilder
	Configure(options ...Option) IServiceProviderBuilder
	Build()
}

type ServiceProviderBuilder struct {
	options      []Option
	constructors ServiceProviderContructors
}

func (b *ServiceProviderBuilder) ProvideService(constructors ServiceProviderContructors) IServiceProviderBuilder {
	b.constructors = constructors
	return b
}

func (b *ServiceProviderBuilder) Configure(options ...Option) IServiceProviderBuilder {
	b.options = options
	return b
}

func (b *ServiceProviderBuilder) Build() {
	serviceProvider := b.constructors()
	service := serviceProvider.New()
	serviceProvider.Configure(service, b.options...)
}

type ServiceProvider interface {
	IServiceConfiguration
	New() IService
}

type IApplicationBuilder interface {
	Use(name string) IServiceProviderBuilder
}

type ApplicationBuilder struct {
}

func (b *ApplicationBuilder) Use(name string) IServiceProviderBuilder {
	return new(ServiceProviderBuilder)
}
