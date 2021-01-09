package genesis

import (
	. "github.com/liujh2010/genesis/builders"
	. "github.com/liujh2010/genesis/environment"
)

func init() {

}

func StartUp() {}

func GetEnvironment() IEnvironment { return nil }

func Configure(fn func(builder IContainerBuilder)) {
	builder := new(ContainerBuilder)
	fn(builder)
}

func ConfigureApplication(fn func(app IApplicationBuilder, env IEnvironment)) {
	app := new(ApplicationBuilder)
	env := new(Environment)
	fn(app, env)
}
