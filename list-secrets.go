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
)

/*
 * yaaaaaaaaaaaay pointers! Needed here to make the JSON response behave
 * correctly, though.
 */

type secretListItem struct {
	Arn                    *string             `json:"ARN"`
	DeletedDate            *int64              `json:"DeletedDate"`
	Description            *string             `json:"Description"`
	KmsKeyId               *string             `json:"KmsKeyId"`
	LastAccessedDate       *int64              `json:"LastAccessedDate"`
	LastChangedDate        *int64              `json:"LastChangedDate"`
	LastRotatedDate        *int64              `json:"LastRotatedDate"`
	Name                   *string             `json:"Name"`
	RotationEnabled        *bool               `json:"RotationEnabled"`
	RotationLambdaArn      *string             `json:"RotationLambdaARN"`
	RotationRules          map[string]int      `json:"RotationRules"`
	SecretVersionsToStages map[string][]string `json:"SecretVersionsToStages"`
	Tags                   map[string]string   `json:"Tags"`
}

type secretList struct {
	NextToken *string
	Items     []*secretListItem `json:"SecretList"`
}

func listSecrets(data map[string]interface{}) (*secretList, smerror.Error) {
	// for now, don't worry about paging or limits. That can be implemented
	// later if the need arises.
	_ = data

	listOfSecrets := new(secretList)
	listOfSecrets.Items = make([]*secretListItem, len(secretMap))

	stages := []string{"AWSCURRENT"}

	i := 0
	for k, _ := range secretMap {
		arn := makeArn(k)
		desc := fmt.Sprintf("A pretend secret, id '%s'", k)
		vId := makeVersionId(k)

		item := new(secretListItem)
		item.Arn = &arn
		item.Name = &k
		item.LastChangedDate = &setTimestamp
		item.Description = &desc
		item.SecretVersionsToStages = map[string][]string{vId: stages}

		listOfSecrets.Items[i] = item
		i++
	}
	return listOfSecrets, nil
}
