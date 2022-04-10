package main

import (
	"fmt"
	"time"
)

//main function here
//Each solution owns one file
var a int = 10

func main() {
	ret := rob([]int{2, 7, 9, 3, 1})
	fmt.Println(ret)
}

func printNum(num int, c *chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println(num)
		time.Sleep(100 * time.Millisecond)
	}
	*c <- 0
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
