package UI

import (
	"github.com/gotk3/gotk3/gtk"
)

type producer struct {
	mainBox  *gtk.Box
	partiton *partition
	input    *input
	launcher *launcher
}

type partition struct {
	box      *gtk.Box
	frame    *gtk.Frame
	checkBtn *gtk.CheckButton
	spinBtn  *gtk.SpinButton
}

type input struct {
	frame    *gtk.Frame
	keyEntry *gtk.Entry
	box      *gtk.Box
}

type launcher struct {
	box    *gtk.Box
	button *gtk.Button
}

func newProducer() *producer {
	producer := new(producer)
	producer.mainBox, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	producer.partitionBox()
	producer.inputWindow()
	producer.sendButton()
	producer.pack()
	return producer
}

func (p *producer) pack() {
	box, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)

	box.PackStart(p.partiton.box, false, false, 0)
	box.PackStart(p.input.box, true, true, 0)
	box.PackStart(p.launcher.box, false, false, 0)
	p.mainBox.PackStart(box, true, true, 0)
}

func (p *producer) partitionBox() {
	p.partiton = new(partition)

	frame, _ := gtk.FrameNew("Partition")
	frame.SetShadowType(gtk.SHADOW_NONE)

	checkBtn, _ := gtk.CheckButtonNew()
	checkBtn.SetLabel("Auto")
	checkBtn.SetTooltipText("Assign automatically")

	adjust, _ := gtk.AdjustmentNew(0, 0, 999, 1, 0, 0)
	spinBtn, _ := gtk.SpinButtonNew(adjust, 0, 0)
	box, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	box.SetVAlign(gtk.ALIGN_START)

	// GtkFrame can only contain one widget, we have to pack it into box
	insiteBox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	insiteBox.PackStart(checkBtn, false, false, 2)
	insiteBox.PackStart(spinBtn, false, false, 2)

	frame.Add(insiteBox)
	box.PackStart(frame, false, false, 0)

	p.partiton.box = box
	p.partiton.frame = frame
	p.partiton.checkBtn = checkBtn
	p.partiton.spinBtn = spinBtn
}

func (p *producer) inputWindow() {
	p.input = new(input)

	keyEntry, _ := gtk.EntryNew()
	keyEntry.SetMarginStart(5)
	keyEntry.SetMarginEnd(5)

	valueEntry, _ := gtk.EntryNew()
	valueEntry.SetMarginStart(5)
	valueEntry.SetMarginEnd(5)

	// GtkFrame can only contain one widget, we have to pack it into box
	KVBox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	KVBox.PackStart(keyEntry, true, true, 2)
	KVBox.PackStart(valueEntry, true, true, 2)

	textFrame, _ := gtk.FrameNew("Message Key/Value")
	textFrame.SetMarginStart(2)
	textFrame.SetMarginEnd(2)
	textFrame.SetMarginBottom(2)
	textFrame.Add(KVBox)
	textFrame.SetShadowType(gtk.SHADOW_NONE)

	box, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	box.PackStart(textFrame, true, true, 0)
	box.SetVAlign(gtk.ALIGN_START)

	p.input.keyEntry = keyEntry
	p.input.frame = textFrame

	p.input.box = box
}

func (p *producer) sendButton() {
	p.launcher = new(launcher)

	button, _ := gtk.ButtonNew()
	button.SetLabel("Send")
	button.SetMarginTop(18)

	box, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	box.PackStart(button, false, false, 0)

	p.launcher.box = box
	p.launcher.button = button
}
