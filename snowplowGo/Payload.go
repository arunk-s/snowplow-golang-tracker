
/*
Payload.go
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

Payload := map[string]int64{}

/**
* Initialize a Payload instance, contains an array in which event parameters are stored (type int64)
*
*/

func (p *Payload) InitPayload( TimeStamp string ) {
	var paraValue int64
	if TimeStamp != nil {
		paraValue = (int64)(TimeStamp)
	} else {
		paraValue = ((int64)(time.Now()) - http.Server.ReadTimeout) * 1000
	}
	p.Add("dtm", paraValue)
}

/**
* Adds a single map (string[int64]) to the payload
*
* @param string name - string for map
* @param int64 value - int64 for map
*/

func (p *Payload) Add(name string, value int64) {
	if value != nil && value != "" {
		p.Payload[name] = value
	}
}

/**
* Adds an array of name-value pairs to the payload
*
* @param array dict - Single level array of name:value pairs
*/

func (p *Payload) AddDict(dict) {
	for name, element := range dict {
		p.Add(name, element)
	}
}

/**
* Adds a JSON formatted array to the payload
* Json encodes the array first (turns it into a string) and then will encode (or not) the string in base64
*
* @param array json - Custom context for the event
* @param bool Base64 - If the payload is base64 encoded
* @param string NameEncoded
* @param string NameNotEncode
*/

func (p *Payload) AddJson(json map[string]string, Base64 bool, NameEncoded string, NameNotEncode string) {
	if json != nil {
		if Base64 {
			p.Add(NameEncoded, b64.StdEncoding.EncodeToString(json.Marshal(json)))
		}
	} else {
		p.Add(NameNotEncode, json.Marshal(json))

	}

}
