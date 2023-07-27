package leetcode

/*
Implement a first in first out (FIFO) queue using only two stacks. The implemented queue should support all the functions of a normal queue (push, peek, pop, and empty).

Implement the MyQueue class:

    void push(int x) Pushes element x to the back of the queue.
    int pop() Removes the element from the front of the queue and returns it.
    int peek() Returns the element at the front of the queue.
    boolean empty() Returns true if the queue is empty, false otherwise.
*/

type MyQueue struct {
	q []int
}

func Constructor() MyQueue {
	return MyQueue{
		q: []int{},
	}
}

func (queue *MyQueue) Push(x int) {
	queue.q = append(queue.q, x)
}

func (queue *MyQueue) Pop() int {
	ret := queue.q[0]
	queue.q = queue.q[1:]
	return ret
}

func (queue *MyQueue) Peek() int {
	return queue.q[0]
}

func (queue *MyQueue) Empty() bool {
	return len(queue.q) == 0
}

/*
you did this in 10 mins congrats,
	but it could have been faster if you just new syntax
*/
