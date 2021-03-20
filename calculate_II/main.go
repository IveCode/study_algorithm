package main

import "fmt"

func main() {
	//fmt.Println(calculate("3+2*2"))     //7
	//fmt.Println(calculate(" 3/2 "))     //1
	//fmt.Println(calculate(" 3+5 / 2 ")) //5
	fmt.Println(calculate("1+1+1")) //3
}

func calculate(s string) int {
	var result int
	ops := []byte{}
	values := []int{}
	for i := 0; i < len(s); {
		switch s[i] {
		case ' ':
			i++
		case '+':
			if len(ops) > 0 {
				switch ops[len(ops)-1] {
				case '+':
					values[len(values)-2] += values[len(values)-1]
					values = values[:len(values)-1]
					ops = ops[:len(ops)-1]
				case '-':
					values[len(values)-2] -= values[len(values)-1]
					values = values[:len(values)-1]
					ops = ops[:len(ops)-1]
				}
			}
			ops = append(ops, '+')
			i++
		case '-':
			if len(ops) > 0 {
				switch ops[len(ops)-1] {
				case '+':
					values[len(values)-2] += values[len(values)-1]
					values = values[:len(values)-1]
					ops = ops[:len(ops)-1]
				case '-':
					values[len(values)-2] -= values[len(values)-1]
					values = values[:len(values)-1]
					ops = ops[:len(ops)-1]
				}
			}
			ops = append(ops, '-')
			i++
		case '*':
			ops = append(ops, '*')
			i++
		case '/':
			ops = append(ops, '/')
			i++
		default:
			num := 0
			for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			if len(ops) > 0 {
				switch ops[len(ops)-1] {
				case '*':
					values[len(values)-1] *= num
					ops = ops[:len(ops)-1]
				case '/':
					values[len(values)-1] /= num
					ops = ops[:len(ops)-1]
				default:
					values = append(values, num)
				}
			} else {
				values = append(values, num)
			}
		}
	}
	if len(ops) > 0 {
		switch ops[len(ops)-1] {
		case '+':
			if len(values) == 2 {
				result = values[0] + values[1]
			} else if len(values) == 1 {
				result = values[0]
			}
		case '-':
			if len(values) == 2 {
				result = values[0] - values[1]
			} else if len(values) == 1 {
				result = -values[0]
			}
		}
	} else if len(values) == 1 {
		result = values[0]
	}

	return result
}
