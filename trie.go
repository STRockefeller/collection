package collection

type TrieNode struct {
	Value       rune
	Children    map[rune]*TrieNode
	IsEndOfWord bool
}

type Trie struct {
	Root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		Root: &TrieNode{
			Children: make(map[rune]*TrieNode),
		},
	}
}

func (t *Trie) Insert(word string) {
	node := t.Root
	for _, char := range word {
		if _, ok := node.Children[char]; !ok {
			node.Children[char] = &TrieNode{
				Value:    char,
				Children: make(map[rune]*TrieNode),
			}
		}
		node = node.Children[char]
	}
	node.IsEndOfWord = true
}

func (t *Trie) Search(word string) bool {
	node := t.Root
	for _, char := range word {
		if _, ok := node.Children[char]; !ok {
			return false
		}
		node = node.Children[char]
	}
	return node.IsEndOfWord
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.Root
	for _, char := range prefix {
		if _, ok := node.Children[char]; !ok {
			return false
		}
		node = node.Children[char]
	}
	return true
}
