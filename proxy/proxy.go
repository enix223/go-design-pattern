package proxy

import (
	"errors"
	"fmt"
	"strings"
)

// Proxy wrapper of the subject to gain more access control
// to the subject instance
type Proxy struct {
	subject Subject
	access  string
}

// Operation implement the Subject interface
func (p *Proxy) Operation() error {
	if strings.Index(p.access, "r") == -1 {
		return errors.New("can not read")
	}
	fmt.Println("authorized to perform operation")
	p.subject.Operation()
	return nil
}
