package main

/*
给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。



示例 1:

输入: 12258
输出: 5
解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"


提示：

0 <= num < 231

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func translateNum(num int) int {
	if num <= 9 {
		return 1
	}

	tmp := num % 100
	if tmp <= 9 || tmp >= 26 {
		return translateNum(num / 10)
	} else {
		return translateNum(num/10) + translateNum(num/100)
	}

}

/*
青蛙跳台阶升级版，稍微麻烦一点
递归终止条件 num <= 9
然后是两位数以上的情况：
10~25有两种组合方式，其余1种
*/

// 没用的
func num2string(n int) string {
	return string([]byte{byte(n) + 'a'})
}
