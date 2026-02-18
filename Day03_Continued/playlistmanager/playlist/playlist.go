package playlist

import (
	"fmt"
	"playlistmanager/songnode"
)

type PlayList struct {
	Name     string
	Duration int

	head *songnode.SongNode
}

func New(name string) *PlayList {
	return &PlayList{
		Name:     name,
		Duration: 0,
		head:     nil,
	}
}

func (pl *PlayList) AddSongAtEnd(song *songnode.SongNode) {
	if pl.head == nil {
		pl.head = song
		pl.Duration = song.Duration
		return
	}

	current := pl.head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = song
	pl.Duration += song.Duration
}

func (pl *PlayList) AddSongAtFront(song *songnode.SongNode) {
	song.Next = pl.head
	pl.head = song

	pl.Duration += song.Duration
}

func (pl *PlayList) DeleteSong(name string) {
	if pl.head == nil {
		return
	}

	if pl.head.Name == name {
		pl.Duration -= pl.head.Duration
		pl.head = pl.head.Next
		return
	}

	current := pl.head
	for current.Next != nil {
		if current.Next.Name == name {
			pl.Duration -= current.Next.Duration
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

func (pl *PlayList) Traverse() {
	current := pl.head
	i := 1
	fmt.Printf("Listing playlist: %s; Total Duration: %d seconds.\n\n", pl.Name, pl.Duration)
	for current != nil {
		fmt.Printf("Song %d:\nName: %s; Artist: %s; Duration: %d seconds\n\n", i, current.Name, current.Artist, current.Duration)

		current = current.Next
		i++
	}
}
