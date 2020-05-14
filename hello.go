package main

import "fmt"

//main function here
//Each solution owns one file
func main() {
	stack := Constructor()
	stack.Push(2)
	stack.Push(1)
	stack.Push(3)
	tmp := stack.next
	for tmp != nil {
		fmt.Println(tmp)
		tmp = tmp.next
	}
}
