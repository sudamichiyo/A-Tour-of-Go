package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	//Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	//Extend itd length.
	s = s[:4]
	printSlice(s)

	//Drop its first two values.
	s = s[2:]
	printSlice(s)

	// s = s[1:3]
	// printSlice(s)

	s = append(s, 1, 2, 3)
	printSlice(s)

	s = append(s, 4, 5, 6, 7)
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
