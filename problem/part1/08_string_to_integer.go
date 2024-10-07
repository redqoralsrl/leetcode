package main

import (
	"math"
	"unicode"
)

func myAtoi(s string) int {
	i, n := 0, len(s)
	for i < n && s[i] == ' ' {
		i++
	}

	sign := 1
	if i < n && (s[i] == '-' || s[i] == '+') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	result := 0
	for i < n && unicode.IsDigit(rune(s[i])) {
		digit := int(s[i] - '0')
		if result > (math.MaxInt32-digit)/10 {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
		result = result*10 + digit
		i++
	}

	return result * sign
}