package main

import (
	"fmt"
	"sort"
)

/**
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.

Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.
Example 2:

Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.
Example 3:

Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.

Constraints:

3 <= nums.length <= 3000
-105 <= nums[i] <= 105
*/

func threeSum(nums []int) [][]int {
    if len(nums) < 3 {
        return [][]int{}
    }

    result := [][]int{}
    for i := 0; i < len(nums)-2; i++ {
        for j := i + 1; j < len(nums)-1; j++ {
            for k := j + 1; k < len(nums); k++ {
                if nums[i]+nums[j]+nums[k] == 0 {
                    triplet := []int{nums[i], nums[j], nums[k]}
                    sort.Ints(triplet)
                    result = append(result, triplet)
                }
            }
        }
    }

    unique := make(map[string]bool)
    answer := [][]int{}

    for _, triplet := range result {
        key := fmt.Sprintf("%d,%d,%d", triplet[0], triplet[1], triplet[2])
        if !unique[key] {
            unique[key] = true
            answer = append(answer, triplet)
        }
    }

    return answer
}