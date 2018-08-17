/*
 * Copyright (c) 2018, Apex Learning, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package smerror

import (
	"fmt"
	"net/http"
)

type smerror struct {
	msg    string
	status int
}

type Error interface {
	String() string
	Error() string
	Status() int
	SetStatus(int)
}

func New(text string) Error {
	return &smerror{msg: text,
		status: http.StatusBadRequest,
	}
}

func Errorf(format string, a ...interface{}) Error {
	return New(fmt.Sprintf(format, a...))
}

func CastErr(err error) Error {
	return Errorf(err.Error())
}

func (e *smerror) Error() string {
	return e.msg
}

func (e *smerror) String() string {
	return e.msg
}

func (e *smerror) SetStatus(s int) {
	e.status = s
}

func (e *smerror) Status() int {
	return e.status
}
