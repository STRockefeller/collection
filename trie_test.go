package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()

	// Test case 1: Insert multiple words and ensure they are found
	words := []string{"hello", "world", "apple", "banana"}
	for _, word := range words {
		trie.Insert(word)
	}

	for _, word := range words {
		assert.True(t, trie.Search(word), "Expected word '%s' to be found in the trie", word)
	}

	// Test case 2: Search for words that exist and don't exist
	assert.True(t, trie.Search("hello"), "Expected 'hello' to be found in the trie")
	assert.False(t, trie.Search("cat"), "Expected 'cat' not to be found in the trie")

	// Test case 3: Test StartsWith with valid and invalid prefixes
	assert.True(t, trie.StartsWith("he"), "Expected 'he' to be a valid prefix")
	assert.False(t, trie.StartsWith("ca"), "Expected 'ca' not to be a valid prefix")
}
