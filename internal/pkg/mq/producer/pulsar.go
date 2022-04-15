package producer

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"time"
)

type PulsarProducer struct {
	pro pulsar.Producer
}

func NewPulsarProducer(topic string, cli pulsar.Client) (*PulsarProducer, error) {
	pul := &PulsarProducer{}
	// producer
	producer, err := cli.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
	})
	if err != nil {
		return nil, err
	}
	pul.pro = producer
	return pul, nil
}

func (p *PulsarProducer) SendMsg(ctx context.Context, msg []byte) error {
	_, err := p.pro.Send(ctx, &pulsar.ProducerMessage{
		Payload: msg,
	})
	return err
}

func (p *PulsarProducer) DelayAtSendMsg(ctx context.Context, msg []byte, date *time.Time) error {
	_, err := p.pro.Send(ctx, &pulsar.ProducerMessage{
		Payload:   msg,
		DeliverAt: *date,
	})
	return err
}

func (p *PulsarProducer) DelayAfterSendMsg(ctx context.Context, msg []byte, dur time.Duration) error {
	_, err := p.pro.Send(ctx, &pulsar.ProducerMessage{
		Payload:      msg,
		DeliverAfter: dur,
	})
	return err
}

func (p *PulsarProducer) CloseFunc() {
	p.pro.Close()
}
