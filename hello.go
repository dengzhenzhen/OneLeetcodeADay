package main

import "fmt"

//main function here
//Each solution owns one file
func main() {
	l1 := ListNode{0, nil}
	l2 := ListNode{0, nil}
	fmt.Println(mergeTwoLists(&l1, &l2))
}
