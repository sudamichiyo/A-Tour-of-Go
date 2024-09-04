package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	var s string = fmt.Sprint(z) //string型の場合、string(z)では型変換ができない
	fmt.Println(x, y, f, z, s)
}
