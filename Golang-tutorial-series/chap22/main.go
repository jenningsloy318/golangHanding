package main

import (  
		"fmt"
		"time"
)

func hello(done chan bool) {  
    fmt.Println("Hello world goroutine")
    done <- true
}


func hello2(done chan bool) {  
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}

// 

func calcSquares(number int, squareop chan int) {  
	sum := 0
	for number != 0 {
			digit := number % 10
			sum += digit * digit
			number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {  
	sum := 0 
	for number != 0 {
			digit := number % 10
			sum += digit * digit * digit
			number /= 10
	}
	cubeop <- sum
} 

// chan conversion
func sendData(sendch chan<- int) {  
	sendch <- 10
}
// get data from chan 
func producer(chnl chan int) {  
	for i := 0; i < 10; i++ {
			chnl <- i
	}
	close(chnl)
}
func main() {  
    done := make(chan bool)
    go hello(done)
    <- done
		fmt.Println("main function")
		// more details explanation
		fmt.Println("Main going to call hello2 go goroutine")
    go hello2(done)
    <- done
    fmt.Println("Main received data")
	 // more details of chan
	 number := 589
	 sqrch := make(chan int)
	 cubech := make(chan int)
	 go calcSquares(number, sqrch)
	 go calcCubes(number, cubech)
	 squares, cubes := <-sqrch, <-cubech
	 fmt.Println("Final output", squares + cubes)

	 // chan conversion 
	 cha1 := make(chan int)
	 go sendData(cha1)
	 fmt.Println(<-cha1)
	// for rang chan
	 ch := make(chan int)
	 go producer(ch)
	 for v := range ch {
			 fmt.Println("Received ",v)
	 }
}