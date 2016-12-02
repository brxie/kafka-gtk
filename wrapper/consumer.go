package wrapper

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/brxie/kafka-gtk/UI"
	"github.com/brxie/kafka-gtk/kafka"
	"github.com/gotk3/gotk3/gtk"
	"github.com/mattn/go-gtk/glib"
)

type consumer struct {
	kafkaConsumer *kafka.KafkaConsumer
	sigINT        chan interface{}
	UI            *UI.UI
	statusChan    *chan interface{}
	mu            sync.Mutex
}

func newConsumer(kafka *kafka.KafkaConsumer, UI *UI.UI, statusChan *chan interface{}) *consumer {
	cons := new(consumer)
	cons.kafkaConsumer = kafka
	cons.UI = UI
	cons.sigINT = make(chan interface{}, 1)
	cons.statusChan = statusChan
	cons.mu = sync.Mutex{}

	cons.onClickRead()
	cons.onClickStop()
	cons.onClickClear()

	return cons
}

func (c *consumer) onClickRead() {
	c.UI.Widgets.WorkArea.Consumer.ReadButton.Connect("clicked", func() {
		partConsumer, err := c.kafkaConsumer.NewPartitionConsumer(c.readFromOffset())
		if err != nil {
			c.setStatus(err)
			return
		}

		c.UI.Widgets.WorkArea.Consumer.ReadButton.SetSensitive(false)
		c.UI.Widgets.WorkArea.Consumer.StopButton.SetSensitive(true)
		c.UI.Widgets.WorkArea.Consumer.Offset.Frame.SetSensitive(false)
		c.UI.Widgets.WorkArea.Consumer.MsgConfig.Box.SetSensitive(false)
		c.UI.Widgets.TopBar.BtnDisct.SetSensitive(false)

		c.setStatus("Start reading")

		buff, _ := c.UI.Widgets.WorkArea.Consumer.TextArea.GetBuffer()
		buff.SetText("")

		var msgCnt int64
		// GTK is not threadsave, we can not use goroutine here.
		// Use glib and add function to default main loop,
		// pass function callback, executed until false is return
		glib.IdleAdd(func() bool {
			select {
			case msg := <-(*partConsumer).Messages():
				msgCnt++

				line := c.makeLine(msg, &msgCnt)

				c.mu.Lock()
				buff.Insert(buff.GetEndIter(), (*line).String())
				c.mu.Unlock()

				if c.UI.ConsumerAutoscroll() {
					c.UI.Widgets.WorkArea.Consumer.ScrollDown()
				}
				// dust off buffer when grows up to ~500MB
				if buff.GetCharCount() > 0x1DFFFFFF {
					c.drainBuffer(buff)
				}
				return true
			case <-c.sigINT:
				c.setStatus("Stopped")
				(*partConsumer).Close()
				return false
			default:
				return true
			}
		})

	})
}

func (c *consumer) drainBuffer(buff *gtk.TextBuffer) {
	start := buff.GetStartIter()
	end := buff.GetIterAtOffset(0xFFFFF) // ~1MB
	buff.Delete(start, end)
}

func (c *consumer) makeLine(msg *sarama.ConsumerMessage, count *int64) *bytes.Buffer {
	var line bytes.Buffer
	msgCfg := c.UI.Widgets.WorkArea.Consumer.MsgConfig
	if msgCfg.BtnCount.GetActive() {
		line.WriteString(fmt.Sprintf("%d:\t", *count))
	}
	if msgCfg.BtnOffset.GetActive() {
		line.WriteString(fmt.Sprintf("[offset]: %d\t", msg.Offset))
	}
	if msgCfg.BtnTimestamp.GetActive() {
		line.WriteString(fmt.Sprintf("[timestamp]: %s\t", msg.Timestamp))
	}
	if msgCfg.BtnKey.GetActive() {
		line.WriteString(fmt.Sprintf("[key]: %s\t", msg.Key))
	}
	if msgCfg.BtnValue.GetActive() {
		line.WriteString(fmt.Sprintf("[value]: %s\t", msg.Value))
	}
	if msgCfg.BtnTopic.GetActive() {
		line.WriteString(fmt.Sprintf("[topic]: %s\t", msg.Topic))
	}
	if msgCfg.BtnPartition.GetActive() {
		line.WriteString(fmt.Sprintf("[partition]: %d", msg.Partition))
	}
	line.WriteString("\n")
	return &line
}

func (c *consumer) readFromOffset() int64 {
	if c.UI.Widgets.WorkArea.Consumer.Offset.NewestBtn.GetActive() {
		return kafka.OFFSET_NEWEST
	}
	return kafka.OFFSET_OLDEST
}

func (c *consumer) onClickStop() {
	c.UI.Widgets.WorkArea.Consumer.StopButton.Connect("clicked", func() {
		c.UI.Widgets.WorkArea.Consumer.ReadButton.SetSensitive(true)
		c.UI.Widgets.WorkArea.Consumer.StopButton.SetSensitive(false)
		c.UI.Widgets.WorkArea.Consumer.Offset.Frame.SetSensitive(true)
		c.UI.Widgets.WorkArea.Consumer.MsgConfig.Box.SetSensitive(true)
		c.UI.Widgets.TopBar.BtnDisct.SetSensitive(true)
		c.sigINT <- true
	})

}

func (c *consumer) onClickClear() {
	c.UI.Widgets.WorkArea.Consumer.ClearButton.SetName("foo")
	c.UI.Widgets.WorkArea.Consumer.ClearButton.Connect("clicked", func() {
		defer c.mu.Unlock()
		c.mu.Lock()
		buff, _ := c.UI.Widgets.WorkArea.Consumer.TextArea.GetBuffer()
		buff.SetText("")
	})
}

func (c *consumer) setStatus(status interface{}) {
	*c.statusChan <- status
}
