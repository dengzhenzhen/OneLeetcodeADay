package main

/*
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func lengthOfLongestSubstring(s string) int {
	freq := make([]int, 128)
	var res = 0
	start, end := 0, -1
	for start < len(s) {
		if end+1 < len(s) && freq[s[end+1]] == 0 {
			end++
			freq[s[end]]++
		} else {
			freq[s[start]]--
			start++
		}
		res = max(res, end-start+1)
	}
	return res
}
func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
