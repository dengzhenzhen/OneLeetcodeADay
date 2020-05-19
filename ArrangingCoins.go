package main

/*
You have a total of n coins that you want to form in a staircase shape, where every k-th row must have exactly k coins.

Given n, find the total number of full staircase rows that can be formed.

n is a non-negative integer and fits within the range of a 32-bit signed integer.

Example 1:

n = 5

The coins can form the following rows:
¤
¤ ¤
¤ ¤

Because the 3rd row is incomplete, we return 2.
Example 2:

n = 8

The coins can form the following rows:
¤
¤ ¤
¤ ¤ ¤
¤ ¤

Because the 4th row is incomplete, we return 3.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/arranging-coins
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 迭代解法 o(n)

func arrangeCoins(n int) int {
	var i int
	for i = 1; i*(i+1) <= 2*n; i++ {

	}
	return i - 1
}

//二分法明天补上！
/*
func arrangeCoins(n int) int {
	var ret_val int
	for {
		continue
	}
}
*/
