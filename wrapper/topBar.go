package wrapper

import (
	"log"

	"github.com/brxie/kafka-gtk/UI"
	"github.com/brxie/kafka-gtk/kafka"
	"github.com/mattn/go-gtk/glib"
)

type topBar struct {
	kafkaConsumer *kafka.KafkaConsumer
	kafkaProducer *kafka.KafkaProducer
	UI            *UI.UI
	statusChan    *chan interface{}
}

func newTopBar(kafkaConsumer *kafka.KafkaConsumer,
	kafkaProduecer *kafka.KafkaProducer, UI *UI.UI,
	statusChan *chan interface{}) *topBar {

	tb := new(topBar)
	tb.kafkaConsumer = kafkaConsumer
	tb.kafkaProducer = kafkaProduecer
	tb.UI = UI
	tb.statusChan = statusChan

	tb.onClickConnect()
	tb.onClickDisconnect()

	return tb
}

func (t *topBar) onClickConnect() {
	t.UI.Widgets.TopBar.BtnConnect.Connect("clicked", func() {
		var err error
		connecting := true
		go func() {
			defer func() { connecting = false }()
			t.setStatus("Connecting...")
			t.kafkaConsumer.Address = t.GetAddr()
			t.kafkaConsumer.Topic = t.GetTopic()
			t.kafkaConsumer.ClientID = t.GetClientID()
			err = t.kafkaConsumer.Connect()
			if err != nil {
				return
			}

			t.kafkaProducer.Address = t.GetAddr()
			t.kafkaProducer.Topic = t.GetTopic()
			t.kafkaProducer.ClientID = t.GetClientID()
			err = t.kafkaProducer.Connect()
		}()

		t.UI.Widgets.TopBar.BtnConnect.SetSensitive(false)
		glib.IdleAdd(func() bool {
			if connecting {
				return true
			}
			t.UI.Widgets.TopBar.BtnConnect.SetSensitive(true)
			if err != nil {
				log.Println(err)
				t.setStatus(err)
				return false
			}

			t.UI.Widgets.TopBar.HostEntry.SetSensitive(false)
			t.UI.Widgets.TopBar.BtnConnect.SetSensitive(false)
			t.UI.Widgets.TopBar.BtnDisct.SetSensitive(true)
			t.UI.Widgets.TopBar.TopicEntry.SetSensitive(false)
			t.UI.Widgets.TopBar.ClientIDEntry.SetSensitive(false)
			t.UI.Widgets.WorkArea.Consumer.ReadButton.SetSensitive(true)
			t.UI.Widgets.WorkArea.Producer.Launcher.Button.SetSensitive(true)
			t.setStatus("Connected")

			return false
		})
	})
}

func (t *topBar) onClickDisconnect() {
	t.UI.Widgets.TopBar.BtnDisct.Connect("clicked", func() {
		err := t.kafkaConsumer.Close()
		if err != nil {
			t.setStatus(err)
			log.Println(err)
		}

		t.UI.Widgets.TopBar.HostEntry.SetSensitive(true)
		t.UI.Widgets.TopBar.BtnConnect.SetSensitive(true)
		t.UI.Widgets.TopBar.BtnDisct.SetSensitive(false)
		t.UI.Widgets.TopBar.TopicEntry.SetSensitive(true)
		t.UI.Widgets.TopBar.ClientIDEntry.SetSensitive(true)
		t.UI.Widgets.WorkArea.Consumer.ReadButton.SetSensitive(false)
		t.UI.Widgets.WorkArea.Consumer.StopButton.SetSensitive(false)
		t.UI.Widgets.WorkArea.Producer.Launcher.Button.SetSensitive(false)
		t.setStatus("Disconnected")
	})
}

func (t *topBar) GetAddr() string {
	text, _ := t.UI.Widgets.TopBar.HostEntry.GetText()
	return text
}

func (t *topBar) GetTopic() string {
	text, _ := t.UI.Widgets.TopBar.TopicEntry.GetText()
	return text
}

func (t *topBar) GetClientID() string {
	text, _ := t.UI.Widgets.TopBar.TopicEntry.GetText()
	return text
}

func (t *topBar) setStatus(status interface{}) {
	*t.statusChan <- status
}
