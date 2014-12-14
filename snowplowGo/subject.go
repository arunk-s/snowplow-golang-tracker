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

// Initializes the Subject
func (s *Subject) InitSubject() {
	s.p = DEFAULT_PLATFORM
}

// Sets the platform from which the event is fired
func (s *Subject) SetPlatform(platform string) {
	s.p = platform
}

// Sets a custom user identification for the event
func (s *Subject) SetUserId(userId string) {
	s.uid = userId
}

// Sets the screen resolution
func (s *Subject) SetScreenResolution(width int, height int) {
	s.res = strconv.Itoa(width) + "x" + strconv.Itoa(height)
}

// Sets the view port resolution
func (s *Subject) SetViewPort(width int, height int) {
	s.vp = strconv.Itoa(width) + "x" + strconv.Itoa(height)
}

// Sets the screen's color depth
func (s *Subject) SetColorDepth(depth int) {
	s.cd = strconv.Itoa(depth)
}

// Sets the event timezone
func (s *Subject) SetTimeZone(timezone string) {
	s.tz = timezone
}

// Sets the language used
func (s *Subject) SetLanguage(language string) {
	s.lang = language
}
