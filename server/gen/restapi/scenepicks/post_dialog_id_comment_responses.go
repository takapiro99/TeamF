// Code generated by go-swagger; DO NOT EDIT.

package scenepicks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/shortintern2020-C-cryptograph/TeamF/server/gen/models"
)

// PostDialogIDCommentBadRequestCode is the HTTP code returned for type PostDialogIDCommentBadRequest
const PostDialogIDCommentBadRequestCode int = 400

/*PostDialogIDCommentBadRequest request error

swagger:response postDialogIdCommentBadRequest
*/
type PostDialogIDCommentBadRequest struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPostDialogIDCommentBadRequest creates PostDialogIDCommentBadRequest with default headers values
func NewPostDialogIDCommentBadRequest() *PostDialogIDCommentBadRequest {

	return &PostDialogIDCommentBadRequest{}
}

// WithPayload adds the payload to the post dialog Id comment bad request response
func (o *PostDialogIDCommentBadRequest) WithPayload(payload models.Error) *PostDialogIDCommentBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post dialog Id comment bad request response
func (o *PostDialogIDCommentBadRequest) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostDialogIDCommentBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostDialogIDCommentInternalServerErrorCode is the HTTP code returned for type PostDialogIDCommentInternalServerError
const PostDialogIDCommentInternalServerErrorCode int = 500

/*PostDialogIDCommentInternalServerError internal serever error

swagger:response postDialogIdCommentInternalServerError
*/
type PostDialogIDCommentInternalServerError struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPostDialogIDCommentInternalServerError creates PostDialogIDCommentInternalServerError with default headers values
func NewPostDialogIDCommentInternalServerError() *PostDialogIDCommentInternalServerError {

	return &PostDialogIDCommentInternalServerError{}
}

// WithPayload adds the payload to the post dialog Id comment internal server error response
func (o *PostDialogIDCommentInternalServerError) WithPayload(payload models.Error) *PostDialogIDCommentInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post dialog Id comment internal server error response
func (o *PostDialogIDCommentInternalServerError) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostDialogIDCommentInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*PostDialogIDCommentDefault generic error response

swagger:response postDialogIdCommentDefault
*/
type PostDialogIDCommentDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPostDialogIDCommentDefault creates PostDialogIDCommentDefault with default headers values
func NewPostDialogIDCommentDefault(code int) *PostDialogIDCommentDefault {
	if code <= 0 {
		code = 500
	}

	return &PostDialogIDCommentDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post dialog ID comment default response
func (o *PostDialogIDCommentDefault) WithStatusCode(code int) *PostDialogIDCommentDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post dialog ID comment default response
func (o *PostDialogIDCommentDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post dialog ID comment default response
func (o *PostDialogIDCommentDefault) WithPayload(payload models.Error) *PostDialogIDCommentDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post dialog ID comment default response
func (o *PostDialogIDCommentDefault) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostDialogIDCommentDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
