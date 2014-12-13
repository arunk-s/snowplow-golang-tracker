/*
Subject.go
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
	"strconv"
)

const (
	DEFAULT_PLATFORM = "srv"
)

var TrackerSettings map[string]string

type Subject struct {
	p string
	res string
	vp string
	uid string
	cd string
	tz string
	lang string
}
//Initialize Subject.go
func (s *Subject) InitSubject() {
	s.p = DEFAULT_PLATFORM
}

//Remember about getSubject as golang variable type access
 /**
* Sets the platform from which the event is fired
*
* @param string platform
*/
func (s *Subject) SetPlatform(platform string) {
	s.p = platform
}

 /**
* Sets a custom user identification for the event
*
* @param string userId
*/

func (s *Subject) SetUserId(userId string) {
	s.uid = userId
}

/**
* Sets the screen resolution
*
* @param int width
* @param int height
*/
func (s *Subject) SetScreenResolution(width int, height int) {
	s.res = strconv.Itoa(width) + "x" + strconv.Itoa(height)
}

/**
* Sets the view port resolution
*
* @param int width
* @param int height
*/
func (s *Subject) SetViewPort(width int, height int) {
	s.vp = strconv.Itoa(width) + "x" + strconv.Itoa(height)
}

 /**
* Sets the colour depth
*
* @param int depth
*/
func (s *Subject) SetColorDepth(depth int) {
	s.cd = strconv.Itoa(depth)
}

/**
* Sets the event timezone
*
* @param string timezone
*/
func (s *Subject) SetTimeZone(timezone string) {
	s.tz = timezone
}

 /**
* Sets the language used
*
* @param string language
*/
func (s *Subject) SetLanguage(language string) {
	s.lang = language
}
