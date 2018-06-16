//https://studygolang.com/articles/11902
// Go 系列教程 —— 8. if-else 语句
package main

import (  
    "fmt"
)

func main() {  
    
    if num := 99; num <= 50 {
        fmt.Println("number is less than or equal to 50")
    } else if num >= 51 && num <= 100 {
        fmt.Println("number is between 51 and 100")
    } else {
        fmt.Println("number is greater than 100")
    }

}