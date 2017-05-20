package tries

type Trie struct {
	value    string
	children map[byte]*Trie
	final    bool
}

func NewTrie(value string) *Trie {
	return &Trie{value, make(map[byte]*Trie), false}
}

func (t *Trie) Add(input string) {
	c := input[0]
	if child, ok := t.children[c]; ok {
		child.Add(input[1:])
	} else {
		t.children[c] = NewTrie(t.value + input[0:1])
		if len(input) > 1 {
			t.children[c].Add(input[1:])
		} else {
			t.children[c].final = true
		}
	}
}

func (t *Trie) Find(input string) int {
	if len(input) == 0 {
		return t.Count()
	}

	if child, ok := t.children[input[0]]; ok {
		return child.Find(input[1:])
	} else {
		return 0
	}
}

func (t *Trie) Count() int {
	count := 0
	for _, child := range t.children {
		count += child.Count()
	}

	if t.final {
		count++
	}

	return count
}
