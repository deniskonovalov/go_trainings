package documentstore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsValidNumber(t *testing.T) {
	assert.True(t, IsValidNumber(123))
	assert.True(t, IsValidNumber(int8(8)))
	assert.True(t, IsValidNumber(int16(16)))
	assert.True(t, IsValidNumber(int32(32)))
	assert.True(t, IsValidNumber(int64(64)))
	assert.True(t, IsValidNumber(float32(32.5)))

	assert.False(t, IsValidNumber("string"))
	assert.False(t, IsValidNumber(true))
	assert.False(t, IsValidNumber([]int{1, 2, 3}))
	assert.False(t, IsValidNumber(map[string]int{"Key": 1}))
	assert.False(t, IsValidNumber(nil))
}

func Test_IsValidMap(t *testing.T) {
	assert.True(t, IsValidMap(map[string]int{"Key": 1}))
	assert.True(t, IsValidMap(map[string]int32{"Key": 1}))
	assert.True(t, IsValidMap(map[string]int64{"Key": 64}))
	assert.True(t, IsValidMap(map[string]float64{"Key": 3.14}))
	assert.True(t, IsValidMap(map[string]string{"Key": "value"}))
	assert.True(t, IsValidMap(map[string]bool{"Key": true}))
	assert.True(t, IsValidMap(map[string]any{"Key": 1}))
	assert.True(t, IsValidMap(map[string]byte{"Key": 255}))

	assert.False(t, IsValidMap(nil))
	assert.False(t, IsValidMap("string"))
	assert.False(t, IsValidMap([]int{1, 2, 3}))
	assert.False(t, IsValidMap(map[int]string{1: "value"}))
	assert.False(t, IsValidMap(map[any]string{"Key": "value"}))
}

func Test_IsValidSlice(t *testing.T) {
	assert.True(t, IsValidSlice([]int{1, 2, 3}))
	assert.True(t, IsValidSlice([]int64{64}))
	assert.True(t, IsValidSlice([]float64{3.14, 2.71}))
	assert.True(t, IsValidSlice([]string{"value1", "value2"}))
	assert.True(t, IsValidSlice([]bool{true, false}))
	assert.True(t, IsValidSlice([]byte{255, 128, 64}))
	assert.True(t, IsValidSlice([]rune{'a', 'b', 'c'}))
	assert.True(t, IsValidSlice([]any{"value", 1, true}))

	assert.False(t, IsValidSlice(nil))
	assert.False(t, IsValidSlice("string"))
	assert.False(t, IsValidSlice(map[string]int{"Key": 1}))
}
