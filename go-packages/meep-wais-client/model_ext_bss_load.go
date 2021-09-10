/*
 * Copyright (c) 2020  InterDigital Communications, Inc
 *
 * Licensed under the Apache License, Version 2.0 (the \"License\");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an \"AS IS\" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * AdvantEDGE WLAN Access Information API
 *
 * WLAN Access Information Service is AdvantEDGE's implementation of [ETSI MEC ISG MEC028 WAI API](http://www.etsi.org/deliver/etsi_gs/MEC/001_099/028/02.02.01_60/gs_MEC028v020201p.pdf) <p>[Copyright (c) ETSI 2020](https://forge.etsi.org/etsi-forge-copyright-notice.txt) <p>**Micro-service**<br>[meep-wais](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-wais) <p>**Type & Usage**<br>Edge Service used by edge applications that want to get information about WLAN access information in the network <p>**Note**<br>AdvantEDGE supports a selected subset of WAI API subscription types. <p>Supported subscriptions: <p> - AssocStaSubscription <p> - StaDataRateSubscription
 *
 * API version: 2.2.1
 * Contact: AdvantEDGE@InterDigital.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type ExtBssLoad struct {
	// Indicates the total number of STAs currently associated with this BSS that have a 1 in the MU Beamformee Capable field of their VHT Capabilities element.
	MuMimoStaCount int32 `json:"muMimoStaCount"`
	// Observable loading on each of the secondary 20 MHz channel.
	ObsSec20MhzUtil int32 `json:"obsSec20MhzUtil"`
	// Observable loading on each of the secondary 40 MHz channel.
	ObsSec40MhzUtil int32 `json:"obsSec40MhzUtil"`
	// Observable loading on each of the secondary 80 MHz channel.
	ObsSec80MhzUtil int32 `json:"obsSec80MhzUtil"`
	// The percentage of time, linearly scaled with 255 representing 100 %, that the AP has underutilized spatial domain resources for given busy time of the medium.
	SpatStreamUnderUtil int32 `json:"spatStreamUnderUtil"`
}
