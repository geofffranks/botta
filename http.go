package goapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var client = DefaultClient()

func httpRequest(method string, url string, data interface{}) (*http.Request, error) {
	marshaled, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(marshaled)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	return req, nil
}

func Get(url string) (*http.Request, error) {
	return httpRequest("GET", url, nil)
}

func Post(url string, data interface{}) (*http.Request, error) {
	return httpRequest("POST", url, data)
}

func Put(url string, data interface{}) (*http.Request, error) {
	return httpRequest("PUT", url, data)
}

func Patch(url string, data interface{}) (*http.Request, error) {
	return httpRequest("PATCH", url, data)
}

func Delete(url string) (*http.Request, error) {
	return httpRequest("DELETE", url, nil)
}

func Issue(req *http.Request) (*Response, error) {
	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return ParseResponse(r)
}

func ParseResponse(r *http.Response) (*Response, error) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	// Decode json first to include in Response if possible
	var o interface{}
	jsonErr := json.Unmarshal(body, o)

	resp := &Response{
		HTTPResponse: r,
		Raw:          body,
		Data:         o,
	}

	// Return an error for failure status codes, include
	// the Response object, so end-users can decode JSON error messages
	// If the body wasn't json, we complain about the code anyway, and return
	// nil Data, but raw response + http message
	if r.StatusCode >= 400 {
		return resp, BadResponseCode{
			URL:        r.Request.URL.String(),
			StatusCode: r.StatusCode,
			Message:    string(body),
		}
	}

	// If we had a successful request, but invalid json, return an error
	// with the Response obj, so end-users can debug as they see fit
	if jsonErr != nil {
		return resp, jsonErr
	}

	// All went well, Return the Response
	return resp, nil
}

func Client(c *http.Client) {
	client = c
}

func DefaultClient() *http.Client {
	return &http.Client{}
}

type BadResponseCode struct {
	StatusCode int
	Message    string
	URL        string
}

func (e BadResponseCode) Error() string {
	return fmt.Sprintf("%s returned %d: %s", e.URL, e.StatusCode, e.Message)
}
