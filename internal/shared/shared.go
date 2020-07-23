package shared

import (
	"github.com/godeezer/dot/internal/player"
	"github.com/godeezer/lib/deezer"
)

type ModuleShare struct {
	DeezerClient *deezer.Client
	Player       *player.Player

	QueryResult   *deezer.SearchResponse
	MusicQueue    []deezer.Song
	MusicCurrent  deezer.Song
	MusicProgress float64
}

func NewModuleShare(client *deezer.Client) *ModuleShare {
	return &ModuleShare{
		DeezerClient: client,
		Player:       player.NewPlayer(client),
	}
}
