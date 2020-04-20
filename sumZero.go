package main

/*
Given an integer n, return any array containing n unique integers such that they add up to 0.



Example 1:

Input: n = 5
Output: [-7,-1,1,3,4]
Explanation: These arrays also are accepted [-5,-1,1,2,3] , [-3,-1,2,-2,4].
Example 2:

Input: n = 3
Output: [-1,0,1]
Example 3:

Input: n = 1
Output: [0]


Constraints:

1 <= n <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-n-unique-integers-sum-up-to-zero
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func sumZero(n int) []int {
	ret_lst := make([]int, n)
	if n&1 == 0 {
		for i := range ret_lst {
			if i < n/2 {
				ret_lst[i] = i - (n / 2)
			} else {
				ret_lst[i] = i - (n / 2) + 1
			}
		}
	} else {
		for i := range ret_lst {
			ret_lst[i] = i - (n / 2)
		}
	}
	return ret_lst
}
