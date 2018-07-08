package main

import (
    "fmt"
)

type order struct {
    ordId      int
    customerId int
}

func main() {
    o := order{
        ordId:      1234,
        customerId: 567,
    }
    fmt.Println(o)
}
