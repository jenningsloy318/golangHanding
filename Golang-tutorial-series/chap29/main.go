package main

import (  
    "fmt"
)

func finished() {  
    fmt.Println("Finished finding largest")
}


func largest(nums []int) {  
    defer finished()
    fmt.Println("Started finding largest")
    max := nums[0]
    for _, v := range nums {
        if v > max {
            max = v
        }
    }
    fmt.Println("Largest number in", nums, "is", max)
}


type person struct {  
    firstName string
    lastName string
}

func (p person) fullName() {  
    fmt.Printf("%s %s\n",p.firstName,p.lastName)
}

func printA(a int) {  
    fmt.Println("value of a in deferred function", a)
}

func main() {  
    nums := []int{78, 109, 2, 563, 300}
    largest(nums)
    // value para  the defer function 
    a := 5
    defer printA(a)
    a = 10
    fmt.Println("value of a before deferred function call", a)

    // defer method/struct 

    p := person {
        firstName: "John",
        lastName: "Smith",
    }
    defer p.fullName()
    fmt.Printf("Welcome ")  

    //
    name := "Naveen"
    fmt.Printf("Orignal String: %s\n", string(name))
    fmt.Printf("Reversed String: ")
    for _, v := range []rune(name) {
        defer fmt.Printf("%c", v)

}
}