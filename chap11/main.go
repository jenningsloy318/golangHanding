package main

import (
    "fmt"
)

func main() {

		a := [3]int{12, 78, 50} // short hand declaration to create array
		fmt.Println(a)

		b := [...]int{44, 88, 52} // ... makes the compiler determine the length
    fmt.Println(b)


		numa := [3]int{78, 79 ,80}
    nums1 := numa[:] // creates a slice which contains all elements of the array
    nums2 := numa[:]
    fmt.Println("array before change", numa)
    nums1[0] = 100
    fmt.Println("array after modification to slice nums1(set nums1[0] = 100) ", numa)
    nums2[1] = 101
		fmt.Println("array after modification to slice nums2(set nums2[1] = 101)", numa)
		

    darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
    dslice := darr[2:5]
    fmt.Println("array before", darr)
    for i := range dslice {  // range can also iterate over just the keys of a map.
        dslice[i]++
    }
		fmt.Println("array after dslice[i]++ " , darr)
		



}
