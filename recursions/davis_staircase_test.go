package recursions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDavisStaircase(t *testing.T) {
	assert.Equal(t, 1, DavisStaircase(1))
	assert.Equal(t, 4, DavisStaircase(3))
	assert.Equal(t, 44, DavisStaircase(7))

	assert.Equal(t, 2082876103, DavisStaircase(36))
}
