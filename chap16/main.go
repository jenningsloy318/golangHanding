package main

import (  
    "fmt"
    "./computer"
)

type Employee struct {  
    firstName, lastName string
    age, salary         int
}


type Person struct {  
    string
    int
}

type Address struct {  
    city, state string
}
type Person2 struct {  
    name string
    age int
    address Address
}

type Person3 struct {  
    name string
    age int
    Address
}


type name struct {  
    firstName string
    lastName string
}

func main() {

    //creating structure using field names
    emp1 := Employee{
        firstName: "Sam",
        age:       25,
        salary:    500,
        lastName:  "Anderson",
    }

    //creating structure without using field names
    emp2 := Employee{"Thomas", "Paul", 29, 800}

    fmt.Println("Employee 1", emp1)
    fmt.Println("Employee 2", emp2)

    // access the value of struct 
    emp6 := Employee{"Sam", "Anderson", 55, 6000}
    fmt.Println("First Name:", emp6.firstName)
    fmt.Println("Last Name:", emp6.lastName)
    fmt.Println("Age:", emp6.age)
    fmt.Printf("Salary: $%d\n", emp6.salary)
    
    // pointer to struct 
    emp9 := &Employee{"Sam", "Anderson", 55, 6000}
    fmt.Println("First Name:", (*emp9).firstName)
    fmt.Println("Age:", (*emp9).age)
    // pointer to struct simplified reference

    emp10 := &Employee{"Sam", "Anderson", 55, 6000}
    fmt.Println("First Name:", emp10.firstName)
    fmt.Println("Age:", emp10.age)

    // Anonymous struct
    p := Person{"Naveen", 50}
    fmt.Println(p)

    //  Anonymous struct  filed name is its type name
    var p1 Person
    p1.string = "naveen"
    p1.int = 50
    fmt.Println(p1)
    // nested struct 

    var p3 Person2
    p3.name = "Naveen"
    p3.age = 50
    p3.address = Address {
        city: "Chicago",
        state: "Illinois",
    }
    fmt.Println("Name:", p3.name)
    fmt.Println("Age:",p3.age)
    fmt.Println("City:",p3.address.city)
    fmt.Println("State:",p3.address.state)

    // promoted fileds
    var p4 Person3
    p4.name = "Naveen"
    p4.age = 50
    p4.Address = Address{
        city:  "Chicago",
        state: "Illinois",
    }
    fmt.Println("Name:", p4.name)
    fmt.Println("Age:", p4.age)
    fmt.Println("City:", p4.city) //city is promoted field
    fmt.Println("State:", p4.state) //state is promoted field

    // exported struct
    var spec computer.Spec
    spec.Maker = "apple"
    spec.Price = 50000
    fmt.Println("Spec:", spec)

    // struct equality 
    name1 := name{"Steve", "Jobs"}
    name2 := name{"Steve", "Jobs"}
    if name1 == name2 {
        fmt.Println("name1 and name2 are equal")
    } else {
        fmt.Println("name1 and name2 are not equal")
    }

    name3 := name{firstName:"Steve", lastName:"Jobs"}
    name4 := name{}
    name4.firstName = "Steve"
    if name3 == name4 {
        fmt.Println("name3 and name4 are equal")
    } else {
        fmt.Println("name3 and name4 are not equal")
    }
}   
