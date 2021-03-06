package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/NikolayGenov/sofiatraffic-api/models"
)

/*PostStationsNoContent Succesfully created a station

swagger:response postStationsNoContent
*/
type PostStationsNoContent struct {
}

// NewPostStationsNoContent creates PostStationsNoContent with default headers values
func NewPostStationsNoContent() *PostStationsNoContent {
	return &PostStationsNoContent{}
}

// WriteResponse to the client
func (o *PostStationsNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

/*PostStationsDefault Unexpected error

swagger:response postStationsDefault
*/
type PostStationsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostStationsDefault creates PostStationsDefault with default headers values
func NewPostStationsDefault(code int) *PostStationsDefault {
	if code <= 0 {
		code = 500
	}

	return &PostStationsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post stations default response
func (o *PostStationsDefault) WithStatusCode(code int) *PostStationsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post stations default response
func (o *PostStationsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post stations default response
func (o *PostStationsDefault) WithPayload(payload *models.Error) *PostStationsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post stations default response
func (o *PostStationsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostStationsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
