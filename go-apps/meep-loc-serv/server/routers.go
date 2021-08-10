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
 * AdvantEDGE Location Service REST API
 *
 * Location Service is AdvantEDGE's implementation of [ETSI MEC ISG MEC013 Location API](http://www.etsi.org/deliver/etsi_gs/MEC/001_099/013/02.01.01_60/gs_mec013v020101p.pdf) <p>The API is based on the Open Mobile Alliance's specification RESTful Network API for Zonal Presence <p>[Copyright (c) ETSI 2017](https://forge.etsi.org/etsi-forge-copyright-notice.txt) <p>**Micro-service**<br>[meep-loc-serv](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-loc-serv) <p>**Type & Usage**<br>Edge Service used by edge applications that want to get information about Users (UE) and Zone locations <p>**Details**<br>API details available at _your-AdvantEDGE-ip-address/api_ <p>AdvantEDGE supports a selected subset of Location API endpoints (see below)
 *
 * API version: 2.1.1
 * Contact: AdvantEDGE@InterDigital.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package server

import (
	"fmt"
	"net/http"
	"strings"

	httpLog "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-http-logger"
	met "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-metrics"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	var handler http.Handler
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		handler = met.MetricsHandler(handler, sandboxName, serviceName)
		handler = httpLog.LogRx(handler, "")
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	// Path prefix router order is important
	// Service Api files
	handler = http.StripPrefix("/location/v2/api/", http.FileServer(http.Dir("./api/")))
	router.
		PathPrefix("/location/v2/api/").
		Name("Api").
		Handler(handler)
	// User supplied service API files
	handler = http.StripPrefix("/location/v2/user-api/", http.FileServer(http.Dir("./user-api/")))
	router.
		PathPrefix("/location/v2/user-api/").
		Name("UserApi").
		Handler(handler)

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/location/v2/",
		Index,
	},

	Route{
		"ApByIdGET",
		strings.ToUpper("Get"),
		"/location/v2/queries/zones/{zoneId}/accessPoints/{accessPointId}",
		ApByIdGET,
	},

	Route{
		"ApGET",
		strings.ToUpper("Get"),
		"/location/v2/queries/zones/{zoneId}/accessPoints",
		ApGET,
	},

	Route{
		"AreaCircleSubDELETE",
		strings.ToUpper("Delete"),
		"/location/v2/subscriptions/area/circle/{subscriptionId}",
		AreaCircleSubDELETE,
	},

	Route{
		"AreaCircleSubGET",
		strings.ToUpper("Get"),
		"/location/v2/subscriptions/area/circle/{subscriptionId}",
		AreaCircleSubGET,
	},

	Route{
		"AreaCircleSubListGET",
		strings.ToUpper("Get"),
		"/location/v2/subscriptions/area/circle",
		AreaCircleSubListGET,
	},

	Route{
		"AreaCircleSubPOST",
		strings.ToUpper("Post"),
		"/location/v2/subscriptions/area/circle",
		AreaCircleSubPOST,
	},

	Route{
		"AreaCircleSubPUT",
		strings.ToUpper("Put"),
		"/location/v2/subscriptions/area/circle/{subscriptionId}",
		AreaCircleSubPUT,
	},

	Route{
		"DistanceGET",
		strings.ToUpper("Get"),
		"/location/v2/queries/distance",
		DistanceGET,
	},

	Route{
		"DistanceSubDELETE",
		strings.ToUpper("Delete"),
		"/location/v2/subscriptions/distance/{subscriptionId}",
		DistanceSubDELETE,
	},

	Route{
		"DistanceSubGET",
		strings.ToUpper("Get"),
		"/location/v2/subscriptions/distance/{subscriptionId}",
		DistanceSubGET,
	},

	Route{
		"DistanceSubListGET",
		strings.ToUpper("Get"),
		"/location/v2/subscriptions/distance",
		DistanceSubListGET,
	},

	Route{
		"DistanceSubPOST",
		strings.ToUpper("Post"),
		"/location/v2/subscriptions/distance",
		DistanceSubPOST,
	},

	Route{
		"DistanceSubPUT",
		strings.ToUpper("Put"),
		"/location/v2/subscriptions/distance/{subscriptionId}",
		DistanceSubPUT,
	},

	Route{
		"Mec011AppTerminationPOST",
		strings.ToUpper("Post"),
		"/location/v2/notifications/mec011/appTermination",
		Mec011AppTerminationPOST,
	},

	Route{
		"PeriodicSubDELETE",
		strings.ToUpper("Delete"),
		"/location/v2/subscriptions/periodic/{subscriptionId}",
		PeriodicSubDELETE,
	},

	Route{
		"PeriodicSubGET",
		strings.ToUpper("Get"),
		"/location/v2/subscriptions/periodic/{subscriptionId}",
		PeriodicSubGET,
	},

	Route{
		"PeriodicSubListGET",
		strings.ToUpper("Get"),
		"/location/v2/subscriptions/periodic",
		PeriodicSubListGET,
	},

	Route{
		"PeriodicSubPOST",
		strings.ToUpper("Post"),
		"/location/v2/subscriptions/periodic",
		PeriodicSubPOST,
	},

	Route{
		"PeriodicSubPUT",
		strings.ToUpper("Put"),
		"/location/v2/subscriptions/periodic/{subscriptionId}",
		PeriodicSubPUT,
	},

	Route{
		"UserTrackingSubDELETE",
		strings.ToUpper("Delete"),
		"/location/v2/subscriptions/userTracking/{subscriptionId}",
		UserTrackingSubDELETE,
	},

	Route{
		"UserTrackingSubGET",
		strings.ToUpper("Get"),
		"/location/v2/subscriptions/userTracking/{subscriptionId}",
		UserTrackingSubGET,
	},

	Route{
		"UserTrackingSubListGET",
		strings.ToUpper("Get"),
		"/location/v2/subscriptions/userTracking",
		UserTrackingSubListGET,
	},

	Route{
		"UserTrackingSubPOST",
		strings.ToUpper("Post"),
		"/location/v2/subscriptions/userTracking",
		UserTrackingSubPOST,
	},

	Route{
		"UserTrackingSubPUT",
		strings.ToUpper("Put"),
		"/location/v2/subscriptions/userTracking/{subscriptionId}",
		UserTrackingSubPUT,
	},

	Route{
		"UsersGET",
		strings.ToUpper("Get"),
		"/location/v2/queries/users",
		UsersGET,
	},

	Route{
		"ZonalTrafficSubDELETE",
		strings.ToUpper("Delete"),
		"/location/v2/subscriptions/zonalTraffic/{subscriptionId}",
		ZonalTrafficSubDELETE,
	},

	Route{
		"ZonalTrafficSubGET",
		strings.ToUpper("Get"),
		"/location/v2/subscriptions/zonalTraffic/{subscriptionId}",
		ZonalTrafficSubGET,
	},

	Route{
		"ZonalTrafficSubListGET",
		strings.ToUpper("Get"),
		"/location/v2/subscriptions/zonalTraffic",
		ZonalTrafficSubListGET,
	},

	Route{
		"ZonalTrafficSubPOST",
		strings.ToUpper("Post"),
		"/location/v2/subscriptions/zonalTraffic",
		ZonalTrafficSubPOST,
	},

	Route{
		"ZonalTrafficSubPUT",
		strings.ToUpper("Put"),
		"/location/v2/subscriptions/zonalTraffic/{subscriptionId}",
		ZonalTrafficSubPUT,
	},

	Route{
		"ZoneStatusSubDELETE",
		strings.ToUpper("Delete"),
		"/location/v2/subscriptions/zoneStatus/{subscriptionId}",
		ZoneStatusSubDELETE,
	},

	Route{
		"ZoneStatusSubGET",
		strings.ToUpper("Get"),
		"/location/v2/subscriptions/zoneStatus/{subscriptionId}",
		ZoneStatusSubGET,
	},

	Route{
		"ZoneStatusSubListGET",
		strings.ToUpper("Get"),
		"/location/v2/subscriptions/zoneStatus",
		ZoneStatusSubListGET,
	},

	Route{
		"ZoneStatusSubPOST",
		strings.ToUpper("Post"),
		"/location/v2/subscriptions/zoneStatus",
		ZoneStatusSubPOST,
	},

	Route{
		"ZoneStatusSubPUT",
		strings.ToUpper("Put"),
		"/location/v2/subscriptions/zoneStatus/{subscriptionId}",
		ZoneStatusSubPUT,
	},

	Route{
		"ZonesGET",
		strings.ToUpper("Get"),
		"/location/v2/queries/zones",
		ZonesGET,
	},

	Route{
		"ZonesGetById",
		strings.ToUpper("Get"),
		"/location/v2/queries/zones/{zoneId}",
		ZonesGetById,
	},
}
