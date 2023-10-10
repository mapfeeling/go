package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type SlidingWindowMute struct {
	windowSize       int        // 滑动窗口大小
	failureThreshold int        // 失败阈值
	window           [12]int64  // 环形数组，存储12个时间窗口内的失败次数
	lock             sync.Mutex // 互斥锁，用于保护环形数组
	startTime        time.Time  // 当前窗口的起始时间
	failureCount     int64      // 累计失败次数
}

func NewSlidingWindowMute() *SlidingWindowMute {
	return &SlidingWindowMute{
		windowSize:       12,
		failureThreshold: 5, // 在1分钟内最多允许失败5次
		window:           [12]int64{},
	}
}

func (swm *SlidingWindowMute) AllowRequest() bool {
	swm.lock.Lock()
	defer swm.lock.Unlock()

	now := time.Now()
	index := now.Minute() / int(time.Duration(swm.windowSize)) % int(time.Hour/time.Minute)

	// 更新窗口的起始时间
	swm.startTime = now.Truncate(time.Duration(swm.windowSize) * time.Minute).Add(time.Duration(index) * time.Minute)
	fmt.Println(swm.startTime)
	// 将当前时间窗口内的失败次数清零
	for i := 0; i < swm.windowSize; i++ {
		if index-i >= 0 && index-i < int(time.Hour/time.Minute) {
			swm.window[index-i] = 0
		}
	}

	// 检查窗口内总失败次数是否超过阈值
	totalFailureCount := int64(0)
	for _, count := range swm.window {
		totalFailureCount += count
	}

	if totalFailureCount >= int64(swm.failureThreshold) {
		// 熔断器触发，禁止后续请求
		return false
	}

	// 在当前时间窗口内增加失败次数计数
	swm.window[index]++
	swm.failureCount++

	return true
}

func Test2demo(t *testing.T) {
	mute := NewSlidingWindowMute()
	defer time.Sleep(11 * time.Minute) // 等待11分钟，以便观察滑动窗口的变化

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			allowed := mute.AllowRequest()
			fmt.Printf("Request %d allowed: %v\n", n, allowed)
		}(i)
	}

	wg.Wait()
	fmt.Printf("Total failure count: %d\n", mute.failureCount)
}
