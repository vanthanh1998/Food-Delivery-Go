package pubsub

import "context"

type Topic string

type Pubsub interface {
	Publish(ctx context.Context, channel Topic, data *Message) error
	Subscribe(ctx context.Context, channel Topic) (ch <-chan *Message, close func())
	// *** return ch <-chan *Message
	// 1. Bên phía nhận chỉ được rút ra, k đươ nhét vào
	// 2. Tai sao dùng channels: vì channel có thể trym data từ đầu này sang đầu khác
	// 3. Nếu như 1 đầu này or đầu kia đang bận thì channel nó sẽ bị block lại
	// *** close func())
	// 1. Muốn unsubscribe gọi hàm này
	// 2. Func này call thì chỉ unsubscribe hàm của nó thui, k unsubscribe hết toàn bộ
}
