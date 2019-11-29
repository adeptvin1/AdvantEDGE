/*
 * Copyright (c) 2019  InterDigital Communications, Inc
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

package metricstore

import (
	"fmt"
	"testing"

	log "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-logger"
)

const eventsStoreName string = "eventsStore"
const eventsStoreAddr string = "http://localhost:30986"

func TestEventsMetricsGetSet(t *testing.T) {
	fmt.Println("--- ", t.Name())
	log.MeepTextLogInit(t.Name())

	fmt.Println("Create valid Metric Store")
	ms, err := NewMetricStore(eventsStoreName, eventsStoreAddr)
	if err != nil {
		t.Errorf("Unable to create Metric Store")
	}

	fmt.Println("Flush store metrics")
	ms.Flush()

	fmt.Println("Set event metric")
	err = ms.SetEventMetric("MOBILITY", "event1")
	if err != nil {
		t.Errorf("Unable to set event metric")
	}
	err = ms.SetEventMetric("NETWORK-CHARACTERISTIC-UPDATE", "event2")
	if err != nil {
		t.Errorf("Unable to set event metric")
	}
	err = ms.SetEventMetric("POAS-IN-RANGE", "event3")
	if err != nil {
		t.Errorf("Unable to set event metric")
	}
	err = ms.SetEventMetric("MOBILITY", "event4")
	if err != nil {
		t.Errorf("Unable to set event metric")
	}
	err = ms.SetEventMetric("NETWORK-CHARACTERISTIC-UPDATE", "event5")
	if err != nil {
		t.Errorf("Unable to set event metric")
	}
	err = ms.SetEventMetric("POAS-IN-RANGE", "event6")
	if err != nil {
		t.Errorf("Unable to set event metric")
	}

	fmt.Println("Get event metrics")
	event, err := ms.GetLastEventMetric("MOBILITY")
	if err != nil {
		t.Errorf("Event metric should exist")
	} else if event != "event4" {
		t.Errorf("Invalid metric values")
	}
	event, err = ms.GetLastEventMetric("NETWORK-CHARACTERISTIC-UPDATE")
	if err != nil {
		t.Errorf("Event metric should exist")
	} else if event != "event5" {
		t.Errorf("Invalid metric values")
	}
	event, err = ms.GetLastEventMetric("POAS-IN-RANGE")
	if err != nil {
		t.Errorf("Event metric should exist")
	} else if event != "event6" {
		t.Errorf("Invalid metric values")
	}

	fmt.Println("Get event metrics")

	_, err = ms.GetEventMetrics("MOBILITY", "1ms", 0)
	if err == nil {
		t.Errorf("No metrics should be found in the last 1 ms")
	}
	result, err := ms.GetEventMetrics("MOBILITY", "", 1)
	if err != nil || len(result) != 1 {
		t.Errorf("Failed to get metric")
	}
	if !validateEventsMetric(result[0], "event4") {
		t.Errorf("Invalid result")
	}
	result, err = ms.GetEventMetrics("MOBILITY", "", 0)
	if err != nil || len(result) != 2 {
		t.Errorf("Failed to get metric")
	}
	if !validateEventsMetric(result[0], "event4") {
		t.Errorf("Invalid result")
	}
	if !validateEventsMetric(result[1], "event1") {
		t.Errorf("Invalid result")
	}

	// t.Errorf("DONE")
}

func validateEventsMetric(result map[string]interface{}, v1 string) bool {
	return !(result["event"] != v1)
}