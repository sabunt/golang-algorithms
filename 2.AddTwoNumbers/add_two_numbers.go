package addtwonumbers

//ListNode struct resolve problem
type ListNode struct {
	Val  int
	Next *ListNode
}

func add(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	resultTail := result
	var carry, value1, value2 int
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			value1 = l1.Val
		} else {
			value1 = 0
		}
		if l2 != nil {
			value2 = l2.Val
		} else {
			value2 = 0
		}
		sum := value1 + value2 + carry
		carry = sum / 10
		out := sum % 10

		resultTail.Next = &ListNode{out, nil}
		resultTail = resultTail.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

	}
	return result.Next
}