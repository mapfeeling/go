package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func workerWithDeadline(ctx context.Context) {
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Printf("worker: deadline set to %s\n", deadline.Format(time.RFC3339))
	}

	select {
	case <-time.After(time.Second * 2):
		fmt.Println("worker completed")
	case <-ctx.Done():
		fmt.Println("worker cancelled")
	}
}

// 使用 WithDeadline 设置任务的截止时间
func TestWithDeadline(t *testing.T) {
	d := time.Now().Add(time.Second * 3)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	go workerWithDeadline(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("main cancelled")
	case <-time.After(time.Second * 4):
		fmt.Println("main completed")
	}
}

/*
我们使用 WithDeadline 派生了一个新的 context,并将其传递给了一个 Goroutine
在 worker 函数中,我们使用 ctx.Deadline 函数获取任务的截止时间，并将其格式化后打印出来
在 select 语句中,我们使用 time.After 函数模拟了 2 秒钟的工作,另一个是 ctx.Done() 信号
如果 ctx 被取消,我们就结束 Goroutine 的执行
在 main 函数中,我们使用 select 语句监听两个 channel,一个是 ctx.Done() 信号,一个是通过 time.After 函数模拟 4 秒钟的执行时间
这样,如果 worker Goroutine 能在 3 秒钟之内完成工作,程序就会输出 "main completed",否则程序就会输出 "main cancelled"

*/
