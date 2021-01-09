package example_test

import (
	"testing"

	. "github.com/liujh2010/genesis"
	. "github.com/liujh2010/genesis/builders"
	. "github.com/liujh2010/genesis/configuration"
	. "github.com/liujh2010/genesis/environment"
	. "github.com/liujh2010/genesis/service"
)

type Logging struct{}
type ILogging interface{}

type HTTPServer struct {
	UseHTTPS bool
}

func (s *HTTPServer) Run()  {}
func (s *HTTPServer) Stop() {}

type HTTPServerProvider struct {
	ServiceConfiguration
}

func (p *HTTPServerProvider) New() IService {
	return new(HTTPServer)
}
func UseHTTPS() Option {
	return func(configStruct interface{}) {
		configStruct.(*HTTPServer).UseHTTPS = true
	}
}

func Test_Example_2(t *testing.T) {

	t.Log("Example_2")
	// env := GetEnvironment()

	// Configure(func(builder IContainerBuilder) {

	// 	if env.IsDev() {
	// 		builder.
	// 			RegisterComponent("logging").
	// 			Constructors(func() ILogging {
	// 				return *&Logging{}
	// 			}).
	// 			SingleInstence()
	// 	}

	// 	builder.
	// 		RegisterComponent("logging").
	// 		Constructors(func() ILogging {
	// 			return *&Logging{}
	// 		}).
	// 		SingleInstence()

	// })

	ConfigureApplication(func(app IApplicationBuilder, env IEnvironment) {
		app.
			Use("HTTPService").
			ProvideService(func() ServiceProvider {
				return new(HTTPServerProvider)
			}).
			Configure(UseHTTPS()).
			Build()
	})

	StartUp()
}
