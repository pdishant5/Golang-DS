package window

import (
	"fmt"
	"tabsmanager/tabs"
	"tabsmanager/trie"
	"tabsmanager/trienode"
)

type Window struct {
	head *tabs.Tab
	Trie *trie.Trie
}

func New() *Window {
	return &Window{
		head: nil,
		Trie: trie.New(),
	}
}

func (w *Window) OpenNewTab(tab *tabs.Tab) {
	if w.head == nil {
		w.head = tab
		w.Trie.Insert(tab.Name)
		return
	}
	current := w.head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = tab
	tab.Prev = current
	w.Trie.Insert(tab.Name)
}

func (w *Window) CloseTab(name string) {
	if w.head == nil {
		return
	}

	if w.head.Name == name {
		w.Trie.Erase(name)
		w.head = w.head.Next
		if w.head != nil {
			w.head.Prev = nil
		}
		return
	}

	current := w.head
	for current != nil {
		if current.Name == name {
			// if current.Prev != nil { // only head can have prev as nil - already checked above
			w.Trie.Erase(name)
			current.Prev.Next = current.Next
			// }
			if current.Next != nil {
				current.Next.Prev = current.Prev
			}
			current.Next = nil
			current.Prev = nil
			return
		}
		current = current.Next
	}
}

func (w *Window) CloseAllTabsToRight(name string) {
	if w.head == nil {
		return
	}

	current := w.head
	for current != nil {
		if current.Name == name {
			node := current.Next
			for node != nil {
				w.Trie.Erase(node.Name)
				node = node.Next
			}

			if current.Next != nil {
				current.Next.Prev = nil
			}
			current.Next = nil
			return
		}
		current = current.Next
	}
}

func (w *Window) CloseAllTabs() {
	w.Trie.Root.Children = map[rune]*trienode.TrieNode{}
	w.Trie.Root.EndWith = 0
	w.Trie.Root.CountPreffix = 0
	w.head = nil
}

func (w *Window) DisplayTabs() {
	current := w.head
	i := 1
	for current != nil {
		fmt.Printf("Tab %d:\nName: %s; URL: %s\n\n", i, current.Name, current.URL)
		current = current.Next
		i++
	}
	if i == 1 {
		fmt.Println("It is an empty window! There are no tabs to display!")
	}
}

func (w *Window) SearchTab(url string) {
	current := w.head
	for current != nil {
		if current.URL == url {
			fmt.Printf("Found the Tab %s: %s\n", current.Name, current.URL)

			return
		}
		current = current.Next
	}

	fmt.Printf("Tab %s doesn't exist!\n", url)
}
