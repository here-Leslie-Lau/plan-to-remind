package consumer

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
)

type PulsarConsumer struct {
	con pulsar.Consumer
}

func (p *PulsarConsumer) Consume(ctx context.Context) ([]byte, error) {
	msg, err := p.con.Receive(ctx)
	if err != nil {
		return nil, err
	}
	return msg.Payload(), err
}

func (p *PulsarConsumer) CloseFunc() {
	p.con.Close()
}

func NewPulsarConsumer(topic, subName string, cli pulsar.Client) (*PulsarConsumer, error) {
	con, err := cli.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: subName,
		Type:             pulsar.Shared,
	})
	if err != nil {
		return nil, err
	}
	return &PulsarConsumer{con: con}, nil
}
