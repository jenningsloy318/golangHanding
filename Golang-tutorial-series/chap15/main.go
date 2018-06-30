package main

import (
    "fmt"
)


func change(val *int) {  
    *val = 55
}
func modify3(sls []int) {
	sls[0] = 90
}

func main() {
    //declare pointer a 
    b := 255
    var a *int 
    fmt.Printf("Type of a is %T\n", a)
    fmt.Println("a is", a)
    // initialize pointer a 
    a  = &b
    fmt.Println("a after initialization is", a)
    fmt.Println("value of b is", *a)
    // modify value of  b via pointer a 
    *a++
    fmt.Println("new value of b is", b)
    fmt.Println("new value of b is", *a)


   // pass pointer to function
   c := 58
   fmt.Println("value of a before function call is",c)
   d := &c  // d is a pointer
   change(d)
   fmt.Println("value of a after function call is", c)
    //	//slice as an argument to a function
   o := [3]int{89, 90, 91}
   fmt.Println("before modified", o)
   modify3(o[:])
   fmt.Println("modify3 using slice", o)
}