package binary_tree

import (
	"fmt"
	"testing"
)

var IsNotCompleteBt bool = true

type BinaryTree struct {
	Value     int
	LeftNode  *BinaryTree
	RightNode *BinaryTree
}

func (n *BinaryTree) Insert(val int) {
	if n == nil {
		return
	}
}

// 先序遍历
func RecursionPreOrderTraversal(node *BinaryTree) {
	//如果当前节点为nil，则结束遍历
	if node == nil {
		return
	}
	// 先输出自身节点的值
	fmt.Println("前序遍历", node.Value)
	// 递归遍历左子树
	RecursionPreOrderTraversal(node.LeftNode)
	// 递归遍历右子树
	RecursionPreOrderTraversal(node.RightNode)
}

// 中序遍历
func RecursionMiddleOrderTraversal(node *BinaryTree) {
	//如果当前节点为nil，则结束遍历
	if node == nil {
		return
	}

	// 递归遍历左子树
	RecursionMiddleOrderTraversal(node.LeftNode)
	// 输出自身节点的值
	fmt.Println("中序遍历", node.Value)
	// 递归遍历右子树
	RecursionMiddleOrderTraversal(node.RightNode)
}

// 后序遍历
func RecursionPostOrderTraversal(node *BinaryTree) {
	//如果当前节点为nil，则结束遍历
	if node == nil {
		return
	}

	// 递归遍历左子树
	RecursionPostOrderTraversal(node.LeftNode)
	// 递归遍历右子树
	RecursionPostOrderTraversal(node.RightNode)
	// 输出自身节点的值
	fmt.Println("后序遍历", node.Value)
}

// 层次遍历
func LevelTraversal(rootNode *BinaryTree) {
	//如果当前节点为nil，则结束遍历
	if rootNode == nil {
		return
	}
	var nodeSlice []*BinaryTree
	nodeSlice = append(nodeSlice, rootNode)

	// 开始递归遍历
	RecursionTraversal(nodeSlice)
}

func RecursionTraversal(nodeSlice []*BinaryTree) {
	if len(nodeSlice) == 0 {
		return
	}
	// 创建新的节点slice,春初下一层需要遍历的node
	var nextSlice []*BinaryTree

	for i := 0; i < len(nodeSlice); i++ {
		node := nodeSlice[i]
		fmt.Println("当前node的值:", node.Value)
		if node.LeftNode != nil {
			nextSlice = append(nextSlice, node.LeftNode)
		}
		if node.RightNode != nil {
			nextSlice = append(nextSlice, node.RightNode)
		}
	}
	//递归遍历下一层的node slice
	RecursionTraversal(nextSlice)
}

// 翻转二叉树
func InvertBinaryTree(root *BinaryTree) *BinaryTree {
	if root == nil {
		return nil
	}
	InvertBinaryTree(root.LeftNode)
	InvertBinaryTree(root.RightNode)
	root.LeftNode, root.RightNode = root.RightNode, root.LeftNode
	return root
}

// 二叉树的最大深度
func MaxDepthBinaryTree(root *BinaryTree) int {
	if root == nil {
		return 0
	}
	leftMaxDepthBinaryTree := MaxDepthBinaryTree(root.LeftNode)
	rightMaxDepthBinaryTree := MaxDepthBinaryTree(root.RightNode)
	if leftMaxDepthBinaryTree > rightMaxDepthBinaryTree {
		return leftMaxDepthBinaryTree + 1
	}
	return rightMaxDepthBinaryTree + 1
}

// 判断一个树是否为完全二叉树
func IsCompleteBt(root *BinaryTree) bool {
	if root == nil {
		return true
	}
	if root.LeftNode != nil && root.RightNode == nil {
		IsNotCompleteBt = false
		return false
	}
	// 1.当一个节点存在右子节点但是不存在左子节点这颗树视为非完全二叉树
	// 2.当一个节点的左子节点存在但是右子节点不存在视为完全二叉树
	IsCompleteBt(root.LeftNode)
	IsCompleteBt(root.RightNode)
	return true
}

func TestBinaryTree(t *testing.T) {
	// 创建二叉树根节点
	//根节点
	var root *BinaryTree = &BinaryTree{
		Value: 1,
		LeftNode: &BinaryTree{
			Value: 2,
			LeftNode: &BinaryTree{
				Value: 4,
				RightNode: &BinaryTree{
					Value:     6,
					LeftNode:  &BinaryTree{Value: 7},
					RightNode: &BinaryTree{Value: 8},
				},
			},
		},
		RightNode: &BinaryTree{
			Value:     3,
			RightNode: &BinaryTree{Value: 5},
		},
	}

	RecursionPreOrderTraversal(root)
	fmt.Println("----------------------------------")
	RecursionMiddleOrderTraversal(root)
	fmt.Println("----------------------------------")
	RecursionPostOrderTraversal(root)
	fmt.Println("-------------层次遍历---------------------")
	LevelTraversal(root)
	fmt.Println("------二叉树翻转后-------层次遍历------------")
	InvertBinaryTree(root)
	LevelTraversal(root)
	fmt.Println("------判断一个二叉树是否为完全二叉树-------")
	IsCompleteBt(root)
	fmt.Println("------是否为完全二叉树:-------", IsNotCompleteBt)
	fmt.Println("-------该二叉树的最大深度:--------", MaxDepthBinaryTree(root))

}
