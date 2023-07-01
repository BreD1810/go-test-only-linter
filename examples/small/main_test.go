package main

import "testing"

func TestUsedFunction(t *testing.T) {
	res := usedFunction()
	if res != 1 {
		t.Fatal("usedFunction did not return 1")
	}
}

func TestNotUsedFunction(t *testing.T) {
	res := notUsedFunction()
	if res != 2 {
		t.Fatal("notUsedFunction did not return 1")
	}
}
