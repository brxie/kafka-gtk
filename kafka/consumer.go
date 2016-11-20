package kafka

import (
	"github.com/Shopify/sarama"
)

type KafkaConsumer struct {
	consumer *sarama.Consumer
	Address  string
	Topic    string
	ClientID string
}

const (
	OFFSET_NEWEST = sarama.OffsetNewest
	OFFSET_OLDEST = sarama.OffsetOldest
)

func NewKafkaConsumer() *KafkaConsumer {
	kcons := new(KafkaConsumer)
	kcons.ClientID = "KafkaGTK"
	return kcons
}

func (k *KafkaConsumer) Connect() error {
	config := sarama.NewConfig()
	config.ClientID = k.ClientID

	consumer, err := sarama.NewConsumer([]string{k.Address}, config)
	if err != nil {
		return err
	}
	k.consumer = &consumer
	return nil
}

func (k *KafkaConsumer) NewPartitionConsumer(offset int64) (*sarama.PartitionConsumer, error) {
	partitionConsumer, err := (*k.consumer).ConsumePartition(k.Topic, 0, offset)
	if err != nil {
		return nil, err
	}
	return &partitionConsumer, nil
}

func (k *KafkaConsumer) Close() error {
	if err := (*k.consumer).Close(); err != nil {
		return err
	}
	return nil
}
