package main

/*
给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

注意：

num1 和num2 的长度都小于 5100.
num1 和num2 都只包含数字 0-9.
num1 和num2 都不包含任何前导零。
你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func addStrings(num1 string, num2 string) string {
	length := len(num1) - len(num2)
	if length < 0 {
		for i := 0; i < -length; i++ {
			num1 = "0" + num1
		}
	}
	if length > 0 {
		for i := 0; i < length; i++ {
			num2 = "0" + num2
		}
	}

	carry := byte(0)
	var tmp_sum string
	ret_str := ""
	for index := len(num1) - 1; index >= 0; index-- {
		tmp_sum, carry = AddByDigit(num1[index], num2[index], carry)
		ret_str = tmp_sum + ret_str
	}
	if carry != 0 {
		ret_str = string(carry|0x30) + ret_str
	}
	return ret_str
}

func AddByDigit(a byte, b byte, c byte) (string, byte) {
	int_ret := (a & 0xCF) + (b & 0xCF) + c
	sum := int_ret % 10
	carry := int_ret / 10
	return string(sum | 0x0000030), carry
}
