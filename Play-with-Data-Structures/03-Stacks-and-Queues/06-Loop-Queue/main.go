package main

import (
	"fmt"
	"queue"
)

func main() {

	var loopQueue = queue.NewDefaultQueue()

	for i := 0; i <= 15; i++ {
		loopQueue.Enqueue(i)
		fmt.Println(loopQueue.ToString())
		if i%3 == 2 {
			loopQueue.Dequeue()
			fmt.Println(loopQueue.ToString())
		}
	}
}
