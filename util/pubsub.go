package util

import (
	"sync"
)

type Message struct {
	Message string `json:"message"`
}

type PubSubChannel struct {
	subscribers map[string]chan Message
	mu          sync.RWMutex
	closed      bool
}

// This pubsub implementation is shit, the point of this package is to get familiar with datastar.
// In short, call Subscribe() to get a channel you can read messages from.
// This obviously falls apart when distributed across multiple machines.
func NewPubSub() *PubSubChannel {
	return &PubSubChannel{
		subscribers: make(map[string]chan Message),
	}
}

// Subscribe creates a new subscription and returns a channel to receive messages
func (c *PubSubChannel) Subscribe(subscriberID string) <-chan Message {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		ch := make(chan Message)
		close(ch)
		return ch
	}

	ch := make(chan Message, 100)
	c.subscribers[subscriberID] = ch
	return ch
}

// Unsubscribe removes a subscriber and closes their channel
func (c *PubSubChannel) Unsubscribe(subscriberID string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if ch, exists := c.subscribers[subscriberID]; exists {
		close(ch)
		delete(c.subscribers, subscriberID)
	}
}

// PushMessage broadcasts a message to all subscribers
func (c *PubSubChannel) PushMessage(content string) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.closed {
		return
	}

	message := Message{Message: content}

	// Send to all subscribers (non-blocking)
	for subscriberID, ch := range c.subscribers {
		select {
		case ch <- message:
		default:

			go func(id string) {
				c.Unsubscribe(id)
			}(subscriberID)
		}
	}
}
