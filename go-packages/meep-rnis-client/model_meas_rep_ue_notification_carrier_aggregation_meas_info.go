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
 * AdvantEDGE Radio Network Information Service REST API
 *
 * Radio Network Information Service is AdvantEDGE's implementation of [ETSI MEC ISG MEC012 RNI API](http://www.etsi.org/deliver/etsi_gs/MEC/001_099/012/02.01.01_60/gs_MEC012v020101p.pdf) <p>[Copyright (c) ETSI 2017](https://forge.etsi.org/etsi-forge-copyright-notice.txt) <p>**Micro-service**<br>[meep-rnis](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-rnis) <p>**Type & Usage**<br>Edge Service used by edge applications that want to get information about radio conditions in the network <p>**Details**<br>API details available at _your-AdvantEDGE-ip-address/api_ <p>AdvantEDGE supports a selected subset of RNI API endpoints (see below) and a subset of subscription types. <p>Supported subscriptions: <p> - CellChangeSubscription <p> - RabEstSubscription <p> - RabRelSubscription
 *
 * API version: 2.1.1
 * Contact: AdvantEDGE@InterDigital.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type MeasRepUeNotificationCarrierAggregationMeasInfo struct {
	CellIdNei string `json:"cellIdNei,omitempty"`
	CellIdSrv string `json:"cellIdSrv,omitempty"`
	// Reference Signal Received Power as defined in ETSI TS 136 214 [i.5].
	RsrpNei int32 `json:"rsrpNei,omitempty"`
	// Extended Reference Signal Received Power, with value mapping defined in ETSI TS 136 133 [i.16].
	RsrpNeiEx int32 `json:"rsrpNeiEx,omitempty"`
	// Reference Signal Received Power as defined in ETSI TS 136 214 [i.5].
	RsrpSrv int32 `json:"rsrpSrv,omitempty"`
	// Extended Reference Signal Received Power, with value mapping defined in ETSI TS 136 133 [i.16].
	RsrpSrvEx int32 `json:"rsrpSrvEx,omitempty"`
	// Reference Signal Received Quality as defined in ETSI TS 136 214 [i.5].
	RsrqNei int32 `json:"rsrqNei,omitempty"`
	// Extended Reference Signal Received Quality, with value mapping defined in ETSI TS 136 133 [i.16].
	RsrqNeiEx int32 `json:"rsrqNeiEx,omitempty"`
	// Reference Signal Received Quality as defined in ETSI TS 136 214 [i.5].
	RsrqSrv int32 `json:"rsrqSrv,omitempty"`
	// Extended Reference Signal Received Quality, with value mapping defined in ETSI TS 136 133 [i.16].
	RsrqSrvEx int32 `json:"rsrqSrvEx,omitempty"`
	// Reference Signal \"Signal to Interference plus Noise Ratio\", with value mapping defined in ETSI TS 136 133 [i.16].
	SinrNei int32 `json:"sinrNei,omitempty"`
	// Reference Signal \"Signal to Interference plus Noise Ratio\", with value mapping defined in ETSI TS 136 133 [i.16].
	SinrSrv int32 `json:"sinrSrv,omitempty"`
}
