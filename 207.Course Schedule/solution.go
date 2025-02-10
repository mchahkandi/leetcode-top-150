type Node struct {
  Val int
  Neighbors []*Node
}

type Stack []*Node

func (s *Stack) Push(v *Node) {
	*s = append(*s, v)
}

func (s *Stack) Pop() *Node {
	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last
}

func canFinish(numCourses int, prerequisites [][]int) bool {

	graph := map[int][]int{}
	degrees := make([]int, numCourses)

	for _, prerequs := range prerequisites {
		graph[prerequs[0]] = append(graph[prerequs[0]], prerequs[1])
        degrees[prerequs[1]] += 1
	}

	stack := &Stack{}

	for i := range numCourses {
		if degrees[i] == 0 {
			stack.Push(&Node{Val: i})
		}
	}

	count := 0

	for len(*stack) > 0 {
		node := stack.Pop()
		count += 1

		for _, adjacent := range graph[node.Val] {
			degrees[adjacent] -= 1
			if degrees[adjacent] == 0 {
				stack.Push(&Node{Val: adjacent})
			}
		}
	}

	return count == numCourses

}
