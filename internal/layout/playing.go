package layout

type Playing struct {
	share   *moduleshare
	Playing *widgets.paragraph
}

func newPlaying(share *moduleshare) *playing {
	playing := widgets.newparagraph()
	playing.border = true
	playing.title = "playing"
	playing.text = "song name 1"
	playing.textstyle.fg = ui.colorwhite
	playing.textstyle.bg = ui.colorblack

	return &playing{
		share:   share,
		playing: playing,
	}
}

func (self *Playing) render() {
	progress := ""
	rect := self.playing.getrect().bounds()
	cols := rect.max.x - rect.min.x
	for i := 0; i < int(float64(cols)*self.share.musicprogress); i++ {
		progress += "█"
	}
	self.Playing.text = self.share.musiccurrent + "\nmore information\n" + progress
	ui.render(self.Playing)
}

func (self *Playing) resize(cols, rows int) {
	self.Playing.setrect(0, rows-5, cols, rows)
}
