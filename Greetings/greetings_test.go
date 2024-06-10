package greetings

import (
	"regexp"
	"testing"
)

// First test checks to see if passing a name results in a failure
func TestHelloName(t *testing.T) {
	name := "Akshay"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Akshay")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Akshay") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// Second test checks to see if NOT passing a name provides a successful response
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
