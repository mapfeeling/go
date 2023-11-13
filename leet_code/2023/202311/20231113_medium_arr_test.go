package _02311

import "testing"

type NumArray struct {
	n    int
	tree []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	tree := make([]int, n<<1)
	for i := n; i < n<<1; i++ {
		tree[i] = nums[i-n]
	}
	for i := n - 1; i > 0; i-- {
		tree[i] = tree[i<<1] + tree[i<<1|1]
	}
	return NumArray{tree: tree, n: n}
}

func (this *NumArray) Update(index int, val int) {
	index += this.n
	this.tree[index] = val
	for ; index > 0; index >>= 1 {
		this.tree[index>>1] = this.tree[index] + this.tree[index^1]
	}
}

func (this *NumArray) SumRange(left int, right int) (ret int) {
	left, right = left+this.n, right+this.n
	for ; left <= right; left, right = left>>1, right>>1 {
		if (left & 1) == 1 {
			ret, left = ret+this.tree[left], left+1
		}
		if (right & 1) == 0 {
			ret, right = ret+this.tree[right], right-1
		}
	}
	return
}

func Test20231113(t *testing.T) {

}

/*
给你一个数组 nums ,请你完成两类查询
其中一类查询要求 更新 数组 nums 下标对应的值
另一类查询要求返回数组 nums 中索引 left 和索引 right 之间（ 包含 ）的nums元素的 和 ，其中 left <= right
实现 NumArray 类：
NumArray(int[] nums) 用整数数组 nums 初始化对象
void update(int index, int val) 将 nums[index] 的值 更新 为 val
int sumRange(int left, int right) 返回数组 nums 中索引 left 和索引 right 之间（ 包含 ）的nums元素的 和 （即，nums[left] + nums[left + 1], ..., nums[right]）
*/
