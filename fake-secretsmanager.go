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
	"github.com/pborman/uuid"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	region               = "us-west-2"
	contentType          = "application/x-amz-json-1.1"
	arnBase              = "arn:aws:secretsmanager:%s:%012d:secret:%s"
	getSecretTarget      = "secretsmanager.GetSecretValue"
	describeSecretTarget = "secretsmanager.DescribeSecret"
)

const accountId = 123456

var secretMap map[string]string
var setTimestamp int64

func init() {
	secretMap = make(map[string]string)
	setTimestamp = time.Now().Unix()
}

func main() {
	opts, err := parseOptions()
	if err != nil {
		log.Fatal(err)
	}

	secretDataFile, err := os.Open(opts.SecretsJson)
	if err != nil {
		log.Fatal(err)
	}
	defer secretDataFile.Close()

	b, _ := ioutil.ReadAll(secretDataFile)
	err = json.Unmarshal(b, &secretMap)
	if err != nil {
		log.Fatal(err)
	}
	keys := make([]string, len(secretMap))
	i := 0
	for k, _ := range secretMap {
		keys[i] = k
		i++
	}
	log.Printf("Loading secrets with keys %s...", strings.Join(keys, ", "))
	http.HandleFunc("/", rootHandler)
	err = http.ListenAndServe(opts.Addr, nil)
	if err != nil {
		log.Println(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", contentType)
	reqId := uuid.New()
	w.Header().Set("X-Amz-RequestId", reqId)

	req, err := parseJSON(r.Body)
	if err != nil {
		jsonErrorReport(w, r, err.Error(), err.Status())
		return
	}

	switch r.Header.Get("x-amz-target") {
	case getSecretTarget:
		val, err := getSecret(req)
		if err != nil {
			jsonErrorReport(w, r, err.Error(), err.Status())
			return
		}
		enc := json.NewEncoder(w)
		if perr := enc.Encode(&val); perr != nil {
			log.Println(perr)
		}
	case describeSecretTarget:
		jsonErrorReport(w, r, "Target secretsmanager.DescribeSecret hasn't been implemented yet.", http.StatusInternalServerError)
	default:
		t := r.Header.Get("x-amz-target")
		jsonErrorReport(w, r, fmt.Sprintf("Unimplemented target %s", t), http.StatusInternalServerError)
	}
	return
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
		return "ResourceNotFoundException"
	case http.StatusInternalServerError:
		return "InternalServiceError"
	default:
		return "UnknownException"
	}
}
