package main

// 1
func twoSum(nums []int, target int) []int {
	numberMap := make(map[int]int)
	for index, value := range nums {
		if otherIndex, ok := numberMap[target-value]; ok {
			return []int{otherIndex, index}
		}
	}
	return []int{}
}

// 2. Add two numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

// loop is better
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node := mergeTwoNode(l1, l2, 0)
	return node
}
func mergeTwoNode(l1 *ListNode, l2 *ListNode, val int) *ListNode {
	if l1 == nil && l2 == nil {
		if val == 0 {
			return nil
		}
		return &ListNode{Val: val, Next: nil}
	}

	if l1 == nil {
		return mergeTwoNode(&ListNode{Val: 0, Next: nil}, l2, val)
	}
	if l2 == nil {
		return mergeTwoNode(l1, &ListNode{Val: 0, Next: nil}, val)
	}

	index := new(ListNode)
	cal := l1.Val + l2.Val + val
	index.Val = cal % 10
	index.Next = mergeTwoNode(l1.Next, l2.Next, cal/10)

	return index
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	index := &ListNode{}
	head := index

	cal := 0
	for cal != 0 || l1 != nil || l2 != nil {
		sum := cal
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l1.Val
			l2 = l2.Next
		}

		node := &ListNode{}
		node.Val = sum % 10
		cal = sum / 10
		index.Next = node
		index = index.Next

	}
	return head

}
