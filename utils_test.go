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

	assert.True(t, *StrPtr(s) == s)
	assert.True(t, *BoolPtr(b) == b)
	assert.True(t, *IntPtr(i) == i)
	assert.True(t, *Uint8Ptr(ui8) == ui8)
	assert.True(t, *Uint16Ptr(ui16) == ui16)
	assert.True(t, *Float32Ptr(f32) == f32)
	assert.True(t, *Float64Ptr(f64) == f64)
}
