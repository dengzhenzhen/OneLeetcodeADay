package main

/*
Given an integer, return its base 7 string representation.

Example 1:
Input: 100
Output: "202"
Example 2:
Input: -7
Output: "-10"
Note: The input will be in range of [-1e7, 1e7].

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/base-7
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func convertToBase7(num int) string {
	var n int
	if num < 0 {
		n = -num
	} else {
		n = num
	}
	ret_str := ""
	for n > 0 {
		tmp := n % 7
		n /= 7
		int_tmp := tmp | '0'
		ret_str = string(int_tmp) + ret_str
	}

	if num < 0 {
		return "-" + ret_str
	}
	if num == 0 {
		return "0"
	}
	return ret_str
}
