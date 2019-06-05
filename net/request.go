package net

import (
	"errors"
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

func getMethodFromString(s string) (HttpMethod, error) {
	switch s {
	case "GET":
		return GET, nil
	case "HEAD":
		return HEAD, nil
	case "POST":
		return POST, nil
	case "PUT":
		return PUT, nil
	case "DELETE":
		return DELETE, nil
	case "UPDATE":
		return UPDATE, nil
	}
	return GET, errors.New("invalid http verb string")
}

type Request struct {
	Path    string
	Method  HttpMethod
	Headers map[string]string
	Content []uint8
}

func NewRequest(path string, method HttpMethod) *Request {
	return &Request{Path: path, Method: method, Headers: map[string]string{}, Content: make([]uint8, 0)}
}

func CreateRequestFromBytes(bytes []uint8) (*Request, error) {
	bytesAsString := string(bytes)
	lines := strings.Split(bytesAsString, "\r\n")
	firstLineTokens := strings.Split(lines[0], " ")
	requestMethod, err := getMethodFromString(firstLineTokens[0])
	if err != nil {
		return nil, err
	}
	r := NewRequest(firstLineTokens[1], requestMethod)
	for _, line := range lines {
		fmt.Println(line)
		if strings.Contains(line, ": ") {
			// Then it is a header?
			header := strings.Split(line, ": ")
			r.Headers[header[0]] = header[1]
		}
		if line == "" {
			fmt.Println("Got an empty newline, that must mean we are finished parsing the headers")
		}
	}
	lastLineIndex := len(lines) - 1
	copy(r.Content, lines[lastLineIndex])
	return r, nil
}
