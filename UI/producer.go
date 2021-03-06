package UI

import "github.com/gotk3/gotk3/gtk"

type producer struct {
	mainBox  *gtk.Box
	Partiton *partition
	Input    *input
	Sender   *sender
}

type partition struct {
	box     *gtk.Box
	frame   *gtk.Frame
	AutoBtn *gtk.CheckButton
	SpinBtn *gtk.SpinButton
}

type input struct {
	frame       *gtk.Frame
	KeyEntry    *gtk.Entry
	box         *gtk.Box
	textFrame   *gtk.Frame
	ValueWindow *gtk.TextView
}

type sender struct {
	box           *gtk.Box
	Button        *gtk.Button
	SplitLinesBtn *gtk.CheckButton
}

func newProducer() *producer {
	producer := new(producer)
	producer.mainBox, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	producer.partitionBox()
	producer.inputWindow()
	producer.sender()
	producer.pack()
	return producer
}

func (p *producer) pack() {
	box, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	box.SetMarginTop(5)
	box.SetMarginStart(5)
	box.SetMarginEnd(5)

	box.PackStart(p.Partiton.box, false, false, 0)
	box.PackStart(p.Input.box, false, false, 0)

	box.PackStart(p.Sender.box, false, false, 0)
	p.mainBox.PackStart(box, true, true, 0)
}

func (p *producer) partitionBox() {
	p.Partiton = new(partition)

	frame, _ := gtk.FrameNew("Partition")
	frame.SetShadowType(gtk.SHADOW_NONE)

	adjust, _ := gtk.AdjustmentNew(0, 0, 999, 1, 0, 0)
	SpinBtn, _ := gtk.SpinButtonNew(adjust, 0, 0)

	autoBtn, _ := gtk.CheckButtonNew()
	autoBtn.SetLabel("Auto")
	autoBtn.SetTooltipText("Assign automatically")
	autoBtn.Connect("toggled", func() {
		if autoBtn.GetActive() {
			SpinBtn.SetSensitive(false)
		} else {
			SpinBtn.SetSensitive(true)
		}
	})
	autoBtn.SetActive(true)

	box, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	box.SetVAlign(gtk.ALIGN_START)

	// GtkFrame can only contain one widget, we have to pack it into box
	insiteBox, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	insiteBox.PackStart(autoBtn, false, false, 2)
	insiteBox.PackStart(SpinBtn, false, false, 2)

	frame.Add(insiteBox)
	box.PackStart(frame, false, false, 0)
	box.SetMarginBottom(10)

	p.Partiton.box = box
	p.Partiton.frame = frame
	p.Partiton.AutoBtn = autoBtn
	p.Partiton.SpinBtn = SpinBtn
}

func (p *producer) inputWindow() {
	p.Input = new(input)

	keyEntry, _ := gtk.EntryNew()
	keyEntry.SetMarginStart(2)
	keyEntry.SetMarginEnd(2)

	p.valueWindow()

	KVBox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	KVBox.PackStart(keyEntry, true, true, 2)
	KVBox.PackStart(p.Input.textFrame, true, true, 2)

	frame, _ := gtk.FrameNew("Message Key")
	frame.SetMarginStart(2)
	frame.SetMarginEnd(2)
	frame.SetMarginBottom(2)
	frame.Add(KVBox)
	frame.SetShadowType(gtk.SHADOW_NONE)

	box, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	box.PackStart(frame, true, true, 0)
	box.SetVAlign(gtk.ALIGN_START)

	p.Input.KeyEntry = keyEntry
	p.Input.frame = frame
	p.Input.box = box
}

func (p *producer) valueWindow() {
	textArea, _ := gtk.TextViewNew()
	textArea.SetMarginStart(5)
	textArea.SetMarginEnd(5)
	textArea.SetLeftMargin(5)
	textArea.SetRightMargin(5)
	textArea.SetEditable(true)

	scrollWin, _ := gtk.ScrolledWindowNew(nil, nil)
	scrollWin.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	scrollWin.SetHasWindow(true)
	scrollWin.Add(textArea)
	scrollWin.SetSizeRequest(0, 200)

	textFrame, _ := gtk.FrameNew("Values")
	textFrame.SetMarginStart(2)
	textFrame.SetMarginEnd(2)
	textFrame.SetMarginBottom(2)
	textFrame.Add(scrollWin)
	p.Input.textFrame = textFrame
	p.Input.ValueWindow = textArea
}

func (p *producer) sender() {
	p.Sender = new(sender)

	splitLinesBtn, _ := gtk.CheckButtonNew()
	splitLinesBtn.SetLabel("Split new lines")
	splitLinesBtn.SetTooltipText("Split message value on \\n character")
	splitLinesBtn.SetMarginEnd(5)

	button, _ := gtk.ButtonNew()
	button.SetLabel("Send")
	button.SetSensitive(false)
	button.SetHExpand(true)
	button.SetMarginStart(5)
	button.SetMarginEnd(25)

	box, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	box.PackStart(button, true, true, 0)
	box.PackStart(splitLinesBtn, false, false, 0)

	p.Sender.SplitLinesBtn = splitLinesBtn
	p.Sender.box = box
	p.Sender.Button = button
}
