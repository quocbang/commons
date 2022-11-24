package conn

import (
	"errors"

	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/mes"
)

// Mes implements backend services with Mes.
type Mes struct {
	client *Client
}

// NewMes returns a new instance of Mes services.
func NewMes() *Mes {
	return &Mes{}
}

// Name returns the name of the services implementation.
func (s *Mes) Name() string {
	return "Mes"
}

// Connect establishes backend connections.
func (s *Mes) Connect(opts *Options) error {
	if s.client != nil {
		return errors.New("conn error:Mes client already initialized")
	}

	client, err := NewClient(opts)
	if err != nil {
		return err
	}
	s.client = client

	return nil
}

// Close closes client connections to all backends.
func (s *Mes) Close() error {
	return s.client.Close()
}

// API returns API service client.
func (s *Mes) API() mes.APIClient {
	return mes.NewAPIClient(s.client.conn)
}
