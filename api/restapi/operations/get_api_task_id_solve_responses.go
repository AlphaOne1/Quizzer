// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetAPITaskIDSolveOKCode is the HTTP code returned for type GetAPITaskIDSolveOK
const GetAPITaskIDSolveOKCode int = 200

/*
GetAPITaskIDSolveOK Check if solution is correct

swagger:response getApiTaskIdSolveOK
*/
type GetAPITaskIDSolveOK struct {

	/*
	  In: Body
	*/
	Payload bool `json:"body,omitempty"`
}

// NewGetAPITaskIDSolveOK creates GetAPITaskIDSolveOK with default headers values
func NewGetAPITaskIDSolveOK() *GetAPITaskIDSolveOK {

	return &GetAPITaskIDSolveOK{}
}

// WithPayload adds the payload to the get Api task Id solve o k response
func (o *GetAPITaskIDSolveOK) WithPayload(payload bool) *GetAPITaskIDSolveOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Api task Id solve o k response
func (o *GetAPITaskIDSolveOK) SetPayload(payload bool) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPITaskIDSolveOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
