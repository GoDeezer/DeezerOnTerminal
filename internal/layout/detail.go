package layout

import (
	"github.com/gizak/termui/v3/widgets"
)

type Detail struct {
	Share       *ModuleShare
	SubModule   []Module
	Information *widgets.Paragraph
	Cover       *widgets.Paragraph
}

func NewDetail(share *ModuleShare, submodule ...Module) *Detail {
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

	return &Detail{
		Share:       share,
		SubModule:   submodule,
		Information: information,
		Cover:       cover,
	}
}

func (self *Detail) Render() {
	ui.Render(self.Information, self.Cover)
	for _, m := range self.SubModule {
		m.Render()
	}
}

func (self *Detail) Resize(cols, rows int) {
	self.Information.SetRect(0, 0, cols/2, rows-5)
	self.Cover.SetRect(cols/2, 0, cols, rows-5)
	for _, m := range self.SubModule {
		m.Resize(cols, rows)
	}
}

func (self *Detail) HandleEvent(ev ui.Event) {
}
