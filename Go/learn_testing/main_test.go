package main

import (
	"regexp"
	"testing"
)

func TestPrintName(t *testing.T) {
	name := "Kelly"
	want := regexp.MustCompile("Kelly")
	msg, err := printName(name)
	if !want.MatchString(msg) || err != nil {
		t.Errorf("printName(%s) failed", name)
	}
}
