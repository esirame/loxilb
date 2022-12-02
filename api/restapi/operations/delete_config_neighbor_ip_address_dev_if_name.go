// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteConfigNeighborIPAddressDevIfNameHandlerFunc turns a function with the right signature into a delete config neighbor IP address dev if name handler
type DeleteConfigNeighborIPAddressDevIfNameHandlerFunc func(DeleteConfigNeighborIPAddressDevIfNameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteConfigNeighborIPAddressDevIfNameHandlerFunc) Handle(params DeleteConfigNeighborIPAddressDevIfNameParams) middleware.Responder {
	return fn(params)
}

// DeleteConfigNeighborIPAddressDevIfNameHandler interface for that can handle valid delete config neighbor IP address dev if name params
type DeleteConfigNeighborIPAddressDevIfNameHandler interface {
	Handle(DeleteConfigNeighborIPAddressDevIfNameParams) middleware.Responder
}

// NewDeleteConfigNeighborIPAddressDevIfName creates a new http.Handler for the delete config neighbor IP address dev if name operation
func NewDeleteConfigNeighborIPAddressDevIfName(ctx *middleware.Context, handler DeleteConfigNeighborIPAddressDevIfNameHandler) *DeleteConfigNeighborIPAddressDevIfName {
	return &DeleteConfigNeighborIPAddressDevIfName{Context: ctx, Handler: handler}
}

/*
	DeleteConfigNeighborIPAddressDevIfName swagger:route DELETE /config/neighbor/{ip_address}/dev/{if_name} deleteConfigNeighborIpAddressDevIfName

# Delete IPv4 neighbor in the device

Delete IPv4 neighbor in the device
*/
type DeleteConfigNeighborIPAddressDevIfName struct {
	Context *middleware.Context
	Handler DeleteConfigNeighborIPAddressDevIfNameHandler
}

func (o *DeleteConfigNeighborIPAddressDevIfName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteConfigNeighborIPAddressDevIfNameParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}