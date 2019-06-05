package net

import (
	"fmt"
	"strings"
)

type HttpMethod int

const (
	GET HttpMethod = iota
	HEAD
	POST
	PUT
	DELETE
	UPDATE
)

func getMethodFromString(s string) HttpMethod {
	switch s {
	case "GET":
		return GET
	case "HEAD":
		return HEAD
	case "POST":
		return POST
	case "PUT":
		return PUT
	case "DELETE":
		return DELETE
	case "UPDATE":
		return UPDATE
	}
	return GET //FIXME: that's just plain stupid. Should have a nil value.
}

type Request struct {
	Path    string
	Method  HttpMethod
	Headers map[string]string
	Content []uint8
}

func NewRequest(path string, method HttpMethod) *Request {
	return &Request{Path: path, Method: method, Headers: map[string]string{}, Content: []uint8{}}
}

func CreateRequestFromBytes(bytes []uint8) *Request {
	bytesAsString := string(bytes)
	lines := strings.Split(bytesAsString, "\r\n")
	firstLineTokens := strings.Split(lines[0], " ")
	r := NewRequest(firstLineTokens[1], getMethodFromString(firstLineTokens[0]))
	for i := 1; i < len(lines)-1; i++ {
		if strings.Contains(lines[i], ": ") {
			// Then it is a header?
			header := strings.Split(lines[i], ": ")
			r.Headers[header[0]] = header[1]
		}
		if lines[i] == "" {
			fmt.Println("Got an empty newline, that must mean we are finished parsing the headers")
		}
	}
	lastLineIndex := len(lines) - 1
	copy(r.Content, lines[lastLineIndex])
	return r
}
