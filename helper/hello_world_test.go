package helper

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Fadhil")

	if result != "Hello Fadhil" {
		//error
		panic("Result is not 'Hello Fadhil'")
	}
}
