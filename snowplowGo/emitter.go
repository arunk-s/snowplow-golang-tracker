/*
Copyright (c) 2014-2015 Snowplow Analytics Ltd. All rights reserved.

This program is licensed to you under the Apache License Version 2.0,
and you may not use this file except in compliance with the Apache License
Version 2.0. You may obtain a copy of the Apache License Version 2.0 at

    http://www.apache.org/licenses/LICENSE-2.0.

Unless required by applicable law or agreed to in writing,
software distributed under the Apache License Version 2.0 is distributed on
an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
express or implied. See the Apache License Version 2.0 for the specific
language governing permissions and limitations there under.
*/
package snowplowGo

import (
	"fmt"
	"net/url"
	// "net/http"
)

// Constants for sending a payload
const (
	DEFAULT_REQ_TYPE    = "POST"
	DEFAULT_PROTOCOL    = "http"
	DEFAULT_BUFFER_SIZE = 10

	SCHEMA_VENDOR = "com.snowplowanalytics.snowplow"
	SCHEMA_FORMAT = "jsonschema"
)

type Emitter struct {
	PostRequestSchema string
	ReqType           string
	Protocol          string
	CollectorUrl      url.URL
	BufferSize        int
	Buffer            []string
	RequestsResult    []string
}

// Intitialize emitter to send event data to a collector
func (e *Emitter) InitEmitter(collectorUri string, reqType string, protocol string, bufferSize int) Emitter {

	e.PostRequestSchema = fmt.Sprintf("iglu:%s/payload_data/%s/1-0-2", SCHEMA_VENDOR, SCHEMA_FORMAT)
	if reqType != "" {
		e.ReqType = reqType
	} else {
		e.ReqType = DEFAULT_REQ_TYPE
	}
	if protocol != "" {
		e.Protocol = protocol
	} else {
		e.Protocol = DEFAULT_PROTOCOL
	}
	var err error
	e.CollectorUrl, err = e.ReturnCollectorUrl(collectorUri)
	if err != nil {
		panic("URL Parse failed")
	}
	if bufferSize == 0 {
		if e.ReqType == "POST" {
			e.BufferSize = DEFAULT_BUFFER_SIZE
		} else {
			e.BufferSize = 1
		}
	} else {
		e.BufferSize = (int)(bufferSize)
	}
	e.BufferSize = 0 // TODO(alexanderdean): won't this overwrite the assignments above?
	e.RequestsResult = nil
	return *e
}

// Returns the collector URL based on: request type, protocol and host given
// If a bad type is given in emitter creation, returns nil.
func (e *Emitter) ReturnCollectorUrl(host string) (url.URL, error) {
	var urlEncoded *url.URL
	switch e.ReqType {
	case "POST":
		rawurl := e.Protocol + "://" + host + "/com.snowplowanalytics.snowplow/tp2"
		urlEncoded, err := url.Parse(rawurl)
		if err != nil {
			return *urlEncoded, err
		}
		return *urlEncoded, nil
	case "GET":
		rawurl := e.Protocol + "://" + host + "/i?"
		urlEncoded, err := url.Parse(rawurl)
		if err != nil {
			return *urlEncoded, err
		}
		return *urlEncoded, nil
	default:
		return *urlEncoded, nil
	}

}

// Pushes the event payload into the emitter buffer.
// When buffer is full it flushes the buffer.
func (e *Emitter) SendEvent(finalPayload []string) {
	e.Buffer = append(e.Buffer, finalPayload...)
	if len(e.Buffer) >= e.BufferSize {
		e.Buffer = make([]string, e.BufferSize) //Flushed
	}

}

// Flushes the event buffer of the emitter
// Checks which send type the emitter is using and forwards data accordingly
// Resets the buffer to nothing after flushing
func (e *Emitter) Flush(emitter Emitter) {
	if len(emitter.Buffer) != 0 {
		if emitter.ReqType == "POST" {
			data := emitter.ReturnPostRequest()
			e.PostRequest(data["data"])
		} else if emitter.ReqType == "GET" {

			// for _, value := range emitter.Buffer {
			// 	e.GetRequest(value)
			// }
			// Maybe trying to do the following
			e.GetRequest(e.Buffer)
		}
		emitter.Buffer = nil
	}
}

// Sends the payload to the collector via a GET request
func (e *Emitter) GetRequest(data []string) {
	// r := url.Get(HttpBuildQuery(data)) Doesn't Work this way. Maybe using http.Get
	// r := http.Get(HttpBuildQuery(data)) //HttpBuildQuery() Undefined.
	// e.StoreRequestResults(r)
	// currently doing nothing
}

// Sends the payload to the collector via a POST request
func (e *Emitter) PostRequest(data []string) {
	m := make(map[string]string)
	m["Content-Type"] = "application/json; charset=utf-8"
	//post method to be made properly here
	// r := http.Post(e.CollectorUrl) //Need more arguements
	// e.StoreRequestResults(r)
}

// Returns an array formatted to be ready for a POST request
func (e *Emitter) ReturnPostRequest() map[string][]string {
	// dataPostRequest := make(map[string][]map[string]string)
	dataPostRequest := make(map[string][]string)
	dataPostRequest["schema"] = append(dataPostRequest["schema"], e.PostRequestSchema)
	for _, element := range e.Buffer {
		dataPostRequest["data"] = append(dataPostRequest["data"], element)
	}
	return dataPostRequest
}

// Stores all of the parameters of the request's response
// into a dynamic array for use in unit testing
// TODO(alexanderdean): is there a cleaner way of doing this?
// // func (e *Emitter) StoreRequestResults(r RequestsResponse) {
// // 	storeArray = make(map[string]string)
// // 	storeArray["url"] = r.url
// // 	storeArray["code"] = r.StatusCode
// // 	storeArray["headers"] = r.headers
// // 	storeArray["body"] = r.body
// // 	storeArray["raw"] = r.raw
// // 	append(emitter.RequestsResult, storeArray)
// }
