package _024_04

import (
	"context"
	"fmt"
	"testing"
	"time"
)

/*
给你一个整数 n ，请你找出所有可能含 n 个节点的 真二叉树 ，并以列表形式返回。答案中每棵树的每个节点都必须符合 Node.val == 0
答案的每个元素都是一棵真二叉树的根节点。你可以按 任意顺序 返回最终的真二叉树列表
真二叉树 是一类二叉树，树中每个节点恰好有 0 或 2 个子节点
*/

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func allPossibleFBT(n int) []*TreeNode {
//
//}
//
//func allTreeNode(n int) []*TreeNode {
//
//}

func checkCondition(ctx context.Context, condition func() bool, result chan<- bool, done func()) {
	select {
	case <-ctx.Done():
		// 如果context已经被取消，则不执行后续操作
		return
	default:
		// 执行条件判断
		if condition() {
			result <- true
			done() // 如果条件满足，调用done通知其他goroutine停止执行
		}
	}
}

func TestAllPossibleFBT(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 确保所有路径上都调用了cancel

	result := make(chan bool, 1) // 使用缓冲为1的channel避免潜在的goroutine泄漏

	// 假设有四个条件判断函数
	conditions := []func() bool{
		func() bool { time.Sleep(2); fmt.Println("a"); return false },
		func() bool { time.Sleep(2); fmt.Println("b"); return false },
		func() bool { time.Sleep(2); fmt.Println("c"); return true }, // 这个条件将会满足
		func() bool { time.Sleep(2); fmt.Println("d"); return false },
	}

	for _, cond := range conditions {
		go checkCondition(ctx, cond, result, cancel)
	}

	select {
	case <-result:
		fmt.Println("条件满足，其余goroutine将停止执行")
	case <-ctx.Done():
		fmt.Println("Context被取消")
	}

	// 等待足够的时间来确保其他goroutine已经接收到取消信号并退出
	// 实际使用中应该避免使用time.Sleep来同步goroutines
	time.Sleep(5 * time.Second)
}
