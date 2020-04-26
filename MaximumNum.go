package main

/*
编写一个方法，找出两个数字a和b中最大的那一个。不得使用if-else或其他比较运算符。

示例：

输入： a = 1, b = 2
输出： 2
*/

func maximum(a int, b int) int {
	tmpArray := [2]int{a, b}
	return tmpArray[((a-b)>>63)&1]
}

func getSymbolDigit(n int) int {
	return (n >> 63) & 1
}

/*
if a-b>0 : return a
if a-b<0 : return b
a-b>0 => The first digit of a-b equals to 1
a-b<0 => The first digit of a-b equals to 0
So make an array [a, b]. Use the SymbolDigit as the key
*/
