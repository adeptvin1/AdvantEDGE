/*
 * AdvantEDGE Radio Network Information Service REST API
 *
 * Radio Network Information Service is AdvantEDGE's implementation of [ETSI MEC ISG MEC012 RNI API](http://www.etsi.org/deliver/etsi_gs/MEC/001_099/012/02.01.01_60/gs_MEC012v020101p.pdf) <p>[Copyright (c) ETSI 2017](https://forge.etsi.org/etsi-forge-copyright-notice.txt) <p>**Micro-service**<br>[meep-rnis](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-rnis) <p>**Type & Usage**<br>Edge Service used by edge applications that want to get information about radio conditions in the network <p>**Details**<br>API details available at _your-AdvantEDGE-ip-address/api_
 *
 * API version: 2.1.1
 * Contact: AdvantEDGE@InterDigital.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package server

// The QoS information for the E-RAB.
type RabEstNotificationErabQosParametersQosInformation struct {
	// This attribute indicates the guaranteed downlink E-RAB Bit Rate as defined in ETSI TS 123 401 [i.4] for this bearer.
	ErabGbrDl int32 `json:"erabGbrDl"`
	// This attribute indicates the guaranteed uplink E-RAB Bit Rate as defined in ETSI TS 123 401 [i.4] for this bearer.
	ErabGbrUl int32 `json:"erabGbrUl"`
	// This attribute indicates the maximum downlink E-RAB Bit Rate as defined in ETSI TS 123 401 [i.4] for this bearer.
	ErabMbrDl int32 `json:"erabMbrDl"`
	// This attribute indicates the maximum uplink E-RAB Bit Rate as defined in ETSI TS 123 401 [i.4] for this bearer.
	ErabMbrUl int32 `json:"erabMbrUl"`
}