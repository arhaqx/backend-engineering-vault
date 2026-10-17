package trie

type Node struct {
	children map[rune]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func New() *Trie {
	return &Trie{root: &Node{children: make(map[rune]*Node)}}
}

func (t *Trie) Insert(word string) {
	curr := t.root
	for _, ch := range word {
		if _, ok := curr.children[ch]; !ok {
			curr.children[ch] = &Node{children: make(map[rune]*Node)}
		}
		curr = curr.children[ch]
	}
	curr.isEnd = true
}

func (t *Trie) Search(word string) bool {
	curr := t.root
	for _, ch := range word {
		if next, ok := curr.children[ch]; ok {
			curr = next
		} else {
			return false
		}
	}
	return curr.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	curr := t.root
	for _, ch := range prefix {
		if next, ok := curr.children[ch]; ok {
			curr = next
		} else {
			return false
		}
	}
	return true
}
