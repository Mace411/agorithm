package main

import "fmt"

func main() {
	n := 10
	fmt.Printf("斐波那契数列第%d项为%d", n, fibonacci(n))
}

//返回斐波那契数列的第n项
func fibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	var (
		prePre = 1
		pre    = 1
		cur    int
	)
	for i := 3; i <= n; i++ {
		cur = prePre + pre
		prePre = pre
		pre = cur
	}
	return cur
}
