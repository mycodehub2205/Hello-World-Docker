package helloworld_test

import (
	"testing"

	"github.com/synoa/helloworld"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		greeter helloworld.Greeter
	}{
		{"Spanish", "Kevin", "Hola Kevin!", helloworld.NewGreeter("Hola")},
		{"German", "Kevin", "Hallo Kevin!", helloworld.NewGreeter("Hallo")},
		{"English", "Kevin", "Hello Kevin!", helloworld.NewGreeter("Hello")},
		{"French", "Kevin", "Bonjour Kevin!", helloworld.NewGreeter("Bonjour")},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if res := tc.greeter.Greet(tc.input); res != tc.want {
				t.Errorf("expected %s, got %s", tc.want, res)
			}
		})
	}
}
