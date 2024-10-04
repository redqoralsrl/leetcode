package main

/**
Given an array of integers arr, return true if and only if it is a valid mountain array.

Recall that arr is a mountain array if and only if:

arr.length >= 3
There exists some i with 0 < i < arr.length - 1 such that:
arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
arr[i] > arr[i + 1] > ... > arr[arr.length - 1]

Example 1:

Input: arr = [2,1]
Output: false
Example 2:

Input: arr = [3,5,5]
Output: false
Example 3:

Input: arr = [0,3,2,1]
Output: true

Constraints:

1 <= arr.length <= 104
0 <= arr[i] <= 104
*/

/*
*
시간복잡도 0(n) 공간 복잡도 0(1)
*/
// func validMountainArray(arr []int) bool {
// 	if len(arr) <= 2 || arr[0] > arr[1] {
// 		return false
// 	}
// 	answer := true
// 	leftHeight := -1
// 	rightHeight := 0

// 	isChange := false
// 	// 0321
// 	for _, data := range arr {
// 		if isChange {
// 			// 내려갈때
// 			if rightHeight > data {
// 				rightHeight = data
// 			} else {
// 				return false
// 			}
// 		} else {
// 			// 올라갈때
// 			if leftHeight < data {
// 				leftHeight = data
// 			} else if leftHeight == data {
// 				return false
// 			} else {
// 				rightHeight = data
// 				isChange = true
// 			}
// 		}
// 	}
// 	if !isChange {
// 		// ex) 0,1,2,3,4,5,6,7
// 		return false
// 	}
// 	return answer
// }

// 코드는 간단하나 시간은 느림
// 단 공간복잡도는 우수
func validMountainArray(arr []int) bool {
	if len(arr) <= 2 || arr[0] > arr[1] { // 시간 상승을 위한 트릭
		return false
	}
	i := 1

	for i < len(arr) && arr[i] > arr[i-1] {
		i++

	}

	if i == 1 || i == len(arr) {
		return false
	}

	for i < len(arr) && arr[i] < arr[i-1] {
		i++
	}

	return i == len(arr)
}
