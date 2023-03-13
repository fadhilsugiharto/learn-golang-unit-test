package helper

import "testing"

func TestHelloFadhil(t *testing.T) {
	result := Hello("Fadhil")

	if result != "Hello Fadhil" {
		// error
		t.Error("Result is not Hello Fadhil")
	}
}

func TestHelloKunK(t *testing.T) {
	result := Hello("KunK")

	if result != "Hello KunK" {
		// error
		t.Fatal("Result is not Hello KunK")
	}
}
