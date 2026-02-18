package playlist

import (
	"fmt"
	"playlistmanager/songnode"
)

type PlayList struct {
	Name       string
	Duration   int
	TotalSongs int

	head *songnode.SongNode
}

func New(name string) *PlayList {
	return &PlayList{
		Name:       name,
		Duration:   0,
		TotalSongs: 0,
		head:       nil,
	}
}

func (pl *PlayList) AddSongAtEnd(song *songnode.SongNode) {
	if pl.head == nil {
		pl.head = song
		pl.Duration = song.Duration
		pl.TotalSongs++
		return
	}

	current := pl.head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = song
	pl.Duration += song.Duration
	pl.TotalSongs++
}

func (pl *PlayList) AddSongAtFront(song *songnode.SongNode) {
	song.Next = pl.head
	pl.head = song

	pl.Duration += song.Duration
	pl.TotalSongs++
}

func (pl *PlayList) DeleteSong(name string) {
	if pl.head == nil {
		return
	}

	if pl.head.Name == name {
		pl.Duration -= pl.head.Duration
		pl.head = pl.head.Next
		pl.TotalSongs--
		return
	}

	current := pl.head
	for current.Next != nil {
		if current.Next.Name == name {
			pl.Duration -= current.Next.Duration
			pl.TotalSongs--
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

func (pl *PlayList) Traverse() {
	current := pl.head
	i := 1
	fmt.Printf("Listing playlist: %s; Total Songs: %d; Total Duration: %d seconds.\n\n", pl.Name, pl.TotalSongs, pl.Duration)
	for current != nil {
		fmt.Printf("Song %d:\nName: %s; Artist: %s; Duration: %d seconds\n\n", i, current.Name, current.Artist, current.Duration)

		current = current.Next
		i++
	}
}

// merges two playlists by appending the longer one at the end of the shorter one in termas of number pf songs
func MergePlaylists1(pl1, pl2 *PlayList) *PlayList {
	mergedPlaylist := New(pl1.Name + " + " + pl2.Name)
	mergedPlaylist.Duration = pl1.Duration + pl2.Duration
	mergedPlaylist.TotalSongs = pl1.TotalSongs + pl2.TotalSongs

	// if playlist 1 is shorter
	if pl1.TotalSongs <= pl2.TotalSongs {
		head1 := pl1.head
		mergedPlaylist.head = pl1.head

		for head1.Next != nil {
			head1 = head1.Next
		}

		head1.Next = pl2.head
		return mergedPlaylist
	}

	// if playlist 2 is shorter
	head2 := pl2.head
	mergedPlaylist.head = pl2.head

	for head2.Next != nil {
		head2 = head2.Next
	}

	head2.Next = pl1.head
	return mergedPlaylist
}

func MergePlaylists2(pl1, pl2 *PlayList) *PlayList {
	mergedPlaylist := New(pl1.Name + " + " + pl2.Name)
	mergedPlaylist.head = pl1.head
	mergedPlaylist.Duration = pl1.Duration + pl2.Duration
	mergedPlaylist.TotalSongs = pl1.TotalSongs + pl2.TotalSongs

	head1 := pl1.head
	head2 := pl2.head

	for head1.Next != nil && head2.Next != nil {
		headNext := head2.Next
		head2.Next = head1.Next
		head1.Next = head2

		head1 = head1.Next.Next
		head2 = headNext
	}

	// if playlist 1 is exhausted
	if head1.Next == nil {
		head1.Next = head2
		return mergedPlaylist
	}

	// if playlist 2 is exhausted
	head2.Next = head1.Next
	head1.Next = head2
	return mergedPlaylist

}
