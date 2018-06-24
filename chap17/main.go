
package main

import (
    "fmt"
    "math"
)

type Employee struct {
    name     string
    salary   int
    currency string
}

/*
  displaySalary() 方法将 Employee 做为接收器类型
*/
func (e Employee) displaySalary() {
    fmt.Printf("working with method,Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}


func displaySalary2(e Employee) {
    fmt.Printf("working with function,Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}


type Rectangle struct {
    length int
    width  int
}

type Circle struct {
    radius float64
}

func (r Rectangle) Area() int {
    return r.length * r.width
}




func (c Circle) Area() float64 {
    return math.Pi * c.radius * c.radius
}


type Employee2 struct {
    name string
    age  int
}

/*
使用值接收器的方法。
*/
func (e Employee2) changeName(newName string) {
    e.name = newName
}

/*
使用指针接收器的方法。
*/
func (e *Employee2) changeAge(newAge int) {
    e.age = newAge
}


type address struct {
    city  string
    state string
}

func (a address) fullAddress() {
    fmt.Printf("Full address: %s, %s\n", a.city, a.state)
}

type person struct {
    firstName string
    lastName  string
    address
}



type rectangle struct {
    length int
    width  int
}
func area(r rectangle) {
    fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

func (r rectangle) area() {
    fmt.Printf("Area Method result: %d\n", (r.length * r.width))
}
func perimeter(r *rectangle) {
    fmt.Println("perimeter function output:", 2*(r.length+r.width))

}

func (r *rectangle) perimeter() {
    fmt.Println("perimeter method output:", 2*(r.length+r.width))
}


type myInt int

func (a myInt) add(b myInt) myInt {
    return a + b
}

func main() {
    
    emp1 := Employee {
        name:     "Sam Adolf",
        salary:   5000,
        currency: "$",
    }
    //调用 Employee 类型的 displaySalary() 方法
    emp1.displaySalary() 

    //使用函数
    displaySalary2(emp1)

    // different struct, can define same name of method, but the content of method can be different

    r := Rectangle{
        length: 10,
        width:  5,
    }
    fmt.Printf("Area of rectangle %d\n", r.Area())

    c := Circle{
        radius: 12,
    }
    fmt.Printf("Area of circle %f\n", c.Area())

    //指针接收器和值接收器

    e := Employee2{
        name: "Mark Andrew",
        age:  50,
    }
    fmt.Printf("Employee name before change: %s", e.name)
    e.changeName("Michael Andrew")
    fmt.Printf("\nEmployee name after change: %s\n", e.name)

    fmt.Printf("\n\nEmployee age before change: %d", e.age)
    e.changeAge(51)
    fmt.Printf("\nEmployee age after change: %d\n", e.age)

    //anonymous fileds
    p := person{
        firstName: "Elon",
        lastName:  "Musk",
        address: address {
            city:  "Los Angeles",
            state: "California",
        },
    }

    p.fullAddress() //访问 address 结构体的 fullAddress 方法
     

    // 当一个函数有一个值参数，它只能接受一个值参数。当一个方法有一个值接收器，它可以接受值接收器和指针接收器。
    r2 := rectangle{
        length: 10,
        width:  5,
    }
    area(r2)
    r2.area()

    p2 := &r2
    /*
       compilation error, cannot use p (type *rectangle) as type rectangle
       in argument to area
    */
    //area(p)

    p2.area()//通过指针调用值接收器



    // 在方法中使用指针接收器 与 在函数中使用指针参数
    r3 := rectangle{
        length: 10,
        width:  5,
    }
    p3 := &r3 //pointer to r
    perimeter(p3)
    p3.perimeter()

    /*
        cannot use r (type rectangle) as type *rectangle in argument to perimeter
    */
    //perimeter(r)

    r3.perimeter()//使用值来调用指针接收器


    // 使用其他类型来定义方法

    num1 := myInt(5)
    num2 := myInt(10)
    sum := num1.add(num2)
    fmt.Println("Sum is", sum)
}
