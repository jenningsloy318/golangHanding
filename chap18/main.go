package main

import (  
    "fmt"
)

type SalaryCalculator interface {  
    CalculateSalary() int
}

type Permanent struct {  
    empId    int
    basicpay int
    pf       int
}

type Contract struct {  
    empId  int
    basicpay int
}

//salary of permanent employee is sum of basic pay and pf
func (p Permanent) CalculateSalary() int {  
    return p.basicpay + p.pf
}

//salary of contract employee is the basic pay alone
func (c Contract) CalculateSalary() int {  
    return c.basicpay
}

/*
total expense is calculated by iterating though the SalaryCalculator slice and summing  
the salaries of the individual employees  
*/
func totalExpense(s []SalaryCalculator) {  
    expense := 0
    for _, v := range s {
        expense = expense + v.CalculateSalary()
    }
    fmt.Printf("Total Expense Per Month $%d\n", expense)
}

//internal of interface 
type Test interface {  
    Tester()
}

type MyFloat float64

func (m MyFloat) Tester() {  
    fmt.Println(m)
}

func describe(t Test) {  
    fmt.Printf("Interface type %T value %v\n", t, t)
}

// interface assert
func assert(i interface{}) {  
    s := i.(int) //get the underlying int value from i
    fmt.Println(s)
}

func assert2(i interface{}) {  
    v, ok := i.(int)
    fmt.Println(v, ok)
}
//type switch 
func findType(i interface{}) {  
    switch i.(type) {
    case string:
        fmt.Printf("I am a string and my value is %s\n", i.(string))
    case int:
        fmt.Printf("I am an int and my value is %d\n", i.(int))
    default:
        fmt.Printf("Unknown type\n")
    }
}


// compare type and interface 

type Describer interface {  
    Describe()
}
type Person struct {  
    name string
    age  int
}

func (p Person) Describe() {  
    fmt.Printf("%s is %d years old", p.name, p.age)
}

func findType2(i interface{}) {  
    switch v := i.(type) {
    case Describer:
        v.Describe()
    default:
        fmt.Printf("unknown type\n")
    }
}
func main() {  
    // interface usage
    pemp1 := Permanent{1, 5000, 20}
    pemp2 := Permanent{2, 6000, 30}
    cemp1 := Contract{3, 3000}
    employees := []SalaryCalculator{pemp1, pemp2, cemp1}
    totalExpense(employees)
    // internal of interface 
    var t Test
    f := MyFloat(89.7)
    t = f
    describe(t)
    describe(f)
    t.Tester()
    f.Tester()

    //interface assert 
    var s interface{} = 56
    assert(s)

    var s2 interface{} = 56
    assert2(s2)
    var i2 interface{} = "Steven Paul"
    assert2(i2)

    // type switch 
    findType("Naveen")
    findType(77)
    findType(89.98)

    // compare type and interface 
    findType2("Naveen")
    p := Person{
        name: "Naveen R",
        age:  25,
    }
    findType2(p)
}   
