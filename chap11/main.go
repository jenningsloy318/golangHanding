package main

import (
    "fmt"
)

// 
func changeLocal(num [5]int) {
  num[0] = 55
  fmt.Println("inside function ", num)
}

//
func subtactOne(numbers []int) {
  for i := range numbers {
      numbers[i] -= 2
  }
}
// 内存优化
func countries() []string {
  countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
  neededCountries := countries[:len(countries)-2]
  countriesCpy := make([]string, len(neededCountries))
  copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
  return countriesCpy
}

func main() {
    //初始化并赋值
		a := [3]int{12, 78, 50} // short hand declaration to create array
		fmt.Println(a)
    //编译器为你自动计算长度
		b := [...]int{44, 88, 52} // ... makes the compiler determine the length
    fmt.Println(b)

    //数组 值传递
    num := [...]int{5, 6, 7, 8, 8}
    fmt.Println("before passing to function ", num)
    changeLocal(num) //num is passed by value
    fmt.Println("after passing to function ", num)

    // 修改slice其实修改的是底层的数组
		numa := [3]int{78, 79 ,80}
    nums1 := numa[:] // creates a slice which contains all elements of the array
    nums2 := numa[:]
    fmt.Println("array before change", numa)
    nums1[0] = 100
    fmt.Println("array after modification to slice nums1(set nums1[0] = 100) ", numa)
    nums2[1] = 101
		fmt.Println("array after modification to slice nums2(set nums2[1] = 101)", numa)
		
  // range 和slice, 修改了slice即是修改了数组
    darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
    dslice := darr[2:5]
    fmt.Println("array before", darr)
    for i := range dslice {  // range can also iterate over just the keys of a map.
        dslice[i]++
    }
		fmt.Println("array after dslice[i]++ " , darr)
		

    // slice的长度和容量

    fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
    fruitslice := fruitarray[1:3]
    fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice))

    // make 创建slice
    i := make([]int, 5, 5)
    fmt.Printf("The slice is %d\n",i)


    // append slice 元素

    cars := []string{"Ferrari", "Honda", "Ford"}
    fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) // capacity of cars is 3
    cars = append(cars, "Toyota")
    fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) // capacity of cars is doubled to 6
    // nil slice 

    var names []string //zero value of a slice is nil
    if names == nil {
        fmt.Println("slice is nil going to append")
        names = append(names, "John", "Sebastian", "Vinay")
        fmt.Println("names contents:",names)
    }

    //... append a slice    to another
    veggies := []string{"potatoes", "tomatoes", "brinjal"}
    fruits := []string{"oranges", "apples"}
    food := append(veggies, fruits...)
    fmt.Println("food:",food)

    // slice 作为函数参数，

    nos := []int{8, 7, 6}
    fmt.Println("slice before function call", nos)
    subtactOne(nos)                               // function modifies the slice
    fmt.Println("slice after function call", nos) // modifications are visible outside

    // 内存优化
    countriesNeeded := countries()
    fmt.Println(countriesNeeded)


}
