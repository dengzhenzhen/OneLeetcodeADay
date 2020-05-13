package main

/*
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var cursor *ListNode
	cursor = new(ListNode)
	*cursor = ListNode{0, nil}
	ret_node := ListNode{0, cursor}
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cursor.Next = l1
			l1 = l1.Next
		} else {
			cursor.Next = l2
			l2 = l2.Next
		}
		cursor = cursor.Next
	}
	if l1 != nil {
		cursor.Next = l1
	}
	if l2 != nil {
		cursor.Next = l2
	}
	return ret_node.Next.Next
}

type NodeHead struct {
	head *ListNode
}

func NodeInit() NodeHead {
	return NodeHead{head: nil}
}

func (this *NodeHead) append(n int) {
	if this.head == nil {
		//var tmp *ListNode
		tmp := new(ListNode)
		*tmp = ListNode{
			Val:  n,
			Next: nil,
		}
		this.head = tmp
	} else {
		var cur *ListNode
		cur = new(ListNode)
		cur = this.head
		for cur.Next != nil {
			cur = cur.Next
		}
		var tmp *ListNode
		tmp = new(ListNode)
		*tmp = ListNode{
			Val:  n,
			Next: nil,
		}
		cur.Next = tmp
	}
}

func (this *NodeHead) add(n int) {
	var tmp *ListNode
	tmp = new(ListNode)
	*tmp = ListNode{Val: n, Next: nil}
	tmp.Next = this.head
	this.head = tmp
}

func (this *NodeHead) print() {
	var tmp *ListNode
	tmp = new(ListNode)
	tmp = this.head
	for tmp != nil {
		print(tmp.Val)
		tmp = tmp.Next
		if tmp == nil {
			break
		}
		print("->")
	}
	println()
}

/*
逐个比较，较小的放到要返回的那个链表
*/
