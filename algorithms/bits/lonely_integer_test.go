package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLonelyInteger(t *testing.T) {
	assert.Equal(t, 2, LonelyInteger(0, 0, 1, 1, 2, 3, 3, 4, 4, 5, 5, 100, 100))
	assert.Equal(t, 100, LonelyInteger(0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 100))
	assert.Equal(t, 0, LonelyInteger(0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 100, 100))
	assert.Equal(t, -1, LonelyInteger(1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 100, 100))
}
