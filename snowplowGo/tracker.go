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
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

const (
	DEFAULT_BASE_64  = true
	TRACKER_VERSION  = "golang-0.1.0"

	SCHEMA_VENDOR    = "com.snowplowanalytics.snowplow"
	SCHEMA_FORMAT    = "jsonschema"
)

type Tracker struct {
	EncodeBase64 bool
	StdNvPairs struct {
		tv  string
		tna string
		aid string
	}
	JsonSchema struct {
		ContextSchema       string
		UnstructEventSchema string
		ScreenViewSchema    string
	}
}

// Initializes a new tracker instance with emitter(s) and a subject.
func (t *Tracker) InitTracker(emitterTracker map[string]string, subject Subject, namespace string, AppId string, EncodeBase64 string) {

	if len(emitterTracker) > 0 {
		t.emitter = emitterTracker
	} else {
		t.emitter = emitterTracker
	}
	t.subject = subject
	if t.EncodeBase64 != nil {
		t.EncodeBase64 = StringToBool(t.EncodeBase64)
	} else {
		t.EncodeBase64 = DEFAULT_BASE_64
	}

	t.StdNvPairs.tv = TRACKER_VERSION
	t.StdNvPairs.tna = namespace
	t.StdNvPairs.aid = AppId

	t.JsonSchema.ContextSchema = "iglu:" + SCHEMA_VENDOR + "/contexts/" + SCHEMA_FORMAT + "/1-0-1"
	t.JsonSchema.UnstructEventSchema = "iglu:" + SCHEMA_VENDOR + "/unstruct_event/" + SCHEMA_FORMAT + "/1-0-0"
	t.JsonSchema.ScreenViewSchema = "iglu:" + SCHEMA_VENDOR + "/screen_view/" + SCHEMA_FORMAT + "/1-0-0"
}

// Updates the tracker with a new subject
func (t *Tracker) UpdateSubject(subject Subject) {
	t.Subject = subject
}

// Appends another emitter to the tracker
func (t *Tracker) AddEmitter(emitter Emitter) {
	append(t.Emitter, emitter)
}

// Sends the Payload to the emitter for processing
func (t *Tracker) SendRequest(payload Payload) {
	finalPayload = ReturnArrayStringify("strval", payload)
	for _, element := range t.Emitter {
		element.SendEvent(finalPayload)
	}
}

// Will force-send all events in the emitter(s) buffers.
// This happens irrespective of whether or not buffer limit has been reached
func FlushEmitters() {
	for _, element = range s.Emitter {
		element.Flush()
	}
}

// Takes a Payload object as a parameter and appends all necessary event data to it
func (t *Tracker) ReturnCompletePayload(payload Payload, context string) Payload{
	var contextEnvelope map[string]string
	if context != nil {
		contextEnvelope["schema"] = t.CONTEXT_SCHEMA 
		contextEnvelope["data"] = context
		payload.AddJson(contextEnvelope, t.EncodeBase64,"cx","co")
	}
	payload.AddDict(t.StdNvPairs)
	payload.AddDict(t.s.GetSubject())
	payload.Add("eid", payload.GenerateUuid())
	return payload
}

// Returns a UUID for a time stamp in Nanoseconds
// TODO(alexanderdean): replace with a UUID library
func (t *Tracker) GenerateUuid() [Size]byte{
	convert := time.Nanoseconds()
	return md5.Sum(convert)
}

// Takes a Payload and a Context and forwards the finalised payload
// map [string]string to the sendRequest function.
func (t *Tracker) Track(payload Payload, context string) {
	payload = t.ReturnCompletePayload(payload, context)
	t.SendRequest(payload.Get())	
}

// Tracks a page view.
func (t *Tracker) TrackPageView(pageUrl string, pageTitle string, referrer string, context string, tstamp string) {
	var payloadEp Payload
	payloadEp.InitPayload(tstamp)
	payloadEp.Add("e", "pv")
	payloadEp.Add("url", pageUrl)
	payloadEp.Add("page", pageTitle)
	payloadEp.Add("refr", referrer)
	t.Track(payloadEp, context)
}

// Tracks a structured event with the aforementioned metrics
func (t *Tracker) TrackStructEvent(category string, action string, label string, property string, value string, context string, tstamp string) {
	ep := InitPayload(tstamp)
	ep.Add("e", "se")
	ep.Add("se_ca", category)
	ep.Add("se_ac", action)
	ep.Add("se_la", label)
	ep.Add("se_pr", property)
	ep.Add("se_va", value)
	t.Track(ep, context)
}

// Creates an event for each item in the ecommerceTransaction item 
func (t *Tracker) TrackEcommerceTransactionItems(orderId string, sku string, price float, quantity string, name string, category string, currency string, context string, tstamp string) {
	ep := InitPayload(tstamp)
	ep.Add("e", "ti")
	ep.Add("ti_id", order_id)
	ep.Add("ti_pr", price)
	ep.Add("ti_sk", sku)
	ep.Add("ti_qu", quantity)
	ep.Add("ti_nm", name)
	ep.Add("ti_ca", category)
	ep.Add("ti_cu", currency)
	t.Track(ep, context)
}
