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

func (self *Playing) Update() {
	progress := ""
	for i := 0; i < int(float64(self.Share.Cols)*self.Share.Player.Progress); i++ {
		progress += "█"
	}
	self.Playing.Text = self.Share.Player.CurrentSong.Title + "\nmore information\n" + progress
}

func (self *Playing) Render() {
	self.Update()
	ui.Render(self.Playing)
}

func (self *Playing) GetDrawable() []ui.Drawable {
	return []ui.Drawable{
		self.Playing,
	}
}

func (self *Playing) Resize(cols, rows int) {
	self.Playing.SetRect(0, rows-5, cols, rows)
}
