package config

import "flag"

var (
	BrokerConfig = &brokerConfig{}
)

func init() {
	BrokerConfig = newConfig()
}

type brokerConfig struct {
	brokerAddr *string
	queueLen   *int
	verbose    *bool
}

func newConfig() *brokerConfig {
	c := &brokerConfig{}
	c.brokerAddr = flag.String("broker", "127.0.0.1:9999", "address of broker to connect to")
	c.queueLen = flag.Int("qlen", 256, "size of an internal queue")
	c.verbose = flag.Bool("verbose", false, "configure verbosity")
	return c
}

func (c *brokerConfig) BrokerAddr() string {
	return *c.brokerAddr
}

func (c *brokerConfig) QueueLen() int {
	return *c.queueLen
}

func (c *brokerConfig) Verbose() bool {
	return *c.verbose
}
