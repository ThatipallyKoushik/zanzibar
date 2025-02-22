// Code generated by zanzibar
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package workflow

import (
	"context"
	"net/textproto"
	"strconv"

	"github.com/uber/zanzibar/config"

	zanzibar "github.com/uber/zanzibar/runtime"

	clientsIDlClientsBazBase "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients-idl/clients/baz/base"
	clientsIDlClientsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients-idl/clients/baz/baz"
	endpointsIDlEndpointsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints-idl/endpoints/baz/baz"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/baz/module"
	"go.uber.org/zap"
)

// SimpleServiceTransHeadersNoReqWorkflow defines the interface for SimpleServiceTransHeadersNoReq workflow
type SimpleServiceTransHeadersNoReqWorkflow interface {
	Handle(
		ctx context.Context,
		reqHeaders zanzibar.Header,
	) (context.Context, *endpointsIDlEndpointsBazBaz.TransHeader, zanzibar.Header, error)
}

// NewSimpleServiceTransHeadersNoReqWorkflow creates a workflow
func NewSimpleServiceTransHeadersNoReqWorkflow(deps *module.Dependencies) SimpleServiceTransHeadersNoReqWorkflow {
	var whitelistedDynamicHeaders []string
	if deps.Default.Config.ContainsKey("clients.baz.alternates") {
		var alternateServiceDetail config.AlternateServiceDetail
		deps.Default.Config.MustGetStruct("clients.baz.alternates", &alternateServiceDetail)
		for _, routingConfig := range alternateServiceDetail.RoutingConfigs {
			whitelistedDynamicHeaders = append(whitelistedDynamicHeaders, textproto.CanonicalMIMEHeaderKey(routingConfig.HeaderName))
		}
	}

	return &simpleServiceTransHeadersNoReqWorkflow{
		Clients:                   deps.Client,
		Logger:                    deps.Default.Logger,
		whitelistedDynamicHeaders: whitelistedDynamicHeaders,
	}
}

// simpleServiceTransHeadersNoReqWorkflow calls thrift client Baz.TransHeadersNoReq
type simpleServiceTransHeadersNoReqWorkflow struct {
	Clients                   *module.ClientDependencies
	Logger                    *zap.Logger
	whitelistedDynamicHeaders []string
}

// Handle calls thrift client.
func (w simpleServiceTransHeadersNoReqWorkflow) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
) (context.Context, *endpointsIDlEndpointsBazBaz.TransHeader, zanzibar.Header, error) {
	clientRequest := propagateHeadersTransHeadersNoReqClientRequests(nil, reqHeaders)

	clientHeaders := map[string]string{}

	var ok bool
	var h string
	var k string

	k = textproto.CanonicalMIMEHeaderKey("x-uber-foo")
	h, ok = reqHeaders.Get(k)
	if ok {
		clientHeaders[k] = h
	}
	k = textproto.CanonicalMIMEHeaderKey("x-uber-bar")
	h, ok = reqHeaders.Get(k)
	if ok {
		clientHeaders[k] = h
	}

	h, ok = reqHeaders.Get("B3")
	if ok {
		clientHeaders["B3"] = h
	}
	h, ok = reqHeaders.Get("I2")
	if ok {
		clientHeaders["I2"] = h
	}
	h, ok = reqHeaders.Get("S1")
	if ok {
		clientHeaders["S1"] = h
	}
	h, ok = reqHeaders.Get("X-Deputy-Forwarded")
	if ok {
		clientHeaders["X-Deputy-Forwarded"] = h
	}
	for _, whitelistedHeader := range w.whitelistedDynamicHeaders {
		headerVal, ok := reqHeaders.Get(whitelistedHeader)
		if ok {
			clientHeaders[whitelistedHeader] = headerVal
		}
	}

	ctx, clientRespBody, cliRespHeaders, err := w.Clients.Baz.TransHeadersNoReq(
		ctx, clientHeaders, clientRequest,
	)

	if err != nil {
		switch errValue := err.(type) {

		case *clientsIDlClientsBazBaz.AuthErr:
			serverErr := convertTransHeadersNoReqAuthErr(
				errValue,
			)

			return ctx, nil, nil, serverErr

		default:
			w.Logger.Warn("Client failure: could not make client request",
				zap.Error(errValue),
				zap.String("client", "Baz"),
			)

			return ctx, nil, nil, err

		}
	}

	// Filter and map response headers from client to server response.
	resHeaders := zanzibar.ServerHTTPHeader{}

	response := convertSimpleServiceTransHeadersNoReqClientResponse(clientRespBody)
	if val, ok := cliRespHeaders[zanzibar.ClientResponseDurationKey]; ok {
		resHeaders.Set(zanzibar.ClientResponseDurationKey, val)
	}

	resHeaders.Set(zanzibar.ClientTypeKey, "tchannel")
	return ctx, response, resHeaders, nil
}

func convertTransHeadersNoReqAuthErr(
	clientError *clientsIDlClientsBazBaz.AuthErr,
) *endpointsIDlEndpointsBazBaz.AuthErr {
	// TODO: Add error fields mapping here.
	serverError := &endpointsIDlEndpointsBazBaz.AuthErr{}
	return serverError
}

func convertSimpleServiceTransHeadersNoReqClientResponse(in *clientsIDlClientsBazBase.TransHeaders) *endpointsIDlEndpointsBazBaz.TransHeader {
	out := &endpointsIDlEndpointsBazBaz.TransHeader{}

	return out
}

func propagateHeadersTransHeadersNoReqClientRequests(in *clientsIDlClientsBazBaz.SimpleService_TransHeadersNoReq_Args, headers zanzibar.Header) *clientsIDlClientsBazBaz.SimpleService_TransHeadersNoReq_Args {
	if in == nil {
		in = &clientsIDlClientsBazBaz.SimpleService_TransHeadersNoReq_Args{}
	}
	if key, ok := headers.Get("b3"); ok {
		if v, err := strconv.ParseBool(key); err == nil {
			in.B4 = &v
		}

	}
	if key, ok := headers.Get("i2"); ok {
		if v, err := strconv.ParseInt(key, 10, 32); err == nil {
			val := int32(v)
			in.I3 = val
		}

	}
	if key, ok := headers.Get("i2"); ok {
		if in.Req == nil {
			in.Req = &clientsIDlClientsBazBase.NestedStruct{}
		}
		if v, err := strconv.ParseInt(key, 10, 32); err == nil {
			val := int32(v)
			in.Req.Check = &val
		}

	}
	if key, ok := headers.Get("s1"); ok {
		if in.Req == nil {
			in.Req = &clientsIDlClientsBazBase.NestedStruct{}
		}
		in.Req.Msg = key

	}
	if key, ok := headers.Get("s1"); ok {
		in.S2 = &key

	}
	return in
}
