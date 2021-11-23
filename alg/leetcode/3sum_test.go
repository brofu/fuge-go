package leetcode

import "testing"

func TestTwoSumWithSkip(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := twoSumValueWithOrder(nums, 1, 0)
	t.Log(result)
}

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	t.Log(result)
}
func TestThreeSumV1(t *testing.T) {
	nums := []int{0, 0, 0}
	result := threeSum(nums)
	t.Log(result)
}
