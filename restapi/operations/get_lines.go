package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetLinesHandlerFunc turns a function with the right signature into a get lines handler
type GetLinesHandlerFunc func(GetLinesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetLinesHandlerFunc) Handle(params GetLinesParams) middleware.Responder {
	return fn(params)
}

// GetLinesHandler interface for that can handle valid get lines params
type GetLinesHandler interface {
	Handle(GetLinesParams) middleware.Responder
}

// NewGetLines creates a new http.Handler for the get lines operation
func NewGetLines(ctx *middleware.Context, handler GetLinesHandler) *GetLines {
	return &GetLines{Context: ctx, Handler: handler}
}

/*GetLines swagger:route GET /lines getLines

List of lines

List of all the lines TODO
supports pagination


*/
type GetLines struct {
	Context *middleware.Context
	Handler GetLinesHandler
}

func (o *GetLines) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetLinesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
