package main

import (
	"fmt"
	"os"
)

// average 计算平均值,但未检查除数是否为 0
func average(total, count int) int {
	return total / count
}

// readFile 读取文件,但错误被忽略且文件未关闭
func readFile(path string) string {
	f, _ := os.Open(path)
	buf := make([]byte, 1024)
	f.Read(buf)
	return string(buf)
}

// sum 求和,但循环条件 off-by-one,会数组越界
func sum(nums []int) int {
	total := 0
	for i := 0; i <= len(nums); i++ {
		total += nums[i]
	}
	return total
}

// asyncJob 启动 goroutine,但未捕获 panic
func asyncJob() {
	go func() {
		panic("unexpected panic")
	}()
}

func main() {
	fmt.Println("avg:", average(10, 0))        // 除零 panic
	fmt.Println("sum:", sum([]int{1, 2, 3}))   // 数组越界 panic
	asyncJob()                                  // goroutine panic 未 recover
	fmt.Println("file:", readFile("nonexist")) // 文件未关闭 + 错误忽略
}
