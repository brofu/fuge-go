package leetcode

// 1
// key point. map to reduce time complexity

func twoSum(nums []int, target int) []int {

	dataMap := make(map[int]int)

	for index, num := range nums {

		nextOne := target - num

		nextIndex, ok := dataMap[nextOne]

		if ok {
			return []int{index, nextIndex}
		}

		dataMap[num] = index
	}

	return []int{}
}
