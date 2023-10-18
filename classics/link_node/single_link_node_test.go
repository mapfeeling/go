package link_node

import (
	"fmt"
	"testing"
)

type LinkNode struct {
	Val  int
	Next *LinkNode
}

// 链表赋值初始化
func NewLinkNode() *LinkNode {
	return &LinkNode{
		Val: 1,
		Next: &LinkNode{
			Val: 2,
			Next: &LinkNode{
				Val: 3,
				Next: &LinkNode{
					Val: 4,
					Next: &LinkNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
}

// 链表的遍历
func fmtRangeLinkNode(head *LinkNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

}

// 链表翻转
func reverseLinkNode(head *LinkNode) *LinkNode {
	var (
		prev *LinkNode
		curr = head
	)
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// 链表区间反转
// 1 ->[ -> 2 -> 3 -> 4 -> 5 -> ] -> 6        1 -> [ -> 5 -> 4 -> 3 -> 2 -> ] -> 6
func reverseBetween(head *LinkNode, left, right int) *LinkNode {
	if head == nil || head.Next == nil {
		return head
	}
	//dummy := &LinkNode{Next: head}
	//prev := dummy
	prev := new(LinkNode)
	cur := head
	fmt.Println(cur.Val)
	// 0 1 2 3 4 5 6
	for i := 0; i < left; i++ {
		prev = cur
		cur = cur.Next
		fmt.Println("cur.Val", cur.Val)
	}

	for i := 0; i < right-left-1; i++ {
		next := cur.Next
		fmt.Println()
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
	return cur
}

// 链表插入
func insertLinkNode(node *LinkNode, val int) *LinkNode {
	head := node
	for head != nil {
		if head.Val == val {
			var prev *LinkNode
			prev.Val = val + 1
			box := head.Next
			head.Next = prev
			prev.Next = box
			break
		}
		head = head.Next
	}
	return head
}

func TestLinkNode(t *testing.T) {
	{
		// 创建带头节点的链表
		fmt.Println("----------创建带头节点的链表----------")
		var head = new(LinkNode)
		head.Next = NewLinkNode()
		fmtRangeLinkNode(head)
		fmt.Println("----------带头节点的链表翻转----------")
		fmtRangeLinkNode(reverseLinkNode(head))
		fmt.Println("----------带头节点的链表区间翻转----------")
		fmtRangeLinkNode(reverseBetween(head, 2, 5))
	}

	{
		// 创建不带头节点的链表
		fmt.Println("----------创建不带头节点的链表----------")
		var linkNode = NewLinkNode()
		fmtRangeLinkNode(linkNode)
		fmt.Println("----------不带头节点的链表翻转----------")
		fmtRangeLinkNode(reverseLinkNode(linkNode))
		fmt.Println("----------不带头节点的链表区间翻转----------")
	}

	fmt.Println("----------链表插入----------")
	//fmtRangeLinkNode(insertLinkNode(linkNode, 3))
}
