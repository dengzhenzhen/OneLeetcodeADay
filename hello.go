package main

import (
	"fmt"
)

//main function here
//Each solution owns one file
func main() {
	fmt.Println(constructRectangle(2))
}

func printbinary(n int) {
	ret_str := ""
	for n > 0 {
		if n&1 == 1 {
			ret_str = "1" + ret_str
		} else {
			ret_str = "0" + ret_str
		}
		n >>= 1
	}
	fmt.Println(ret_str)
}
