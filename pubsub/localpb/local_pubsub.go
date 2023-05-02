package localpb

import (
	"Food-Delivery/common"
	"Food-Delivery/pubsub"
	"context"
	"log"
	"sync"
)

// A pubsub run locally (in-memory)
// It has a queue (buffer channel) at it's core and many group of subscribers
// B/c we want to sand a message with a specific topic for many subscribers in a group can

type localPubSub struct {
	messageQueue chan *pubsub.Message
	mapChannel   map[pubsub.Topic][]chan *pubsub.Message // map[key]value
	locker       *sync.RWMutex
}

func NewPubSub() *localPubSub {
	pb := &localPubSub{
		messageQueue: make(chan *pubsub.Message, 1000),
		mapChannel:   make(map[pubsub.Topic][]chan *pubsub.Message),
		locker:       new(sync.RWMutex),
	}

	pb.run()

	return pb
}

func (ps *localPubSub) Publish(ctx context.Context, topic pubsub.Topic, data *pubsub.Message) error {
	data.SetChannel(topic)

	go func() {
		defer common.AppRecover()
		ps.messageQueue <- data
		log.Println("New event published:", data.String(), "with data", data.Data())
	}()
	return nil
}

func (ps *localPubSub) Subscribe(ctx context.Context, topic pubsub.Topic) (ch <-chan *pubsub.Message, close func()) {
	c := make(chan *pubsub.Message)

	ps.locker.Lock()

	if val, ok := ps.mapChannel[topic]; ok {
		val = append(ps.mapChannel[topic], c)
		ps.mapChannel[topic] = val
	} else {
		ps.mapChannel[topic] = []chan *pubsub.Message{c}
	}

	ps.locker.Unlock()

	return c, func() {
		log.Println("Unsubscribe")

		if chans, ok := ps.mapChannel[topic]; ok {
			for i := range chans {
				if chans[i] == c {
					// rm element at index in chans
					chans = append(chans[:i], chans[i+1:]...)

					ps.locker.Lock()
					ps.mapChannel[topic] = chans
					ps.locker.Unlock()
					break
				}
			}
		}
	}

}

func (ps *localPubSub) run() error {
	log.Println("Pubsub started")

	go func() {
		defer common.AppRecover()
		for {
			mess := <-ps.messageQueue // rút message
			log.Println("Message dequeue", mess.String())

			if subs, ok := ps.mapChannel[mess.Channel()]; ok {
				for i := range subs {
					go func(c chan *pubsub.Message) {
						defer common.AppRecover()
						c <- mess
					}(subs[i])
				}
			}
		}
	}()

	return nil
}
