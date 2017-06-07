package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLify(t *testing.T) {
	buffer := []byte("Mr John Smith    ")
	URLify(buffer, 13)
	assert.Equal(t, "Mr%20John%20Smith", string(buffer))

	buffer = []byte("Mr John Smith         ")
	URLify(buffer, 13)
	assert.Equal(t, "Mr%20John%20Smith     ", string(buffer))
}
