package binary_tree

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"testing"
)

type compare interface {
	compare(data interface{}) bool
	equals(data interface{}) bool
}

// 二叉排序树 or 二叉搜索树
type BinarySearchTree struct {
	Data      compare
	LeftNode  *BinarySearchTree
	RightNode *BinarySearchTree
}

type data struct {
	key   int
	value interface{}
}

func (d *data) equals(dOther interface{}) bool {
	if d.key == dOther.(*data).key {
		return true
	}
	return false
}

func (d *data) compare(dOther interface{}) bool {
	if d.key > dOther.(*data).key {
		return true
	}
	return false
}

func (root *BinarySearchTree) insert(content compare) {
	var parent *BinarySearchTree = root
	if root.Data == nil {
		root.Data = content
		return
	}

	for root != nil {
		if root.Data.compare(content) {
			root = root.LeftNode
		} else {
			root = root.RightNode
		}
	}

	if parent.Data.compare(content) {
		parent.LeftNode = &BinarySearchTree{
			Data:      content,
			LeftNode:  nil,
			RightNode: nil,
		}
	} else {
		parent.RightNode = &BinarySearchTree{
			Data:      content,
			LeftNode:  nil,
			RightNode: nil,
		}
	}
}

func (root *BinarySearchTree) query(node compare) compare {
	for true {
		if root == nil {
			return nil
		}
		if root.Data.equals(node) {
			return root.Data
		}
		if root.Data.compare(node) {
			root = root.LeftNode
		} else {
			root = root.RightNode
		}
	}
	return nil
}

func (root *BinarySearchTree) update(node compare) error {
	for true {
		if root == nil {
			return errors.New("not found")
		}
		if root.Data.equals(node) {
			root.Data = node
			return nil
		}
		if root.Data.compare(node) {
			root = root.LeftNode
		} else {
			root = root.RightNode
		}
	}
	return nil
}

// 二叉树的前序遍历
func (root *BinarySearchTree) preOrderTraverse() (res []compare) {
	if root == nil || root.Data == nil {
		return nil
	}
	res = append(res, root.Data)
	root.LeftNode.preOrderTraverse()
	root.RightNode.preOrderTraverse()
	return
}

// 二叉树的中序遍历
func (root *BinarySearchTree) midOrderTraverse() (res []compare) {
	if root == nil || root.Data == nil {
		return nil
	}
	root.LeftNode.midOrderTraverse()
	res = append(res, root.Data)
	fmt.Println(root.Data)
	root.RightNode.midOrderTraverse()
	return
}

// 二叉树的后序遍历
func (root *BinarySearchTree) postOrderTraverse() (res []compare) {
	if root == nil || root.Data == nil {
		return nil
	}
	root.LeftNode.postOrderTraverse()
	root.RightNode.postOrderTraverse()
	res = append(res, root.Data)
	return
}

// 层次遍历
func (root *BinarySearchTree) LevelTraversal() {
	if root == nil {
		return
	}
	var nodeSlice []*BinarySearchTree
	nodeSlice = append(nodeSlice, root)
	// 开始遍历
	root.RecursionTraversal(nodeSlice)
}

func (root *BinarySearchTree) RecursionTraversal(nodeSlice []*BinarySearchTree) {
	if len(nodeSlice) == 0 {
		return
	}
	// 创建新的节点slice
	var nextSlice []*BinarySearchTree
	for _, node := range nodeSlice {
		fmt.Println("当前节点的值:", node.Data)
		if node.LeftNode != nil {
			nextSlice = append(nextSlice, node.LeftNode)
		}
		if node.RightNode != nil {
			nextSlice = append(nextSlice, node.RightNode)
		}
	}
	// 递归下一层的叶子节点
	root.RecursionTraversal(nextSlice)
}

func TestBinarySearchTree(t *testing.T) {
	bst := new(BinarySearchTree)
	bst.insert(&data{
		key:   3,
		value: "123",
	})

	bst.insert(&data{
		key:   2,
		value: "234",
	})
	bst.insert(&data{
		key:   1,
		value: "345",
	})
	bst.insert(&data{
		key:   4,
		value: "456",
	})
	bst.insert(&data{
		key:   5,
		value: "567",
	})
	bst.insert(&data{
		key:   6,
		value: "678",
	})
	fmt.Println(bst)
	bst.LevelTraversal()
	fmt.Println(bst.query(&data{key: 4}))
	fmt.Println(bst.preOrderTraverse())
	fmt.Println(bst.midOrderTraverse())
	fmt.Println(bst.postOrderTraverse())
	fmt.Println(splitNum(4325))
}

func splitNum(num int) int {
	strNum := []byte(strconv.Itoa(num))
	sort.Slice(strNum, func(i, j int) bool {
		return strNum[i] < strNum[j]
	})
	num1, num2 := 0, 0
	for i := 0; i < len(strNum); i++ {
		if i%2 == 0 {
			num1 = num1*10 + int(strNum[i]-'0')
		} else {
			num2 = num2*10 + int(strNum[i]-'0')
		}
	}
	return num1 + num2
}
