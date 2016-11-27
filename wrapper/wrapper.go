package wrapper

import (
	"github.com/brxie/kafka-gtk/UI"
	"github.com/brxie/kafka-gtk/kafka"
)

type wrapper struct {
	kafkaConsumer *kafka.KafkaConsumer
	kafkaProducer *kafka.KafkaProducer
	statusChan    *chan interface{}
	UI            *UI.UI
	topBar        *topBar
	consumer      *consumer
	producer      *producer
}

func NewWrapper(ui *UI.UI) *wrapper {
	wrap := new(wrapper)
	wrap.UI = ui

	wrap.kafkaConsumer = kafka.NewKafkaConsumer()
	wrap.kafkaProducer = kafka.NewKafkaProducer()
	schan := make(chan interface{}, 1)
	wrap.statusChan = &schan

	wrap.topBar = newTopBar(wrap.kafkaConsumer, wrap.kafkaProducer, wrap.UI, wrap.statusChan)
	wrap.consumer = newConsumer(wrap.kafkaConsumer, wrap.UI, wrap.statusChan)
	wrap.producer = newProducer(wrap.kafkaProducer, wrap.UI, wrap.statusChan)
	newStatusBar(wrap.UI, wrap.statusChan)

	return wrap
}

func (w *wrapper) setStatus(status interface{}) {
	*w.statusChan <- status
}
