/**
* @Author:hawrkchen
* @Date:2021/4/2 10:01 上午
* @desc:
 */

package main

import (
	"fmt"
	"strconv"
	"testing"
)
/*
*   beachmark test 基准测试：
*   1. 编写 XXXX_test.go 文件， 必须以 _test.go 结尾
*   2. 函数名必须以 Beachmark开判断， 带上参数 b *testing.B 函数可导出，无返回值
*   3. b.ResetTimer是重置计时器，这样可以避免for循环之前的初始化代码的干扰
*   4.go test beachmark_test.go -bench=. -benchmem -run=none
*   -bench=. -benchtime=3s -benchmem    指定运行3秒 和 打印分配 内存次数和分配的总内存
*   =======output=======
*   goos: darwin
*   goarch: amd64
*   cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
*   BenchmarkSprintf-12     18144073                63.15 ns/op            2 B/op          1 allocs/op
*   BenchmarkFormat-12      513897085                2.317 ns/op           0 B/op          0 allocs/op
*   BenchmarkItoa-12        494323935                2.415 ns/op           0 B/op          0 allocs/op
*   PASS
*   ok      command-line-arguments  4.147s

*   5. 参数解析：-12 cpu核数，也即 GOMAXPROCS  每秒执行次数  每次执行所需时间  每次执行分配空间大小  每次内存分配次数
*
 */

func BenchmarkSprintf(b *testing.B) {
	num := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", num)
	}
}

func BenchmarkFormat(b *testing.B) {
	num := int64(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.FormatInt(num, 10)
	}
}

func BenchmarkItoa(b *testing.B) {
	num := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.Itoa(num)
	}
}


