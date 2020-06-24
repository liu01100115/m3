package arrow

import (
	"testing"
	"time"

	"github.com/apache/arrow/go/arrow/math"
	"github.com/apache/arrow/go/arrow/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestArrowBlockIterator(t *testing.T) {
	start := time.Now().Truncate(time.Hour)
	pool := memory.NewGoAllocator()

	recorder := newDatapointRecorder(pool)
	bl := newBlockIterator(start.UnixNano(),
		int(time.Second*10), int(time.Minute*5))
	it := newArrowBlockIterator(start, time.Minute*2, recorder, bl)

	step := 0
	exSums := []float64{66 /* 0..11 */, 210 /* 12..23 */, 189 /* 24..30 */}
	counts := []int{12, 12, 7}

	exVal := 0.0
	exTime := start.UnixNano()
	for it.next() {
		require.True(t, step < len(exSums))

		rec := it.current()

		vals := rec.values()
		assert.Equal(t, exSums[step], math.Float64.Sum(vals))
		require.Equal(t, counts[step], vals.Len())
		for i := 0; i < counts[step]; i++ {
			assert.Equal(t, exVal, vals.Value(i))
			exVal++
		}

		times := rec.timestamps()
		require.Equal(t, counts[step], times.Len())
		for i := 0; i < counts[step]; i++ {
			assert.Equal(t, exTime, times.Value(i))
			exTime = exTime + int64(time.Second*10)
		}

		step++
	}

	assert.NoError(t, it.close())
}