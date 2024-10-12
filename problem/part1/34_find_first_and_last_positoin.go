package main

/**
Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.

Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
Example 3:

Input: nums = [], target = 0
Output: [-1,-1]

Constraints:

0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums is a non-decreasing array.
-109 <= target <= 109
*/

// 답은 맞으나 시간 복잡도는 O(n)
func searchRange(nums []int, target int) []int {
    left := 0
    right := len(nums) - 1

    l, r := -1, -1
    for (l == -1 || r == -1) && left <= right {
        if l == -1 {
            if nums[left] == target {
                l = left
            } else {
                left++
            }
        }
        if r == -1 {
            if nums[right] == target {
                r = right
            } else {
                right--
            }
        }
    }

    return []int{l, r}
}

/**
func binar_search(nums []int, target int, move_left bool) int {
	l := 0
	r := len(nums) - 1
	found_at := -1
	for l <= r {
		m := int(float32(l+r) / 2)
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			found_at = m
			if move_left {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}
	return found_at
}

func searchRange(nums []int, target int) []int {
	left_most := -1
	right_most := -1
	left_most = binar_search(nums, target, true)
	right_most = binar_search(nums, target, false)
	return []int{left_most, right_most}
}
*/