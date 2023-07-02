package stuff_test

import (
	"test-only-example-receiver/pkg/stuff"
	"testing"
)

func TestUsedFunction(t *testing.T) {
	s := stuff.NewStuffStruct()
	res := s.UsedFunction()
	if res != "testStruct" {
		t.Fatal("UsedFunction did not return testStruct")
	}
}

func TestNotUsedFunction(t *testing.T) {
	s := stuff.NewStuffStruct()
	res := s.NotUsedFunction()
	if res != "testStruct" {
		t.Fatal("NotUsedFunction did not return testStruct")
	}
}
