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
	"fmt"
	"github.com/apexlearning/fake-secretsmanager/internal/smerror"
	"net/http"
)

type secretVersion struct {
	Arn           string   `json:"ARN"`
	CreatedDate   int64    `json:"CreatedDate"`
	Name          string   `json:"Name"`
	SecretBinary  []byte   `json:"SecretBinary"`
	SecretString  string   `json:"SecretString"`
	VersionId     string   `json:"VersionId"`
	VersionStages []string `json:"VersionStages"`
}

func getSecret(data map[string]interface{}) (*secretVersion, smerror.Error) {
	var secretId string
	if sid, ok := data["SecretId"]; ok {
		secretId = sid.(string)
	} else {
		smerr := smerror.Errorf("no SecretId found in request!")
		smerr.SetStatus(http.StatusInternalServerError)
	}

	s := new(secretVersion)
	s.Arn = fmt.Sprintf(arnBase, region, accountId, secretId)
	s.CreatedDate = setTimestamp
	s.Name = secretId
	if scr, ok := secretMap[secretId]; !ok {
		return nil, smerror.Errorf("secret %s not found", secretId)
	} else {
		s.SecretString = scr
	}
	s.VersionId = makeVersionId(secretId)
	s.VersionStages = []string{"AWSCURRENT"}

	return s, nil
}
