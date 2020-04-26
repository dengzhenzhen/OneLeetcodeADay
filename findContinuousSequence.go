package main

/*输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。



示例 1：

输入：target = 9
输出：[[2,3,4],[4,5]]
示例 2：

输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]]


限制：

1 <= target <= 10^5


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findContinuousSequence(target int) [][]int {
	begin := 1
	end := 2
	var ret [][]int
	for begin < end {
		if serialSum(begin, end) == target {
			ret = append(ret, arrayGen(begin, end))
			//fmt.Println(ret)
		}
		if serialSum(begin, end) >= target {
			begin++
		} else {
			end++
		}
	}
	return ret
}

// 等差数列求和
func serialSum(begin int, end int) int {
	return (begin + end) * (end - begin + 1) / 2
}

// 生成 begin 到 end 的数列
func arrayGen(begin int, end int) []int {
	ret := make([]int, end-begin+1)
	for i := begin; i <= end; i++ {
		ret[i-begin] = i
	}
	return ret
}

/*
用begin和end两个在数轴上滑动
如果begin和end之间的和比target小，那end需要增加，反之begin增加
begin超过end之后循环结束
*/
