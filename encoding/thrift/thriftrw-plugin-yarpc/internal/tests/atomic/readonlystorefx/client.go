// Code generated by thriftrw-plugin-yarpc
// @generated

package readonlystorefx

import (
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/encoding/thrift"
	"go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/atomic/readonlystoreclient"
)

// Client provides a ReadOnlyStore client to an Fx application using the given name
// for routing.
//
// 	fx.Provide(
// 		readonlystorefx.Client("..."),
// 		newHandler,
// 	)
func Client(name string, opts ...thrift.ClientOption) interface{} {
	return func(d *yarpc.Dispatcher) readonlystoreclient.Interface {
		return readonlystoreclient.New(d.ClientConfig(name), opts...)
	}
}