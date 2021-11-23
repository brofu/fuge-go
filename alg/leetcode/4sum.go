package leetcode

import (
	"sort"
)

// 18. 4sum

// Time O(n ^ (k-1)). For k=4, O(n^3). (k -2 loops) * O(n)
// Space O(n). We need O(k)O(k) space for the recursion. kk can be the same as nn in the worst case for the generalized algorithm. Note that, for the purpose of complexity analysis, we ignore the memory required for the output. ?? ==> why?

// Key Point. Regression to N-1 numbers
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return FindNumber(0, len(nums)-1, nums, 4, target)
}

// use pointer of nums
func fourSumLessMem(nums []int, target int) [][]int {
	sort.Ints(nums)
	return FindNumberLessM(0, len(nums)-1, &nums, 4, target)
}

// recursively resolve the number of numbers
// issue: nums is passed as value, which may lead to high memory? there would be so much difference.
// refer to implemention of `slice`
func FindNumber(start, end int, nums []int, number int, target int) (result [][]int) {

	if end-start+1 < number || number < 2 || target < nums[start]*number || target > nums[end]*number {
		return
	}

	if number == 2 {

		for start < end {
			sum := nums[start] + nums[end]
			if sum == target {
				result = append(result, []int{nums[start], nums[end]})
				start++
				end--
				for start < end && nums[start] == nums[start-1] {
					start++
				}
				for start < end && nums[end] == nums[end+1] {
					end--
				}
			} else if sum < target {
				start++
				for start < end && nums[start] == nums[start-1] {
					start++
				}
			} else {
				end--
				for start < end && nums[end] == nums[end+1] {
					end--
				}
			}
		}
		return

	} else {
		for index := start; index < end+1; index++ {
			if index == start || nums[index] != nums[index-1] {
				temp := FindNumber(index+1, end, nums, number-1, target-nums[index])
				for _, t := range temp {
					t = append(t, nums[index])
					result = append(result, t)
				}
			}
		}
		return
	}

}

func FindNumberLessM(start, end int, nums *[]int, number int, target int) (result [][]int) {

	if end-start+1 < number || number < 2 || target < (*nums)[start]*number || target > (*nums)[end]*number {
		return
	}

	if number == 2 {

		for start < end {
			sum := (*nums)[start] + (*nums)[end]
			if sum == target {
				result = append(result, []int{(*nums)[start], (*nums)[end]})
				start++
				end--
				for start < end && (*nums)[start] == (*nums)[start-1] {
					start++
				}
				for start < end && (*nums)[end] == (*nums)[end+1] {
					end--
				}
			} else if sum < target {
				start++
				for start < end && (*nums)[start] == (*nums)[start-1] {
					start++
				}
			} else {
				end--
				for start < end && (*nums)[end] == (*nums)[end+1] {
					end--
				}
			}
		}
		return

	} else {
		for index := start; index < end+1; index++ {
			if index == start || (*nums)[index] != (*nums)[index-1] {
				temp := FindNumberLessM(index+1, end, nums, number-1, target-(*nums)[index])
				for _, t := range temp {
					t = append(t, (*nums)[index])
					result = append(result, t)
				}
			}
		}
		return
	}

}
