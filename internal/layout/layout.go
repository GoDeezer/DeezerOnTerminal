package layout

import (
	ui "github.com/gizak/termui/v3"
)

type Module interface {
	Render()
	Resize(int, int)
}

// interface for module with event handle
type Layout interface {
	Module
	HandleEvent(ui.Event)
}

type LayoutList struct {
	CurrentLayout int
	Layout        []Layout
}

func (self *LayoutList) Next() {
	self.CurrentLayout++
	if self.CurrentLayout >= len(self.Layout) {
		self.CurrentLayout = 0
	}
}

func (self *LayoutList) Render() {
	self.Layout[self.CurrentLayout].Render()
}

func (self *LayoutList) HandleEvent(ev ui.Event) {
	self.Layout[self.CurrentLayout].HandleEvent(ev)
}
