package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(fib1(20))
	fmt.Println(time.Now().Sub(start))

	start = time.Now()
	values := map[int]int{}
	fmt.Println(fib2(20, values))
	fmt.Println(time.Now().Sub(start))

	start = time.Now()
	fmt.Println(fib3(20))
	fmt.Println(time.Now().Sub(start))

	start = time.Now()
	fmt.Println(fib4(20))
	fmt.Println(time.Now().Sub(start))
}

func fib1(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	return fib1(n-1) + fib1(n-2)
}

func fib2(n int, values map[int]int) int {
	if n == 1 || n == 2 {
		return 1
	}

	if v, ok := values[n]; ok {
		return v
	}
	values[n] = fib2(n-1, values) + fib2(n-2, values)

	return values[n]
}

func fib3(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 1
	for i := 3; i < len(dp); i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func fib4(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	v1, v2 := 1, 1
	for i := 3; i <= n; i++ {
		sum := v1 + v2
		v1, v2 = v2, sum
	}
	return v2
}
