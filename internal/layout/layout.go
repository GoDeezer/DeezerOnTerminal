package layout

import (
	ui "github.com/gizak/termui/v3"
	"github.com/godeezer/dot/internal/player"
	"github.com/godeezer/lib/deezer"
)

type ModuleShare struct {
	DeezerClient *deezer.Client

	Cols        int
	Rows        int
	Popup       Layout
	Player      *player.Player
	QueryResult *deezer.SearchResponse
}

func NewModuleShare(client *deezer.Client) *ModuleShare {
	return &ModuleShare{
		DeezerClient: client,
		Popup:        nil,
		Player:       player.NewPlayer(client),
	}
}

type Module interface {
	Render()
	Resize(int, int)
}

// interface for module with event handle
type Layout interface {
	Render()
	Resize(int, int)
	HandleEvent(ui.Event)
}

type LayoutList struct {
	Share         *ModuleShare
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
	if self.Share.Popup != nil {
		self.Share.Popup.Render()
	} else {
		self.Layout[self.CurrentLayout].Render()
	}
}

func (self *LayoutList) HandleEvent(ev ui.Event) {
	if self.Share.Popup != nil {
		self.Share.Popup.HandleEvent(ev)
		return
	}
	self.Layout[self.CurrentLayout].HandleEvent(ev)
}
