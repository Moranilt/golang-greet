package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Danil"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Danil")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Danil) = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TEstHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
