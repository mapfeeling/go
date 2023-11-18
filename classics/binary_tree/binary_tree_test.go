package binary_tree

import (
	"fmt"
	"math"
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
	return max(MaxDepthBinaryTree(root.LeftNode), MaxDepthBinaryTree(root.RightNode)) + 1
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

func findCompleteBtRightNode(root *BinaryTree) *BinaryTree {
	if root == nil {
		return nil
	}
	rightNode := findCompleteBtRightNode(root.RightNode)
	if rightNode != nil {
		return rightNode
	}
	return root

}

// LCA
func lowestCommonAncestor(root, p, q *BinaryTree) *BinaryTree {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.LeftNode, p, q)
	right := lowestCommonAncestor(root.RightNode, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	return right
}

func findAllContainPaths(root *BinaryTree, target int, paths []int) [][]int {
	if root == nil {
		return [][]int{}
	}
	paths = append(paths, root.Value)
	if root.Value == target {
		return [][]int{paths}
	}
	leftPaths := findAllContainPaths(root.LeftNode, target, paths)
	rightPaths := findAllContainPaths(root.RightNode, target, paths)
	return append(leftPaths, rightPaths...)
}

func findAllPathsList(root *BinaryTree) [][]int {
	var res [][]int

	if root == nil {
		return res
	}

	if root.LeftNode == nil || root.RightNode == nil {
		return [][]int{{root.Value}}
	}

	leftPaths := findAllPathsList(root.LeftNode)
	rightPaths := findAllPathsList(root.RightNode)

	paths := make([][]int, 0)
	for _, path := range leftPaths {
		paths = append(paths, append([]int{root.Value}, path...))
	}
	for _, path := range rightPaths {
		paths = append(paths, append([]int{root.Value}, path...))
	}

	return paths
}

func dfsPathSum(root *BinaryTree, sum int, arr []int, ret *[][]int) {
	if root == nil {
		return
	}
	arr = append(arr, root.Value)
	if root.Value == sum && root.LeftNode == nil && root.RightNode == nil {
		box := make([]int, len(arr))
		copy(box, arr)
		*ret = append(*ret, box)
	}
	dfsPathSum(root.LeftNode, sum-root.Value, arr, ret)
	dfsPathSum(root.RightNode, sum-root.Value, arr, ret)
	arr = arr[:len(arr)-1]
}

// 二叉树中和为某一值的路径
func pathSum(root *BinaryTree, sum int) [][]int {
	if root == nil {
		return nil
	}
	var ret [][]int
	dfsPathSum(root, sum, []int{}, &ret)
	return ret
}

func dfsMaxPathSum(root *BinaryTree, maxSum int) int {
	if root == nil {
		return 0
	}
	left := max(0, dfsMaxPathSum(root.LeftNode, maxSum))
	right := max(0, dfsMaxPathSum(root.RightNode, maxSum))
	maxSum = max(maxSum, root.Value+left+right)
	return maxSum
}

// 二叉树最大路径之和
func maxPathSum(root *BinaryTree) (maxSum int) {
	maxSum = math.MinInt32
	dfsMaxPathSum(root, maxSum)
	return
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

	fmt.Println("-------------先序遍历---------------------")
	RecursionPreOrderTraversal(root)
	fmt.Println("-------------中序遍历---------------------")
	RecursionMiddleOrderTraversal(root)
	fmt.Println("-------------后序遍历---------------------")
	RecursionPostOrderTraversal(root)
	fmt.Println("-------------层次遍历---------------------")
	LevelTraversal(root)
	fmt.Println("------二叉树翻转后-------层次遍历------------")
	InvertBinaryTree(root)
	LevelTraversal(root)
	fmt.Println("------判断一个二叉树是否为完全二叉树-------")
	IsCompleteBt(root)
	fmt.Println("------一个完全二叉树的最右子节点-------")
	fmt.Println(findCompleteBtRightNode(root).Value)
	fmt.Println("------是否为完全二叉树:-------", IsNotCompleteBt)
	fmt.Println("-------该二叉树的最大深度:--------", MaxDepthBinaryTree(root))
	// LCA
	{
		root := &BinaryTree{Value: 3}
		root.LeftNode = &BinaryTree{Value: 5}
		root.RightNode = &BinaryTree{Value: 1}
		root.LeftNode.LeftNode = &BinaryTree{Value: 6}
		root.LeftNode.RightNode = &BinaryTree{Value: 2}
		root.RightNode.LeftNode = &BinaryTree{Value: 0}
		root.RightNode.RightNode = &BinaryTree{Value: 8}
		p, q := root.LeftNode.LeftNode, root.LeftNode.RightNode
		fmt.Println("------LCA:-------", lowestCommonAncestor(root, p, q))
		fmt.Println("------findAllPathsList:-------", findAllPathsList(root))
		fmt.Println("------findAllContainPaths:-------", findAllContainPaths(root, 5, []int{}))
	}
	{

	}
	fmt.Println("------二叉树中和为某一值的路径:-------", pathSum(root, 20))

}
