// Code generated by go-swagger; DO NOT EDIT.

package number

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetNumbersHandlerFunc turns a function with the right signature into a get numbers handler
type GetNumbersHandlerFunc func(GetNumbersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetNumbersHandlerFunc) Handle(params GetNumbersParams) middleware.Responder {
	return fn(params)
}

// GetNumbersHandler interface for that can handle valid get numbers params
type GetNumbersHandler interface {
	Handle(GetNumbersParams) middleware.Responder
}

// NewGetNumbers creates a new http.Handler for the get numbers operation
func NewGetNumbers(ctx *middleware.Context, handler GetNumbersHandler) *GetNumbers {
	return &GetNumbers{Context: ctx, Handler: handler}
}

/*GetNumbers swagger:route GET /numbers number getNumbers

Returns numbers

Returns all stored numbers

*/
type GetNumbers struct {
	Context *middleware.Context
	Handler GetNumbersHandler
}

func (o *GetNumbers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetNumbersParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
