package layout

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Popup struct {
	Share *ModuleShare

	Box *widgets.Paragraph
}

func NewPopup(share *ModuleShare) *Popup {
	box := widgets.NewParagraph()
	box.Border = true
	box.Title = "error"
	box.TextStyle.Fg = ui.ColorWhite
	box.TextStyle.Bg = ui.ColorBlack
	ui.Render(box)
	ui.Clear()

	return &Popup{
		Share: share,
		Box:   box,
	}
}

// interface

func (self *Popup) Render() {
	ui.Render(self.Box)
}

func (self *Popup) Resize(cols, rows int) {
	self.Box.SetRect(2, 2, cols-2, rows-2)
}

func (self *Popup) HandleEvent(ev ui.Event) {
	switch ev.ID {
	case "<Enter>":
		fallthrough
	case "<Backspace>":
		self.Share.Popup = nil
	}
}
