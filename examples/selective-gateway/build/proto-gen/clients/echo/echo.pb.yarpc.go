// Code generated by protoc-gen-yarpc-go. DO NOT EDIT.
// source: clients/echo/echo.proto

package echo

import (
	"context"
	"io/ioutil"
	"reflect"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"go.uber.org/fx"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/api/x/restriction"
	"go.uber.org/yarpc/encoding/protobuf"
	"go.uber.org/yarpc/encoding/protobuf/reflection"
)

var _ = ioutil.NopCloser

// EchoYARPCClient is the YARPC client-side interface for the Echo service.
type EchoYARPCClient interface {
	Echo(context.Context, *Request, ...yarpc.CallOption) (*Response, error)
}

func newEchoYARPCClient(clientConfig transport.ClientConfig, anyResolver jsonpb.AnyResolver, options ...protobuf.ClientOption) EchoYARPCClient {
	return &_EchoYARPCCaller{protobuf.NewStreamClient(
		protobuf.ClientParams{
			ServiceName:  "echo.Echo",
			ClientConfig: clientConfig,
			AnyResolver:  anyResolver,
			Options:      options,
		},
	)}
}

// NewEchoYARPCClient builds a new YARPC client for the Echo service.
func NewEchoYARPCClient(clientConfig transport.ClientConfig, options ...protobuf.ClientOption) EchoYARPCClient {
	return newEchoYARPCClient(clientConfig, nil, options...)
}

// EchoYARPCServer is the YARPC server-side interface for the Echo service.
type EchoYARPCServer interface {
	Echo(context.Context, *Request) (*Response, error)
}

type buildEchoYARPCProceduresParams struct {
	Server      EchoYARPCServer
	AnyResolver jsonpb.AnyResolver
}

func buildEchoYARPCProcedures(params buildEchoYARPCProceduresParams) []transport.Procedure {
	handler := &_EchoYARPCHandler{params.Server}
	return protobuf.BuildProcedures(
		protobuf.BuildProceduresParams{
			ServiceName: "echo.Echo",
			UnaryHandlerParams: []protobuf.BuildProceduresUnaryHandlerParams{
				{
					MethodName: "Echo",
					Handler: protobuf.NewUnaryHandler(
						protobuf.UnaryHandlerParams{
							Handle:      handler.Echo,
							NewRequest:  newEchoServiceEchoYARPCRequest,
							AnyResolver: params.AnyResolver,
						},
					),
				},
			},
			OnewayHandlerParams: []protobuf.BuildProceduresOnewayHandlerParams{},
			StreamHandlerParams: []protobuf.BuildProceduresStreamHandlerParams{},
		},
	)
}

// BuildEchoYARPCProcedures prepares an implementation of the Echo service for YARPC registration.
func BuildEchoYARPCProcedures(server EchoYARPCServer) []transport.Procedure {
	return buildEchoYARPCProcedures(buildEchoYARPCProceduresParams{Server: server})
}

// FxEchoYARPCClientParams defines the input
// for NewFxEchoYARPCClient. It provides the
// paramaters to get a EchoYARPCClient in an
// Fx application.
type FxEchoYARPCClientParams struct {
	fx.In

	Provider    yarpc.ClientConfig
	AnyResolver jsonpb.AnyResolver  `name:"yarpcfx" optional:"true"`
	Restriction restriction.Checker `optional:"true"`
}

// FxEchoYARPCClientResult defines the output
// of NewFxEchoYARPCClient. It provides a
// EchoYARPCClient to an Fx application.
type FxEchoYARPCClientResult struct {
	fx.Out

	Client EchoYARPCClient

	// We are using an fx.Out struct here instead of just returning a client
	// so that we can add more values or add named versions of the client in
	// the future without breaking any existing code.
}

// NewFxEchoYARPCClient provides a EchoYARPCClient
// to an Fx application using the given name for routing.
//
//  fx.Provide(
//    echo.NewFxEchoYARPCClient("service-name"),
//    ...
//  )
func NewFxEchoYARPCClient(name string, options ...protobuf.ClientOption) interface{} {
	return func(params FxEchoYARPCClientParams) FxEchoYARPCClientResult {
		cc := params.Provider.ClientConfig(name)

		if params.Restriction != nil {
			if namer, ok := cc.GetUnaryOutbound().(transport.Namer); ok {
				if err := params.Restriction.Check(protobuf.Encoding, namer.TransportName()); err != nil {
					panic(err.Error())
				}
			}
		}

		return FxEchoYARPCClientResult{
			Client: newEchoYARPCClient(cc, params.AnyResolver, options...),
		}
	}
}

// FxEchoYARPCProceduresParams defines the input
// for NewFxEchoYARPCProcedures. It provides the
// paramaters to get EchoYARPCServer procedures in an
// Fx application.
type FxEchoYARPCProceduresParams struct {
	fx.In

	Server      EchoYARPCServer
	AnyResolver jsonpb.AnyResolver `name:"yarpcfx" optional:"true"`
}

// FxEchoYARPCProceduresResult defines the output
// of NewFxEchoYARPCProcedures. It provides
// EchoYARPCServer procedures to an Fx application.
//
// The procedures are provided to the "yarpcfx" value group.
// Dig 1.2 or newer must be used for this feature to work.
type FxEchoYARPCProceduresResult struct {
	fx.Out

	Procedures     []transport.Procedure `group:"yarpcfx"`
	ReflectionMeta reflection.ServerMeta `group:"yarpcfx"`
}

// NewFxEchoYARPCProcedures provides EchoYARPCServer procedures to an Fx application.
// It expects a EchoYARPCServer to be present in the container.
//
//  fx.Provide(
//    echo.NewFxEchoYARPCProcedures(),
//    ...
//  )
func NewFxEchoYARPCProcedures() interface{} {
	return func(params FxEchoYARPCProceduresParams) FxEchoYARPCProceduresResult {
		return FxEchoYARPCProceduresResult{
			Procedures: buildEchoYARPCProcedures(buildEchoYARPCProceduresParams{
				Server:      params.Server,
				AnyResolver: params.AnyResolver,
			}),
			ReflectionMeta: reflection.ServerMeta{
				ServiceName:     "echo.Echo",
				FileDescriptors: yarpcFileDescriptorClosure6caa5dd77ae15c91,
			},
		}
	}
}

type _EchoYARPCCaller struct {
	streamClient protobuf.StreamClient
}

func (c *_EchoYARPCCaller) Echo(ctx context.Context, request *Request, options ...yarpc.CallOption) (*Response, error) {
	responseMessage, err := c.streamClient.Call(ctx, "Echo", request, newEchoServiceEchoYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*Response)
	if !ok {
		return nil, protobuf.CastError(emptyEchoServiceEchoYARPCResponse, responseMessage)
	}
	return response, err
}

type _EchoYARPCHandler struct {
	server EchoYARPCServer
}

func (h *_EchoYARPCHandler) Echo(ctx context.Context, requestMessage proto.Message) (proto.Message, error) {
	var request *Request
	var ok bool
	if requestMessage != nil {
		request, ok = requestMessage.(*Request)
		if !ok {
			return nil, protobuf.CastError(emptyEchoServiceEchoYARPCRequest, requestMessage)
		}
	}
	response, err := h.server.Echo(ctx, request)
	if response == nil {
		return nil, err
	}
	return response, err
}

func newEchoServiceEchoYARPCRequest() proto.Message {
	return &Request{}
}

func newEchoServiceEchoYARPCResponse() proto.Message {
	return &Response{}
}

var (
	emptyEchoServiceEchoYARPCRequest  = &Request{}
	emptyEchoServiceEchoYARPCResponse = &Response{}
)

// EchoInternalYARPCClient is the YARPC client-side interface for the EchoInternal service.
type EchoInternalYARPCClient interface {
	Echo(context.Context, *Request, ...yarpc.CallOption) (*Response, error)
}

func newEchoInternalYARPCClient(clientConfig transport.ClientConfig, anyResolver jsonpb.AnyResolver, options ...protobuf.ClientOption) EchoInternalYARPCClient {
	return &_EchoInternalYARPCCaller{protobuf.NewStreamClient(
		protobuf.ClientParams{
			ServiceName:  "echo.EchoInternal",
			ClientConfig: clientConfig,
			AnyResolver:  anyResolver,
			Options:      options,
		},
	)}
}

// NewEchoInternalYARPCClient builds a new YARPC client for the EchoInternal service.
func NewEchoInternalYARPCClient(clientConfig transport.ClientConfig, options ...protobuf.ClientOption) EchoInternalYARPCClient {
	return newEchoInternalYARPCClient(clientConfig, nil, options...)
}

// EchoInternalYARPCServer is the YARPC server-side interface for the EchoInternal service.
type EchoInternalYARPCServer interface {
	Echo(context.Context, *Request) (*Response, error)
}

type buildEchoInternalYARPCProceduresParams struct {
	Server      EchoInternalYARPCServer
	AnyResolver jsonpb.AnyResolver
}

func buildEchoInternalYARPCProcedures(params buildEchoInternalYARPCProceduresParams) []transport.Procedure {
	handler := &_EchoInternalYARPCHandler{params.Server}
	return protobuf.BuildProcedures(
		protobuf.BuildProceduresParams{
			ServiceName: "echo.EchoInternal",
			UnaryHandlerParams: []protobuf.BuildProceduresUnaryHandlerParams{
				{
					MethodName: "Echo",
					Handler: protobuf.NewUnaryHandler(
						protobuf.UnaryHandlerParams{
							Handle:      handler.Echo,
							NewRequest:  newEchoInternalServiceEchoYARPCRequest,
							AnyResolver: params.AnyResolver,
						},
					),
				},
			},
			OnewayHandlerParams: []protobuf.BuildProceduresOnewayHandlerParams{},
			StreamHandlerParams: []protobuf.BuildProceduresStreamHandlerParams{},
		},
	)
}

// BuildEchoInternalYARPCProcedures prepares an implementation of the EchoInternal service for YARPC registration.
func BuildEchoInternalYARPCProcedures(server EchoInternalYARPCServer) []transport.Procedure {
	return buildEchoInternalYARPCProcedures(buildEchoInternalYARPCProceduresParams{Server: server})
}

// FxEchoInternalYARPCClientParams defines the input
// for NewFxEchoInternalYARPCClient. It provides the
// paramaters to get a EchoInternalYARPCClient in an
// Fx application.
type FxEchoInternalYARPCClientParams struct {
	fx.In

	Provider    yarpc.ClientConfig
	AnyResolver jsonpb.AnyResolver  `name:"yarpcfx" optional:"true"`
	Restriction restriction.Checker `optional:"true"`
}

// FxEchoInternalYARPCClientResult defines the output
// of NewFxEchoInternalYARPCClient. It provides a
// EchoInternalYARPCClient to an Fx application.
type FxEchoInternalYARPCClientResult struct {
	fx.Out

	Client EchoInternalYARPCClient

	// We are using an fx.Out struct here instead of just returning a client
	// so that we can add more values or add named versions of the client in
	// the future without breaking any existing code.
}

// NewFxEchoInternalYARPCClient provides a EchoInternalYARPCClient
// to an Fx application using the given name for routing.
//
//  fx.Provide(
//    echo.NewFxEchoInternalYARPCClient("service-name"),
//    ...
//  )
func NewFxEchoInternalYARPCClient(name string, options ...protobuf.ClientOption) interface{} {
	return func(params FxEchoInternalYARPCClientParams) FxEchoInternalYARPCClientResult {
		cc := params.Provider.ClientConfig(name)

		if params.Restriction != nil {
			if namer, ok := cc.GetUnaryOutbound().(transport.Namer); ok {
				if err := params.Restriction.Check(protobuf.Encoding, namer.TransportName()); err != nil {
					panic(err.Error())
				}
			}
		}

		return FxEchoInternalYARPCClientResult{
			Client: newEchoInternalYARPCClient(cc, params.AnyResolver, options...),
		}
	}
}

// FxEchoInternalYARPCProceduresParams defines the input
// for NewFxEchoInternalYARPCProcedures. It provides the
// paramaters to get EchoInternalYARPCServer procedures in an
// Fx application.
type FxEchoInternalYARPCProceduresParams struct {
	fx.In

	Server      EchoInternalYARPCServer
	AnyResolver jsonpb.AnyResolver `name:"yarpcfx" optional:"true"`
}

// FxEchoInternalYARPCProceduresResult defines the output
// of NewFxEchoInternalYARPCProcedures. It provides
// EchoInternalYARPCServer procedures to an Fx application.
//
// The procedures are provided to the "yarpcfx" value group.
// Dig 1.2 or newer must be used for this feature to work.
type FxEchoInternalYARPCProceduresResult struct {
	fx.Out

	Procedures     []transport.Procedure `group:"yarpcfx"`
	ReflectionMeta reflection.ServerMeta `group:"yarpcfx"`
}

// NewFxEchoInternalYARPCProcedures provides EchoInternalYARPCServer procedures to an Fx application.
// It expects a EchoInternalYARPCServer to be present in the container.
//
//  fx.Provide(
//    echo.NewFxEchoInternalYARPCProcedures(),
//    ...
//  )
func NewFxEchoInternalYARPCProcedures() interface{} {
	return func(params FxEchoInternalYARPCProceduresParams) FxEchoInternalYARPCProceduresResult {
		return FxEchoInternalYARPCProceduresResult{
			Procedures: buildEchoInternalYARPCProcedures(buildEchoInternalYARPCProceduresParams{
				Server:      params.Server,
				AnyResolver: params.AnyResolver,
			}),
			ReflectionMeta: reflection.ServerMeta{
				ServiceName:     "echo.EchoInternal",
				FileDescriptors: yarpcFileDescriptorClosure6caa5dd77ae15c91,
			},
		}
	}
}

type _EchoInternalYARPCCaller struct {
	streamClient protobuf.StreamClient
}

func (c *_EchoInternalYARPCCaller) Echo(ctx context.Context, request *Request, options ...yarpc.CallOption) (*Response, error) {
	responseMessage, err := c.streamClient.Call(ctx, "Echo", request, newEchoInternalServiceEchoYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*Response)
	if !ok {
		return nil, protobuf.CastError(emptyEchoInternalServiceEchoYARPCResponse, responseMessage)
	}
	return response, err
}

type _EchoInternalYARPCHandler struct {
	server EchoInternalYARPCServer
}

func (h *_EchoInternalYARPCHandler) Echo(ctx context.Context, requestMessage proto.Message) (proto.Message, error) {
	var request *Request
	var ok bool
	if requestMessage != nil {
		request, ok = requestMessage.(*Request)
		if !ok {
			return nil, protobuf.CastError(emptyEchoInternalServiceEchoYARPCRequest, requestMessage)
		}
	}
	response, err := h.server.Echo(ctx, request)
	if response == nil {
		return nil, err
	}
	return response, err
}

func newEchoInternalServiceEchoYARPCRequest() proto.Message {
	return &Request{}
}

func newEchoInternalServiceEchoYARPCResponse() proto.Message {
	return &Response{}
}

var (
	emptyEchoInternalServiceEchoYARPCRequest  = &Request{}
	emptyEchoInternalServiceEchoYARPCResponse = &Response{}
)

var yarpcFileDescriptorClosure6caa5dd77ae15c91 = [][]byte{
	// clients/echo/echo.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0xce, 0xc9, 0x4c,
		0xcd, 0x2b, 0x29, 0xd6, 0x4f, 0x4d, 0xce, 0xc8, 0x07, 0x13, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9,
		0x42, 0x2c, 0x20, 0xb6, 0x92, 0x32, 0x17, 0x7b, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90,
		0x04, 0x17, 0x7b, 0x6e, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67,
		0x10, 0x8c, 0xab, 0xa4, 0xc2, 0xc5, 0x11, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x8a, 0x5b,
		0x95, 0x91, 0x2e, 0x17, 0x8b, 0x6b, 0x72, 0x46, 0xbe, 0x90, 0x2a, 0x94, 0xe6, 0xd5, 0x03, 0xdb,
		0x06, 0x35, 0x5e, 0x8a, 0x0f, 0xc6, 0x85, 0x18, 0x64, 0x64, 0xca, 0xc5, 0x03, 0x52, 0xe6, 0x99,
		0x57, 0x92, 0x5a, 0x94, 0x97, 0x98, 0x43, 0xa4, 0xb6, 0x24, 0x36, 0xb0, 0xeb, 0x8d, 0x01, 0x01,
		0x00, 0x00, 0xff, 0xff, 0xfb, 0xd2, 0x16, 0xa4, 0xd8, 0x00, 0x00, 0x00,
	},
}

func init() {
	yarpc.RegisterClientBuilder(
		func(clientConfig transport.ClientConfig, structField reflect.StructField) EchoYARPCClient {
			return NewEchoYARPCClient(clientConfig, protobuf.ClientBuilderOptions(clientConfig, structField)...)
		},
	)
	yarpc.RegisterClientBuilder(
		func(clientConfig transport.ClientConfig, structField reflect.StructField) EchoInternalYARPCClient {
			return NewEchoInternalYARPCClient(clientConfig, protobuf.ClientBuilderOptions(clientConfig, structField)...)
		},
	)
}
