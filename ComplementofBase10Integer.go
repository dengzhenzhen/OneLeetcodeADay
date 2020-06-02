package main

/*
Every non-negative integer N has a binary representation.  For example, 5 can be represented as "101" in binary, 11 as "1011" in binary, and so on.  Note that except for N = 0, there are no leading zeroes in any binary representation.

The complement of a binary representation is the number in binary you get when changing every 1 to a 0 and 0 to a 1.  For example, the complement of "101" in binary is "010" in binary.

For a given number N in base-10, return the complement of it's binary representation as a base-10 integer.



Example 1:

Input: 5
Output: 2
Explanation: 5 is "101" in binary, with complement "010" in binary, which is 2 in base-10.
Example 2:

Input: 7
Output: 0
Explanation: 7 is "111" in binary, with complement "000" in binary, which is 0 in base-10.
Example 3:

Input: 10
Output: 5
Explanation: 10 is "1010" in binary, with complement "0101" in binary, which is 5 in base-10.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/complement-of-base-10-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func bitwiseComplement(N int) int {
	ret_num := 0
	tmp := 1
	if N == 0 {
		return 1
	}
	for N != 0 {
		if N&1 == 0 {
			ret_num += tmp
		}
		N >>= 1
		tmp <<= 1
	}
	return ret_num
}
