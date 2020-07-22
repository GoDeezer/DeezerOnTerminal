package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	internel "github.com/godeezer/dot/internal"
	deezer "github.com/godeezer/lib/deezer"
)

const (
	SearchTrack = iota
	SearchAlbum
	SearchArtist
	SearchPlaylist
)

const (
	TabSearch = iota
	TabQueue
	TabLast
)

// interface for display module
type Module interface {
	Render()
	Resize(int, int)
}

type App struct {
	Stop  bool
	Share *internal.ModuleShare

	Layout *internal.LayoutList
}

func NewApp() *App {
	shared := &ModuleShare{}

	// Modules
	playing := internel.NewPlaying(shared)

	app := &App{
		Stop:  false,
		Share: shared,

		Layout: &LayoutList{
			CurrentLayout: 0,
			Layout: []EventModule{
				// Event Modules
				internel.NewSearch(shared, playing),
				internel.NewQueue(shared, playing),
				internel.NewCurrentSong(shared, playing),
			},
		},
	}

	app.HandleResize()
	app.Render()
	return app
}

func (self *App) Render() {
	ui.Clear()
	self.Layout.Render()
}

func (self *App) HandleEvent(ev ui.Event) {
	// across tabs
	switch ev.Type {
	case ui.KeyboardEvent:
		switch ev.ID {
		case "<C-c>":
			self.Stop = true
			return
		case "<Tab>":
			self.Tab.NextTab()
			self.Render()
			return
		}
	case ui.ResizeEvent:
		self.HandleResize()
		return
	default:
	}

	// tab specific
	self.Tab.Tabs[self.Tab.CurrentTab].HandleEvent(ev)
	self.Render()
}

func (self *App) HandleResize() {
	// Setting layout here
	cols, rows := ui.TerminalDimensions()
	for _, m := range self.Tab.Tabs {
		m.Resize(cols, rows)
	}
}

func (self *App) Run() error {
	c, err := deezer.NewClient("7579cd89a4d2ab3d6dc2b418446e35c7bd11ba7e62b11d7a2034d888b73864f16a7bc3088a5087a00f53d079eefce6821b0a5e2f746bd9ca8161789a4da11ff7ece21cfbbf692eb7e749c256b1df5bfd4be1e0b1bbc8a441b769d51daea39212")
	if err != nil {
		return err
	}
	self.Share.DeezerClient = c

	ev := ui.PollEvents()
	for !self.Stop {
		e := <-ev
		self.HandleEvent(e)
		self.Render()
	}
	return nil
}

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	if err := NewApp().Run(); err != nil {
		panic(err)
	}
}
