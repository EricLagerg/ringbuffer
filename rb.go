// Copyright 2015 Eric Lagergren.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package ring

// Buffer is a ring buffer.
type Buffer struct {
	data  []int
	front int
	back  int
}

// NewBuffer returns a Buffer with the given size - 1
// filled with defaultVal.
func NewBuffer(defaultVal, size int) *Buffer {
	data := make([]int, size)

	// Go automatically zeros memory.
	if defaultVal != 0 {
		for i := range data {
			data[i] = defaultVal
		}
	}
	return &Buffer{data: data}
}

// Empty returns true if the Buffer is empty.
func (b *Buffer) Empty() bool {
	return b.front == b.back
}

// Full returns true if the Buffer is full.
func (b *Buffer) Full() bool {
	return b.front == (b.front+len(b.data)-1)%len(b.data)
}

// Clear clears the Buffer.
func (b *Buffer) Clear() {
	b.front = 0
	b.back = 0
}

// Push pushes a new int into the Buffer.
func (b *Buffer) Push(x int) (old int) {
	empty := 1
	if b.Empty() {
		empty = 0
	}

	idx := (b.front + empty) % len(b.data)
	old = b.data[idx]
	b.data[idx] = x
	b.front = idx
	if idx == b.back {
		b.back = (b.back + empty) % len(b.data)
	}
	return old
}

// Pop removes an int from the Buffer.
func (b *Buffer) Pop() (front int) {
	if b.Empty() {
		panic("ring.Buffer.Pop: cannot pop an empty Buffer")
	}

	front = b.data[b.front]
	if !b.Empty() {
		b.front = (b.front + len(b.data) - 1) % len(b.data)
	}
	return front
}

// Peek returns the first element of the Buffer without
// removing it.
func (b *Buffer) Peek() (front int) {
	if b.Empty() {
		panic("ring.Buffer.Peek: cannot pop an empty Buffer")
	}
	return b.data[b.front]
}
