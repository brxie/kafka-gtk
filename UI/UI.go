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
	Widgets     *widgets
}

type widgets struct {
	TopBar    *topBar
	WorkArea  *workArea
	StatusBar *statusBar
	box       *gtk.Box
}

func NewUI(windowTitle string, width, height int) *UI {
	window := initGTK(windowTitle)

	return &UI{
		windowTitle: windowTitle,
		width:       width,
		height:      height,
		window:      window,
		Widgets:     newWidgets(),
	}
}

func (u *UI) Render() {
	u.window.Add(u.Widgets.box)
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
	u.Widgets.TopBar.Box.SetSensitive(sensitive)
	u.Widgets.WorkArea.Notebook.SetSensitive(sensitive)
}

func initGTK(winTitle string) *gtk.Window {
	gtk.Init(nil)
	window, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	window.SetTitle(winTitle)
	window.Connect("destroy", gtk.MainQuit)
	return window
}

func newWidgets() *widgets {
	widgets := new(widgets)
	widgets.TopBar = newTopBar()
	widgets.WorkArea = newWorkArea()
	widgets.box, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 1)
	widgets.StatusBar = newStatusBar()
	widgets.pack()
	return widgets
}

func (i *widgets) pack() {
	i.box.PackStart(i.TopBar.Box, false, false, 5)
	sep, _ := gtk.SeparatorNew(gtk.ORIENTATION_HORIZONTAL)
	i.box.PackStart(sep, false, false, 0)
	i.box.PackStart(i.WorkArea.Notebook, true, true, 0)
	i.box.PackStart(i.StatusBar.box, false, true, 0)
}

func (u *UI) ConsumerAutoscroll() bool {
	return u.Widgets.WorkArea.Consumer.AutoScroll.Switch.GetActive()
}
