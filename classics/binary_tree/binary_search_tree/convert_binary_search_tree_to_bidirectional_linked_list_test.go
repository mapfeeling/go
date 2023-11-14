package binary_search_tree

type BinarySearchTree struct {
	Val       int
	LeftNode  *BinarySearchTree
	RightNode *BinarySearchTree
}

type BidirectionalLinkedList struct {
	Val  int
	Prev *BidirectionalLinkedList
	Next *BidirectionalLinkedList
}

type Stack[T any] struct {
	values []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{values: make([]T, 0)}
}

func (s *Stack[T]) Peak() T {
	return s.values[s.LastIndex()]
}

func (s *Stack[T]) Pop() T {
	ans := s.values[s.LastIndex()]
	s.values = s.values[:s.LastIndex()]
	return ans
}

func (s *Stack[T]) Push(n T) {
	s.values = append(s.values, n)
}

func (s *Stack[T]) LastIndex() int {
	return len(s.values) - 1
}

func (s *Stack[T]) Size() int {
	return len(s.values)
}

func inorderTraversal(root *BinarySearchTree) []int {
	var (
		ans   []int
		stack = NewStack[*BidirectionalLinkedList]()
		curr  = root
	)
	for curr != nil || stack.Size() != 0 {

	}

}

// 用go实现将二叉搜索树转成有序双向链表，要求空间复杂度O(1)
