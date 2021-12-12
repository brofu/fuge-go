package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func ConstructListNodeFromArray(input []int) *ListNode {
	var (
		root = &ListNode{}
		p    = root
		f    = root
	)

	for _, v := range input {
		p.Val = v

		f = p
		p.Next = &ListNode{}
		p = p.Next
	}

	if p.Val == 0 && p.Next == nil {
		f.Next = nil
	}

	return root
}

func ConstructArrayFromListNode(node *ListNode) []int {
	if node == nil {
		return nil
	}

	array := []int{}

	for node != nil {
		array = append(array, node.Val)
		node = node.Next
	}

	return array
}

func TestConstructList(t *testing.T) {
	input := []int{1, 2, 8}
	node := ConstructListNodeFromArray(input)
	output := ConstructArrayFromListNode(node)

	assert.Equal(t, len(input), len(output))

	for index := 0; index < len(input); index++ {
		assert.Equal(t, input[index], output[index])
	}
}

type AddTwoNumbersCase struct {
	*BaseCase
	name           string
	skipped        bool
	debug          bool
	input1         *ListNode
	input2         *ListNode
	output         *ListNode
	expectedResult []int
	fun            func(*ListNode, *ListNode) *ListNode
}

func (c *AddTwoNumbersCase) Execute(t *testing.T) {

	if c.fun != nil {
		c.output = c.fun(c.input1, c.input2)
	}
}

func (c *AddTwoNumbersCase) DebugInfo(t *testing.T) {
	if c.debug {
		t.Log("input1: ", ConstructArrayFromListNode(c.input1),
			"\ninput2: ", ConstructArrayFromListNode(c.input2),
			"\noutput: ", ConstructArrayFromListNode(c.output),
			"\nexpected: ", c.expectedResult)
	}
}

func (c *AddTwoNumbersCase) Verify(t *testing.T) bool {

	array := ConstructArrayFromListNode(c.output)
	if !assert.Equal(t, len(c.expectedResult), len(array)) {
		return false
	}
	for index := 0; index < len(array); index++ {
		if !assert.Equal(t, c.expectedResult[index], array[index]) {
			return false
		}
	}
	return true
}

func TestAddTwoNumbers(t *testing.T) {
	cases := []TestCase{
		&AddTwoNumbersCase{
			BaseCase:       baseCase,
			name:           "case1",
			fun:            addTwoNumbersV2,
			input1:         ConstructListNodeFromArray([]int{1, 2, 3, 6}),
			input2:         ConstructListNodeFromArray([]int{1, 2, 3, 6}),
			expectedResult: []int{2, 4, 6, 2, 1},
			skipped:        false,
		},
		&AddTwoNumbersCase{
			BaseCase:       baseCase,
			name:           "case1",
			fun:            addTwoNumbersV2,
			input1:         ConstructListNodeFromArray([]int{1, 2, 3, 4}),
			input2:         ConstructListNodeFromArray([]int{1, 2, 3, 4}),
			expectedResult: []int{2, 4, 6, 8},
			skipped:        false,
		},
		&AddTwoNumbersCase{
			BaseCase:       baseCase,
			name:           "case1",
			fun:            addTwoNumbersV1,
			input1:         ConstructListNodeFromArray([]int{1, 2, 3, 4}),
			input2:         ConstructListNodeFromArray([]int{1, 2, 3, 4}),
			expectedResult: []int{2, 4, 6, 8},
			skipped:        false,
		},

		&AddTwoNumbersCase{
			BaseCase:       baseCase,
			name:           "case1",
			fun:            addTwoNumbersV1,
			input1:         ConstructListNodeFromArray([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
			input2:         ConstructListNodeFromArray([]int{5, 6, 4}),
			expectedResult: []int{6, 5, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			skipped:        false,
			debug:          true,
		},
	}

	RunTestCases(t, cases)
}
