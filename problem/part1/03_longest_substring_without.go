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
    if len(s) <= 1 {
        return len(s)
    }

    value := make(map[byte]int)
    max_len := 0

    left := 0
    right := 0
    for left < len(s) && right < len(s) {
        if index, ok := value[s[right]]; ok {
            left = max(left, index+1)
        }

        value[s[right]] = right
        max_len = max(max_len, right - left + 1)
        right++
    }

    return max_len
}