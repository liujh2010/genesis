package builders

type IContainerBuilder interface {
	RegisterComponent(componentName string) IComponentBuilder
	RegisterService(serviceName string) IComponentBuilder
}

type ContainerBuilder struct {
}

func (b *ContainerBuilder) RegisterComponent(componentName string) IComponentBuilder {
	return new(ComponentBuilder)
}
func (b *ContainerBuilder) RegisterService(serviceName string) IComponentBuilder {
	return new(ComponentBuilder)
}

type IComponentBuilder interface {
	Constructors(function interface{}) IComponentBuilder
	SingleInstence()
	PerOwner()
	PerCaller()
}

type ComponentBuilder struct {
}

func (b *ComponentBuilder) Constructors(function interface{}) IComponentBuilder {
	return b
}
func (b *ComponentBuilder) SingleInstence() {}
func (b *ComponentBuilder) PerOwner()       {}
func (b *ComponentBuilder) PerCaller()      {}
