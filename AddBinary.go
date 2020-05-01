package main

import "fmt"

/*
Given two binary strings, return their sum (also a binary string).

The input strings are both non-empty and contains only characters 1 or 0.

Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"


Constraints:

Each string consists only of '0' or '1' characters.
1 <= a.length, b.length <= 10^4
Each string is either "0" or doesn't contain any leading zero.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-binary
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func harfAdd(A byte, B byte, C byte) (byte, byte) {
	var S byte
	var Carry byte
	S = A ^ B ^ C
	Carry = (A & B) | (A & C) | (B & C)
	return S, Carry
}

func addBinary(a string, b string) string {
	length := len(b) - len(a)
	if length > 0 {
		for i := 0; i < length; i++ {
			a = "0" + a
		}
	}
	if length < 0 {
		for i := 0; i < (-length); i++ {
			b = "0" + b
		}
	}
	//println(a, b)
	carry := byte(0)
	length = len(b)
	var S byte
	ret_str := ""
	for i := length - 1; i >= 0; i-- {
		A := a[i] - 48
		B := b[i] - 48
		//fmt.Println(A, B, carry)
		S, carry = harfAdd(A, B, carry)
		//fmt.Println(S, carry)
		ret_str = fmt.Sprintf("%d", S) + ret_str
	}
	if carry == 1 {
		ret_str = fmt.Sprintf("%d", carry) + ret_str
	}
	fmt.Println(ret_str)
	return ret_str
}

/*
想象一下数电里学过的累加器
每一位等于两位相加加上上一位的进位
注意关注最后一位有进位的情况
*/
