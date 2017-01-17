package greeting

type Greeter interface {
	Greet(name string) string
}