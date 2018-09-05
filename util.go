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
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/apexlearning/fake-secretsmanager/internal/smerror"
	"io"
	"log"
	"net/http"
)

/* Various utility functions that may be useful in multiple places. */

const arnBase = "arn:aws:secretsmanager:%s:%012d:secret:%s"

func makeArn(secretId string) string {
	return fmt.Sprintf(arnBase, region, accountId, secretId)
}

func makeVersionId(secretId string) string {
	b := []byte(secretId)
	idLen := 16
	if len(b) < 16 {
		k := make([]byte, 16-len(b))
		b = append(k, b...)
	}
	str := hex.EncodeToString(b)
	vId := fmt.Sprintf("SECRET-%s", str[len(str)-(idLen*2):])
	return vId
}

func parseJSON(data io.ReadCloser) (map[string]interface{}, smerror.Error) {
	reqData := make(map[string]interface{})
	dec := json.NewDecoder(data)
	if err := dec.Decode(&reqData); err != nil {
		smerr := smerror.CastErr(err)
		// 500 I *think* is probably most appropriate here, but 400
		// might be more appropriate.
		smerr.SetStatus(http.StatusInternalServerError)
		return nil, smerr
	}
	return reqData, nil
}

func jsonErrorReport(w http.ResponseWriter, r *http.Request, errorStr string, status int) {
	log.Println(errorStr)

	jsonError := make(map[string]string)
	jsonError["Message"] = errorStr
	jsonError["__type"] = exceptionType(status)

	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	if err := enc.Encode(&jsonError); err != nil {
		log.Println(err)
	}
	return
}

func exceptionType(status int) string {
	switch status {
	case http.StatusBadRequest:
		return resourceNotFound
	case http.StatusInternalServerError:
		return internalServiceErr
	default:
		return unknownException
	}
}
