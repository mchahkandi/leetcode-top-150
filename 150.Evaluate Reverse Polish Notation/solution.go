type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last
}

func evalRPN(tokens []string) int {
	stack := make(Stack, 0)

	for _, val := range tokens {
		if value, err := strconv.Atoi(val); err == nil {
			stack.Push(value)
		} else {
			var sum int
			last := stack.Pop()
			secondLast := stack.Pop()
			switch val {
			case "+":
				sum = secondLast + last
			case "*":
				sum = secondLast * last
			case "/":
				sum = secondLast / last
			case "-":
				sum = secondLast - last
			}
			stack.Push(sum)
		}
	}
	return stack.Pop()
}
