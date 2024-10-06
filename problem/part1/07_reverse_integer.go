package main

/**
Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

Example 1:

Input: x = 123
Output: 321
Example 2:

Input: x = -123
Output: -321
Example 3:

Input: x = 120
Output: 21

Constraints:

-231 <= x <= 231 - 1 // 31승임

*/

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	stringInt := strconv.Itoa(x)
	result := ""
	for i := len(stringInt) - 1; i >= 0; i-- {
		if i == 0 {
			if string(stringInt[i]) == "-" {
				result = string(stringInt[i]) + result
			} else {
				result += string(stringInt[i])
			}
		} else {
			result += string(stringInt[i])
		}
	}

	res, _ := strconv.Atoi(result)

	if res >= math.MinInt32 && res <= math.MaxInt32 {
		return res
	} else {
		return 0
	}
}