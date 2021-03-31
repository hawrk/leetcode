/**
* @Author:hawrkchen
* @Date:2021/3/31 9:56 上午
* @desc:
 */

package main

import (
	"context"
	"fmt"
	"time"
)

func hardWork(job interface{}) error {
	time.Sleep(time.Second * 10)
	return nil
}

func requestWork(ctx context.Context, job interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second * 2)
	defer cancel()

	done := make(chan error, 1)   // 关键：设置成不阻塞的， 如果是阻塞，超时后，
								 // 函数退出， done 一直没有接收者，会一直等待，资源无法释放
	go func() {
		done <- hardWork(job)
	}()

	select {
		case err := <-done: {
			fmt.Println("normal done")
			return err
		}
		case <- ctx.Done(): {
			fmt.Println("time out")
			return ctx.Err()
		}
	}
}


