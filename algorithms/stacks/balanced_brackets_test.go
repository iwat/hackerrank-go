package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBalanceBrackets(t *testing.T) {
	assert.True(t, BalancedBrackets("{[()]}"))
	assert.False(t, BalancedBrackets("{[(])}"))
	assert.True(t, BalancedBrackets("{{[[(())]]}}"))
	assert.False(t, BalancedBrackets("{{[[(())]]}}{{{"))
}
