/**
* @Author:hawrkchen
* @Date:2021/3/6 2:43 下午
* @desc:
 */

package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// test 1_two_sum.go
	input := []int{3,2,4 }
	target := 6
	out := twoSum(input, target)
	fmt.Println(out)

	//test timeout_demo.go
	const total = 1000
	var wg sync.WaitGroup
	wg.Add(total)
	now := time.Now()
	for i := 0; i < total; i++ {
		go func() {
			defer wg.Done()
			requestWork(context.Background(), "any")
		}()
	}
	wg.Wait()
	fmt.Println("elapsed:", time.Since(now))
	time.Sleep(time.Second * 10)
	fmt.Println("number of goroutines:", runtime.NumGoroutine())
}