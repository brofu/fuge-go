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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersV2(l1, l2)
}

// 8ms, 87.7%
// 4.8MB, 94.91%
// key point: link list operation
func addTwoNumbersV2(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		sum      = 0
		carry    = 0
		rootNode = &ListNode{Val: 0}
		p        = rootNode
		f        = rootNode
	)

	for l1 != nil && l2 != nil {

		sum = l1.Val + l2.Val + carry
		carry = sum / 10
		p.Val = sum % 10

		l1 = l1.Next
		l2 = l2.Next

		f = p
		newNode := &ListNode{}
		p.Next = newNode
		p = p.Next
	}

	for l1 != nil {
		sum = l1.Val + carry
		carry = sum / 10
		p.Val = sum % 10

		l1 = l1.Next

		f = p
		newNode := &ListNode{}
		p.Next = newNode
		p = p.Next

	}

	for l2 != nil {
		sum = l2.Val + carry
		carry = sum / 10
		p.Val = sum % 10

		l2 = l2.Next

		f = p
		newNode := &ListNode{}
		p.Next = newNode
		p = p.Next

	}

	if carry == 0 {
		f.Next = nil
	} else {
		p.Val = carry
	}

	return rootNode
}

func PowInt(x, y int) int64 {

	var result int64 = 1

	for ; y > 0; y-- {
		result *= int64(x)
	}

	return result
}

// convert to number, add then convert back.
// issue: may overflow
func addTwoNumbersV1(l1 *ListNode, l2 *ListNode) *ListNode {

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
