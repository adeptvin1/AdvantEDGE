/*
 * Copyright (c) 2019  InterDigital Communications, Inc
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
 * Radio Network Information Service is AdvantEDGE's implementation of [ETSI MEC ISG MEC012 RNI API](http://www.etsi.org/deliver/etsi_gs/MEC/001_099/012/01.01.01_60/gs_MEC012v010101p.pdf) <p>[Copyright (c) ETSI 2017](https://forge.etsi.org/etsi-forge-copyright-notice.txt) <p>**Micro-service**<br>[meep-rnis](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-rnis) <p>**Type & Usage**<br>Edge Service used by edge applications that want to get information about radio conditions in the network <p>**Details**<br>API details available at _your-AdvantEDGE-ip-address/api_
 *
 * API version: 1.1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Trigger : Unique identifier for the mobile edge application instance
type Trigger string

// List of Trigger
const (
	NOT_AVAILABLE_Trigger                             Trigger = "NOT_AVAILABLE"
	PERIODICAL_REPORT_STRONGEST_CELLS_Trigger         Trigger = "PERIODICAL_REPORT_STRONGEST_CELLS"
	PERIODICAL_REPORT_STRONGEST_CELLS_FOR_SON_Trigger Trigger = "PERIODICAL_REPORT_STRONGEST_CELLS_FOR_SON"
	PERIODICAL_REPORT_CGI_Trigger                     Trigger = "PERIODICAL_REPORT_CGI"
	EVENT_A1_Trigger                                  Trigger = "EVENT_A1"
	EVENT_A2_Trigger                                  Trigger = "EVENT_A2"
	EVENT_A3_Trigger                                  Trigger = "EVENT_A3"
	EVENT_A4_Trigger                                  Trigger = "EVENT_A4"
	EVENT_A5_Trigger                                  Trigger = "EVENT_A5"
	EVENT_A6_Trigger                                  Trigger = "EVENT_A6"
	EVENT_B1_Trigger                                  Trigger = "EVENT_B1"
	EVENT_B2_Trigger                                  Trigger = "EVENT_B2"
	EVENT_C1_Trigger                                  Trigger = "EVENT_C1"
	EVENT_C2_Trigger                                  Trigger = "EVENT_C2"
	EVENT_W1_Trigger                                  Trigger = "EVENT_W1"
	EVENT_W2_Trigger                                  Trigger = "EVENT_W2"
	EVENT_W3_Trigger                                  Trigger = "EVENT_W3"
)
