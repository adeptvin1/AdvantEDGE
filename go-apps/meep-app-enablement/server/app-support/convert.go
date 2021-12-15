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

package server

import (
	"encoding/json"

	log "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-logger"
)

func convertAppTerminationNotifSubToJson(obj *AppTerminationNotificationSubscription) string {
	jsonInfo, err := json.Marshal(*obj)
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	return string(jsonInfo)
}

func convertAppTerminationNotifToJson(obj *AppTerminationNotification) string {
	jsonInfo, err := json.Marshal(*obj)
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	return string(jsonInfo)
}

func convertProblemDetailsToJson(obj *ProblemDetails) string {
	jsonInfo, err := json.Marshal(*obj)
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	return string(jsonInfo)
}

func convertSubscriptionLinkListToJson(obj *SubscriptionLinkList) string {
	jsonInfo, err := json.Marshal(*obj)
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	return string(jsonInfo)
}
