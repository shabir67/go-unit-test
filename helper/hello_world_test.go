package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(SBY)",
			request:  "SBY",
			expected: "Hello SBY",
		},
		{
			name:     "HelloWorld(AHY)",
			request:  "AHY",
			expected: "Hello AHY",
		},
		{
			name:     "HelloWorld(IBAS)",
			request:  "IBAS",
			expected: "Hello IBAS",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}

}

func TestSubTest(t *testing.T) {
	t.Run("Puan", func(t *testing.T) {
		result := HelloWorld("Puan")
		require.Equal(t, "Hello Puan", result, "Result Must be 'Hello Puan")
	})
	t.Run("Mega", func(t *testing.T) {
		result := HelloWorld("Mega")
		require.Equal(t, "Hello Mega", result, "Result Must be 'Hello Mega")
	})

}

func TestMain(m *testing.M) {
	//before(this can include database connection test)
	fmt.Println("Before Unit Test")

	m.Run()

	//after
	fmt.Println("After unit test")

}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot unit test run in Mac")
	}

	result := HelloWorld("Ganjar")
	require.Equal(t, "Hello Ganjar", result)
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Prabowo")
	require.Equal(t, "Hello Prabowo", result, "Result must be 'Hello Prabowo'")
	fmt.Println("Test HelloWorldRequire with Require Done")

}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Prabowo")
	assert.Equal(t, "Hello Prabowo", result, "Result must be 'Hello Prabowo'")
	fmt.Println("Test HelloWorldAssert with Assert Done")

}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Shobir")
	if result != "Hello Shobir" {
		// error
		t.Error("Result mus be 'Hello Shobir'")
	}
	fmt.Println("Test HelloWorld Shobir Done")
}

func TestHelloWorldJoko(t *testing.T) {
	result := HelloWorld("Joko")
	if result != "Hello Joko" {
		// error
		t.Fatal("Result mus be 'Hello Joko'")
	}
	fmt.Println("Test HelloWorld Joko Done")
}
