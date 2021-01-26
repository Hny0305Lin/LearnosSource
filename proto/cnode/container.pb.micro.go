// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/cnode/container.proto

package proto_cnode_service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Node service

type NodeService interface {
	Service(ctx context.Context, in *CallOpt, opts ...client.CallOption) (*CallRsp, error)
	CreateMsg(ctx context.Context, in *CreateNotice, opts ...client.CallOption) (*CreateRsp, error)
}

type nodeService struct {
	c    client.Client
	name string
}

func NewNodeService(name string, c client.Client) NodeService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "proto.cnode.service"
	}
	return &nodeService{
		c:    c,
		name: name,
	}
}

func (c *nodeService) Service(ctx context.Context, in *CallOpt, opts ...client.CallOption) (*CallRsp, error) {
	req := c.c.NewRequest(c.name, "Node.Service", in)
	out := new(CallRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeService) CreateMsg(ctx context.Context, in *CreateNotice, opts ...client.CallOption) (*CreateRsp, error) {
	req := c.c.NewRequest(c.name, "Node.CreateMsg", in)
	out := new(CreateRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Node service

type NodeHandler interface {
	Service(context.Context, *CallOpt, *CallRsp) error
	CreateMsg(context.Context, *CreateNotice, *CreateRsp) error
}

func RegisterNodeHandler(s server.Server, hdlr NodeHandler, opts ...server.HandlerOption) error {
	type node interface {
		Service(ctx context.Context, in *CallOpt, out *CallRsp) error
		CreateMsg(ctx context.Context, in *CreateNotice, out *CreateRsp) error
	}
	type Node struct {
		node
	}
	h := &nodeHandler{hdlr}
	return s.Handle(s.NewHandler(&Node{h}, opts...))
}

type nodeHandler struct {
	NodeHandler
}

func (h *nodeHandler) Service(ctx context.Context, in *CallOpt, out *CallRsp) error {
	return h.NodeHandler.Service(ctx, in, out)
}

func (h *nodeHandler) CreateMsg(ctx context.Context, in *CreateNotice, out *CreateRsp) error {
	return h.NodeHandler.CreateMsg(ctx, in, out)
}
