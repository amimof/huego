package huego

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointers(t *testing.T) {
	s := "test"
	b := true
	i := int(123)
	ui8 := uint8(123)
	ui16 := uint16(123)
	f32 := float32(1.234)
	f64 := float64(1.234)

	assert.True(t, *strPtr(s) == s)
	assert.True(t, *boolPtr(b) == b)
	assert.True(t, *intPtr(i) == i)
	assert.True(t, *uint8Ptr(ui8) == ui8)
	assert.True(t, *uint16Ptr(ui16) == ui16)
	assert.True(t, *float32Ptr(f32) == f32)
	assert.True(t, *float64Ptr(f64) == f64)
}
