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
	"testing"
)

func TestListSecrets(t *testing.T) {
	r := make(map[string]interface{}) // empty - not using any of the values
					  // that could be passed in.
	list, _ := listSecrets(r)
	
	expectedLen := len(secretMap)
	if len(list.Items) != expectedLen {
		t.Errorf("There should have been %d items in the secrets list, but there were %d instead.", expectedLen, len(list.Items))
	}

	if list.NextToken != nil {
		t.Errorf("Strange - list.NextToken is not nil, but it should always be currently: %v", *list.NextToken)
	}

	i := 0
	for k, _ := range secretMap {
		if k != *list.Items[i].Name {
			t.Errorf("Item %d's name should have been '%s', but was '%s'.", i, k, *list.Items[i].Name)
		}
		testArn := makeArn(k)
		if testArn != *list.Items[i].Arn {
			t.Errorf("Item %d ARN should be '%s', but is '%s'", i, testArn, *list.Items[i].Arn)
		}
		expectedDesc := fmt.Sprintf("A pretend secret, id '%s'", k)
		if expectedDesc != *list.Items[i].Description {
			t.Errorf("Expected item %d description:\n%s\n\nActual Description:\n%s\n\n", i, expectedDesc, *list.Items[i].Description)
		}
		i++
	}
}
