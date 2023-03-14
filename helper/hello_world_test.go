package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMain(m *testing.M) {
	// Before Test
	fmt.Println("UNIT TEST STARTED")

	m.Run()

	// After Test
	fmt.Println("UNIT TEST FINISHED")
}

func TestTableHello(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Hello(Fadhil)",
			request:  "Fadhil",
			expected: "Hello Fadhil",
		},
		{
			name:     "Hello(KunK)",
			request:  "KunK",
			expected: "Hello KunK",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Hello(test.request)
			assert.Equal(t, result, test.expected, "Result is not matched")
		})
	}

}

func TestSubTest(t *testing.T) {
	t.Run("TestNameFadhil", func(t *testing.T) {
		result := Hello("Fadhil")
		assert.Equal(t, result, "Hello Fadhil", "Result must be 'Hello Fadhil'")
	})

	t.Run("TestNameKunK", func(t *testing.T) {
		result := Hello("KunK")
		assert.Equal(t, result, "Hello KunK", "Result must be 'Hello KunK'")
	})
}

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
