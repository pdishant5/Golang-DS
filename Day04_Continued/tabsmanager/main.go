package main

import (
	"fmt"
	"tabsmanager/tabs"
	"tabsmanager/window"
)

func main() {
	window1 := window.New()
	window1.OpenNewTab(tabs.New("https://example1.com", "Example1"))
	// window1.DisplayTabs()

	window1.OpenNewTab(tabs.New("https://example.com", "Example"))
	window1.OpenNewTab(tabs.New("https://example.com", "Example"))
	window1.OpenNewTab(tabs.New("https://example2.com", "Example2"))
	window1.OpenNewTab(tabs.New("https://example4.com", "Example4"))

	window1.OpenNewTab(tabs.New("https://example3.com", "Example3"))
	window1.OpenNewTab(tabs.New("https://example4.com", "Example4"))
	// window1.DisplayTabs()

	window1.CloseTab("Example")
	window1.DisplayTabs()
	// window1.CloseTab("Example1")
	// window1.DisplayTabs()
	// window1.CloseTab("Example")
	// window1.DisplayTabs()

	tabName := "Exam"
	fmt.Printf("Does any tabs start with %s? - %v\n", tabName, window1.Trie.StartsWith(tabName))

	tabName = "Exa"
	results := window1.Trie.FindAllMatchingWords(tabName)
	count := window1.Trie.CountWordSatrtingWith(tabName)
	// for _, result := range results {
	// 	fmt.Println(result)
	// }
	// fmt.Printf("All matching tabs with %s: %v\n", tabName, results)
	fmt.Printf("Total %d tabs matches with %s: %v\n", count, tabName, results)

	// adding two new tabs with same name
	window1.OpenNewTab(tabs.New("https://dishant.com", "Dishant"))
	window1.OpenNewTab(tabs.New("https://dishant.com", "Dishant"))
	window1.DisplayTabs()

	tabName = "Dishant"
	results = window1.Trie.FindAllMatchingWords(tabName)
	count = window1.Trie.CountWordSatrtingWith(tabName)
	fmt.Printf("Total %d tabs matches with %s: %v\n", count, tabName, results)

	window1.CloseTab("Dishant")
	window1.DisplayTabs()
	results = window1.Trie.FindAllMatchingWords(tabName)
	count = window1.Trie.CountWordSatrtingWith(tabName)
	fmt.Printf("Total %d tabs matches with %s: %v\n", count, tabName, results)

	window1.SearchTab("https://example.com")
	window1.SearchTab("https://example5.com")

	window1.CloseAllTabsToRight("Example4")
	window1.DisplayTabs()
	results = window1.Trie.FindAllMatchingWords(tabName)
	count = window1.Trie.CountWordSatrtingWith(tabName)
	fmt.Printf("Total %d tabs matches with %s: %v\n", count, tabName, results)

	window1.CloseAllTabs()
	results = window1.Trie.FindAllMatchingWords(tabName)
	count = window1.Trie.CountWordSatrtingWith(tabName)
	fmt.Printf("Total %d tabs matches with %s: %v\n", count, tabName, results)
	// window1.DisplayTabs()
}
