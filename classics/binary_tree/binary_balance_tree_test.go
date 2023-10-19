package binary_tree

import (
	"fmt"
	"testing"
)

type BinaryBalanceTree struct {
	Data      interface{}
	LeftNode  *BinaryBalanceTree
	RightNode *BinaryBalanceTree
}

// 在计算机科学中，二叉树是一种常用的数据结构
// 它由节点组成，每个节点最多有两个子节点：左子节点和右子节点
// 当左子节点和右子节点的深度相差不超过1时，我们称之为平衡二叉树

func (root *BinaryBalanceTree) height() int {
	if root == nil {
		return 0
	}
	leftHeight := root.LeftNode.height()
	rightHeight := root.RightNode.height()
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func (root *BinaryBalanceTree) IsBalanceTree() bool {
	if root == nil {
		return true
	}
	leftHeight := root.LeftNode.height()
	rightHeight := root.RightNode.height()
	var fAbs = func(i, j int) int {
		if i > j {
			return i - j
		}
		return -i + j
	}
	if fAbs(leftHeight, rightHeight) > 1 {
		return false
	}
	return root.LeftNode.IsBalanceTree() && root.RightNode.IsBalanceTree()
}

func TestBinaryBalanceTree(t *testing.T) {
	// 创建一个平衡二叉树
	var balancedTree = &BinaryBalanceTree{
		Data: 1,
		LeftNode: &BinaryBalanceTree{
			Data: 2,
			LeftNode: &BinaryBalanceTree{
				Data: 3,
			},
			RightNode: &BinaryBalanceTree{
				Data: 4,
			},
		},
		RightNode: &BinaryBalanceTree{
			Data: 5,
			LeftNode: &BinaryBalanceTree{
				Data: 6,
			},
			RightNode: &BinaryBalanceTree{
				Data: 7,
			},
		},
	}
	// 创建一个非平衡二叉树
	var unBalancedTree = &BinaryBalanceTree{
		Data: 1,
		LeftNode: &BinaryBalanceTree{
			Data: 2,
			LeftNode: &BinaryBalanceTree{
				Data: 3,
				LeftNode: &BinaryBalanceTree{
					Data: 4,
				},
			},
		},
		RightNode: &BinaryBalanceTree{
			Data: 5,
		},
	}

	fmt.Println("平衡二叉树是否平衡：", balancedTree.IsBalanceTree())
	fmt.Println("不平衡二叉树是否平衡：", unBalancedTree.IsBalanceTree())
}
