package main

import (
	"testing"
)

func TestArray1(t *testing.T) {
	var a [3]int
	for i := 0; i <= 2; i++ {
		a[i] = i
	}
	t.Log("array a is", a)

	var b = [3]int{1, 2, 3}
	for i := 0; i < 3; i++ {
		t.Logf("b[%d] is %d", i, b[i])
	}

	for index, value := range b {

		t.Logf("array b , index is %d, and value is %d\n", index, value)

	}
}
