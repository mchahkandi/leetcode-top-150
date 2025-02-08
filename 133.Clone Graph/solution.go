type Stack []*Node

func (s *Stack) Push(v *Node) {
	*s = append(*s, v)
}

func (s *Stack) Pop() *Node {
	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := map[*Node]*Node{}

	stack := Stack{}

	cop := &Node{Val: node.Val}

	visited[node] = cop

	stack.Push(node)

	for len(stack) > 0 {
		curr := stack.Pop()
		for _, n := range curr.Neighbors {
			if _, ok := visited[n]; ok {
				visited[curr].Neighbors = append(visited[curr].Neighbors, visited[n])
			}else{
				deep := &Node{Val: n.Val}
				visited[n] = deep
				visited[curr].Neighbors = append(visited[curr].Neighbors, deep)
				stack.Push(n)
			}
		}

	}
	return cop
}
