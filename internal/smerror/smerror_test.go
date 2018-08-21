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
	"testing"
)

func TestNewError(t *testing.T) {
	errMsg := "hi I'm a test"
	smerr := Errorf(errMsg)
	if smerr.Error() != errMsg {
		t.Errorf("smerr.Error() should have been '%s', but was '%s' instead.", errMsg, smerr.Error())
	}
	if smerr.Status() != http.StatusBadRequest {
		t.Errorf("smerr.Status() should have been %d, but was %d instead", http.StatusBadRequest, smerr.Status())
	}
}

func TestSetStatus(t *testing.T) {
	errMsg := "status err"
	smerr := Errorf(errMsg)
	smerr.SetStatus(http.StatusInternalServerError)
	if smerr.Status() != http.StatusInternalServerError {
		t.Errorf("smerr status should have been %d, but was %d instead.", http.StatusInternalServerError, smerr.Status())
	}
}

func TestCastError(t *testing.T) {
	err := fmt.Errorf("yo dawg")
	smerr := CastErr(err)
	if smerr.Status() != http.StatusBadRequest {
		t.Errorf("somehow casting an err failed, what on earth")
	}
}
