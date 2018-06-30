package main

import (
    "fmt"
)

func find(num int, nums ...int) {
    fmt.Printf("type of nums is %T\n", nums)
    found := false
    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            found = true
        }
    }
    if !found {
        fmt.Println(num, "not found in ", nums)
    }
    fmt.Printf("\n")
}



func change(s ...string) {  
  s[0] = "Go"
}

func changeappend(s ...string) {
  s[0] = "Go"
  s = append(s, "playground") // 当新的元素被添加到切片时，会创建一个新的数组。现有数组的元素被复制到这个新数组中，并返回这个新数组的新切片引用,但是老的slice引用的数组却没有发生改变
  fmt.Println(s)
}
func main() {
    find(89, 89, 90, 95)
    find(45, 56, 67, 45, 90, 109)
    find(78, 38, 56, 98)
    find(87)

    // 传递slice
    nums := []int{89, 90, 95}
    find(89, nums...)

    // 可变函数和slice传参

    welcome := []string{"hello", "world"}
    change(welcome...)
    fmt.Println(welcome)

    changeappend(welcome...)
    fmt.Println(welcome)
}