package pubsub_ext

import (
	"context"
	"log"
	"sync/atomic"
	"time"

	"cloud.google.com/go/pubsub"
)

type PubsubWrapper struct {
	c *pubsub.Client

	counter int64

	ctx context.Context

	cancel context.CancelFunc

	watchInterval time.Duration
}
type PubsubWrapperOptions struct {
	//better not less than 5sec
	WatchInterval time.Duration
}

func NewPubsubWrapper(c *pubsub.Client, opts ...PubsubWrapperOptions) *PubsubWrapper {

	ctx, cancel := context.WithCancel(context.Background())

	w := &PubsubWrapper{

		ctx: ctx,

		cancel: cancel,

		c: c,
	}
	for _, v := range opts {
		w.watchInterval = v.WatchInterval
	}
	if w.watchInterval == 0 {
		w.watchInterval = 10 * time.Second
	}

	return w
}

// Receive is used for receiving subscription message. if u willing to keep the program alive when receive is canceled u need to use ur own context on f.
func (p *PubsubWrapper) Receive(subs *pubsub.Subscription, f func(context.Context, *pubsub.Message)) error {

	return subs.Receive(p.ctx, func(ctx context.Context, m *pubsub.Message) {

		atomic.AddInt64(&p.counter, 1)

		f(ctx, m)

		atomic.AddInt64(&p.counter, -1)

	})

}

// close pubsub gracefully without timeout
func (p *PubsubWrapper) Close() {

	//cancelling the receiver

	log.Println("closing subscription receiver")

	p.cancel()

	log.Println("subscription receiver canceled")

	//wait counter until 0

	log.Println("waiting untill all task done")

	for p.counter != 0 {

		time.Sleep(p.watchInterval)

	}

	//close the pubsub client

	log.Println("closing pubsub client")

	p.c.Close()

	log.Println("pubsub client closed")

}
