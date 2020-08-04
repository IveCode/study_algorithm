package main

import (
	"fmt"
	"strconv"
)

//https://leetcode-cn.com/problems/add-strings/

func main() {
	//fmt.Println(addStrings("0", "0"))
	//fmt.Println(addStrings("123456789", "987654321"))
	//fmt.Println(addStrings("123456789", "87654321"))
	fmt.Println(addStrings("9", "99"))
}

func addStrings(num1 string, num2 string) string {
	numLength1, numLength2 := len(num1), len(num2)

	var result string
	var carry int

	numLength1--
	numLength2--
	for numLength1 >= 0 && numLength2 >= 0 {
		r1 := int(num1[numLength1] - '0')
		r2 := int(num2[numLength2] - '0')
		r := r1 + r2 + carry

		if r >= 10 {
			carry = r / 10
			r = r % 10
		} else {
			carry = 0
		}

		result = strconv.Itoa(r) + result
		numLength1--
		numLength2--
	}

	for numLength1 >= 0 {
		r := int(num1[numLength1]-'0') + carry

		if r >= 10 {
			carry = r / 10
			r = r % 10
		} else {
			carry = 0
		}

		result = strconv.Itoa(r) + result
		numLength1--
	}

	for numLength2 >= 0 {
		r := int(num2[numLength2]-'0') + carry

		if r >= 10 {
			carry = r / 10
			r = r % 10
		} else {
			carry = 0
		}

		result = strconv.Itoa(r) + result
		numLength2--
	}

	if carry > 0 {
		result = strconv.Itoa(carry) + result
	}

	return result
}
