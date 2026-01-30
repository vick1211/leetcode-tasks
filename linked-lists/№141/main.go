package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	rabbit, turtle := head, head 
	
	for rabbit != nil {
		rabbit = rabbit.Next.Next 
		if rabbit == turtle {
			return true
		}
		turtle = turtle.Next
	}
	return false
}