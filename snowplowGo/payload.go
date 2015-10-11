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
	// "net/http"
	"strconv"
	"time"
)

// TODO(alexanderdean): why are the values integers? Shouldn't they be strings?
type Payload map[string]int64

// Initializes a Payload instance, containing an array in which event
// parameters are stored (type int64)
// TODO(alexanderdean): why int64 for event parameters?
func (p *Payload) InitPayload(TimeStamp string) {
	var paraValue int64

	if len(TimeStamp) > 0 {
		paraValue, _ := strconv.ParseInt(TimeStamp, 10, 64) //
		p.Add("dtm", paraValue)
	} else {
		// TODO(alexanderdean): what is http.Server.ReadTimeout?
		paraValue = time.Now().Unix() //- http.Server.ReadTimeout) * 1000
		p.Add("dtm", paraValue)
	}

}

// Adds a single map (string[int64]) to the payload
// TODO(alexanderdean): why int64 for event parameters?
func (p Payload) Add(name string, value int64) {
	// if value != nil && value != "" {  //Cannot compare value(int) to nil
	//TODO: Add another type of checks
	p[name] = value
	// }
}

// Adds an array of name-value pairs to the payload
// TODO(alexanderdean): how does this work without a type for dict?
// func (p *Payload) AddDict(dict) {
// 	for name, element := range dict {
// 		p.Add(name, element)
// 	}
// }

// Adds a JSON formatted array to the payload.
// Json encodes the array first (turns it into a string) and then will encode (or not) the string in base64
// TODO(alexanderdean): JSON as a string->string map is very narrow (should be string->interface?)
func (p *Payload) AddJson(instance map[string]string, Base64 bool, NameEncoded string, NameNotEncode string) {
	if instance != nil {
		if Base64 {
			val, err := json.Marshal(instance)
			if err != nil {
				panic("JSON Parse Failed!")
			}
			// val =
			intval, err := strconv.ParseInt((string)(base64.StdEncoding.EncodeToString(val)), 10, 64)
			if err != nil {
				panic("String Parse Failed!")
			}
			p.Add(NameEncoded, intval)
		}
	} else {
		val, err := json.Marshal(instance)
		if err != nil {
			panic("JSON Parse Failed!")
		}
		intval, err := strconv.ParseInt((string)(val), 10, 64)
		if err != nil {
			panic("String Parse Failed!")
		}
		p.Add(NameNotEncode, intval)

	}
}
