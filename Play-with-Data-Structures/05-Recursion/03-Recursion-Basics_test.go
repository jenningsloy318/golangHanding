package recursion

import "testing"

func Sum(array []int) int {
	return sum(array, 0)
}

//comupte sum of  array[lindex,n), as sum from lindex to n-1
func sum(array []int, lindex int) int {

	// when lindex == lenth, it is out of the array index, since max index is lenth-1,
	// first step is a basic problem, here this is the very basic problem
	if lindex == len(array) {
		return 0
	}
	// start recursion
	return array[lindex] + sum(array, lindex+1)
}

func TestRecursion(t *testing.T) {
	var num []int

	for i := 0; i <= 10000; i++ {
		num = append(num, i)
	}
	t.Logf("Sum of the array is: %d", Sum(num))
}
