package main

/*
Write an algorithm to determine if a number n is "happy".

A happy number is a number defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.

Return True if n is a happy number, and False if not.

Example:

Input: 19
Output: true
Explanation:
1^2 + 9^2 = 82
8^2 + 2^2 = 68
6^2 + 8^2 = 100
1^2 + 0^2 + 0^2 = 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/happy-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isHappy(n int) bool {
	record := map[int]bool{}
	for n != 1 {
		_, exist := record[n]
		if exist {
			return false
		} else {
			record[n] = true
			n = change(n)
		}
	}
	return true
}

func change(n int) int {
	var tmp int
	ret := 0
	for tmp = 10; n/tmp > 0; tmp *= 10 {
		x := ((n % tmp) / (tmp / 10)) * ((n % tmp) / (tmp / 10))
		ret += x

	}
	ret += ((n % tmp) / (tmp / 10)) * ((n % tmp) / (tmp / 10))
	return ret
}
