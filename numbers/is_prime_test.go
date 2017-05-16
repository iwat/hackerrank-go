package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrime(t *testing.T) {
	assert.False(t, IsPrime(1))
	assert.False(t, IsPrime(4))
	assert.False(t, IsPrime(9))
	assert.False(t, IsPrime(16))
	assert.False(t, IsPrime(25))
	assert.False(t, IsPrime(36))
	assert.False(t, IsPrime(49))
	assert.False(t, IsPrime(64))
	assert.False(t, IsPrime(81))
	assert.False(t, IsPrime(100))
	assert.False(t, IsPrime(121))
	assert.False(t, IsPrime(144))
	assert.False(t, IsPrime(169))
	assert.False(t, IsPrime(196))
	assert.False(t, IsPrime(225))
	assert.False(t, IsPrime(256))
	assert.False(t, IsPrime(289))
	assert.False(t, IsPrime(324))
	assert.False(t, IsPrime(361))
	assert.False(t, IsPrime(400))
	assert.False(t, IsPrime(441))
	assert.False(t, IsPrime(484))
	assert.False(t, IsPrime(529))
	assert.False(t, IsPrime(576))
	assert.False(t, IsPrime(625))
	assert.False(t, IsPrime(676))
	assert.False(t, IsPrime(729))
	assert.False(t, IsPrime(784))
	assert.False(t, IsPrime(841))
}
