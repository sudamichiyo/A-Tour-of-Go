package main

import "fmt"

func def() {
	defer fmt.Println("世界")

	fmt.Println("こんにちは")
}

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
	def()
}
