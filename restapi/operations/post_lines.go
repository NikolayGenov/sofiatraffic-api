package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostLinesHandlerFunc turns a function with the right signature into a post lines handler
type PostLinesHandlerFunc func(PostLinesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostLinesHandlerFunc) Handle(params PostLinesParams) middleware.Responder {
	return fn(params)
}

// PostLinesHandler interface for that can handle valid post lines params
type PostLinesHandler interface {
	Handle(PostLinesParams) middleware.Responder
}

// NewPostLines creates a new http.Handler for the post lines operation
func NewPostLines(ctx *middleware.Context, handler PostLinesHandler) *PostLines {
	return &PostLines{Context: ctx, Handler: handler}
}

/*PostLines swagger:route POST /lines postLines

Create a line

*/
type PostLines struct {
	Context *middleware.Context
	Handler PostLinesHandler
}

func (o *PostLines) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPostLinesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
