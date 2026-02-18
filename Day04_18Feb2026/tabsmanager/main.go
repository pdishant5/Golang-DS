package main

import (
	"tabsmanager/tabs"
	"tabsmanager/window"
)

func main() {
	window1 := window.New()
	window1.OpenNewTab(tabs.New("https://example1.com", "Example1"))
	window1.DisplayTabs()

	window1.OpenNewTab(tabs.New("https://example.com", "Example"))
	window1.OpenNewTab(tabs.New("https://example.com", "Example"))
	window1.DisplayTabs()

	window1.CloseTab("Example")
	window1.DisplayTabs()
	window1.CloseTab("Example1")
	window1.DisplayTabs()
	window1.CloseTab("Example")
	window1.DisplayTabs()
}
