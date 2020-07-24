package player

import (
	"sync"

	"github.com/godeezer/lib/deezer"
)

type PlayerQueue struct {
	DeezerClient *deezer.Client

	Queue []deezer.Song
}

func NewPlayerQueue(client *deezer.Client) *PlayerQueue {
	return &PlayerQueue{
		DeezerClient: client,
		Queue:        []deezer.Song{},
	}
}

func (self *PlayerQueue) IsInQueue(song deezer.Song) bool {
	for _, s := range self.Queue {
		if s.ID == song.ID {
			return true
		}
	}
	return false
}

func (self *PlayerQueue) AddSong(songs ...deezer.Song) {
	for _, song := range songs {
		if self.IsInQueue(song) {
			continue
		}
		self.Queue = append(self.Queue, songs...)
	}
}

func (self *PlayerQueue) AddAlbum(albums ...deezer.Album) {
	var wg sync.WaitGroup
	for _, a := range albums {
		wg.Add(1)
		go func() {
			songs, err := self.DeezerClient.SongsByAlbum(a.ID, -1)
			if err == nil {
				self.AddSong(songs...)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func (self *PlayerQueue) AddArtist(artists ...deezer.Artist) {
	var wg sync.WaitGroup
	for _, a := range artists {
		wg.Add(1)
		go func() {
			albums, err := self.DeezerClient.AlbumsByArtist(a.ID)
			if err == nil {
				self.AddAlbum(albums...)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func (self *PlayerQueue) Delete(index int) {
	if index >= len(self.Queue) || len(self.Queue)-1 < 0 {
		return
	}
	self.Queue = append(self.Queue[:index], self.Queue[index+1:]...)
}
