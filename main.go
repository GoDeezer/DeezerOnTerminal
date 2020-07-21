package main

import (
	"fmt"
	"log"

	ui "github.com/gizak/termui/v3"
	widgets "github.com/gizak/termui/v3/widgets"
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

// interface for module with event handle
type EventModule interface {
	Module
	HandleEvent(ui.Event)
}

type TabView struct {
	CurrentTab int
	Tabs       []EventModule
}

func (self *TabView) NextTab() {
	self.CurrentTab++
	if self.CurrentTab >= len(self.Tabs) {
		self.CurrentTab = 0
	}
}

// Search
type Search struct {
	Share     *ModuleShare
	SubModule []Module

	SearchBarMode int
	SearchBar     *widgets.Paragraph
	SearchResult  *widgets.List
}

func NewSearch(share *ModuleShare, submodule ...Module) *Search {
	searchbar := widgets.NewParagraph()
	searchbar.Text = ""
	searchbar.Title = "search"
	searchbar.Border = true
	searchbar.TextStyle.Fg = ui.ColorWhite
	searchbar.TextStyle.Bg = ui.ColorBlack

	searchresult := widgets.NewList()
	searchresult.Border = true
	searchresult.Title = "result"
	searchresult.Rows = []string{}
	searchresult.SelectedRow = 0
	searchresult.SelectedRowStyle.Fg = ui.ColorBlack
	searchresult.SelectedRowStyle.Bg = ui.ColorWhite
	searchresult.TextStyle.Fg = ui.ColorWhite
	searchresult.TextStyle.Bg = ui.ColorBlack

	return &Search{
		Share:         share,
		SubModule:     submodule,
		SearchBarMode: SearchTrack,
		SearchBar:     searchbar,
		SearchResult:  searchresult,
	}
}

func (self *Search) Render() {
	ui.Render(self.SearchBar, self.SearchResult)
	for _, m := range self.SubModule {
		m.Render()
	}
}

func (self *Search) Resize(cols, rows int) {
	self.SearchBar.SetRect(0, 0, cols, 3)
	self.SearchResult.SetRect(0, 3, cols, rows-5)
	for _, m := range self.SubModule {
		m.Resize(cols, rows)
	}
}

func (self *Search) HandleEvent(ev ui.Event) {
	switch ev.ID {
	case "<Enter>":
		if self.SearchBar.Text != "" {
			// Load song here
			self.SearchResult.Rows = []string{}
			self.SearchBar.Text = ""
			for i := 0; i < 50; i++ {
				self.SearchResult.Rows = append(self.SearchResult.Rows, fmt.Sprint(i)+":song")
			}
		} else {
			// Add to queue
			self.Share.MusicQueue = append(self.Share.MusicQueue, self.SearchResult.Rows[self.SearchResult.SelectedRow])
		}
	case "<Backspace>":
		if self.SearchBar.Text != "" {
			self.SearchBar.Text = self.SearchBar.Text[:len(self.SearchBar.Text)-1]
		}
	case "<Up>":
		if self.SearchResult.SelectedRow > 0 {
			self.SearchResult.SelectedRow--
		}
	case "<Down>":
		if self.SearchResult.SelectedRow < len(self.SearchResult.Rows)-1 {
			self.SearchResult.SelectedRow++
		}
	default:
		if len(ev.ID) == 1 {
			self.SearchBar.Text += ev.ID
		}
	}
}

// Queue
type Queue struct {
	Share     *ModuleShare
	SubModule []Module

	QueueList *widgets.List
}

func NewQueue(share *ModuleShare, submodule ...Module) *Queue {
	queue := widgets.NewList()
	queue.Border = true
	queue.Title = "queue"
	queue.Rows = []string{}
	queue.SelectedRow = 0
	queue.SelectedRowStyle.Fg = ui.ColorBlack
	queue.SelectedRowStyle.Bg = ui.ColorWhite
	queue.TextStyle.Fg = ui.ColorWhite
	queue.TextStyle.Bg = ui.ColorBlack
	ui.Render(queue)
	ui.Clear()

	return &Queue{
		Share:     share,
		SubModule: submodule,
		QueueList: queue,
	}
}

func (self *Queue) Render() {
	self.QueueList.Rows = self.Share.MusicQueue
	ui.Render(self.QueueList)
	for _, m := range self.SubModule {
		m.Render()
	}
}

func (self *Queue) Resize(cols, rows int) {
	self.QueueList.SetRect(0, 0, cols, rows-5)
	for _, m := range self.SubModule {
		m.Resize(cols, rows)
	}
}

func (self *Queue) HandleEvent(ev ui.Event) {
	switch ev.ID {
	case "<Backspace>":
	case "<Up>":
		if self.QueueList.SelectedRow > 0 {
			self.QueueList.SelectedRow--
		}
	case "<Down>":
		if self.QueueList.SelectedRow < len(self.QueueList.Rows)-1 {
			self.QueueList.SelectedRow++
		}
	case "<Enter>":
		self.Share.MusicCurrent = self.Share.MusicQueue[self.QueueList.SelectedRow]
		self.Share.MusicProgress = 0.5
	case "x":
		index := self.QueueList.SelectedRow
		if index < len(self.Share.MusicQueue) {
			copy(self.Share.MusicQueue[index:], self.Share.MusicQueue[index+1:])
			self.Share.MusicQueue[len(self.Share.MusicQueue)-1] = ""
			self.Share.MusicQueue = self.Share.MusicQueue[:len(self.Share.MusicQueue)-1]
			if index == len(self.Share.MusicQueue) && index != 0 {
				self.QueueList.SelectedRow--
			}
		}
	}
}

// Current Song
type CurrentSong struct {
	Share       *ModuleShare
	SubModule   []Module
	Information *widgets.Paragraph
	Cover       *widgets.Paragraph
}

func NewCurrentSong(share *ModuleShare, submodule ...Module) *CurrentSong {
	information := widgets.NewParagraph()
	information.Border = true
	information.Title = "Currently Playing In Detail..."
	information.Text = "song name 1\nBlaBlaBlaBlaBlaBlaBlaBlBlaBlaBlaBlaa\nBlaBlaBlaBlaBlaBlaBlaBlaBlaBlaBla"
	information.TextStyle.Fg = ui.ColorWhite
	information.TextStyle.Bg = ui.ColorBlack

	cover := widgets.NewParagraph()
	cover.Border = true
	cover.Title = "Cover"
	cover.Text = "*--------*\n*--------*\n*--------*\n*--------*\n*--------*\n"
	cover.TextStyle.Fg = ui.ColorWhite
	cover.TextStyle.Bg = ui.ColorBlack

	return &CurrentSong{
		Share:       share,
		SubModule:   submodule,
		Information: information,
		Cover:       cover,
	}
}

func (self *CurrentSong) Render() {
	ui.Render(self.Information, self.Cover)
	for _, m := range self.SubModule {
		m.Render()
	}
}

func (self *CurrentSong) Resize(cols, rows int) {
	self.Information.SetRect(0, 0, cols/2, rows-5)
	self.Cover.SetRect(cols/2, 0, cols, rows-5)
	for _, m := range self.SubModule {
		m.Resize(cols, rows)
	}
}

func (self *CurrentSong) HandleEvent(ev ui.Event) {
}

// Playing
type Playing struct {
	Share   *ModuleShare
	Playing *widgets.Paragraph
}

func NewPlaying(share *ModuleShare) *Playing {
	playing := widgets.NewParagraph()
	playing.Border = true
	playing.Title = "playing"
	playing.Text = "song name 1"
	playing.TextStyle.Fg = ui.ColorWhite
	playing.TextStyle.Bg = ui.ColorBlack

	return &Playing{
		Share:   share,
		Playing: playing,
	}
}

func (self *Playing) Render() {
	progress := ""
	rect := self.Playing.GetRect().Bounds()
	cols := rect.Max.X - rect.Min.X
	for i := 0; i < int(float64(cols)*self.Share.MusicProgress); i++ {
		progress += "█"
	}
	self.Playing.Text = self.Share.MusicCurrent + "\nmore information\n" + progress
	ui.Render(self.Playing)
}

func (self *Playing) Resize(cols, rows int) {
	self.Playing.SetRect(0, rows-5, cols, rows)
}

type ModuleShare struct {
	MusicQueue    []string // use []song instead of []string
	MusicCurrent  string
	MusicProgress float64 // 0 to 1
}

type App struct {
	Stop bool

	// Tab
	Tab *TabView
}

func NewApp() *App {
	shared := &ModuleShare{}

	playing := NewPlaying(shared)

	app := &App{
		Stop: false,

		Tab: &TabView{
			CurrentTab: 0,
			Tabs: []EventModule{
				NewSearch(shared, playing),
				NewQueue(shared, playing),
				NewCurrentSong(shared, playing),
			},
		},
	}

	app.HandleResize()
	app.Render()
	return app
}

func (self *App) Render() {
	ui.Clear()
	self.Tab.Tabs[self.Tab.CurrentTab].Render()
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

func (self *App) Run() {
	ev := ui.PollEvents()
	for !self.Stop {
		e := <-ev
		self.HandleEvent(e)
		self.Render()
	}
}

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	NewApp().Run()
}
