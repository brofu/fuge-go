package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type SumCase struct {
	*BaseCase
	name           string
	skipped        bool
	debug          bool
	input          []int
	output         int
	expectedResult int
	fun            func([]int) int
}

func (c *SumCase) Execute(t *testing.T) {

	if c.fun != nil {
		c.output = c.fun(c.input)
	}
}

func (c *SumCase) DebugInfo(t *testing.T) {
	if c.debug {
		t.Log("input: ", c.input, "output: ", c.output, "expected: ", c.expectedResult)
	}
}

func (c *SumCase) Verify(t *testing.T) bool {
	return assert.Equal(t, c.expectedResult, c.output)
}

func TestSumCase(t *testing.T) {

	cases := []TestCase{
		&SumCase{
			BaseCase:       baseCase,
			name:           "case1",
			fun:            countQuadruplets,
			input:          []int{1, 2, 3, 6},
			expectedResult: 1,
			skipped:        false,
		},
		&SumCase{
			BaseCase:       baseCase,
			name:           "case2",
			fun:            countQuadruplets,
			input:          []int{1, 1, 1, 3, 5},
			expectedResult: 4,
			skipped:        false,
		},
		&SumCase{
			BaseCase:       baseCase,
			name:           "case3",
			fun:            countQuadruplets,
			input:          []int{2, 16, 9, 27, 3, 39},
			expectedResult: 2,
			skipped:        false,
		},
		&SumCase{
			BaseCase:       baseCase,
			name:           "case4",
			fun:            countQuadruplets,
			input:          []int{28, 8, 49, 85, 37, 90, 20, 8},
			expectedResult: 1,
			skipped:        false,
		},
		&SumCase{
			BaseCase:       baseCase,
			name:           "case5",
			fun:            countQuadruplets,
			input:          []int{9, 6, 8, 23, 39, 23},
			expectedResult: 2,
			skipped:        false,
			//debug:          true,
		},
	}

	RunTestCases(t, cases)
}

func TestCountQuadruplets(t *testing.T) {

	var (
		nums   []int
		result int
	)

	nums = []int{9, 6, 8, 23, 39, 23}
	t.Log(nums)
	result = countQuadruplets(nums)
	assert.Equal(t, 2, result)
}
