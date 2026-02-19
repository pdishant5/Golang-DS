package trie

import (
	"fmt"
	"tabsmanager/trienode"
)

type Trie struct {
	root *trienode.TrieNode
}

func New() *Trie {
	return &Trie{
		root: trienode.New(),
	}
}

func (t *Trie) Insert(word string) {
	node := t.root

	for _, ch := range word {
		if !node.ContainsKey(ch) {
			node.Put(ch, trienode.New())
		}
		node = node.Get(ch)
	}
	node.SetEnd()
}

func (t *Trie) Search(word string) bool {
	node := t.root

	for _, ch := range word {
		if !node.ContainsKey(ch) {
			fmt.Printf("Word: %s is not in the list!\n", word)
			return false
		}
		node = node.Get(ch)
	}
	return node.IsEnd()
}

func (t *Trie) StartsWith(word string) bool {
	node := t.root

	for _, ch := range word {
		if !node.ContainsKey(ch) {
			fmt.Printf("No word in the list starts with: %s!\n", word)
			return false
		}
		node = node.Get(ch)
	}
	return true
}

func (t *Trie) findPreffixNode(word string) *trienode.TrieNode {
	node := t.root

	for _, ch := range word {
		if !node.ContainsKey(ch) {
			return nil
		}
		node = node.Get(ch)
	}
	return node
}

func (t *Trie) dfs(node *trienode.TrieNode, word string, results *[]string) {
	if node.IsEnd() {
		*results = append(*results, word)
	}

	for ch, child := range node.Children {
		if child != nil {
			t.dfs(child, word+string(ch), results)
		}
	}
}

func (t *Trie) FindAllMatchingWords(word string) []string {
	node := t.findPreffixNode(word)
	if node == nil {
		return []string{}
	}

	var results []string
	t.dfs(node, word, &results)
	return results
}
