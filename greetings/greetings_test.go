package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Gladys"
	message, err := Hello("Gladys")
	want := regexp.MustCompile(`\b` + name + `\b`)
	if !want.MatchString(message) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v. Want match for %#q, nil`, message, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	message, err := Hello("")
	if message != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v. Want "", error`, message, err)
	}
}
