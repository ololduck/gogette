package net

import (
	"fmt"
)

// It's a shame we can't create struct constants with `const`...
var (
	OK          = NewHttpStatus(200, "OK")
	BAD_REQUEST = NewHttpStatus(400, "Bad Request")
	NOT_FOUND   = NewHttpStatus(404, "Not Found")
)

type HttpStatus struct {
	code   uint16
	reason string
}

func NewHttpStatus(code uint16, reason string) *HttpStatus {
	return &HttpStatus{code: code, reason: reason}
}

type Response struct {
	Status  HttpStatus
	Headers map[string]string
	Content []uint8
}

func NewResponse(status HttpStatus) *Response {
	return &Response{Status: status, Headers: make(map[string]string, 5), Content: []uint8{}}
}

func (r Response) ToBytes() *[]uint8 {
	// TODO: use bytes package
	b := make([]uint8, 0)
	b = append(b, []uint8(fmt.Sprintf("HTTP/1.1 %d %s\r\n", r.Status.code, r.Status.reason))...)
	for k, v := range r.Headers {
		b = append(b, []uint8(fmt.Sprintf("%s: %s\r\n", k, v))...)
	}
	b = append(b, []uint8("\r\n")...)
	return &b
}
