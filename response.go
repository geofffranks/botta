package goapi

import (
	"github.com/geofffranks/treefinder"
	"net/http"
)

type Response struct {
	HTTPResponse *http.Response
	Raw          []byte
	Data         interface{}
}

func (r *Response) StringVal(path string) (string, error) {
	return treefinder.FindString(r.Data, path)
}

func (r *Response) NumVal(path string) (treefinder.Number, error) {
	return treefinder.FindNum(r.Data, path)
}

func (r *Response) BoolVal(path string) (bool, error) {
	return treefinder.FindBool(r.Data, path)
}

func (r *Response) MapVal(path string) (map[string]interface{}, error) {
	return treefinder.FindMap(r.Data, path)
}

func (r *Response) ArrayVal(path string) ([]interface{}, error) {
	return treefinder.FindArray(r.Data, path)
}

func (r *Response) Val(path string) (interface{}, error) {
	return treefinder.Find(r.Data, path)
}
