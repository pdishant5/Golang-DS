package trienode

type TrieNode struct {
	Children map[rune]*TrieNode
	End      bool
}

func New() *TrieNode {
	return &TrieNode{
		Children: make(map[rune]*TrieNode),
		End:      false,
	}
}

func (tn *TrieNode) ContainsKey(ch rune) bool {
	return tn.Children[ch] != nil
}

func (tn *TrieNode) Put(ch rune, node *TrieNode) {
	tn.Children[ch] = node
}

func (tn *TrieNode) Get(ch rune) *TrieNode {
	return tn.Children[ch]
}

func (tn *TrieNode) SetEnd() {
	tn.End = true
}

func (tn *TrieNode) IsEnd() bool {
	return tn.End
}
