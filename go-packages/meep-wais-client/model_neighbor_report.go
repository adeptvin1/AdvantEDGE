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

type NeighborReport struct {
	// Relative value indicating the preferred ordering for this BSS as a transition candidate for roaming.  255 indicating the most preferred candidate and 1 indicating the least preferred candidate, as defined in Table 9-152 within IEEE 802.11-2016 [8].
	BssTransitionCandidatePreference int32 `json:"bssTransitionCandidatePreference,omitempty"`
	// BSSID (MAC address) of the Access Point that is being reported.
	Bssid     string     `json:"bssid"`
	BssidInfo *BssidInfo `json:"bssidInfo"`
	// Channel field indicates a channel number, which is interpreted in the context of the indicated operating class. Channel numbers are defined in Annex E within IEEE 802.11-2016 [8].
	Channel int32 `json:"channel"`
	// Measurement ID of the Measurement configuration applied to this Neighbor Report.
	MeasurementId string `json:"measurementId"`
	// Operating Class field indicates an operating class value as defined in Annex E within IEEE 802.11-2016 [8].
	OperatingClass int32 `json:"operatingClass"`
	// PHY type of the AP indicated by this BSSID. It is an integer value coded according to the value of the dot11PHYType, Annex C within IEEE 802.11-2016 [8]. 2 = dsss 4 = ofdm 5 = hrdsss 6 = erp 7 = ht 8 = dmg 9 = vht 10 = tvht
	PhyType int32        `json:"phyType"`
	StaId   *StaIdentity `json:"staId,omitempty"`
}
