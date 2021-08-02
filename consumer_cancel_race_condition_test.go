package cony

import (
	"testing"

	"github.com/streadway/amqp"
	// "github.com/peczenyj/cony"
)

type mockMQDeleter struct{}

func (m *mockMQDeleter) deleteConsumer(_ *Consumer) {}

func (m *mockMQDeleter) deletePublisher(_ *Publisher) {}

type mockMQChannel struct {
	deliveries chan amqp.Delivery
}

func (m *mockMQChannel) Consume(string, string, bool, bool, bool, bool, amqp.Table) (<-chan amqp.Delivery, error) {
	return m.deliveries, nil
}

func (m *mockMQChannel) Close() error {
	return nil
}
func (m *mockMQChannel) NotifyClose(chan *amqp.Error) chan *amqp.Error {
	return nil
}

func (m *mockMQChannel) Publish(string, string, bool, bool, amqp.Publishing) error {
	return nil
}

func (m *mockMQChannel) Qos(int, int, bool) error {
	return nil
}

func TestPanicWhenCloseCustomer(t *testing.T) {
	deliveries := make(chan amqp.Delivery)

	go func() {
		deliveries <- amqp.Delivery{}
		deliveries <- amqp.Delivery{}
		deliveries <- amqp.Delivery{}
		deliveries <- amqp.Delivery{}
		deliveries <- amqp.Delivery{}
	}()

	q := &Queue{}
	consumer := NewConsumer(q)

	client := &mockMQDeleter{}
	ch := &mockMQChannel{deliveries}

	x := make(chan struct{})
	go func() {
		for d := range consumer.Deliveries() {
			_ = d
		}
	}()
	go func() {
		consumer.serve(client, ch)
		close(x)
	}()

	consumer.Cancel()

	<-x
}
