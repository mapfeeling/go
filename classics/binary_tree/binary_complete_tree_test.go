package binary_tree

import (
	"fmt"
	"testing"
)

type BinaryCompleteTree struct {
	Data      int
	LeftNode  *BinaryCompleteTree
	RightNode *BinaryCompleteTree
}

func (root *BinaryCompleteTree) Height() int {

	if root == nil {
		return 0
	}
	leftBinaryCompleteTreeHeight := root.LeftNode.Height()
	rightBinaryCompleteTreeHeight := root.RightNode.Height()
	if leftBinaryCompleteTreeHeight == 0 {
		return leftBinaryCompleteTreeHeight + 1
	}
	return rightBinaryCompleteTreeHeight + 1
}

func (root *BinaryCompleteTree) MaxHeight() int {
	if root == nil {
		return 0
	}

	return maxInt(root.LeftNode.MaxHeight(), root.RightNode.MaxHeight()) + 1
}

func maxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func (root *BinaryCompleteTree) IsBinaryCompleteTree() bool {
	if root == nil {
		return true
	}
	leftBinaryCompleteTreeHeight := root.LeftNode.Height()
	rightBinaryCompleteTreeHeight := root.RightNode.Height()
	if leftBinaryCompleteTreeHeight != rightBinaryCompleteTreeHeight {
		return false
	}
	return root.LeftNode.IsBinaryCompleteTree() && root.RightNode.IsBinaryCompleteTree()
}

func TestBinaryCompleteTree(t *testing.T) {
	var binaryCompleteTree = &BinaryCompleteTree{
		Data: 0,
		LeftNode: &BinaryCompleteTree{
			Data: 1,
			LeftNode: &BinaryCompleteTree{
				Data:      3,
				LeftNode:  nil,
				RightNode: nil,
			},
			RightNode: &BinaryCompleteTree{
				Data:      4,
				LeftNode:  nil,
				RightNode: nil,
			},
		},
		RightNode: &BinaryCompleteTree{
			Data: 2,
			LeftNode: &BinaryCompleteTree{
				Data:      5,
				LeftNode:  nil,
				RightNode: nil,
			},
			RightNode: &BinaryCompleteTree{
				Data:      6,
				LeftNode:  nil,
				RightNode: nil,
			},
		},
	}
	var unBinaryCompleteTree = &BinaryCompleteTree{
		Data: 0,
		LeftNode: &BinaryCompleteTree{
			Data: 1,
			LeftNode: &BinaryCompleteTree{
				Data: 3,
				LeftNode: &BinaryCompleteTree{
					Data:      7,
					LeftNode:  nil,
					RightNode: nil,
				},
				RightNode: nil,
			},
			RightNode: &BinaryCompleteTree{
				Data:      4,
				LeftNode:  nil,
				RightNode: nil,
			},
		},
		RightNode: &BinaryCompleteTree{
			Data: 2,
			LeftNode: &BinaryCompleteTree{
				Data:      5,
				LeftNode:  nil,
				RightNode: nil,
			},
			RightNode: &BinaryCompleteTree{
				Data:     6,
				LeftNode: nil,
				RightNode: &BinaryCompleteTree{
					Data:     8,
					LeftNode: nil,
					RightNode: &BinaryCompleteTree{
						Data:     9,
						LeftNode: nil,
						RightNode: &BinaryCompleteTree{
							Data:      10,
							LeftNode:  nil,
							RightNode: nil,
						},
					},
				},
			},
		},
	}
	fmt.Println(binaryCompleteTree.IsBinaryCompleteTree())
	fmt.Println(binaryCompleteTree.MaxHeight() + 1)
	fmt.Println(unBinaryCompleteTree.IsBinaryCompleteTree())
	fmt.Println(unBinaryCompleteTree.MaxHeight() + 1)
}
