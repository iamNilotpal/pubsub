package pubsub

import "sync"

// Message represents a published message in the PubSub system.
// It contains the topic and the message content.
type Message struct {
	Topic   string // Topic to which the message belongs.
	Message string // Content of the message.
}

// PubSub implements a basic publish-subscribe system.
// It maintains topic-based subscriptions and allows messages to be published
// to subscribers listening to specific topics.
type PubSub struct {
	closed        bool                       // Indicates if the PubSub system is closed.
	config        *config                    // Configuration settings for the PubSub system.
	mu            sync.RWMutex               // Mutex to synchronize access to subscriptions.
	subscriptions map[string][]chan *Message // Map of topic subscriptions to channels.
}

// config holds the configuration settings for the PubSub system.
// It defines parameters such as channel size for message buffering.
type config struct {
	channelSize int // Defines the buffer size for subscriber channels.
}

// Option defines an interface for applying configuration options
// to the PubSub system.
type Option interface {
	apply(*config) // Method to apply the option to the given configuration.
}

// channelSizeOption is a specific implementation of Option,
// used to configure the size of the channels for subscribers.
type channelSizeOption struct {
	size int // Buffer size for subscriber channels.
}

// apply modifies the configuration by setting the channel size.
func (c channelSizeOption) apply(cfg *config) {
	if c.size > 0 {
		cfg.channelSize = c.size
	}
}

// WithChannelSize provides a functional option to set the size
// of the subscriber channels in the PubSub system.
func WithChannelSize(size int) Option {
	return channelSizeOption{size: size}
}
