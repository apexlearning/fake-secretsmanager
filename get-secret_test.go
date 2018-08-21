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

const (
	goodkey = "good/secret/key"
	badkey  = "bad/secret/key"
	goodval = "i'm a secret, shhh"
)

func init() {
	secretMap = make(map[string]string)
	secretMap[goodkey] = goodval
}

func TestGetSecret(t *testing.T) {
	fetching := map[string]interface{}{"SecretId": goodkey}
	s, err := getSecret(fetching)
	if s == nil {
		t.Errorf("secret data was nil!")
	}
	if err != nil {
		t.Errorf("getSecret error is: %+v", err)
	}

	if s != nil && s.SecretString != goodval {
		t.Errorf("SecretString should have been '%s', but was '%s' instead.", goodval, s.SecretString)
	}
}

func TestGetNotExistSecret(t *testing.T) {
	fetching := map[string]interface{}{"SecretId": badkey}
	s, err := getSecret(fetching)
	if s != nil {
		t.Errorf("secret data was not nil when it should have been. Struct contains: %+v", s)
	}
	if err != nil && err.Status() != http.StatusBadRequest {
		t.Errorf("err status should have been %d, but was %d", http.StatusBadRequest, err.Status())
	}
}

func TestEmptySecretId(t *testing.T) {
	fetching := make(map[string]interface{})
	_, err := getSecret(fetching)
	if err == nil {
		t.Errorf("trying to fetch a secret with no supplied SecretId failed without an error")
	} else if err != nil && err.Error() != "no SecretId found in request!" {
		t.Errorf("Trying to fetch a secret without a SecretId failed, but for the wrong reason. Error was: '%s'.", err.Error())
	}

}
