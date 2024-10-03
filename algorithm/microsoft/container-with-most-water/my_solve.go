package main

// 내가 푼거 (time limit 걸림)
/**
푼 방식 : 무차별 대입 방식
원인 : 시간 복잡도 중첩 루프여서 n2의 복잡도를 나타냄
공간 복잡도는 양호
*/
// func maxArea(height []int) int {
// 	res := 0
// 	for p1_index, p1_data := range height {
// 		for p2_index, p2_data := range height {
// 			if p1_index == p2_index {
// 				continue
// 			}
// 			bigger := p1_data
// 			if bigger > p2_data {
// 				bigger = p2_data
// 			}

// 			if res <= bigger*(p1_index-p2_index) {
// 				res = bigger * (p1_index - p2_index)
// 			}
// 		}
// 	}

// 	return res
// }

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0

	for left < right {
		length := min(height[left], height[right])
		width := right - left
		area := width * length
		maxArea = max(area, maxArea)
		if height[left] < height[right] {
			left += 1
		} else {
			right -= 1
		}
	}

	return maxArea
}
