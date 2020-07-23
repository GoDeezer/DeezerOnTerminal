package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/godeezer/dot/internal/layout"
	deezer "github.com/godeezer/lib/deezer"
)

const (
	TabSearch = iota
	TabQueue
	TabLast
)

type App struct {
	Stop  bool
	Share *layout.ModuleShare

	Layout *layout.LayoutList
}

func NewApp() (*App, error) {
	client, err := deezer.NewClient("")
	if err != nil {
		return nil, err
	}
	shared := layout.NewModuleShare(client)

	// Modules
	playing := layout.NewPlaying(shared)

	app := &App{
		Stop:  false,
		Share: shared,

		Layout: &layout.LayoutList{
			CurrentLayout: 0,
			Layout: []layout.Layout{
				// Event Modules
				layout.NewSearch(shared, playing),
				layout.NewQueue(shared, playing),
				layout.NewDetail(shared, playing),
			},
		},
	}

	app.HandleResize()
	app.Render()
	return app, nil
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
			self.Layout.Next()
			self.Render()
			return
		}
	case ui.ResizeEvent:
		self.HandleResize()
		return
	default:
	}

	// tab specific
	self.Layout.HandleEvent(ev)
	self.Render()
}

func (self *App) HandleResize() {
	// Setting layout here
	cols, rows := ui.TerminalDimensions()
	for _, m := range self.Layout.Layout {
		m.Resize(cols, rows)
	}
}

func (self *App) Run() error {
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

	app, err := NewApp()
	if err != nil {
		log.Fatalf("failed to create new application: %v", err)
	}

	if err := app.Run(); err != nil {
		log.Fatalf("failed to run application: %v", err)
	}
}
