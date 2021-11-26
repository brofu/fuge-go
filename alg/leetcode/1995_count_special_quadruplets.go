package leetcode

import (
	"fmt"
	"sort"
)

// first idea. O(n^4)
func countQuadrupletsV1(nums []int) int {
	var (
		result int
		length = len(nums)
	)

	for i := 0; i < length; i++ {
		for j := i + 1; i < length; j++ {
			for k := j + 1; k < length; k++ {
				for l := k + 1; l < length; l++ {
					if nums[i]+nums[j]+nums[k] == nums[l] {
						result++
					}
				}
			}
		}
	}

	return result
}

// wrong version
// think it's same as 4sum
func countQuadrupletsV2(nums []int) int {
	var (
		result int
		length = len(nums)
	)

	sort.Ints(nums)

	for i := length - 1; i > 0; i-- {
		temp := FindNumberWithDuplicate(0, i-1, nums, 3, nums[i])
		result += temp
	}

	return result
}

func FindNumberWithDuplicate(start, end int, nums []int, number int, target int) (result int) {

	if end-start+1 < number || number < 2 || target < nums[start]*number || target > nums[end]*number {
		return
	}

	if number == 2 {

		for start < end {
			sum := nums[start] + nums[end]
			if sum == target {
				fmt.Println(nums[start], nums[end], sum)
				result++
				start++
			} else if sum < target {
				start++
			} else {
				end--
			}
		}
		return

	} else {
		for index := start; index < end+1; index++ {
			temp := FindNumberWithDuplicate(index+1, end, nums, number-1, target-nums[index])
			result += temp
		}
		return
	}

}

// 17ms, 60%. 2.9M, 23.33%
// 8ms, 83.33%. 2.6M, 43.33%
// 24ms, 43.33%, 2.9M, 23.33%
// O(n^3) O(n)
func countQuadrupletsV3(nums []int) int {
	var (
		result int
		length = len(nums)
		m      = make(map[int]int)
	)

	for i := length - 2; i > 0; i-- {
		SetMapAdd(m, nums[i+1], 1)
		for j := i - 1; j > 0; j-- {
			for k := j - 1; k >= 0; k-- {
				result += GetMap(m, nums[i]+nums[j]+nums[k])
			}
		}
	}
	return result
}

// 8ms, 83.33%. 4.3M 10%
// 8ms, 83.33%. 4M 13.3%
// 8ms, 83.33%. 4M 13.3%
// O(n^2) O(n^2)
func countQuadruplets(nums []int) int {
	var (
		result int
		length = len(nums)
		m      = make(map[int]int)
	)

	SetMapAdd(m, nums[length-1]-nums[length-2], 1)

	for i := length - 3; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			result += GetMap(m, nums[i]+nums[j])
		}
		for l := length - 1; l > i; l-- {
			SetMapAdd(m, nums[l]-nums[i], 1)
		}
	}
	return result
}

func GetMap(m map[int]int, key int) int {
	if v, exits := m[key]; exits {
		return v
	}
	return 0
}

func SetMapAdd(m map[int]int, key, value int) {
	if v, exits := m[key]; exits {
		m[key] = v + value
	} else {
		m[key] = value
	}
}
