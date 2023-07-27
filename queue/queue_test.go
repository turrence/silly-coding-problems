package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	myQueue := Constructor()
	myQueue.Push(1)                    // queue is: [1]
	myQueue.Push(2)                    // queue is: [1, 2] (leftmost is front of the queue)
	assert.Equal(t, 1, myQueue.Peek()) // return 1
	assert.Equal(t, 1, myQueue.Pop())  // return 1
	assert.False(t, myQueue.Empty())   // return false
	assert.Equal(t, 2, myQueue.Pop())
	assert.True(t, myQueue.Empty())
}
