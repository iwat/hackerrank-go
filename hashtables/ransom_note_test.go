package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRansomNote(t *testing.T) {
	assert.Equal(t, true, RansomNote(
		"give me one grand today night",
		"give one grand today",
	))
}
