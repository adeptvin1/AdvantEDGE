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
 * ETSI GS MEC 028 - WLAN Access Information API
 *
 * WLAN Access Information Service is AdvantEDGE's implementation of [ETSI MEC ISG MEC028 WAI API](http://www.etsi.org/deliver/etsi_gs/MEC/001_099/028/02.01.01_60/gs_MEC028v020101p.pdf) <p>[Copyright (c) ETSI 2020](https://forge.etsi.org/etsi-forge-copyright-notice.txt) <p>**Micro-service**<br>[meep-wais](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-wais) <p>**Type & Usage**<br>Edge Service used by edge applications that want to get information about WLAN access information in the network <p>**Details**<br>API details available at _your-AdvantEDGE-ip-address/api_
 *
 * API version: 2.1.1
 * Contact: AdvantEDGE@InterDigital.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package server

type BeaconReport struct {
	// The BSSID field indicates the BSSID of the BSS(s) for which a beacon report has been received.
	BssId []string `json:"bssId"`
	// Channel number where the beacon was received.
	ChannelId int32 `json:"channelId"`
	// Measurement ID of the Measurement configuration applied to this Beacon Report.
	MeasurementId string `json:"measurementId"`
	// As in table T9-89 of IEEE 802.11-2012 [8].
	ReportingCondition int32 `json:"reportingCondition"`
	// (Optional) The SSID subelement indicates the ESS(s) or IBSS(s) for which a beacon report is received.
	SsId []string `json:"ssId"`

	StaId *StaIdentity `json:"staId"`
}
