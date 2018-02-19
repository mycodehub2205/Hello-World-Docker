package main

import (
	"fmt"

	"github.com/synoa/helloworld"
)

func main() {
	gr := helloworld.NewGreeter("Hello")
	fmt.Println(gr.Greet("World"))
}
