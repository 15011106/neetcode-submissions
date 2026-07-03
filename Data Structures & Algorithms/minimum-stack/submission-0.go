type MinStack struct {
	valArr []int
	valMinStack []int
}

func Constructor() MinStack {
	var stack MinStack
	return stack
}

func (this *MinStack) Push(val int) {
	this.valArr = append(this.valArr, val)
	if len(this.valMinStack) > 0{
		if this.valMinStack[len(this.valMinStack)-1] > val{
			this.valMinStack = append(this.valMinStack, val)
		}else{
			this.valMinStack = append(this.valMinStack, this.valMinStack[len(this.valMinStack)-1])
		}
	}else{
			this.valMinStack = append(this.valMinStack, val)
		}
}

func (this *MinStack) Pop() {
	this.valArr = this.valArr[0:len(this.valArr)-1]
	this.valMinStack = this.valMinStack[0:len(this.valMinStack)-1]
}

func (this *MinStack) Top() int {
	return this.valArr[len(this.valArr)-1]
}

func (this *MinStack) GetMin() int {
	return this.valMinStack[len(this.valMinStack)-1]
}
