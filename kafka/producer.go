package kafka

import "github.com/Shopify/sarama"

type KafkaProducer struct {
	producer *sarama.SyncProducer
	Address  string
	Topic    string
	ClientID string
}

func NewKafkaProducer() *KafkaProducer {
	prod := new(KafkaProducer)
	prod.ClientID = "KafkaGTK"
	return prod
}

func (k *KafkaProducer) Connect() error {
	config := sarama.NewConfig()
	config.ClientID = k.ClientID
	producer, err := sarama.NewSyncProducer([]string{k.Address}, config)
	if err != nil {
		return err
	}
	k.producer = &producer
	return nil
}

func (k *KafkaProducer) Produce(key, value *string) error {
	var err error
	msg := k.createMessage(key, value)
	_, _, err = (*k.producer).SendMessage(msg)

	if err != nil {
		return err
	}
	return nil
}

func (k *KafkaProducer) createMessage(key, value *string) *sarama.ProducerMessage {
	message := &sarama.ProducerMessage{Topic: k.Topic, Partition: 0}

	if *key != "" {
		message.Key = sarama.StringEncoder(*key)
	}
	if *value != "" {
		message.Value = sarama.StringEncoder(*value)
	}
	return message
}

func (k *KafkaProducer) Close() error {
	if err := (*k.producer).Close(); err != nil {
		return err
	}
	return nil
}
