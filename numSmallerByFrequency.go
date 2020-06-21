package main

import "sort"

/*
我们来定义一个函数 f(s)，其中传入参数 s 是一个非空字符串；该函数的功能是统计 s  中（按字典序比较）最小字母的出现频次。

例如，若 s = "dcce"，那么 f(s) = 2，因为最小的字母是 "c"，它出现了 2 次。

现在，给你两个字符串数组待查表 queries 和词汇表 words，请你返回一个整数数组 answer 作为答案，其中每个 answer[i] 是满足 f(queries[i]) < f(W) 的词的数目，W 是词汇表 words 中的词。



示例 1：

输入：queries = ["cbd"], words = ["zaaaz"]
输出：[1]
解释：查询 f("cbd") = 1，而 f("zaaaz") = 3 所以 f("cbd") < f("zaaaz")。
示例 2：

输入：queries = ["bbb","cc"], words = ["a","aa","aaa","aaaa"]
输出：[1,2]
解释：第一个查询 f("bbb") < f("aaaa")，第二个查询 f("aaa") 和 f("aaaa") 都 > f("cc")。


提示：

1 <= queries.length <= 2000
1 <= words.length <= 2000
1 <= queries[i].length, words[i].length <= 10
queries[i][j], words[i][j] 都是小写英文字母

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/compare-strings-by-frequency-of-the-smallest-character
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func numSmallerByFrequency(queries []string, words []string) []int {
	ret_lst := make([]int, len(queries))
	words_frequency := make([]int, len(words))
	for i, w := range words {
		words_frequency[i] = getFrequency(w)
	}
	sort.Ints(words_frequency)
	length := len(words_frequency)
	for index, _ := range queries {
		tmp := getFrequency(queries[index])

		for i, _ := range words_frequency {
			//用二分法会更快
			if tmp >= words_frequency[length-i-1] {
				break
			}
			ret_lst[index] = i + 1
		}
	}
	return ret_lst
}

func getFrequency(s string) int {
	MinChar := 'z' + 1
	frequency := 0
	for _, c := range s {
		if c == MinChar {
			frequency++
		}

		if c < MinChar {
			MinChar = c
			frequency = 1
		}
	}
	return frequency
}
