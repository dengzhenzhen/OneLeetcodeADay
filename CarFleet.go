package main

import "sort"

/*
N cars are going to the same destination along a one lane road.  The destination is target miles away.

Each car i has a constant speed speed[i] (in miles per hour), and initial position position[i] miles towards the target along the road.

A car can never pass another car ahead of it, but it can catch up to it, and drive bumper to bumper at the same speed.

The distance between these two cars is ignored - they are assumed to have the same position.

A car fleet is some non-empty set of cars driving at the same position and same speed.  Note that a single car is also a car fleet.

If a car catches up to a car fleet right at the destination point, it will still be considered as one car fleet.


How many car fleets will arrive at the destination?



Example 1:

Input: target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]
Output: 3
Explanation:
The cars starting at 10 and 8 become a fleet, meeting each other at 12.
The car starting at 0 doesn't catch up to any other car, so it is a fleet by itself.
The cars starting at 5 and 3 become a fleet, meeting each other at 6.
Note that no other cars meet these fleets before the destination, so the answer is 3.

Note:

0 <= N <= 10 ^ 4
0 < target <= 10 ^ 6
0 < speed[i] <= 10 ^ 6
0 <= position[i] < target
All initial positions are different.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/car-fleet
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func carFleet(target int, position []int, speed []int) int {
	ret := 0
	dct := make(map[int]int)
	for index := range position {
		dct[position[index]] = speed[index]
	}
	sort.Ints(position) //懒得写排序了nlog(n)
	var first int
	for i := len(position) - 1; i >= 0; i-- {
		if i == len(position)-1 {
			ret++
			first = i
			continue
		}
		//println(canCatch([2]int{position[i], position[first]}, [2]int{ dct[position[i]], dct[position[first]] }, target ))
		if canCatch([2]int{position[i], position[first]}, [2]int{dct[position[i]], dct[position[first]]}, target) {
			continue
		} else {
			ret++
			first = i
		}

	}

	return ret
}

func canCatch(position [2]int, speed [2]int, target int) bool {
	return (position[1]-position[0])*speed[1] <= (target-position[1])*(speed[0]-speed[1])
}

/*
按照位置排序
一个车队取决于最前面又最慢的那个
从离终点最近的一辆车开始遍历，如果前一辆能追上，那他们在同一个车队，如果不能追上，那就是一个新的车队
*/
