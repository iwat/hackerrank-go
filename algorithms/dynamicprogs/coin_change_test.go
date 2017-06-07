package dynamicprogs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoinChange(t *testing.T) {
	assert.Equal(t, 2, CoinChange(2, 1, 2))
	assert.Equal(t, 2, CoinChange(3, 1, 2))

	// 1# 1, 1, 1, 1
	// 2# 1, 1, 2
	// 3# 1, 3
	// 4# 2, 2
	assert.Equal(t, 4, CoinChange(4, 1, 2, 3))

	// 1# 2, 2, 2, 2, 2
	// 2# 2, 2, 3, 3
	// 3# 2, 2, 6
	// 4# 2, 3, 5
	// 5# 5, 5
	assert.Equal(t, 5, CoinChange(10, 2, 5, 3, 6))
}
