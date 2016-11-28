package wrapper

import (
	"github.com/brxie/kafka-gtk/UI"
	"github.com/brxie/kafka-gtk/kafka"
	"github.com/gotk3/gotk3/glib"
)

type producer struct {
	kafkaProducer *kafka.KafkaProducer
	UI            *UI.UI
	statusChan    *chan interface{}
}

func newProducer(kafka *kafka.KafkaProducer, UI *UI.UI, statusChan *chan interface{}) *producer {
	producer := new(producer)
	producer.kafkaProducer = kafka
	producer.UI = UI
	producer.onClickSend()
	return producer
}

func (p *producer) onClickSend() {
	p.UI.Widgets.WorkArea.Producer.Launcher.Button.Connect("clicked", func() {
		// GTK is not threadsave, we can not use goroutine here.
		// Use glib and add function to default main loop,
		// pass function callback, executed until false is return
		glib.IdleAdd(func() bool {
			err := p.kafkaProducer.Produce(p.getKey(), p.getValue(), p.getPartitionNumber())
			if err != nil {
				p.setStatus(err)
			}
			p.clearValueWindow()
			return false
		})
	})
}

func (p *producer) getPartitionNumber() *int {
	autoSelect := p.UI.Widgets.WorkArea.Producer.Partiton.CheckBtn.GetActive()
	if autoSelect {
		return nil
	}
	partnr := p.UI.Widgets.WorkArea.Producer.Partiton.SpinBtn.GetValueAsInt()
	return &partnr
}

func (p *producer) clearValueWindow() {
	buff, _ := p.UI.Widgets.WorkArea.Producer.Input.ValueWindow.GetBuffer()
	buff.SetText("")
}

func (p *producer) getKey() *string {
	key, _ := p.UI.Widgets.WorkArea.Producer.Input.KeyEntry.GetText()
	return &key
}

func (p *producer) getValue() *string {
	buff, _ := p.UI.Widgets.WorkArea.Producer.Input.ValueWindow.GetBuffer()
	start, end := buff.GetBounds()
	text, _ := buff.GetText(start, end, true)
	return &text
}

func (p *producer) setStatus(status interface{}) {
	*p.statusChan <- status
}
