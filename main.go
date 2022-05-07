package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world.")

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	back := arr[len(arr)-1] // Top
	arr = arr[:len(arr)-1]  // Pop
	arr = append([]int{back}, arr...)
	println(fmt.Sprintln(arr))
}
