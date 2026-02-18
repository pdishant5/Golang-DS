package window

import (
	"fmt"
	"tabsmanager/tabs"
)

type Window struct {
	head *tabs.Tab
}

func New() *Window {
	return &Window{
		head: nil,
	}
}

func (w *Window) OpenNewTab(tab *tabs.Tab) {
	if w.head == nil {
		w.head = tab
		return
	}
	current := w.head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = tab
	tab.Prev = current
}

func (w *Window) CloseTab(name string) {
	if w.head == nil {
		return
	}

	if w.head.Name == name {
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

func (w *Window) DisplayTabs() {
	current := w.head
	i := 1
	for current != nil {
		fmt.Printf("Tab %d:\nName: %s; URL: %s\n\n", i, current.Name, current.URL)
		current = current.Next
		i++
	}
}
