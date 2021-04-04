package util

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello golang"
	fmt.Println("want :", want)
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestHi(t *testing.T) {
	want := "Hi golang"
	fmt.Println("want :", want)
	if got := Hi(); got != want {
		t.Errorf("Hi() = %q, want %q", got, want)
	}
}
