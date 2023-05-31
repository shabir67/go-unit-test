package helper

import (
	"fmt"
	"testing"
)

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
