package event_stream

import (
    "github.com/Shopify/sarama"
)

var (
    EventStreamConn sarama.SyncProducer
)

func Connect() {
    config := sarama.NewConfig()
    config.Producer.Return.Successes = true
    config.Producer.RequiredAcks = sarama.WaitForAll
    config.Producer.Retry.Max = 5
    brokers := []string{"kafka-service:9092"}

    conn, err := sarama.NewSyncProducer(brokers, config)
    if err != nil {
        panic(err.Error())
    }

    EventStreamConn = conn
}
