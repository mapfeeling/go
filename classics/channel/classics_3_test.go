package main

import (
	"fmt"
	"testing"
	"time"
)

func TestClassics3(t *testing.T) {
	var (
		chAB = make(chan struct{}, 1)
		chBC = make(chan struct{}, 1)
		chCA = make(chan struct{}, 1)
	)
	chCA <- struct{}{}
	go func() {
		for {
			select {
			case <-chCA:
				fmt.Println("我打印A")
				chAB <- struct{}{}
			}
		}
	}()
	go func() {
		for {
			select {
			case <-chAB:
				fmt.Println("我打印B")
				chBC <- struct{}{}
			}
		}
	}()
	go func() {
		for {
			select {
			case <-chBC:
				fmt.Println("我打印C")
				chCA <- struct{}{}
			}
		}
	}()
	time.Sleep(3 * time.Second)
}
