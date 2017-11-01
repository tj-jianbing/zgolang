package main

import (
	"fmt"
	"zgolang/Algorithm"
)


func main() {
	fmt.Println("ok")
	array := []int{12, 3, 45, 12, 44, 56, 89, 97, 87, 66, 57, 85, 63, 69, 27, 54, 54, 36, 83, 346, 34534, 15, 75645, 234235, 67845, 23446, 2346, 13, 3, 7, 2346, 24}
	fmt.Println(array)
	array02 :=Algorithm.OddEven(array)
	fmt.Println(array02)
}
