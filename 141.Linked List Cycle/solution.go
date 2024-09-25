/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    hashMap := map[*ListNode]bool{}
    node := head

    for node != nil {
        _, exists := hashMap[node]
        if exists {
            return true
        }else{
            hashMap[node] = true
        }
        node = node.Next
    }
    return false
}
