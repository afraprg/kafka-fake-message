package kafka

import "github.com/Shopify/sarama"

type Kafka struct {
	Brokers []string
	Topic   string
}

func (k *Kafka) init() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(k.Brokers, config)

	return producer, err
}
