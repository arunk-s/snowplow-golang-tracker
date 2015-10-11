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
// "time"
//    "crypto/md5"
)

const (
	DEFAULT_BASE_64 = true
	TRACKER_VERSION = "golang-0.1.0"

	// SCHEMA_VENDOR    = "com.snowplowanalytics.snowplow"
	// SCHEMA_FORMAT    = "jsonschema"
)

type Tracker struct {
	EncodeBase64 bool
	StdNvPairs   struct {
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

type Items struct {
	sku      string
	price    string
	quantity string
	name     string
	category string
}

// Initializes a new tracker instance with emitter(s) and a subject.
func (t *Tracker) InitTracker(emitterTracker map[string]string, subject Subject, namespace string, AppId string, EncodeBase64 string) {

	// what is the purpose of bounds checking, if doing the same thing
	// if len(emitterTracker) > 0 {
	// 	t.emitter = emitterTracker //Tracker don't have a emitter object
	// } else {
	// 	t.emitter = emitterTracker
	// }
	// t.subject = subject //Tracker don't have a subject object
	// if t.EncodeBase64 {
	// 	t.EncodeBase64 = StringToBool(t.EncodeBase64)
	// } else {
	// 	t.EncodeBase64 = DEFAULT_BASE_64
	// }

	t.EncodeBase64 = DEFAULT_BASE_64

	t.StdNvPairs.tv = TRACKER_VERSION
	t.StdNvPairs.tna = namespace
	t.StdNvPairs.aid = AppId

	t.JsonSchema.ContextSchema = "iglu:" + SCHEMA_VENDOR + "/contexts/" + SCHEMA_FORMAT + "/1-0-1"
	t.JsonSchema.UnstructEventSchema = "iglu:" + SCHEMA_VENDOR + "/unstruct_event/" + SCHEMA_FORMAT + "/1-0-0"
	t.JsonSchema.ScreenViewSchema = "iglu:" + SCHEMA_VENDOR + "/screen_view/" + SCHEMA_FORMAT + "/1-0-0"
}

// Updates the tracker with a new subject
// func (t *Tracker) UpdateSubject(subject Subject) {
// 	t.Subject = subject
// }

// Appends another emitter to the tracker
// func (t *Tracker) AddEmitter(emitter Emitter) {
// 	append(t.Emitter, emitter)
// }

// Sends the Payload to the emitter for processing
func (t *Tracker) SendRequest(payload Payload) {
	// finalPayload := ReturnArrayStringify("strval", payload) //What is ReturnArrayStringify
	// for _, element := range t.Emitter {
	// 	element.SendEvent(finalPayload)
	// }
}

// Will force-send all events in the emitter(s) buffers.
// This happens irrespective of whether or not buffer limit has been reached
// func FlushEmitters() {
// 	for _, element := range s.Emitter {
// 		element.Flush()
// 	}
// }

// Takes a Payload object as a parameter and appends all necessary event data to it
func (t *Tracker) ReturnCompletePayload(payload Payload, context string) Payload {
	var contextEnvelope map[string]string
	if context != "" {
		// contextEnvelope["schema"] = t.CONTEXT_SCHEMA
		contextEnvelope["schema"] = t.JsonSchema.ContextSchema
		contextEnvelope["data"] = context
		payload.AddJson(contextEnvelope, t.EncodeBase64, "cx", "co")
	}
	// payload.AddDict(t.StdNvPairs)
	// payload.AddDict(t.s.GetSubject())
	// payload.Add("eid", payload.GenerateUuid())
	return payload
}

// Returns a UUID for a time stamp in Nanoseconds
// TODO(alexanderdean): replace with a UUID library
// func (t *Tracker) GenerateUuid() []byte {
// 	convert,_ := time.Now().MarshalBinary()

// 	return md5.Sum(convert)
// }

// Takes a Payload and a Context and forwards the finalised payload
// map [string]string to the sendRequest function.
func (t *Tracker) Track(payload Payload, context string) {
	payload = t.ReturnCompletePayload(payload, context)
	// t.SendRequest(payload.Get()) //Get method still undefined
}

// Tracks a page view.
func (t *Tracker) TrackPageView(pageUrl string, pageTitle string, referrer string, context string, tstamp string) {
	var payloadEp Payload
	payloadEp.InitPayload(tstamp)
	// payloadEp.Add("e", pv)  //int64 expected
	// payloadEp.Add("url", pageUrl)
	// payloadEp.Add("page", pageTitle)
	// payloadEp.Add("refr", referrer)
	t.Track(payloadEp, context)
}

// Tracks a structured event with the aforementioned metrics
func (t *Tracker) TrackStructEvent(category string, action string, label string, property string, value string, context string, tstamp string) {
	var ep Payload
	ep.InitPayload(tstamp)
	// ep.Add("e", "se")
	// ep.Add("se_ca", category)
	// ep.Add("se_ac", action)
	// ep.Add("se_la", label)
	// ep.Add("se_pr", property)
	// ep.Add("se_va", value)
	t.Track(ep, context)
}

// Creates an event for each item in the ecommerceTransaction item
func (t *Tracker) TrackEcommerceTransactionItems(orderId string, sku string, price float64, quantity string, name string, category string, currency string, context string, tstamp string) {
	var ep Payload
	ep.InitPayload(tstamp)
	// ep.Add("e", "ti")
	// ep.Add("ti_id", order_id)
	// ep.Add("ti_pr", price)
	// ep.Add("ti_sk", sku)
	// ep.Add("ti_qu", quantity)
	// ep.Add("ti_nm", name)
	// ep.Add("ti_ca", category)
	// ep.Add("ti_cu", currency)
	t.Track(ep, context)
}

//Tracks an unstructured event with the aforementioned metrics
func (t *Tracker) TrackUnstructEvent(eventJson string, context string, tstamp string) {
	var envelope map[string]string
	envelope["schema"] = t.JsonSchema.UnstructEventSchema
	envelope["data"] = eventJson
	var ep Payload
	ep.InitPayload(tstamp)
	// ep.Add("e", "ue")
	ep.AddJson(envelope, t.EncodeBase64, "ue_px", "ue_pr")
	t.Track(ep, context)
}

//Tracks a screen view event with the metrics
func (t *Tracker) TrackScreenView(name string, id string, context string, tstamp string) {
	var ScreenViewProperties, epJson map[string]string
	if name != "" {
		ScreenViewProperties["name"] = name
	}
	if id != "" {
		ScreenViewProperties["id"] = id
	}
	epJson["schema"] = t.JsonSchema.ScreenViewSchema
	// epJson["data"] = ScreenViewProperties
	// t.TrackUnstructEvent(epJson, context, tstamp) //cannot use epJson (map[string]string) as type string
}

//Tracks an ecommerce transaction event, can contain many items
func (t *Tracker) TrackEcommerceTransaction(orderId string, totalValue string, currency string, affiliation string, taxValue string, shipping string, city string, state string, country string, items Items, context string, tstamp string) {
	var ep Payload
	ep.InitPayload(tstamp)
	// ep.Add("e", "tr")
	// ep.Add("tr_id", orderId)
	// ep.Add("tr_tt", totalValue)
	// ep.Add("tr_cu", currency)
	// ep.Add("tr_af", affiliation)
	// ep.Add("tr_tx", taxValue)
	// ep.Add("tr_sh", shipping)
	// ep.Add("tr_ci", city)
	// ep.Add("tr_st", state)
	// ep.Add("tr_co", country)
	t.Track(ep, context)
	// for _, element := range items { //Cannot range over items of type Item
	// t.TrackEcommerceTransactionItems(orderId, element.sku, element.price, element.quantity, element.name, element.category, currency, context, tstamp)
	// }
}
