package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceToBytes(t *testing.T) {
	// int slice
	s1 := []int{1, 2, 3}
	b1 := SliceToBytes(s1)
	assert.Equal(t, "1,2,3", string(b1))

	// int32 slice
	s2 := []int32{1, 2, 3}
	b2 := SliceToBytes(s2)
	assert.Equal(t, "1,2,3", string(b2))

	// int64 slice
	s3 := []int64{1, 2, 3}
	b3 := SliceToBytes(s3)
	assert.Equal(t, "1,2,3", string(b3))

	// uint slice
	s4 := []uint{1, 2, 3}
	b4 := SliceToBytes(s4)
	assert.Equal(t, "1,2,3", string(b4))

	// uint32 slice
	s5 := []uint32{1, 2, 3}
	b5 := SliceToBytes(s5)
	assert.Equal(t, "1,2,3", string(b5))

	// uint64 slice
	s6 := []uint64{1, 2, 3}
	b6 := SliceToBytes(s6)
	assert.Equal(t, "1,2,3", string(b6))

	// string slice
	s7 := []string{"a", "b", "c"}
	b7 := SliceToBytes(s7)
	assert.Equal(t, "'a','b','c'", string(b7))

	// string slice
	s8 := []float64{1.1, 2.2, 3.3}
	b8 := SliceToBytes(s8)
	assert.Equal(t, "", string(b8))
}
