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

package main

import (
	"net/http"
	"testing"
)

func TestExceptionType(t *testing.T) {
	e := exceptionType(http.StatusBadRequest)
	if e != resourceNotFound {
		t.Errorf("Wrong exception type: should have been '%s', got '%s'", resourceNotFound, e)
	}
	e = exceptionType(http.StatusInternalServerError)
	if e != internalServiceErr {
		t.Errorf("Wrong exception type: should have been '%s', got '%s'", internalServiceErr, e)
	}
	e = exceptionType(http.StatusOK)
	if e != unknownException {
		t.Errorf("Wrong exception type: should have been '%s', got '%s'", unknownException, e)
	}
}

func TestMakeVersionId(t *testing.T) {
	secretId := "go-test/foo/bar/baz/bug/bleh"
	expectedId := "SECRET-6261722f62617a2f6275672f626c6568"
	vId := makeVersionId(secretId)

	if vId != expectedId {
		t.Errorf("wrong VersionId returned: expected '%s', got '%s'.", expectedId, vId)
	}
	if len(vId) > 64 || len(vId) < 32 {
		t.Errorf("incorrect VersionId length: it should be less than 64 characters and more than 32, but was %d.", len(vId))
	}
	short := "short"
	shortId := makeVersionId(short)
	if len(shortId) > 64 || len(shortId) < 32 {
		t.Errorf("incorrect VersionId length for short secret key: it should be less than 64 characters and more than 32, but was %d.", len(shortId))
	}

	long := "longlonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglong"
	longId := makeVersionId(long)
	if len(longId) > 64 || len(longId) < 32 {
		t.Errorf("incorrect VersionId length for long secret key: it should be less than 64 characters and more than 32, but was %d.", len(longId))
	}
}
