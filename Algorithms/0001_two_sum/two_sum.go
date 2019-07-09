package leetcode

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	result := make([]int, 0)
	for i, v := range nums {
		numsMap[v] = i
	}
	for i, v := range nums {
		next := target - v
		if ni, exist := numsMap[next]; exist && i != ni {
			result = append(result, i)
			result = append(result, ni)
			break
		}
	}
	return result
}
