package main

import "fmt"

func fibonacci(c, quit chan int) {
	fmt.Println("IN FUNCTION")
	x, y := 0, 1
	for {
		select {
		case c <- x:

			x, y = y, x+y
			fmt.Println("fibonacci", x)
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func log(c chan int, quit chan int) {

	for i := 1; i < 10; i++ {
		fmt.Println("Log function", i)
		fmt.Println(<-c)
	}
	quit <- 0
}
func main() {
	c := make(chan int)
	quit := make(chan int)
	go log(c, quit)
	fibonacci(c, quit)
}
