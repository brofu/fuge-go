package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func PowInt(x, y int) int64 {

	var result int64 = 1

	for ; y > 0; y-- {
		result *= int64(x)
	}

	return result
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var (
		n1, n2, sum int64
	)

	for n := 0; l1 != nil; n++ {
		n1 += int64(l1.Val) * PowInt(10, n)
		l1 = l1.Next
	}

	for n := 0; l2 != nil; n++ {
		n2 += int64(l2.Val) * PowInt(10, n)
		l2 = l2.Next
	}

	sum = n1 + n2

	r := &ListNode{}
	t := r

	for {
		t.Val = int(sum % 10)
		sum = sum / 10
		if sum == 0 {
			break
		}
		t.Next = &ListNode{}
		t = t.Next
	}

	return r
}
