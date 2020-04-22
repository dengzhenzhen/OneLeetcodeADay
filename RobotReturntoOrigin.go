package main

/*
There is a robot starting at position (0, 0), the origin, on a 2D plane. Given a sequence of its moves, judge if this robot ends up at (0, 0) after it completes its moves.

The move sequence is represented by a string, and the character moves[i] represents its ith move. Valid moves are R (right), L (left), U (up), and D (down). If the robot returns to the origin after it finishes all of its moves, return true. Otherwise, return false.

Note: The way that the robot is "facing" is irrelevant. "R" will always make the robot move to the right once, "L" will always make it move left, etc. Also, assume that the magnitude of the robot's movement is the same for each move.

Example 1:

Input: "UD"
Output: true
Explanation: The robot moves up once, and then down once. All moves have the same magnitude, so it ended up at the origin where it started. Therefore, we return true.


Example 2:

Input: "LL"
Output: false
Explanation: The robot moves left twice. It ends up two "moves" to the left of the origin. We return false because it is not at the origin at the end of its moves.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/robot-return-to-origin
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
func judgeCircle(moves string) bool {
	movCount := map[byte]int{68:0, 76:0, 82:0, 85:0}
	for _,c := range []byte(moves){
		movCount[c]++
	}
	return movCount[68] == movCount[85] && movCount[82] == movCount[76]
}
*/
func judgeCircle(moves string) bool {
	x, y := 0, 0
	for _, c := range moves {
		if c == 'U' {
			y++
		}
		if c == 'D' {
			y--
		}
		if c == 'R' {
			x++
		}
		if c == 'L' {
			x--
		}
	}
	return x == 0 && y == 0
}
