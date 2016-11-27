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
			err := p.kafkaProducer.Produce(p.getKey(), p.getValue())
			if err != nil {
				p.setStatus(err)
			}
			return false
		})
	})
}

func (p *producer) getKey() *string {
	key, _ := p.UI.Widgets.WorkArea.Producer.Input.KeyEntry.GetText()
	return &key
}

func (p *producer) getValue() *string {
	//val, _ := p.UI.Widgets.WorkArea.Producer.Input.ValueEntry.GetText()
	val := "aa"
	return &val
}

func (p *producer) setStatus(status interface{}) {
	*p.statusChan <- status
}
