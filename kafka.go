package common

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

// KafkaProducer : Kafka Producer
type KafkaProducer struct {
	SyncProducer sarama.SyncProducer
	Brokers      []string
}

func (prod *KafkaProducer) New() sarama.SyncProducer {

	// For the data collector, we are looking for strong consistency semantics.
	// Because we don't change the flush settings, sarama will try to produce messages
	// as fast as possible to keep latency low.

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	config.Producer.Retry.Max = 10                   // Retry up to 10 times to produce the message
	config.Producer.Return.Successes = true

	// On the broker side, you may want to change the following settings to get
	// stronger consistency guarantees:
	// - For your broker, set `unclean.leader.election.enable` to false
	// - For the topic, you could increase `min.insync.replicas`.

	producer, err := sarama.NewSyncProducer(prod.Brokers, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}

	prod.SyncProducer = producer

	return producer
}

func (prod *KafkaProducer) Close() error {
	if err := prod.SyncProducer.Close(); err != nil {
		log.Println("Failed to shut down data collector cleanly", err)
	}
	return nil
}

func (prod *KafkaProducer) Send(topic string, message string) {

	producerMessage := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := prod.SyncProducer.SendMessage(producerMessage)

	if err != nil {
		fmt.Printf("Failed to store your data:%s, err %s", message, err)
	} else {
		fmt.Printf("Your data is stored with unique identifier important/%d/%d", partition, offset)
	}

}
