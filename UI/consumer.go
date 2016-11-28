package UI

import (
	"github.com/gotk3/gotk3/gtk"
)

type consumer struct {
	mainBox     *gtk.Box
	TextArea    *gtk.TextView
	textFrame   *gtk.Frame
	scrollWin   *gtk.ScrolledWindow
	menuBox     *gtk.Box
	ReadButton  *gtk.Button
	StopButton  *gtk.Button
	ClearButton *gtk.Button
	Offset      *offset
	AutoScroll  *autoScroll
	MsgConfig   *msgConfig
}

type offset struct {
	NewestBtn *gtk.RadioButton
	OldestBtn *gtk.RadioButton
	Frame     *gtk.Frame
}

type autoScroll struct {
	Switch *gtk.Switch
	Box    *gtk.Box
}

type msgConfig struct {
	BtnCount     *gtk.CheckButton
	BtnOffset    *gtk.CheckButton
	BtnTimestamp *gtk.CheckButton
	BtnKey       *gtk.CheckButton
	BtnValue     *gtk.CheckButton
	BtnTopic     *gtk.CheckButton
	BtnPartition *gtk.CheckButton
	Box          *gtk.Box
	Expander     *gtk.Expander
}

func newConsumer() *consumer {
	consumer := new(consumer)
	consumer.mainBox, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	consumer.menuBox, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)

	consumer.offsetBox()
	consumer.autoScrollSwitch()
	consumer.buttons()
	consumer.outputWindow()
	consumer.msgCfg()
	consumer.pack()
	return consumer
}

func (c *consumer) ScrollDown() {
	adj := c.scrollWin.GetVAdjustment()
	adj.SetValue(adj.GetUpper() - adj.GetPageSize())
	c.scrollWin.SetVAdjustment(adj)
}

func (c *consumer) pack() {
	c.mainBox.PackStart(c.menuBox, false, true, 0)
	c.menuBox.PackStart(c.MsgConfig.Expander, false, true, 0)
	c.mainBox.SetMarginTop(5)

	c.mainBox.PackStart(c.textFrame, true, true, 1)
}

func (c *consumer) buttons() {

	btnBox, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)

	readBtn, _ := gtk.ButtonNewFromIconName("media-playback-start", 4)
	readBtn.SetTooltipText("Read")
	readBtn.SetSensitive(false)
	readBtn.SetMarginStart(5)

	stopBtn, _ := gtk.ButtonNewFromIconName("media-playback-stop", 4)
	stopBtn.SetTooltipText("Stop")
	stopBtn.SetSensitive(false)

	clrBtn, _ := gtk.ButtonNewFromIconName("edit-clear-all-symbolic", 4)
	clrBtn.SetTooltipText("Clear")
	clrBtn.SetMarginStart(5)

	btnBox.Add(readBtn)
	btnBox.Add(stopBtn)
	btnBox.Add(clrBtn)
	btnBox.Add(c.AutoScroll.Box)
	btnBox.Add(c.Offset.Frame)

	c.menuBox.PackStart(btnBox, true, false, 0)

	c.ReadButton = readBtn
	c.StopButton = stopBtn
	c.ClearButton = clrBtn
}

func (c *consumer) outputWindow() {
	textArea, _ := gtk.TextViewNew()
	textArea.SetEditable(false)
	textArea.SetMarginStart(5)
	textArea.SetMarginEnd(5)
	textArea.SetLeftMargin(5)
	textArea.SetRightMargin(5)

	scrollWin, _ := gtk.ScrolledWindowNew(nil, nil)
	scrollWin.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	scrollWin.SetHasWindow(true)
	scrollWin.Add(textArea)

	textFrame, _ := gtk.FrameNew("output")
	textFrame.SetMarginStart(2)
	textFrame.SetMarginEnd(2)
	textFrame.SetMarginBottom(2)
	textFrame.Add(scrollWin)

	c.TextArea = textArea
	c.textFrame = textFrame
	c.scrollWin = scrollWin
}

func (c *consumer) autoScrollSwitch() {
	c.AutoScroll = new(autoScroll)
	box, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	box.SetMarginEnd(5)
	label, _ := gtk.LabelNew("Auto scroll")
	switch_, _ := gtk.SwitchNew()
	switch_.SetActive(true)
	box.Add(label)
	box.Add(switch_)

	box.SetMarginStart(10)
	c.AutoScroll.Box = box
	c.AutoScroll.Switch = switch_
}

func (c *consumer) msgCfg() {
	buttNames := []string{"Counter", "Offset", "Timestamp", "Key", "Value", "Topic", "Partition"}
	buttons := make(map[string]*gtk.CheckButton)
	c.MsgConfig = new(msgConfig)

	exp, _ := gtk.ExpanderNew("")
	exp.SetLabel("Values")
	box, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	box.SetMarginStart(5)

	btnBox, _ := gtk.GridNew()
	btnSize := 30
	var cnt, row int
	for _, btnLabel := range buttNames {
		btn, _ := gtk.CheckButtonNew()
		btn.SetLabel(btnLabel)
		sep, _ := gtk.SeparatorNew(gtk.ORIENTATION_VERTICAL)
		btnBox.Add(sep)
		btnBox.Attach(btn, cnt*btnSize, row*btnSize, btnSize, btnSize)
		buttons[btnLabel] = btn
		cnt++
		if cnt > 3 {
			cnt = 0
			row++
		}
	}

	box.PackStart(btnBox, false, false, 0)
	exp.Add(box)

	c.MsgConfig.Expander = exp
	c.MsgConfig.Box = box
	c.MsgConfig.BtnCount = buttons["Counter"]
	c.MsgConfig.BtnOffset = buttons["Offset"]
	c.MsgConfig.BtnTimestamp = buttons["Timestamp"]
	c.MsgConfig.BtnKey = buttons["Key"]
	c.MsgConfig.BtnValue = buttons["Value"]
	c.MsgConfig.BtnTopic = buttons["Topic"]
	c.MsgConfig.BtnPartition = buttons["Partition"]

	c.MsgConfig.BtnOffset.SetActive(true)
	c.MsgConfig.BtnValue.SetActive(true)
}

func (c *consumer) offsetBox() {
	c.Offset = new(offset)
	frame, _ := gtk.FrameNew("Offset")
	frame.SetMarginStart(35)
	frame.SetShadowType(gtk.SHADOW_NONE)

	box, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	var btn *gtk.RadioButton
	bnew, _ := gtk.RadioButtonNewWithLabelFromWidget(btn, "Newest")
	bold, _ := gtk.RadioButtonNewWithLabelFromWidget(bnew, "Oldest")

	box.PackStart(bnew, false, false, 0)
	box.PackStart(bold, false, false, 0)
	frame.Add(box)

	c.Offset.NewestBtn = bnew
	c.Offset.OldestBtn = bold
	c.Offset.Frame = frame
}
