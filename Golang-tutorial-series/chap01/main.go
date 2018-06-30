//指针学习
// *p 指针代表了变量，对它改值，就相当于对变量赋值了
package main

import (
	"fmt"
	"flag"
	"strings"
)
 
var n = flag.Bool("n",false,"omit the trailling newline")
var sep = flag.String("s"," ","separator")

func main () {
	x := 1
	p := &x 
	fmt.Println("var x address is",p)
	fmt.Println("var x  value is",*p)
	*p = 2// reset the value of x to 2 
	fmt.Println("var x  value is",*p)


	flag.Parse()
	fmt.Println(*sep)
	for _, args := range flag.Args() {

		fmt.Print(args)
	}

	
	fmt.Println(strings.Join(flag.Args(),*sep))
	if !*n {
		fmt.Println()
	}
	}
