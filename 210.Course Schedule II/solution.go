type Stack []any

func (s *Stack) Push(v any) {
	*s = append(*s, v)
}

func (s *Stack) Pop() any {
	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last
}

func findOrder(numCourses int, prerequisites [][]int) []int {

	graph := map[int][]int{}
	indegrees := make([]int, numCourses)

	for _, prerequisite := range prerequisites {
		graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
		indegrees[prerequisite[0]] += 1
	}

	stack := &Stack{}

	for i := range numCourses {
		if indegrees[i] == 0 {
			stack.Push(i)
		}
	}

	result := []int{}

	for len(*stack) > 0 {
		node := stack.Pop()
		result = append(result, node.(int))
		for _, prerequisite := range graph[node.(int)] {
			indegrees[prerequisite] -= 1
			if indegrees[prerequisite] == 0 {
				stack.Push(prerequisite)
			}
		}
	}

	if len(result) == numCourses {
		return result
	}else {
		return []int{}
	}
}
