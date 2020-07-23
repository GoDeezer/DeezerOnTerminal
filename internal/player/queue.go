package player

import (
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

func (self *PlayerQueue) AddSong(songs ...deezer.Song) error {
	self.Queue = append(self.Queue, songs...)
	return nil
}

func (self *PlayerQueue) AddAlbum(albums ...deezer.Album) error {
	for _, a := range albums {
		songs, err := self.DeezerClient.SongsByAlbum(a.ID, -1)
		if err != nil {
			return err
		}
		err = self.AddSong(songs...)
		if err != nil {
			return err
		}
	}
	return nil
}

func (self *PlayerQueue) AddArtist(artists ...deezer.Artist) error {
	for _, a := range artists {
		albums, err := self.DeezerClient.AlbumsByArtist(a.ID)
		if err != nil {
			return err
		}
		err = self.AddAlbum(albums...)
		if err != nil {
			return err
		}
	}
	return nil
}

func (self *PlayerQueue) Delete(index int) {
	if index >= len(self.Queue) || len(self.Queue)-1 < 0 {
		return
	}
	self.Queue = append(self.Queue[:index], self.Queue[index+1:]...)
}
