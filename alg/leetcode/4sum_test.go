package leetcode

import (
	"sort"
	"testing"
)

func TestFindNumber(t *testing.T) {

	result := [][]int{}

	nums := []int{-1, 1}
	result = FindNumber(0, len(nums)-1, nums, 2, 0)
	t.Log(result)

	nums = []int{-1, 1, 1}
	result = FindNumber(0, len(nums)-1, nums, 3, 1)
	t.Log(result)

	nums = []int{-1, 1, 1, 0, 2}
	sort.Ints(nums)
	result = FindNumber(0, len(nums)-1, nums, 3, 1)
	t.Log(result)

	nums = []int{1, 0, -1, 0, -2, 2}
	sort.Ints(nums)
	result = FindNumber(0, len(nums)-1, nums, 4, 0)
	t.Log(result)

}
