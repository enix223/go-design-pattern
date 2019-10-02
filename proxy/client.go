package proxy

import "fmt"

// Client which need to access the subject instance
// but we need more control to the access right to
// the subject instance, so we need a proxy.
type Client struct {
	sub Subject
}

// NewClientWithSubject create client with holding reference
// to the subject instance directly
func NewClientWithSubject() *Client {
	return &Client{
		sub: &RealSubject{},
	}
}

// NewClientWithReadableProxy create client with readable
// access right proxy to the subject instance
func NewClientWithReadableProxy() *Client {
	return &Client{
		sub: &Proxy{
			access:  "r",
			subject: &RealSubject{},
		},
	}
}

// NewClientWithNonReadableProxy create client with non-readable
// access right proxy to the subject instance
func NewClientWithNonReadableProxy() *Client {
	return &Client{
		sub: &Proxy{
			access:  "",
			subject: &RealSubject{},
		},
	}
}

// Run execute subject operations
func (c *Client) Run() {
	if err := c.sub.Operation(); err != nil {
		fmt.Println(err)
	}
}
