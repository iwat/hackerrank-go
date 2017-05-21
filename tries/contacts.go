package tries

type Trie struct {
	value    string
	children map[byte]*Trie
	count    int
	final    bool
}

func NewTrie(value string) *Trie {
	return &Trie{value, make(map[byte]*Trie), 0, false}
}

func (t *Trie) Add(input string) {
	c := input[0]
	if child, ok := t.children[c]; ok {
		child.count++
		child.Add(input[1:])
	} else {
		t.children[c] = NewTrie(t.value + input[0:1])
		t.children[c].count = 1
		if len(input) > 1 {
			t.children[c].Add(input[1:])
		} else {
			t.children[c].final = true
		}
	}
}

func (t *Trie) Find(input string) int {
	if len(input) == 0 {
		return t.count
	}

	if child, ok := t.children[input[0]]; ok {
		return child.Find(input[1:])
	} else {
		return 0
	}
}
