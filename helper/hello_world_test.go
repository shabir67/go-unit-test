package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Shobir")
	if result != "Hello Shobir" {
		// unit test failed
		panic("Result is not Hello Shobir")
	}
}
