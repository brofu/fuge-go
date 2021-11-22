package leetcode

func threeSum(nums []int) [][]int {

	result := [][]int{}

	if len(nums) < 3 {
		return result
	}

	hit := [][]int{}
	for index, num := range nums {

		target := 0 - num

		tempResult := twoSumValueWithOrder(nums, target, index)

		result = append(result, tempResult...)
	}

	for _, pair := range result {
		x, y := pair[0], pair[1]
		tempResult := make([]int, 3)

		num := 0 - (x + y)

		if num <= x {
			tempResult[0] = num
			tempResult[1] = x
			tempResult[2] = y
		}

		if num > y {
			tempResult[0] = x
			tempResult[1] = y
			tempResult[2] = num
		}

		if x < num && num <= y {
			tempResult[0] = x
			tempResult[1] = num
			tempResult[2] = y
		}

		for _, r := range hit {
			if !compareTwoList(tempResult, r) {
				hit = append(hit, tempResult)
			}
		}
	}

	return hit
}

func compareTwoList(nums1, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}

	return nums1[0] == nums2[0] && nums1[1] == nums2[1]
}

func twoSumValueWithOrder(nums []int, target int, index int) [][]int {
	dataMap := make(map[int]bool)
	result := [][]int{}

	for idx, num := range nums {

		if idx == index {
			continue
		}

		nextOne := target - num

		_, exists := dataMap[nextOne]

		if exists {
			var temp []int
			if num <= nextOne {
				temp = []int{num, nextOne}
			} else {
				temp = []int{nextOne, num}
			}
			result = append(result, temp)
		} else {
			dataMap[num] = true
		}
	}

	return result
}

func twoSumWithSkip(nums []int, target int, index int) [][]int {

	dataMap := make(map[int]bool)
	result := [][]int{}

	for idx, num := range nums {

		if idx == index {
			continue
		}

		nextOne := target - num

		_, exists := dataMap[nextOne]

		if exists {
			result = append(result, []int{num, nextOne, 0 - target})
		} else {
			dataMap[num] = true
		}
	}

	return result
}
