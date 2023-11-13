package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		default:
			fmt.Printf("worker %d is running\n", id)
		case <-ctx.Done():
			fmt.Printf("worker %d is cancelled\n", id)
			return
		}
	}
}

// 示例 1: 使用 WithCancel 实现 Goroutine 的取消
func TestContextWithCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	// 启动两个 worker
	go worker(ctx, 1)
	go worker(ctx, 2)

	// 运行一段时间后取消所有 worker
	time.Sleep(time.Second * 3)
	cancel()
	time.Sleep(time.Second)
}

/*
我们通过使用 WithCancel 派生了一个新的 context,并将其传递给了两个 Goroutine
在 main 函数中，我们等待 3 秒钟后取消了所有的 Goroutine
在 worker 函数中,我们使用 select 语句来监听 ctx.Done() 信号，如果 ctx 被取消,我们就结束 Goroutine 的执行
*/
