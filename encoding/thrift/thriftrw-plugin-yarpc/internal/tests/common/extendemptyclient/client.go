// Code generated by thriftrw-plugin-yarpc
// @generated

package extendemptyclient

import (
	context "context"
	wire "go.uber.org/thriftrw/wire"
	yarpc "go.uber.org/yarpc"
	transport "go.uber.org/yarpc/api/transport"
	thrift "go.uber.org/yarpc/encoding/thrift"
	common "go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/common"
	emptyserviceclient "go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/common/emptyserviceclient"
	reflect "reflect"
)

// Interface is a client for the ExtendEmpty service.
type Interface interface {
	emptyserviceclient.Interface

	Hello(
		ctx context.Context,
		opts ...yarpc.CallOption,
	) error
}

// New builds a new client for the ExtendEmpty service.
//
// 	client := extendemptyclient.New(dispatcher.ClientConfig("extendempty"))
func New(c transport.ClientConfig, opts ...thrift.ClientOption) Interface {
	return client{
		c: thrift.New(thrift.Config{
			Service:      "ExtendEmpty",
			ClientConfig: c,
		}, opts...),
		nwc: thrift.NewNoWire(thrift.Config{
			Service:      "ExtendEmpty",
			ClientConfig: c,
		}, opts...),

		Interface: emptyserviceclient.New(
			c,
			append(
				opts,
				thrift.Named("ExtendEmpty"),
			)...,
		),
	}
}

func init() {
	yarpc.RegisterClientBuilder(
		func(c transport.ClientConfig, f reflect.StructField) Interface {
			return New(c, thrift.ClientBuilderOptions(c, f)...)
		},
	)
}

type client struct {
	emptyserviceclient.Interface

	c   thrift.Client
	nwc thrift.NoWireClient
}

func (c client) Hello(
	ctx context.Context,
	opts ...yarpc.CallOption,
) (err error) {

	var result common.ExtendEmpty_Hello_Result
	args := common.ExtendEmpty_Hello_Helper.Args()

	if c.nwc != nil && c.nwc.Enabled() {
		if err = c.nwc.Call(ctx, args, &result, opts...); err != nil {
			return
		}
	} else {
		var body wire.Value
		if body, err = c.c.Call(ctx, args, opts...); err != nil {
			return
		}

		if err = result.FromWire(body); err != nil {
			return
		}
	}

	err = common.ExtendEmpty_Hello_Helper.UnwrapResponse(&result)
	return
}
