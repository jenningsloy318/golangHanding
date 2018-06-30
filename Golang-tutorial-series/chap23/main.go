package main

import (  
		"fmt"
		"time"
)

func write(ch chan int) {  
	for i := 0; i < 5; i++ {
			ch <- i
			fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

func main() {  
    ch := make(chan string, 2)
    ch <- "naveen"
    ch <- "paul"
    fmt.Println(<- ch)
		fmt.Println(<- ch)
		//read/write a chan
		ch2 := make(chan int, 2)
    go write(ch2)
    time.Sleep(2 * time.Second)
    for v := range ch2 {
        fmt.Println("read value", v,"from ch")
        time.Sleep(2 * time.Second)

		}
		// capacity and lenth of chan
		ch3 := make(chan string, 3)
    ch3 <- "naveen"
    ch3 <- "paul"
    fmt.Println("capacity is", cap(ch3))
    fmt.Println("length is", len(ch3))
    fmt.Println("read value", <-ch3)
    fmt.Println("new length is", len(ch3))
}