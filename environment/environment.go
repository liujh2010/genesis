package environment

type IEnvironment interface {
	IsDev() bool
	IsStg() bool
	IsProd() bool
	ReloadFromFile()
	ReloadFromSystemEnvironment()
	GetString(key string) string
	GetSlice(key string) []string
	GetInt(key string) int
	GetFloat(key string) float32
}

type Environment struct {
}

func (env *Environment) IsDev() bool {
	return false
}
func (env *Environment) IsStg() bool {
	return false
}
func (env *Environment) IsProd() bool {
	return false
}
func (env *Environment) GetString(key string) string {
	return "false"
}
func (env *Environment) GetSlice(key string) []string {
	return []string{""}
}
func (env *Environment) GetInt(key string) int {
	return 1
}
func (env *Environment) GetFloat(key string) float32 {
	return 1
}

func (env *Environment) ReloadFromFile()              {}
func (env *Environment) ReloadFromSystemEnvironment() {}

type IDefualt interface {
	SetString(key, value string)
	SetSlice(Key string, value []string)
	SetInt(key string, value int)
	SetFloat(key string, value float32)
}
