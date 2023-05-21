package client

import (
	"context"
	pb "ddk/test/api"
	"errors"
	"time"

	"google.golang.org/grpc"
)

var (
	ErrNoEndpoints = errors.New("client: no endpoints avaliable")
	ErrConnTimeout = errors.New("client: conn timeout")
	ErrParam       = errors.New("client: err param")
)

type Option func(client *TestClient)

type TestClient struct {
	conn     *grpc.ClientConn
	endPoint string
	timeout  time.Duration
	client   *pb.TestClient
}

func WithEndPoint(endpoint string) Option {
	return func(client *TestClient) {
		client.endPoint = endpoint
	}
}

func WithTimeOut(timeout int64) Option {
	return func(client *TestClient) {
		client.timeout = time.Duration(timeout) * time.Second
	}
}

func NewTestClient(ctx context.Context, opts ...Option) (*TestClient, error) {
	client := &TestClient{
		conn:     nil,
		endPoint: "",
		timeout:  0,
	}

	for _, o := range opts {
		o(client)
	}

	if client.endPoint == "" {
		return nil, ErrNoEndpoints
	}

	if client.timeout != 0 {
		ctx, _ = context.WithTimeout(ctx, client.timeout)
	}

	conn, err := grpc.DialContext(ctx, client.endPoint)
	if err != nil {
		return nil, err
	}

	client.conn = conn
	client.client = pb.NewTestClient(client.conn)

	return client, nil
}

func (client *TestClient) Put(context.Context, key, value string)