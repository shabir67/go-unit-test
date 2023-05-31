package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Shobir")
	if result != "Hello Shobir" {
		// error
		t.Fail()
	}
	fmt.Println("Test HelloWorld Shobir Done")
}

func TestHelloWorldJoko(t *testing.T) {
	result := HelloWorld("Joko")
	if result != "Hello Joko" {
		// error
		t.FailNow()
	}
	fmt.Println("Test HelloWorld Joko Done")
}
