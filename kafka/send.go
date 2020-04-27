package kafka

import (
	"github.com/Shopify/sarama"
)

func prepareMessage(topic, message string) *sarama.ProducerMessage {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(message),
	}

	return msg
}

func (k *Kafka) SendMessage(message string) {
	initKafka, err := k.init()
	if err != nil {
		panic("Cannot initialize")
	}

	initKafka.SendMessage(prepareMessage(k.Topic, message))
}
