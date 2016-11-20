package wrapper

import (
	"log"

	"github.com/brxie/kafka-gtk/UI"
	"github.com/brxie/kafka-gtk/kafka"
	"github.com/mattn/go-gtk/glib"
)

type topBar struct {
	kafkaConsumer *kafka.KafkaConsumer
	UI            *UI.UI
	statusChan    *chan interface{}
}

func newTopBar(kafka *kafka.KafkaConsumer, UI *UI.UI, statusChan *chan interface{}) *topBar {
	tb := new(topBar)
	tb.kafkaConsumer = kafka
	tb.UI = UI
	tb.statusChan = statusChan

	tb.onClickConnect()
	tb.onClickDisconnect()

	return tb
}

func (t *topBar) onClickConnect() {
	t.UI.Items.TopBar.BtnConnect.Connect("clicked", func() {
		var err error
		var connecting bool
		go func() {
			t.kafkaConsumer.Address = t.GetAddr()
			t.kafkaConsumer.Topic = t.GetTopic()
			t.kafkaConsumer.ClientID = t.GetClientID()
			err = t.kafkaConsumer.Connect()
			connecting = true
		}()
		t.UI.Items.TopBar.BtnConnect.SetSensitive(false)
		glib.IdleAdd(func() bool {
			if !connecting {
				return true
			}
			t.UI.Items.TopBar.BtnConnect.SetSensitive(true)
			if err != nil {
				log.Println(err)
				t.setStatus(err)
				return false
			}

			t.UI.Items.TopBar.HostEntry.SetSensitive(false)
			t.UI.Items.TopBar.BtnConnect.SetSensitive(false)
			t.UI.Items.TopBar.BtnDisct.SetSensitive(true)
			t.UI.Items.TopBar.TopicEntry.SetSensitive(false)
			t.UI.Items.TopBar.ClientIDEntry.SetSensitive(false)
			t.UI.Items.WorkArea.Consumer.ReadButton.SetSensitive(true)
			t.setStatus("Connected")

			return false
		})
	})
}

func (t *topBar) onClickDisconnect() {
	t.UI.Items.TopBar.BtnDisct.Connect("clicked", func() {
		err := t.kafkaConsumer.Close()
		if err != nil {
			t.setStatus(err)
			log.Println(err)
		}

		t.UI.Items.TopBar.HostEntry.SetSensitive(true)
		t.UI.Items.TopBar.BtnConnect.SetSensitive(true)
		t.UI.Items.TopBar.BtnDisct.SetSensitive(false)
		t.UI.Items.TopBar.TopicEntry.SetSensitive(true)
		t.UI.Items.TopBar.ClientIDEntry.SetSensitive(true)
		t.UI.Items.WorkArea.Consumer.ReadButton.SetSensitive(false)
		t.UI.Items.WorkArea.Consumer.StopButton.SetSensitive(false)
		t.setStatus("Disconnected")
	})
}

func (t *topBar) GetAddr() string {
	text, _ := t.UI.Items.TopBar.HostEntry.GetText()
	return text
}

func (t *topBar) GetTopic() string {
	text, _ := t.UI.Items.TopBar.TopicEntry.GetText()
	return text
}

func (t *topBar) GetClientID() string {
	text, _ := t.UI.Items.TopBar.TopicEntry.GetText()
	return text
}

func (t *topBar) setStatus(status interface{}) {
	*t.statusChan <- status
}
