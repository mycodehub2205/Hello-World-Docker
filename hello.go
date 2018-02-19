package helloworld

// Greeter is a common interface to greet things
type Greeter struct {
	Greeting string
}

// NewGreeter creates a new Greeter
func NewGreeter(greeting string) Greeter {
	return Greeter{Greeting: greeting}
}

// Greet greets a person or thing
func (g *Greeter) Greet(name string) string {
	return g.Greeting + " " + name + "!"
}
