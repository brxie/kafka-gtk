package main

import (
	"github.com/brxie/kafka-gtk/UI"
	"github.com/brxie/kafka-gtk/wrapper"
)

func main() {
	ui := UI.NewUI("KafkaGTK", 600, 700)
	wrapper.NewWrapper(ui)
	ui.Render()
}
