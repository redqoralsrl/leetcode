package main

import "strconv"

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

	result, _ := strconv.Atoi(result)

	return 0
}