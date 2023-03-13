package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHelloAssert(t *testing.T) {
	result := Hello("Fadhil")

	assert.Equal(t, result, "Hello Fadhil", "Result must be 'Hello Fadhil'")
	fmt.Println("TestHelloAssert Finished")
}

func TestHelloRequire(t *testing.T) {
	result := Hello("Fadhil")

	require.Equal(t, result, "Hello Fadhil", "Result must be 'Hello Fadhil'")
	fmt.Println("TestHelloRequire Finished")
}

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
