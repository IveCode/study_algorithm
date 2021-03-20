package main

import "fmt"

//https://leetcode-cn.com/problems/open-the-lock/

func main() {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
	fmt.Println(openLock([]string{"8888"}, "0009"))
	fmt.Println(openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888"))
	fmt.Println(openLock([]string{"0000"}, "8888"))
}

func plusOne(s string, j int) string {
	bts := []byte(s)
	if bts[j] == '9' {
		bts[j] = '0'
	} else {
		bts[j]++
	}
	return string(bts)
}

func minusOne(s string, j int) string {
	bts := []byte(s)
	if bts[j] == '0' {
		bts[j] = '9'
	} else {
		bts[j]--
	}
	return string(bts)
}

func openLock(deadends []string, target string) int {
	var queue []string
	queue = append(queue, "0000")
	length := len(queue)
	deadenedMaps := map[string]struct{}{}
	for i := 0; i < len(deadends); i++ {
		if deadends[i] == target {
			return -1
		}
		deadenedMaps[deadends[i]] = struct{}{}
	}
	visited := map[string]struct{}{
		"0000": {},
	}
	depth := 0
	for length != 0 {
		for i := 0; i < length; i++ {
			value := queue[i]

			if _, ok := deadenedMaps[value]; ok {
				continue
			}

			if value == target {
				return depth
			}

			for j := 0; j < 4; j++ {
				plusValue := plusOne(value, j)
				if _, ok := visited[plusValue]; !ok {
					queue = append(queue, plusValue)
					visited[plusValue] = struct{}{}
				}
				minusValue := minusOne(value, j)
				if _, ok := visited[minusValue]; !ok {
					queue = append(queue, minusValue)
					visited[minusValue] = struct{}{}
				}
			}
		}
		queue = queue[length:]
		length = len(queue)
		depth++
	}
	return -1
}
