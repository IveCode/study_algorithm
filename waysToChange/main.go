package main

import "fmt"

func main() {
	fmt.Println(waysToChange(5))
	fmt.Println(waysToChange(10))
}

//var coins = []int{25, 10, 5, 1}
var coins = []int{1, 5, 10, 25}

func waysToChange(n int) int {
	array := make([]int, n+1)
	array[0] = 1
	for _, coin := range coins {
		for i := coin; i <= n; i++ {
			array[i] = (array[i] + array[i-coin]) % 1000000007
		}
		fmt.Println(coin, array)
	}

	return array[n]
}
