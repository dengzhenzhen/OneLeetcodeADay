package main

/*
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
getMin() -- Retrieve the minimum element in the stack.


Example 1:

Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]

Explanation
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2


Constraints:

Methods pop, top and getMin operations will always be called on non-empty stacks.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/min-stack
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type MinStack struct {
	val     int
	min_val int
	next    *MinStack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	var ret *MinStack
	ret = new(MinStack)
	*ret = MinStack{
		val:     0,
		min_val: 0,
		next:    nil,
	}
	return *ret
}

func (this *MinStack) Push(x int) {
	var tmp *MinStack
	tmp = new(MinStack)
	*tmp = MinStack{
		val:     x,
		min_val: x,
		next:    nil,
	}
	if this.next == nil {
		this.next = tmp
	} else {
		if tmp.val > this.next.min_val {
			tmp.min_val = this.next.min_val
		}
		tmp.next = this.next
		this.next = tmp
	}
}

func (this *MinStack) Pop() {
	this.next = this.next.next
}

func (this *MinStack) Top() int {
	return this.next.val
}

func (this *MinStack) GetMin() int {
	return this.next.min_val
}

/*
用链表实现一个栈
压栈时记录一下当前最小值
*/

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
