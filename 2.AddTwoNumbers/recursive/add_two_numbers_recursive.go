package addtwonumbers

//ListNode struct resolve problem
type ListNode struct {
	Val  int
	Next *ListNode
}

func addRecursive(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil && l2 != nil {
		return l2
	} else if l1 != nil && l2 == nil {
		return l1
	}

	sum := l1.Val + l2.Val
	next := addRecursive(l1.Next, l2.Next)

	if sum >= 10 {
		sum %= 10
		next = addRecursive(next, &ListNode{1, nil})
	}

	return &ListNode{sum,next}

}
