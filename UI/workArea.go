package UI

import (
	"github.com/gotk3/gotk3/gtk"
)

type workArea struct {
	Consumer *consumer
	Producer *producer
	Notebook *gtk.Notebook
}

func newWorkArea() *workArea {
	workArea := new(workArea)
	workArea.Notebook, _ = gtk.NotebookNew()
	workArea.Consumer = newConsumer()
	workArea.Producer = newProducer()
	workArea.pack()
	return workArea
}

func (w *workArea) pack() {
	clabel, _ := gtk.LabelNew("Consumer")
	w.Notebook.AppendPage(w.Consumer.mainBox, clabel)

	plabel, _ := gtk.LabelNew("Producer")
	w.Notebook.AppendPage(w.Producer.mainBox, plabel)
}
