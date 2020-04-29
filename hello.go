package main

import "fmt"

//main function here
//Each solution owns one file
func main() {
	target := 12
	position := []int{10, 8, 0, 5, 3}
	speed := []int{2, 4, 1, 1, 3}
	fmt.Println(carFleet(target, position, speed))
}
