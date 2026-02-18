package main

import (
	"playlistmanager/playlist"
	"playlistmanager/songnode"
)

func main() {
	playlist := playlist.New("My Playlist")
	newSong := songnode.New("Shiva Tandava Stotram", "Shankar Mahadevan", 300)
	playlist.AddSongAtFront(newSong)
	playlist.Traverse()
	playlist.AddSongAtFront(songnode.New("Nagar Nandji Na Lal", "Aditya Gadhvi", 200))
	playlist.Traverse()

	playlist.AddSongAtEnd(songnode.New("Maa Tujhe Salaam", "A.R. Rahman", 240))
	// playlist.Traverse()
	playlist.AddSongAtEnd(songnode.New("Goti Lo", "Aditya Gadhvi", 180))
	// playlist.Traverse()
	playlist.AddSongAtEnd(songnode.New("Krishn Bhagwan Halya", "Aditya Gadhvi", 210))
	playlist.Traverse()

	playlist.DeleteSong("Goti Lo")
	playlist.Traverse()
}
