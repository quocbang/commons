package mail

import (
	"errors"
	"sync"
)

// Agent provides some function for the email service.
type Agent interface {
	Send(ms ...*Message) error
}

var (
	agent      Agent
	agentMutex sync.Mutex
)

// RegisterAgent registers the agent for the email service.
func RegisterAgent(a Agent) error {
	if agent == nil {
		agentMutex.Lock()
		defer agentMutex.Unlock()
		if agent == nil {
			agent = a
			return nil
		}
	}
	return errors.New("the email agent has been registered")
}

// A returns the email agent.
func A() Agent {
	return agent
}
