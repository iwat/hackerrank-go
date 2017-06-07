package tries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	trie := NewTrie("/")
	trie.Add("hack")
	trie.Add("hackerrank")
	trie.Add("hackerrank.com")
	assert.Equal(t, 3, trie.Find("hac"))
	assert.Equal(t, 0, trie.Find("hak"))
	assert.Equal(t, 3, trie.Find("hack"))
	assert.Equal(t, 2, trie.Find("hackerrank"))
	assert.Equal(t, 1, trie.Find("hackerrank."))
}
