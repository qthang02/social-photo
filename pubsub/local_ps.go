package pubsub

import (
	"context"
	"log"
	"social-photo/common"
	"sync"
)

// A pb run locally (in-mem)
// It has a queue (buffer channel) at it's core and many group of subscribers.
// Because we want to send a message with a specific topic for many subscribers in a group can handle.

type localPubSub struct {
	messageQueue chan *Message
	mapChannel   map[Topic][]chan *Message
	locker       *sync.RWMutex
}

func NewPubSub() *localPubSub {
	pb := &localPubSub{
		messageQueue: make(chan *Message, 10000),
		mapChannel:   make(map[Topic][]chan *Message),
		locker:       new(sync.RWMutex),
	}

	err := pb.run()
	if err != nil {
		return nil
	}

	return pb
}

func (ps *localPubSub) Publish(ctx context.Context, topic Topic, data *Message) error {
	data.SetChannel(topic)

	go func() {
		defer common.Recovery()
		ps.messageQueue <- data
		log.Println("New message published:", data.String())
	}()

	return nil
}

func (ps *localPubSub) Subscribe(ctx context.Context, topic Topic) (ch <-chan *Message, unsubscribe func()) {
	c := make(chan *Message)

	ps.locker.Lock()

	if val, ok := ps.mapChannel[topic]; ok {
		val = append(ps.mapChannel[topic], c)
		ps.mapChannel[topic] = val
	} else {
		ps.mapChannel[topic] = []chan *Message{c}
	}

	ps.locker.Unlock()

	return c, func() {
		log.Println("Unsubscribe")

		if chans, ok := ps.mapChannel[topic]; ok {
			for i := range chans {
				if chans[i] == c {
					// remove element at index in chans
					// [1,2,3,4,5] //  i = 3
					// [1,2,3] (arr[:i])
					// [5] (arr[i+1:])
					// [1,2,3,5]
					chans = append(chans[:i], chans[i+1:]...)

					ps.locker.Lock()
					ps.mapChannel[topic] = chans
					ps.locker.Unlock()

					close(c)
					break
				}
			}
		}
	}

}

func (ps *localPubSub) run() error {
	go func() {
		defer common.Recovery()
		for {
			mess := <-ps.messageQueue
			log.Println("Message dequeue:", mess.String())

			ps.locker.RLock()

			if subs, ok := ps.mapChannel[mess.Channel()]; ok {
				for i := range subs {
					go func(c chan *Message) {
						defer common.Recovery()
						c <- mess
					}(subs[i])
				}
			}

			ps.locker.RUnlock()
		}
	}()

	return nil
}
