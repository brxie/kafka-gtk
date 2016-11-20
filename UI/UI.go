package UI

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gotk3/gotk3/gtk"
)

type UI struct {
	windowTitle string
	width       int
	height      int
	window      *gtk.Window
	Items       *items
}

type items struct {
	TopBar    *topBar
	WorkArea  *workArea
	StatusBar *statusBar
	box       *gtk.Box
}

func NewUI(windowTitle string, width, height int) *UI {
	window := initGTK(windowTitle)
	gtkItems := newItems()

	return &UI{
		windowTitle: windowTitle,
		width:       width,
		height:      height,
		window:      window,
		Items:       gtkItems,
	}
}

func (u *UI) Render() {
	u.window.Add(u.Items.box)
	u.window.SetDefaultSize(u.width, u.height)
	u.setAppIco()
	u.window.ShowAll()
	gtk.Main()
}

func (u *UI) setAppIco() {
	cwd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	u.window.SetIconFromFile(cwd + "/ico.png")
}

func (u *UI) SensitiveAll(sensitive bool) {
	u.Items.TopBar.Box.SetSensitive(sensitive)
	u.Items.WorkArea.Notebook.SetSensitive(sensitive)
}

func initGTK(winTitle string) *gtk.Window {
	gtk.Init(nil)
	window, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	window.SetTitle(winTitle)
	window.Connect("destroy", gtk.MainQuit)
	return window
}

func newItems() *items {
	items := new(items)
	items.TopBar = newTopBar()
	items.WorkArea = newWorkArea()
	items.box, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 1)
	items.StatusBar = newStatusBar()
	items.pack()
	return items
}

func (i *items) pack() {
	i.box.PackStart(i.TopBar.Box, false, false, 5)
	sep, _ := gtk.SeparatorNew(gtk.ORIENTATION_HORIZONTAL)
	i.box.PackStart(sep, false, false, 0)
	i.box.PackStart(i.WorkArea.Notebook, true, true, 0)
	i.box.PackStart(i.StatusBar.box, false, true, 0)
}

func (u *UI) ConsumerAutoscroll() bool {
	return u.Items.WorkArea.Consumer.AutoScroll.Switch.GetActive()
}
