package main

import "fmt"

func main() {
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("  hello world!  "))
	fmt.Println(reverseWords("a good   example"))
}

func reverseWords(s string) string {
	var (
		results []string
		tempStr string
	)

	for _, v := range s {
		if v != ' ' {
			tempStr += string(v)
		} else if len(tempStr) > 0 {
			results = append(results, tempStr)
			tempStr = ""
		}
	}

	if len(tempStr) > 0 {
		results = append(results, tempStr)
	}

	for i, j := 0, len(results)-1; i < j; {
		results[i], results[j] = results[j], results[i]
		i++
		j--
	}

	tempStr = ""
	for i, v := range results {
		tempStr += v
		if i != len(results)-1 {
			tempStr += " "
		}
	}

	return tempStr
}
