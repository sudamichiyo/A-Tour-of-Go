package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	// ch := make(chan int, 1)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
