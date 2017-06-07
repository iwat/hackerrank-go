package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountAnagram(t *testing.T) {
	assert.Equal(t, 4, CountAnagram("abc", "cde"))
	assert.Equal(t, 0, CountAnagram("abc", "cba"))
	assert.Equal(t, 5, CountAnagram("abbcccdddd", "dddccbaaa"))
}
