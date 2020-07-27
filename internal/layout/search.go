package layout

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/godeezer/dot/internal/format"
)

const (
	SearchSong = iota
	SearchAlbum
	SearchArtist
	//	SearchPlaylist playlist currently not supported by lib
	SearchLast
)

// TODO Use string builder/bytebuffer for input field
type Search struct {
	Share     *ModuleShare
	SubModule []Module

	SearchBarMode int // Track, Album, Artist
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
		SearchBarMode: SearchSong,
		SearchBar:     searchbar,
		SearchResult:  searchresult,
	}
}

// methods
func (self *Search) NextMode() {
	self.SearchBarMode++
	if self.SearchBarMode >= SearchLast {
		self.SearchBarMode = 0
	}
}

func (self *Search) LoadQuery() {
	query, err := self.Share.DeezerClient.Search(self.SearchBar.Text, "", "", 0, -1)
	if err != nil {
		self.Share.Popup = NewPopup(self.Share)
		self.Share.Popup.(*Popup).Box.Text = "Error\n" + fmt.Sprint(err)
		self.Share.Popup.Resize(ui.TerminalDimensions())
		return
	}

	// Setup list
	switch self.SearchBarMode {
	case SearchSong:
		self.SearchResult.Rows = format.FormatSongs(query.Songs.Data, self.Share.Cols)
	case SearchAlbum:
		self.SearchResult.Rows = format.FormatAlbums(query.Albums.Data, self.Share.Cols)
	case SearchArtist:
		self.SearchResult.Rows = format.FormatArtists(query.Artists.Data, self.Share.Cols)
	}

	// Reset
	self.SearchBar.Text = ""
	self.SearchResult.SelectedRow = 0

	// Storing query to shared
	self.Share.QueryResult = query
}

// Add selected song/album/artist to queue
func (self *Search) AddQueue() {
	if self.Share.QueryResult == nil {
		return
	}
	switch self.SearchBarMode {
	case SearchSong: // Add single song to queue
		self.Share.Player.PlayerQueue.AddSong(self.Share.QueryResult.Songs.Data[self.SearchResult.SelectedRow])
	case SearchAlbum: // Add entire album to queue
		self.Share.Player.PlayerQueue.AddAlbum(self.Share.QueryResult.Albums.Data[self.SearchResult.SelectedRow])
	case SearchArtist: // Add songs of artist to queue
		self.Share.Player.PlayerQueue.AddArtist(self.Share.QueryResult.Artists.Data[self.SearchResult.SelectedRow])
	}
}

// interface

func (self *Search) Update() {
	self.SearchBar.Title = "search"
	switch self.SearchBarMode {
	case SearchSong:
		self.SearchBar.Title += " - song"
	case SearchAlbum:
		self.SearchBar.Title += " - album"
	case SearchArtist:
		self.SearchBar.Title += " - artist"
	}

	self.SearchResult.Title = "result - " + fmt.Sprint(len(self.SearchResult.Rows))
}

func (self *Search) Render() {
	self.Update()
	ui.Clear()
	for _, m := range self.SubModule {
		m.Render()
	}
	ui.Render(self.SearchBar, self.SearchResult)
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
			self.LoadQuery()
		} else {
			self.AddQueue()
		}
	case "<Space>":
		self.SearchBar.Text += " "
	case "<Backspace>":
		if self.SearchBar.Text != "" {
			self.SearchBar.Text = self.SearchBar.Text[:len(self.SearchBar.Text)-1]
		} else {
			self.NextMode()
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
