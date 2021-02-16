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

var emptySample Sample

// Sample represents a sampled value.
type Sample struct {
	value    float64 // sampled value
	numRanks int64   // number of ranks represented
	delta    int64   // delta between min rank and max rank
	prev     *Sample // previous sample
	next     *Sample // next sample
}

// newSample creates a new sample.
func newSample() *Sample {
	return &Sample{}
}

// reset resets a sample.
func (s *Sample) reset() {
	*s = emptySample
}

// setData sets sample data.
// nolint: unparam
func (s *Sample) setData(value float64, numRanks int64, delta int64) {
	s.value = value
	s.numRanks = numRanks
	s.delta = delta
}
