package main

import (
	"testing"
)

func TestArray(t *testing.T) {
	var a [3]int
	for i := 0; i <= 2; i++ {
		a[i] = i
	}
	t.Logf("array a is %d\n", a)

	var b = [3]int{1, 2, 3}
	for i := 0; i < 3; i++ {
		t.Logf("b[%d] is %d\n", i, b[i])
	}

	for index, value := range b {

		t.Logf("index is %d, and value is %d\n", index, value)

	}
}
