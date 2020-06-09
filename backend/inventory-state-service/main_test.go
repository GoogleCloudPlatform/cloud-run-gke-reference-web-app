// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	state "github.com/GoogleCloudPlatform/cloud-run-anthos-reference-web-app/inventory-state-service/src"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestItemInventoryStateToEvent(t *testing.T) {
	transactionId := "123"
	s := state.ItemInventoryState{
		ItemId:     "111",
		TotalCount: 5,
	}

	expectedEvent := cloudevents.NewEvent()
	expectedEvent.SetID(transactionId)
	expectedEvent.SetSource("inventory-state-service")
	expectedEvent.SetType("state.ItemInventoryState")
	expectedEvent.SetData(cloudevents.ApplicationJSON, s)

	event := itemInventoryStateToEvent(transactionId, s)
	if !cmp.Equal(event, expectedEvent) {
		t.Errorf("Expected = %s; got %s", expectedEvent, event)
	}
}
