package main

import (
	"fmt"
	"time"

	"playlistmanager/playlist"
	"playlistmanager/songnode"
)

func main() {
	playlist1 := playlist.New("My Playlist")
	newSong := songnode.New("Shiva Tandava Stotram", "Shankar Mahadevan", 300)
	playlist1.AddSongAtFront(newSong)
	// playlist1.Traverse()
	playlist1.AddSongAtFront(songnode.New("Nagar Nandji Na Lal", "Aditya Gadhvi", 200))
	// playlist1.Traverse()

	playlist1.AddSongAtEnd(songnode.New("Maa Tujhe Salaam", "A.R. Rahman", 240))
	// playlist1.Traverse()
	playlist1.AddSongAtEnd(songnode.New("Goti Lo", "Aditya Gadhvi", 180))
	// playlist1.Traverse()
	playlist1.AddSongAtEnd(songnode.New("Krishn Bhagwan Halya", "Aditya Gadhvi", 210))
	// playlist1.Traverse()

	playlist1.DeleteSong("Goti Lo")
	playlist1.Traverse()

	playlist2 := playlist.New("Your Playlist")
	newSong = songnode.New("Bhajan 1", "Artist 1", 200)
	playlist2.AddSongAtFront(newSong)
	// playlist2.Traverse()
	playlist2.AddSongAtFront(songnode.New("Bhajan 2", "Artist 2", 250))
	// playlist2.Traverse()

	playlist2.AddSongAtEnd(songnode.New("Bhajan 3", "Artist 1", 240))
	// playlist2.Traverse()
	playlist2.AddSongAtFront(songnode.New("Bhajan 4", "Artist 3", 280))
	// playlist2.Traverse()
	playlist2.AddSongAtEnd(songnode.New("Bhajan 5", "Artist 3", 310))
	playlist2.AddSongAtFront(songnode.New("Bhajan 6", "Artist 1", 230))
	playlist2.AddSongAtEnd(songnode.New("Bhajan 7", "Artist 3", 210))
	playlist2.Traverse()

	// playlist2.DeleteSong("Goti Lo")
	// playlist2.Traverse() // same list as the previous, because we deleted the song that was not in the list

	// merging two playlists by appending one list at the tail of the other
	t1 := time.Now()
	mergedPlaylist1 := playlist.MergePlaylists1(playlist1, playlist2)
	fmt.Println("Time taken in method 1:", time.Since(t1))
	mergedPlaylist1.Traverse()

	// niling bothe the playlists to ensure that they will not be used anymore and the merged playlist is a new list
	playlist1, playlist2 = nil, nil

	// merging two playlists by alternating the songs from both the lists
	// mergedPlaylist2 := playlist.MergePlaylists2(playlist1, playlist2)
	// t1 := time.Now()
	// mergedPlaylist2 := playlist.MergePlaylists2(playlist2, playlist1)
	// fmt.Println("Time taken in method 2:", time.Since(t1))
	// mergedPlaylist2.Traverse()
	// playlist1, playlist2 = nil, nil
}
