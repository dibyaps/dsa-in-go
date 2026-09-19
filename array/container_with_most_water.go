package array

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	max_water := 0

	for left <= right {
		width := right - left
		height_limit := min(height[left], height[right])

		max_water = max(max_water, (width * height_limit))

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max_water

}
