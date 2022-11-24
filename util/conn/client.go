package conn

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Options for gRPC server connection.
type Options struct {
	ServerAddress string
	ServerPort    uint
	TLSCert       string
	HostName      string
}

// Client for gRPC server backend service
type Client struct {
	conn *grpc.ClientConn
}

// NewClient returns a new gRPC client.
func NewClient(opts *Options) (*Client, error) {
	dialOpts := []grpc.DialOption{}

	//是否註冊憑證
	if opts.TLSCert != "" {
		creds, err := credentials.NewClientTLSFromFile(opts.TLSCert, opts.HostName)
		if err != nil {
			return nil, err
		}
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(creds))
	} else {
		dialOpts = append(dialOpts, grpc.WithInsecure())
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", opts.ServerAddress, opts.ServerPort), dialOpts...)

	return &Client{conn: conn}, err
}

// Close closes gRPC connection.
func (c *Client) Close() error {
	return c.conn.Close()
}
