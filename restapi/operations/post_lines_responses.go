package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/NikolayGenov/sofiatraffic-api/models"
)

/*PostLinesNoContent Succesfully created a line

swagger:response postLinesNoContent
*/
type PostLinesNoContent struct {
}

// NewPostLinesNoContent creates PostLinesNoContent with default headers values
func NewPostLinesNoContent() *PostLinesNoContent {
	return &PostLinesNoContent{}
}

// WriteResponse to the client
func (o *PostLinesNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

/*PostLinesDefault Unexpected error

swagger:response postLinesDefault
*/
type PostLinesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostLinesDefault creates PostLinesDefault with default headers values
func NewPostLinesDefault(code int) *PostLinesDefault {
	if code <= 0 {
		code = 500
	}

	return &PostLinesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post lines default response
func (o *PostLinesDefault) WithStatusCode(code int) *PostLinesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post lines default response
func (o *PostLinesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post lines default response
func (o *PostLinesDefault) WithPayload(payload *models.Error) *PostLinesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post lines default response
func (o *PostLinesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostLinesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
