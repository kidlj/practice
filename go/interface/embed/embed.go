package embed

type TestInterface interface {
	Put() string
}

type Test struct {
	TestInterface
}
