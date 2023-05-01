package main

import (
	"Food-Delivery/pubsub"
	"Food-Delivery/pubsub/localpb"
	"context"
	"log"
	"time"
)

func main() {
	var localPS pubsub.Pubsub = localpb.NewPubSub()

	var topic pubsub.Topic = "OrdersCreated"

	sub1, close1 := localPS.Subscribe(context.Background(), topic)
	sub2, close2 := localPS.Subscribe(context.Background(), topic)

	localPS.Publish(context.Background(), topic, pubsub.NewMessage(1))
	localPS.Publish(context.Background(), topic, pubsub.NewMessage(2))

	go func() {
		for {
			log.Println("Sub 1:", (<-sub1).Data())
			time.Sleep(time.Millisecond * 400)
		}
	}()

	go func() {
		for {
			log.Println("Sub 2:", (<-sub2).Data())
			time.Sleep(time.Millisecond * 400)
		}
	}()

	time.Sleep(time.Second * 3)
	close1()
	close2()
	localPS.Publish(context.Background(), topic, pubsub.NewMessage(3))
	time.Sleep(time.Second * 2)
}
