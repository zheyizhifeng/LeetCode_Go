// @algorithm @lc id=142 lang=golang 
// @title linked-list-cycle-ii


package main
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    fast, slow := head, head
		for fast != nil && fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
			if fast == slow {
				fast = head
				for fast != slow {
					fast = fast.Next
					slow = slow.Next
				}
				return slow
			}
		}
		return nil
}