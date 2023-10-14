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
	var prev *LinkNode
	for head != nil {
		box := head.Next
		head.Next = prev
		prev = head
		head = box
	}
	return prev
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
	fmt.Println("----------链表初始化----------")
	var linkNode = NewLinkNode()
	fmtRangeLinkNode(linkNode)
	fmt.Println("----------链表翻转----------")
	fmtRangeLinkNode(reverseLinkNode(linkNode))
	fmt.Println("----------链表插入----------")
	fmtRangeLinkNode(insertLinkNode(linkNode, 3))
}
