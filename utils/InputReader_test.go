package utils

import (
	"testing"
)

func TestReadLines(t *testing.T) {
	except := "If you see this line - the parser is working fine"
	result := ReadLines("test", "test")
	if except != result[0] {
		t.Fail()
	}
}
