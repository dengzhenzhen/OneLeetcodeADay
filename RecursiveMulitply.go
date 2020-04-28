package main

/*
Write a recursive function to multiply two positive integers without using the * operator. You can use addition, subtraction, and bit shifting, but you should minimize the number of those operations.

Example 1:

 Input: A = 1, B = 10
 Output: 10
Example 2:

 Input: A = 3, B = 4
 Output: 12
Note:

The result will not overflow.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/recursive-mulitply-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func multiply(A int, B int) int {
	if A > B {
		return multiply(B, A)
	}
	if A == 1 {
		return B
	} else if B == 1 {
		return A
	} else {
		return multiply(A/2, B) + (multiply(A-A/2, B))
	}
}
