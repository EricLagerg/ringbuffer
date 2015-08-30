// Package ring provides a simple ring buffer implementation.
package ring

type Buffer struct {
	Data    []int
	Default int
	Front   int
	Back    int
	Empty   bool
}

// NewBuffer returns a pointer to a Buffer with the given size
// filled with defaultVal.
func NewBuffer(defaultVal, size int) *Buffer {

	ir := &Buffer{
		Data:    make([]int, size),
		Default: defaultVal,
		Empty:   true,
	}

	for i := 0; i < size; i++ {
		ir.Data[i] = defaultVal
	}

	return ir
}

// IsEmpty returns true if the Buffer is empty.
func (b *Buffer) IsEmpty() bool { return b.Empty }

// Push pushes a new int into the Buffer.
func (b *Buffer) Push(x int) int {
	empty := 1
	if b.Empty {
		empty = 0
	}

	idx := (b.Front + empty) % len(b.Data)
	old := b.Data[idx]
	b.Data[idx] = x
	b.Front = idx

	if idx == b.Back {
		b.Back = (b.Back + empty) % len(b.Data)
	}
	b.Empty = false

	return old
}

// Pop removes an int from the Buffer.
func (b *Buffer) Pop() int {
	if b.IsEmpty() {
		panic("Cannot pop empty Buffer")
	}

	top := b.Data[b.Front]
	b.Data[b.Front] = b.Default

	if b.Front == b.Back {
		b.Empty = true
	} else {
		b.Front = (b.Front + len(b.Data) - 1) % len(b.Data)
	}

	return top
}
