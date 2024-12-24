/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Dequeue []*TreeNode

func (q *Dequeue) PushBack(node *TreeNode) {
	*q = append(*q, node)
}

func (q *Dequeue) PushFront(node *TreeNode) {
	*q = append([]*TreeNode{node}, *q...)
}

func (q *Dequeue) PopFront() *TreeNode {
	if len(*q) == 0 {
		return nil
	}
	front := (*q)[0]
	*q = (*q)[1:]
	return front
}

func (q *Dequeue) PopBack() *TreeNode {
	if len(*q) == 0 {
		return nil
	}
	back := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return back
}


func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }

    array := [][]int{}

    queue := make(Dequeue, 0)

    queue.PushBack(root)

    reverse := false

    for len(queue) > 0 {
        level := len(queue)

        arr := []int{}

        for level != 0 {

            if reverse {
                node := queue.PopBack()

                arr = append(arr,node.Val)

                if node.Right != nil {
                    queue.PushFront(node.Right)
                }
                if node.Left != nil {
                    queue.PushFront(node.Left)
                }

            } else {
                node := queue.PopFront()

                arr = append(arr,node.Val)

                if node.Left != nil {
                    queue.PushBack(node.Left)
                }
                if node.Right != nil {
                    queue.PushBack(node.Right)
                }
            }
            level--
        }
        array = append(array,arr)
        reverse = !reverse
    }
    return array
}
