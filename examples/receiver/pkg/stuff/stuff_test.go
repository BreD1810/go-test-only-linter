package stuff_test

import (
	stuff "test-only-example-big/pkg"
	"testing"
)

func TestUsedFunction(t *testing.T) {
	res := stuff.UsedFunction()
	if res != 1 {
		t.Fatal("usedFunction did not return 1")
	}
}

func TestNotUsedFunction(t *testing.T) {
	res := stuff.NotUsedFunction()
	if res != 2 {
		t.Fatal("notUsedFunction did not return 1")
	}
}
