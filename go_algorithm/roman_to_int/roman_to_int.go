package roman_to_int

import (
	"fmt"
)

func RomanToInt(s string) int {
	res := 0
	var previous rune
	for _, c := range reverse(s) {
		//fmt.Println("I is ", reflect.TypeOf(rune(21)))
		//fmt.Println("romanToInt ", 'I', c)
		if c == 'I' && (previous == 'V' || previous == 'X') {
			fmt.Println("romanToInt ", 'I', c)
			res -= 1
		} else if c == 'X' && (previous == 'L' || previous == 'C') {
			res -= 10
		} else if c == 'C' && (previous == 'D' || previous == 'M') {
			res -= 100
		} else {
			res += intValueOf(c)
		}
		previous = c
		fmt.Println("previous ", previous, res)
	}
	return res
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func intValueOf(c rune) int {
	switch c {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}
