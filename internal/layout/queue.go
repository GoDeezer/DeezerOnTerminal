package layout

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"github.com/godeezer/dot/internal/format"
)

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
	queue.Block.PaddingLeft = 1
	ui.Render(queue)
	ui.Clear()

	return &Queue{
		Share:     share,
		SubModule: submodule,
		QueueList: queue,
	}
}

func (self *Queue) Render() {
	cols, _ := ui.TerminalDimensions()
	self.QueueList.Rows = format.FormatSongs(self.Share.MusicQueue, cols-2)
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
			// Delete
		}
	}
}
