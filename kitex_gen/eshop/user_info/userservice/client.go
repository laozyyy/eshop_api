// Code generated by Kitex v0.11.3. DO NOT EDIT.

package userservice

import (
	"context"
	user_info "eshop_api/kitex_gen/eshop/user_info"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetOneUser(ctx context.Context, request *user_info.GetOneUserRequest, callOptions ...callopt.Option) (r *user_info.GetOneUserResponse, err error)
	GetOneUserByName(ctx context.Context, name string, callOptions ...callopt.Option) (r *user_info.GetOneUserResponse, err error)
	InsertOneUser(ctx context.Context, user *user_info.User, callOptions ...callopt.Option) (r *user_info.InsertOneUserResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) GetOneUser(ctx context.Context, request *user_info.GetOneUserRequest, callOptions ...callopt.Option) (r *user_info.GetOneUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetOneUser(ctx, request)
}

func (p *kUserServiceClient) GetOneUserByName(ctx context.Context, name string, callOptions ...callopt.Option) (r *user_info.GetOneUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetOneUserByName(ctx, name)
}

func (p *kUserServiceClient) InsertOneUser(ctx context.Context, user *user_info.User, callOptions ...callopt.Option) (r *user_info.InsertOneUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.InsertOneUser(ctx, user)
}
