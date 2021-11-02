package app

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type Controller struct {
	service Service
}

type Resource struct {
	collection string
	id         string
	props      map[string]string
	path       string
}

func getTokensFromPath(path string) []string {
	tokens := strings.Split(path, "/")
	j := 0
	for _, t := range tokens {
		if t != "" {
			tokens[j] = t
			j++
		}
	}
	return tokens[:j]
}

func getResourceFromPath(path string) Resource {
	tokens := getTokensFromPath(path)

	props := make(map[string]string)
	collection := ""
	id := ""
	key := ""
	for i, token := range tokens {
		if i%2 == 0 {
			collection = token
			key = token
			id = ""
		} else {
			props[key] = token
			id = token
		}
	}
	println("collection:", collection)
	println("id:", id)
	println("props:", props)
	println("path:", path)

	return Resource{collection, id, props, path}
}

func parseBody(bodyStream io.Reader) (*Document, error) {
	body := new(Document)
	err := json.NewDecoder(bodyStream).Decode(body)
	return body, err
}

func (controller Controller) handleRequest(rw http.ResponseWriter, r *http.Request) {
	resource := getResourceFromPath(r.URL.Path)
	println(r.Method + " " + r.URL.Path)

	var responseBody interface{}
	var statusCode int

	service := controller.service

	switch r.Method {
	case "POST":
		body, parseError := parseBody(r.Body)
		if parseError != nil {
			writeResponse(rw, make(map[string]string), 400)
			return
		}
		responseBody, statusCode = service.postDocument(resource, body)
	case "GET":
		if resource.id != "" {
			responseBody, statusCode = service.getDocument(resource)
		} else {
			responseBody, statusCode = service.listDocuments(resource, nil)
		}
	case "PATCH":
		body, parseError := parseBody(r.Body)
		if parseError != nil {
			writeResponse(rw, make(map[string]string), 400)
			return
		}
		responseBody, statusCode = service.patchDocument(resource, body)
	case "PUT":
		body, parseError := parseBody(r.Body)
		if parseError != nil {
			writeResponse(rw, make(map[string]string), 400)
			return
		}
		responseBody, statusCode = service.putDocument(resource, body)
	case "DELETE":
		responseBody, statusCode = service.deleteDocument(resource)
	default:
		rw.WriteHeader(405)
	}

	writeResponse(rw, responseBody, statusCode)
}

func writeResponse(rw http.ResponseWriter, responseBody interface{}, statusCode int) {
	rw.Header().Set("Content-Type", "application/json")
	//write the header
	rw.WriteHeader(statusCode)

	//write the body
	err := json.NewEncoder(rw).Encode(responseBody)
	if err != nil {
		println("failed to write response: " + err.Error())
	}
}

func InitializeController() *Controller {
	return new(Controller)
}
