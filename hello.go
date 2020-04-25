package main

import "fmt"

//main function here
//Each solution owns one file
func main() {
	for i := 0; i <= 30; i++ {
		fmt.Println(i, " : ", fib(i))
	}
}
