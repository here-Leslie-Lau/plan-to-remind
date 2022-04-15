package mq

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"plan-to-remind/internal/pkg/mq/consumer"
	"plan-to-remind/internal/pkg/mq/producer"
	"sync"
	"testing"
)

func TestPulsar(t *testing.T) {
	cli, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://127.0.0.1:6650",
	})
	if err != nil {
		t.Fail()
		return
	}
	defer cli.Close()
	// producer
	pro, err := producer.NewPulsarProducer("demo", cli)
	if err != nil || pro == nil {
		t.Fail()
		return
	}
	defer pro.CloseFunc()
	ctx := context.Background()
	if err := pro.SendMsg(ctx, []byte("hello")); err != nil {
		t.Fail()
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	// consumer
	con, err := consumer.NewPulsarConsumer("demo", "consumer-one", cli)
	if err != nil || con == nil {
		t.Fail()
		return
	}
	defer con.CloseFunc()
	go func() {
		bytes, err := con.Consume(ctx)
		if err != nil {
			t.Fail()
		}
		fmt.Println("bytes:", string(bytes))
		wg.Done()
	}()

	wg.Wait()
}
