package main

/*
*
Given a string s, return the longest
palindromic

substring

	in s.

Example 1:

Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
Example 2:

Input: s = "cbbd"
Output: "bb"

Constraints:

1 <= s.length <= 1000
s consist of only digits and English letters.
*/
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	var start, end int

	expandAroundCenter := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		if right-left-1 > end-start {
			start = left + 1
			end = right - 1
		}
	}

	for i := 0; i < len(s); i++ {
		expandAroundCenter(i, i)
		expandAroundCenter(i, i+1)
	}

	return s[start : end+1]
}
