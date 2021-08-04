package mergetwosortedlists

// ListNode holding array in a linked list manner
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) getList() []int {
	finalList := []int{}
	if l == nil {
		return finalList
	}
	for {
		finalList = append(finalList, l.Val)
		if l.Next == nil {
			break
		} else {
			l = l.Next
		}
	}
	return finalList
}

func makeListNodes(list []int) *ListNode {
	node := &ListNode{}
	if len(list) == 0 {
		return nil
	}
	if len(list) == 1 {
		node.Val = list[0]
	} else {
		node.Val = list[0]
		node.Next = makeListNodes(list[1:])
	}

	return node
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	list := []int{}
	if l1 == nil && l2 == nil {
		return nil
	}
	for {
		if l1 == nil && l2 != nil {
			list = append(list, l2.getList()...)
			break
		}
		if l2 == nil && l1 != nil {
			list = append(list, l1.getList()...)
			break
		}
		if l1.Val <= l2.Val {
			list = append(list, l1.Val)
			l1 = l1.Next
		} else {
			list = append(list, l2.Val)
			l2 = l2.Next
		}
	}
	return makeListNodes(list)
}
