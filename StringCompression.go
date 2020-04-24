package main

/*
Given an array of characters, compress it in-place.

The length after compression must always be smaller than or equal to the original array.

Every element of the array should be a character (not int) of length 1.

After you are done modifying the input array in-place, return the new length of the array.


Follow up:
Could you solve it using only O(1) extra space?


Example 1:

Input:
["a","a","b","b","c","c","c"]

Output:
Return 6, and the first 6 characters of the input array should be: ["a","2","b","2","c","3"]

Explanation:
"aa" is replaced by "a2". "bb" is replaced by "b2". "ccc" is replaced by "c3".


Example 2:

Input:
["a"]

Output:
Return 1, and the first 1 characters of the input array should be: ["a"]

Explanation:
Nothing is replaced.


Example 3:

Input:
["a","b","b","b","b","b","b","b","b","b","b","b","b"]

Output:
Return 4, and the first 4 characters of the input array should be: ["a","b","1","2"].

Explanation:
Since the character "a" does not repeat, it is not compressed. "bbbbbbbbbbbb" is replaced by "b12".
Notice each digit has it's own entry in the array.


Note:

All characters have an ASCII value in [35, 126].
1 <= len(chars) <= 1000.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/string-compression
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func compress(chars []byte) int {
	sameNum := 0
	length := 0
	for i, _ := range chars {
		if i == 0 {
			sameNum = 1
			length = 1
		} else {
			if chars[i] == chars[i-1] {
				sameNum++
				if sameNum == 2 {
					length++
				}
				if sameNum == 10 {
					length += 1
				}
				if sameNum == 100 {
					length += 1
				}
				if sameNum == 1000 {
					length += 1
				}
			} else {
				sameNum = 1
				length++
			}
		}
	}
	return length
}

// 还有问题，没有修改原来数组，只返回了值
