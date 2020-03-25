package main

import "fmt"

func main() {
	//fmt.Println(intToRoman(3))
	//fmt.Println(intToRoman(4))
	//fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(624))
	fmt.Println(intToRoman(2888))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))
}

func intToRoman(num int) string {
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var (
		result string
	)
	for num != 0 {
		for i := 0; i < len(nums); i++ {
			if num >= nums[i] {
				num -= nums[i]
				result += romans[i]
				break
			}
		}
	}
	return result
}
