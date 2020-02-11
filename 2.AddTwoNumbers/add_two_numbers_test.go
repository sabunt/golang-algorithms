package addtwonumbers

import (
	"reflect"
	"testing"
)

func Test1_add(t *testing.T) {
	var l1, l2, l3 ListNode
	l1 = ListNode{1, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}
	l2 = ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	l3 = ListNode{6, &ListNode{6, &ListNode{4, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}

	t.Run("Test for longly list node", func(t *testing.T) {
		if got := add(&l1, &l2); !reflect.DeepEqual(got, &l3) {
			t.Errorf("add() = %v, want %v", got, l3)
		}
	})
}

func Test_add(t *testing.T) {
	var l1, l2, l3 ListNode
	l1 = ListNode{8, &ListNode{5, &ListNode{1, nil}}}
	l2 = ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	l3 = ListNode{3, &ListNode{2, &ListNode{6, nil}}}

	t.Run("Test for shortly list node", func(t *testing.T) {
		if got := add(&l1, &l2); !reflect.DeepEqual(got, &l3) {
			t.Errorf("add() = %v, want %v", got, l3)
		}
	})
}

func Test2_add(t *testing.T) {
	var l1, l2, l3 ListNode
	l1 = ListNode{0, nil}
	l2 = ListNode{0, nil}
	l3 = ListNode{0, nil}

	t.Run("Test for zero list node", func(t *testing.T) {
		if got := add(&l1, &l2); !reflect.DeepEqual(got, &l3) {
			t.Errorf("add() = %v, want %v", got, l3)
		}
	})
}
