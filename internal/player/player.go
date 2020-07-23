package player

import (
	"github.com/godeezer/lib/deezer"
)

type Player struct {
	DeezerClient *deezer.Client

	PlayerQueue *PlayerQueue
	CurrentSong deezer.Song
	Progress    float64
}

func NewPlayer(client *deezer.Client) *Player {
	return &Player{
		DeezerClient: client,
		PlayerQueue:  NewPlayerQueue(client),
	}
}

func (self *Player) SetCurrentSong(song deezer.Song) {
	self.CurrentSong = song
	self.Progress = 0.0
}
