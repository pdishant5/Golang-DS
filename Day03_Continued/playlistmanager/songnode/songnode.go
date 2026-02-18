package songnode

type SongNode struct {
	Name     string
	Artist   string // may also be a list of names in case of multiple artists
	Duration int

	Next *SongNode
}

func New(name, artist string, duration int) *SongNode {
	return &SongNode{
		Name:     name,
		Artist:   artist,
		Duration: duration,
		Next:     nil,
	}
}
