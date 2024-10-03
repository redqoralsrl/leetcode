package main

/*
*
Given a string s, find the length of the longest
substring

	without repeating characters.

Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.
*/
// 슬라이딩 윈도우 기법
func lengthOfLongestSubstring(s string) int {
	answer := make(map[byte]int)
	maxLen := 0
	start := 0

	for i := 0; i < len(s); i++ {
		if lastIdx, ok := answer[s[i]]; ok && lastIdx >= start {
			start = lastIdx + 1
		}

		answer[s[i]] = i

		maxLen = max(maxLen, i-start+1)
	}

	return maxLen
}
