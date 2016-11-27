package UI

import "github.com/gotk3/gotk3/gtk"

type topBar struct {
	Box           *gtk.Box
	HostEntry     *gtk.Entry
	BtnConnect    *gtk.Button
	BtnDisct      *gtk.Button
	separator     *gtk.Separator
	TopicEntry    *gtk.Entry
	ClientIDEntry *gtk.Entry
}

func newTopBar() *topBar {
	topBar := new(topBar)
	box, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	topBar.Box = box
	topBar.Box.SetSizeRequest(0, 20)

	// buttons
	btnBox, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	topBar.BtnConnect, _ = gtk.ButtonNewFromIconName("goa-panel-symbolic", 2)
	topBar.BtnConnect.SetTooltipText("Connect")
	topBar.BtnDisct, _ = gtk.ButtonNewFromIconName("edit-delete-symbolic", 2)
	topBar.BtnDisct.SetTooltipText("Disconnect")
	topBar.BtnDisct.SetSensitive(false)
	btnBox.PackStart(topBar.BtnConnect, false, false, 0)
	btnBox.PackStart(topBar.BtnDisct, false, false, 0)

	topBar.separator, _ = gtk.SeparatorNew(gtk.ORIENTATION_VERTICAL)

	// entries
	topBar.HostEntry, _ = gtk.EntryNew()
	topBar.HostEntry.SetText("localhost:9092")
	topBar.HostEntry.SetTooltipText("Kafka broker host:port")
	topBar.HostEntry.SetMarginStart(5)
	topBar.HostEntry.SetWidthChars(20)
	topBar.HostEntry.SetEditable(true)

	topBar.TopicEntry, _ = gtk.EntryNew()
	topBar.TopicEntry.SetText("topic")
	topBar.TopicEntry.SetTooltipText("Kafka topic")
	topBar.TopicEntry.SetMarginStart(10)
	topBar.TopicEntry.SetEditable(true)
	topBar.TopicEntry.SetWidthChars(25)

	topBar.ClientIDEntry, _ = gtk.EntryNew()
	topBar.ClientIDEntry.SetText("KafkaGTK")
	topBar.ClientIDEntry.SetTooltipText("Kafka client ID")
	topBar.ClientIDEntry.SetMarginEnd(5)
	topBar.ClientIDEntry.SetEditable(true)
	topBar.ClientIDEntry.SetWidthChars(10)

	topBar.Box.PackStart(topBar.HostEntry, false, false, 0)
	topBar.Box.PackStart(btnBox, false, false, 0)
	topBar.Box.PackStart(topBar.separator, false, false, 2)
	topBar.Box.PackStart(topBar.TopicEntry, false, false, 0)
	topBar.Box.PackStart(topBar.ClientIDEntry, false, false, 0)

	return topBar
}
