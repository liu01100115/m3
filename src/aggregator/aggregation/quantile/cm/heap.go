// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cm

// minHeap is a typed min heap for floating point numbers. Unlike the generic
// heap in the container/heap package, pushing data to or popping data off of
// the heap doesn't require conversion between floats and interface{} objects,
// therefore avoiding the memory and GC overhead due to the additional allocations.
type minHeap []float64

// Len returns the number of values in the heap.
func (h minHeap) Len() int { return len(h) }

// Min returns the minimum value from the heap.
func (h minHeap) Min() float64 { return h[0] }

// Push pushes a value onto the heap.
func (h *minHeap) Push(value float64) {
	// append
	*h = append(*h, value)

	// then, shift up if necessary. manually inlined.
	heap := *h
	i := len(heap) - 1
	for {
		parent := (i - 1) / 2
		if parent == i || heap[parent] <= heap[i] {
			break
		}
		heap[parent], heap[i] = heap[i], heap[parent]
		i = parent
	}
}

// Pop pops the minimum value from the heap.
func (h *minHeap) Pop() float64 {
	var (
		old = *h
		n   = old.Len()
		val = old[0]
	)

	old[0], old[n-1] = old[n-1], old[0]
	i, nn := 0, n-1 // heapify, manually inlined.
	for {
		left := i*2 + 1
		right := left + 1
		smallest := i
		if left < nn && old[left] < old[smallest] {
			smallest = left
		}
		if right < nn && old[right] < old[smallest] {
			smallest = right
		}
		if smallest == i {
			break
		}
		old[i], old[smallest] = old[smallest], old[i]
		i = smallest
	}
	*h = (*h)[0 : n-1]
	return val
}
