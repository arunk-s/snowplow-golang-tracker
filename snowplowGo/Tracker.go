/*
Tracker.go
Copyright (c) 2014 Snowplow Analytics Ltd. All rights reserved.
This program is licensed to you under the Apache License Version 2.0,
and you may not use this file except in compliance with the Apache License
Version 2.0. You may obtain a copy of the Apache License Version 2.0 at
http://www.apache.org/licenses/LICENSE-2.0.
Unless required by applicable law or agreed to in writing,
software distributed under the Apache License Version 2.0 is distributed on
an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
express or implied. See the Apache License Version 2.0 for the specific
language governing permissions and limitations there under.
Authors: Aalekh Nigam
Copyright: Copyright (c) 2014 Snowplow Analytics Ltd
License: Apache License Version 2.0
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
	//Tracker Constants
	DEFAULT_BASE_64  = true
	TRACKER_VERSION  = "golang-0.1.0"
	// Schema Constants
	BASE_SCHEMA_PATH = "iglu:com.snowplowanalytics.snowplow"
	SCHEMA_TAG       = "jsonschema"
)

type Tracker struct {
	EncodeBase64 bool
	StdNvPairs struct {
		tv  string  TRACKER_VERSION
	        tna string
	        aid string
	}
	JsonSchema struct {
		ContextSchema       string
		UnstructEventSchema string
		ScreenViewSchema    string
	}
}


/**
* Initialize a new tracker instance with emitter(s) and a subject.
*
* @param map[string]string emitterTracker - Emitter object, used for sending event payloads to for processing
* @param Subject Subject - Subject object, contains extra information which is parcelled with the event
* @param string namespace
* @param string AppId
* @param string EncodeBase64 - Boolean stating whether or not to encode certain values as base64
*/
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

	t.JsonSchema.ContextSchema = BASE_SCHEMA_PATH + "/contexts/" + SCHEMA_TAG + "/1-0-0"
	t.JsonSchema.UnstructEventSchema = BASE_SCHEMA_PATH + "/unstruct_event/" + SCHEMA_TAG + "/1-0-0"
	t.JsonSchema.ScreenViewSchema = BASE_SCHEMA_PATH + "/screen_view/" + SCHEMA_TAG + "/1-0-0"
}

/**
* Updates the subject of the tracker with a new subject
*
* @param Subject subject
*/
func (t *Tracker) UpdateSubject(subject Subject) {
	t.Subject = subject
}

 /**
* Appends another emitter to the tracker
*
* @param Emitter emitter
*/
func (t *Tracker) AddEmitter(emitter Emitter) {
	append(t.Emitter, emitter)
}

 // Emitter Send Functions
/**
* Sends the Payload to the emitter for processing
*
* @param Payload payload
*/
func (t *Tracker) SendRequest(payload Payload) {
	finalPayload = ReturnArrayStringify("strval", payload)
	for _, element := range t.Emitter {
		element.SendEvent(finalPayload)
	}
}

 /**
* Will force send all events in the emitter(s) buffers
* This happens irrespective of whether or not buffer limit has been reached
*/
func FlushEmitters() {
	for _, element = range s.Emitter {
		element.Flush()
	}
}

/**
     * Takes a Payload object as a parameter and appends all necessary event data to it
     *
     * @param Payload payload - Payload instance, contains an array of nv pairs
     * @param map [string]string context - Event context map, contains extra information on the event
     * @return Payload
     */
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

//Return Unique id for a time stamp in Nanoseconds
// @return [Size]byte
func (t *Tracker) GenerateUuid() [Size]byte{
	convert := time.Nanoseconds()
	return md5.Sum(convert)
}

// Tracking Functions
    /**
     * Takes a Payload and a Context and forwards the finalised payload map [string]string to the sendRequest function.
     *
     * @param Payload payload - Payload object as parameter
     * @param string context - Context to be added to the event
     */
func (t *Tracker) Track(payload Payload, context string) {
	payload = t.ReturnCompletePayload(payload, context)
	t.SendRequest(payload.Get())	
}

/**
     * Tracks a page view with the aforementioned metrics
     *
     * @param string pageUrl - Page URL you want to track
     * @param string pageTitle - Page Title
     * @param string referrer - Referral Page
     * @param string context - Event Context
     * @param string tstamp - Event Timestamp
**/
func (t *Tracker) TrackPageView(pageUrl string, pageTitle string, referrer string, context string, tstamp string) {
	var payloadEp Payload
	payloadEp.InitPayload(tstamp)
	payloadEp.Add("e", "pv")
	payloadEp.Add("url", pageUrl)
	payloadEp.Add("page", pageTitle)
	payloadEp.Add("refr", referrer)
	t.Track(payloadEp, context)
}

/**
     * Tracks a structured event with the aforementioned metrics
     *
     * @param string category - Event Category
     * @param string action - Event itself
     * @param string label - Refers to the string the action is performed on
     * @param string property - Property associated with either the action or the object
     * @param int value - A value associated with the user action
     * @param string context - Event Context
     * @param int tstamp - Event Timestamp
     */
    func (t *Tracker) TrackStructEvent(category string, action string, label string, property string, value string, context string, tstamp string) {
        var ep  Payload(tstamp)
        ep.InitPayload(tstamp)
        ep.Add("e", "se")
        ep.Add("se_ca", category)
        ep.Add("se_ac", action)
        ep.Add("se_la", label)
        ep.Add("se_pr", property)
        ep.Add("se_va", value)
        t.Track(ep, context)
    }
