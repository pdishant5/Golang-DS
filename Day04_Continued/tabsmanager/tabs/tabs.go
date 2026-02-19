package tabs

type Tab struct {
	URL  string
	Name string
	// Logo []byte

	Next *Tab
	Prev *Tab
}

func New(url, name string) *Tab {
	return &Tab{
		URL:  url,
		Name: name,
		Next: nil,
		Prev: nil,
	}
}
