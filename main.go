package main

import (
	"fmt"
)

func main() {
	string := "()[]{}"
	string = "([{(}])"
	fmt.Println(isValid(string))

}

func isValid(s string) bool {
	str := make([]int32, 0, len(s))
	for _, v := range s {
		switch v {
		case 40, 91, 123:
			str = append(str, v)
		case 41:
			l := len(str)
			if str[l-1] == 40 {
				str = append(str[:l-1])
			}
		case 93:
			l := len(str)
			if str[l-1] == 91 {
				str = append(str[:l-1])
			}
		case 125:
			l := len(str)
			if str[l-1] == 123 {
				str = append(str[:l-1])
			}
		}
	}
	return len(str) == 0
}
