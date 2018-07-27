package detect

import (
	"testing"
)

func TestMaxLenA(t *testing.T) {
	if res := MaxLen("abc", "abcdefg"); res == 7 {
		t.Log("the test pass.")
	} else {
		t.Error("the test failed.")
	}
}

func TestMaxLenB(t *testing.T) {
	if res := MaxLen("abc", "dfg"); res == 3 {
		t.Log("the test pass.")
	} else {
		t.Error("the test failed.")
	}
}

func TestMaxLenC(t *testing.T) {
	if res := MaxLen("abcdefgh", "dfg"); res == 8 {
		t.Log("the test pass.")
	} else {
		t.Error("the test failed.")
	}
}
