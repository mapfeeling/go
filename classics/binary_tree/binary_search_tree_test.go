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

type BinarySearchTree struct {
	data      compare
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

func (b *BinarySearchTree) insert(content compare) {
	var parent *BinarySearchTree = b
	if b.data == nil {
		b.data = content
		return
	}

	for {
		if b == nil {
			break
		}
		if b.data.compare(content) {
			b = b.LeftNode
		} else {
			b = b.RightNode
		}
	}

	if parent.data.compare(content) {
		parent.LeftNode = &BinarySearchTree{
			data:      content,
			LeftNode:  nil,
			RightNode: nil,
		}
	} else {
		parent.RightNode = &BinarySearchTree{
			data:      content,
			LeftNode:  nil,
			RightNode: nil,
		}
	}
}

func (b *BinarySearchTree) query(node compare) compare {
	for true {
		if b == nil {
			return nil
		}
		if b.data.equals(node) {
			return b.data
		}
		if b.data.compare(node) {
			b = b.LeftNode
		} else {
			b = b.RightNode
		}
	}
	return nil
}

func (b *BinarySearchTree) update(node compare) error {
	for true {
		if b == nil {
			return errors.New("not found")
		}
		if b.data.equals(node) {
			b.data = node
			return nil
		}
		if b.data.compare(node) {
			b = b.LeftNode
		} else {
			b = b.RightNode
		}
	}
	return nil
}

// 二叉树的前序遍历
func (b *BinarySearchTree) preOrderTraverse() (res []compare) {
	if b == nil || b.data == nil {
		return nil
	}
	res = append(res, b.data)
	b.LeftNode.preOrderTraverse()
	b.RightNode.preOrderTraverse()
	return
}

// 二叉树的中序遍历
func (b *BinarySearchTree) midOrderTraverse() (res []compare) {
	if b == nil || b.data == nil {
		return nil
	}
	b.LeftNode.midOrderTraverse()
	res = append(res, b.data)
	fmt.Println(b.data)
	b.RightNode.midOrderTraverse()
	return
}

// 二叉树的后序遍历
func (b *BinarySearchTree) postOrderTraverse() (res []compare) {
	if b == nil || b.data == nil {
		return nil
	}
	b.LeftNode.postOrderTraverse()
	b.RightNode.postOrderTraverse()
	res = append(res, b.data)
	return
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
	fmt.Println(bst)
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
