package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

type key int

const nameKey key = 0

// 使用 WithValue 存储请求特定的值
func workerWithValue(ctx context.Context) {
	if name, ok := ctx.Value(nameKey).(string); ok {
		fmt.Printf("worker: hello, %s!\n", name)
	} else {
		fmt.Println("worker: no name found")
	}
}

func TestWithValue(t *testing.T) {
	ctx := context.WithValue(context.Background(), nameKey, "Alice")

	go workerWithValue(ctx)

	// 等待一段时间，以便让 worker 完成工作
	time.Sleep(time.Second)
}

/*
我们使用 WithValue 函数在 context 中存储了一个值
在 worker 函数中，我们通过 ctx.Value 函数来获取这个值,并将其作为字符串类型打印出来
在 main 函数中,我们使用 fmt.Scanln 函数等待用户的输入,以便让程序保持运行状态,直到 worker Goroutine 完成工作
*/
