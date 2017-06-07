package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsUnique(t *testing.T) {
	assert.False(t, IsUnique("the quick brown fox jumps over the lazy dog"), "Spaces is not unique")
	assert.True(t, IsUnique("thequickbrownfxjmpsvlazydg"), "It must be unique")
}
