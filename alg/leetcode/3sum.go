package leetcode

import (
	"fmt"
	"sort"
)

// Key Points
// 2 pointer
// sort and remove duplicate in loop control

func threeSum(nums []int) [][]int {

	var (
		result = [][]int{}
		n      = len(nums)
	)

	if n < 3 {
		return result
	}

	sort.Ints(nums)

	for index := 0; index < n-2; index++ {

		if index > 0 && nums[index] == nums[index-1] {
			continue
		}

		var (
			target = -nums[index]
			left   = index + 1
			right  = n - 1
		)

		for left < right {

			sum := nums[left] + nums[right]
			if sum == target {
				result = append(result, []int{-target, nums[left], nums[right]}) // collect result
				left++
				right--
				for left < right && (nums[left] == nums[left-1]) { // adjust left pointer
					left++
				}
				for left < right && (nums[right] == nums[right+1]) { // adjust right pointer
					right--
				}
			} else if sum < target {
				left++
				for nums[left] == nums[left-1] { // adjust left pointer
					left++
				}
			} else {
				right--
				for nums[right] == nums[right+1] { // adjust right pointer
					right--
				}
			}

		}
	}

	return result
}

// deprecated version
// exceed running time
// Key point: remove duplicate results

func threeSumDeprecated(nums []int) [][]int {

	var (
		result      = [][]int{}
		finalResult = [][]int{}
		hit         = []map[int]int{}
	)

	if len(nums) < 3 {
		return result
	}

	for index, num := range nums {
		target := 0 - num
		tempResult := twoSumValue(nums, target, index)
		result = append(result, tempResult...)
	}

	for _, pair := range result {

		tr := map[int]int{}

		for _, v := range pair {
			MapAdd(tr, v)
		}
		MapAdd(tr, 0-(pair[0]+pair[1]))

		fmt.Println("flag", tr)

		if len(hit) == 0 {
			hit = append(hit, tr)
		} else {
			exists := false
			for _, r := range hit {
				if compareTwoMap(tr, r) {
					exists = true
				}
			}
			if !exists {
				hit = append(hit, tr)
			}
		}
	}

	for _, v := range hit {
		t := make([]int, 3)
		index := 0
		for k, count := range v {
			for idx := 0; idx < count; idx++ {
				t[index] = k
				index++
			}
		}
		finalResult = append(finalResult, t)
	}

	return finalResult
}

func MapAdd(data map[int]int, ele int) {
	if _, exists := data[ele]; exists {
		data[ele]++
	} else {
		data[ele] = 1
	}
}

func compareTwoMap(nums1, nums2 map[int]int) (equal bool) {
	if len(nums1) != len(nums2) {
		return false
	}

	for k := range nums1 {
		if _, exists := nums2[k]; !exists {
			return false
		}
	}

	return true
}

func compareTwoList(nums1, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}

	return nums1[0] == nums2[0] && nums1[1] == nums2[1]
}

func twoSumValue(nums []int, target int, index int) [][]int {
	dataMap := make(map[int]bool)
	result := [][]int{}

	for idx, num := range nums {

		if idx == index {
			continue
		}

		nextOne := target - num

		_, exists := dataMap[nextOne]

		if exists {
			result = append(result, []int{num, nextOne})
		} else {
			dataMap[num] = true
		}
	}

	return result
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
