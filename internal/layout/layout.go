package layout

import (
	ui "github.com/gizak/termui/v3"
	"github.com/godeezer/dot/internal/player"
	"github.com/godeezer/lib/deezer"
)

type ModuleShare struct {
	DeezerClient *deezer.Client
	ReDraw       chan struct{}

	Cols        int
	Rows        int
	Popup       Layout
	Player      *player.Player
	QueryResult *deezer.SearchResponse
}

func NewModuleShare(client *deezer.Client) *ModuleShare {
	redraw := make(chan struct{}, 10)
	return &ModuleShare{
		DeezerClient: client,
		ReDraw:       redraw,
		Popup:        nil,
		Player:       player.NewPlayer(client, redraw),
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
