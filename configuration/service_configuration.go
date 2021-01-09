package configuration

type Option func(configStruct interface{})

type IServiceConfiguration interface {
	Configure(configStruct interface{}, options ...Option)
}

type ServiceConfiguration struct {
}

func (c *ServiceConfiguration) Configure(configStruct interface{}, options ...Option) {
	for _, option := range options {
		option(configStruct)
	}
}
