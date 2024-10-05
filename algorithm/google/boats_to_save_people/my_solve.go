package main

import "sort"

/**
You are given an array people where people[i] is the weight of the ith person, and an infinite number of boats where each boat can carry a maximum weight of limit. Each boat carries at most two people at the same time, provided the sum of the weight of those people is at most limit.

Return the minimum number of boats to carry every given person.

Example 1:

Input: people = [1,2], limit = 3
Output: 1
Explanation: 1 boat (1, 2)
Example 2:

Input: people = [3,2,2,1], limit = 3
Output: 3
Explanation: 3 boats (1, 2), (2) and (3)
Example 3:

Input: people = [3,5,3,4], limit = 5
Output: 4
Explanation: 4 boats (3), (3), (4), (5)

Constraints:

1 <= people.length <= 5 * 104
1 <= people[i] <= limit <= 3 * 104
*/

/*
*
시간복잡도 0(n log n) - sort 사용해서
공간 복잡도 0(n) - sort 사용해서
*/
func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	boat := 0

	heavy := len(people) - 1
	light := 0

	for light <= heavy {
		if people[heavy]+people[light] > limit {
			heavy -= 1
			boat++
		} else {
			heavy -= 1
			light += 1
			boat++
		}
	}

	return boat
}
