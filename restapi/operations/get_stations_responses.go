package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/NikolayGenov/sofiatraffic-api/models"
)

/*GetStationsOK An array of stations

swagger:response getStationsOK
*/
type GetStationsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Station `json:"body,omitempty"`
}

// NewGetStationsOK creates GetStationsOK with default headers values
func NewGetStationsOK() *GetStationsOK {
	return &GetStationsOK{}
}

// WithPayload adds the payload to the get stations o k response
func (o *GetStationsOK) WithPayload(payload []*models.Station) *GetStationsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get stations o k response
func (o *GetStationsOK) SetPayload(payload []*models.Station) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStationsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Station, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetStationsDefault Unexpected error

swagger:response getStationsDefault
*/
type GetStationsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetStationsDefault creates GetStationsDefault with default headers values
func NewGetStationsDefault(code int) *GetStationsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetStationsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get stations default response
func (o *GetStationsDefault) WithStatusCode(code int) *GetStationsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get stations default response
func (o *GetStationsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get stations default response
func (o *GetStationsDefault) WithPayload(payload *models.Error) *GetStationsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get stations default response
func (o *GetStationsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStationsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
