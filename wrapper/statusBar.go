package wrapper

import (
	"fmt"

	"github.com/brxie/kafka-gtk/UI"
	"github.com/mattn/go-gtk/glib"
)

func newStatusBar(UI *UI.UI, statusChan *chan interface{}) {
	glib.IdleAdd(func() bool {
		select {
		case text := <-*statusChan:
			t := fmt.Sprintf("%s", text)
			UI.Widgets.StatusBar.Push(&t)
			return true
		default:
			return true
		}
	})
}
