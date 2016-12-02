package wrapper

import (
	"fmt"
	"time"

	"github.com/brxie/kafka-gtk/UI"
	"github.com/mattn/go-gtk/glib"
)

const statusRetentionTime = 10 // seconds

func newStatusBar(UI *UI.UI, statusChan *chan interface{}) {
	var lastUpdate time.Time
	glib.IdleAdd(func() bool {
		select {
		case text := <-*statusChan:
			t := fmt.Sprintf("%s", text)
			UI.Widgets.StatusBar.Push(&t)
			lastUpdate = time.Now()
		default:
			// clear status after given time
			if time.Since(lastUpdate).Seconds() > statusRetentionTime {
				UI.Widgets.StatusBar.Push(new(string))
			}
		}
		return true
	})
}
