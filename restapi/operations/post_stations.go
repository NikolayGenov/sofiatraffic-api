package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostStationsHandlerFunc turns a function with the right signature into a post stations handler
type PostStationsHandlerFunc func(PostStationsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostStationsHandlerFunc) Handle(params PostStationsParams) middleware.Responder {
	return fn(params)
}

// PostStationsHandler interface for that can handle valid post stations params
type PostStationsHandler interface {
	Handle(PostStationsParams) middleware.Responder
}

// NewPostStations creates a new http.Handler for the post stations operation
func NewPostStations(ctx *middleware.Context, handler PostStationsHandler) *PostStations {
	return &PostStations{Context: ctx, Handler: handler}
}

/*PostStations swagger:route POST /stations postStations

Create a station

Create a new station TODO

*/
type PostStations struct {
	Context *middleware.Context
	Handler PostStationsHandler
}

func (o *PostStations) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPostStationsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}