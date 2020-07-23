package layout

import (
	ui "github.com/gizak/termui/v3"
	"github.com/godeezer/dot/internal/player"
	"github.com/godeezer/lib/deezer"
)

type ModuleShare struct {
	DeezerClient *deezer.Client

	Player      *player.Player
	QueryResult *deezer.SearchResponse
}

func NewModuleShare(client *deezer.Client) *ModuleShare {
	return &ModuleShare{
		DeezerClient: client,
		Player:       player.NewPlayer(client),
	}
}

type Module interface {
	Render()
	Resize(int, int)
}

// interface for module with event handle
type Layout interface {
	Module
	HandleEvent(ui.Event)
}

type LayoutList struct {
	CurrentLayout int
	Layout        []Layout
}

func (self *LayoutList) Next() {
	self.CurrentLayout++
	if self.CurrentLayout >= len(self.Layout) {
		self.CurrentLayout = 0
	}
}

func (self *LayoutList) Render() {
	self.Layout[self.CurrentLayout].Render()
}

func (self *LayoutList) HandleEvent(ev ui.Event) {
	self.Layout[self.CurrentLayout].HandleEvent(ev)
}
