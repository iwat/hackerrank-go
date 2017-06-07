package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckPermutation(t *testing.T) {
	assert.True(t, CheckPermutation("abcdefg", "gfedcba"))
	assert.True(t, CheckPermutation("abcdefg", "gfedcba"))
	assert.False(t, CheckPermutation("abcdefg", "gfedcb"))
	assert.True(t, CheckPermutation("abcdefgabcdefg", "gfedcbagfedcba"))
	assert.False(t, CheckPermutation("abcdefgabcdefg", "gfedcb"))
}
