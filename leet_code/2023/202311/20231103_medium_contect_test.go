package _02311

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	box := []*Node{root}
	for len(box) > 0 {
		tmp := box
		box = nil
		for index, node := range tmp {
			if index+1 < len(tmp) {
				node.Next = tmp[index+1]
			}
			if node.Left != nil {
				box = append(box, node.Left)
			}
			if node.Right != nil {
				box = append(box, node.Right)
			}
		}
	}

	return root
}

/*
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL 。
初始状态下，所有 next 指针都被设置为 NULL
*/
