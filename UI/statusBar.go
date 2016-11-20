package UI

import (
	"github.com/gotk3/gotk3/gtk"
)

type statusBar struct {
	bar *gtk.Statusbar
	box *gtk.Box
}

func newStatusBar() *statusBar {
	sbar := new(statusBar)
	sbar.bar, _ = gtk.StatusbarNew()
	sbar.box, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 1)
	sbar.pack()
	return sbar
}

func (s *statusBar) pack() {
	s.box.PackStart(s.bar, true, true, 2)
}

func (s *statusBar) Push(text *string) {
	s.bar.Push(0, *text)
}
