type MinStack struct {
    stack [][]int
}


func Constructor() MinStack {
    return MinStack{stack : [][]int{}}
}


func (this *MinStack) Push(val int)  {
    if len(this.stack) == 0 {
        this.stack = append(this.stack, []int{val, val})
    } else {
        last := this.stack[len(this.stack) - 1]
        this.stack = append(this.stack, []int{val, min(val, last[1])})
    }
}


func (this *MinStack) Pop()  {
    this.stack = this.stack[:len(this.stack) - 1]
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack) - 1][0]
}


func (this *MinStack) GetMin() int {
    return this.stack[len(this.stack) - 1][1]
}
