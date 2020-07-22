package layout

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

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
	self.Playing.Text = self.Share.MusicCurrent.Title + "\nmore information\n" + progress
	ui.Render(self.Playing)
}

func (self *Playing) Resize(cols, rows int) {
	self.Playing.SetRect(0, rows-5, cols, rows)
}
