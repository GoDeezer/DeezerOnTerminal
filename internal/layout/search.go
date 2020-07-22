package layout

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

const (
	SearchTrack = iota
	SearchAlbum
	SearchArtist
	SearchPlaylist
)

// TODO Use string builder/bytebuffer for input field
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
			res := []string{}
			query, err := self.Share.DeezerClient.Search(self.SearchBar.Text, "", "", 0, 20)
			if err != nil {
				res = append(res, "error", fmt.Sprint(err))
			} else {
				for _, q := range query.Songs.Data {
					artist := q.ArtistName
					title := q.Title
					if len(artist) > 20 {
						artist = artist[:20]
					}
					if len(title) > 20 {
						title = title[:20]
					}
					res = append(res, fmt.Sprintf("%-20s | %-20s %d", artist, title, q.ExplicitContent.LyricsStatus))
				}
			}
			self.SearchBar.Text = ""
			self.SearchResult.Rows = res
		} else {
			// Add to queue
			self.Share.MusicQueue = append(self.Share.MusicQueue, self.SearchResult.Rows[self.SearchResult.SelectedRow])
		}
	case "<Space>":
		self.SearchBar.Text += " "
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
