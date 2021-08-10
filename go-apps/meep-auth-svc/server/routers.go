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
 * AdvantEDGE Auth Service REST API
 *
 * This API provides microservice API authentication & authorization services <p>**Micro-service**<br>[meep-auth](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-auth) <p>**Type & Usage**<br>Platform interface used by ingress to authenticate & authorize microservice API access <p>**Details**<br>API details available at _your-AdvantEDGE-ip-address/api_
 *
 * API version: 1.0.0
 * Contact: AdvantEDGE@InterDigital.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

import (
	"fmt"
	"net/http"
	"strings"

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
		handler = met.MetricsHandler(handler, "", serviceName)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	// Path prefix router order is important
	// Service Api files
	handler = http.StripPrefix("/auth/v1/api/", http.FileServer(http.Dir("./api/")))
	router.
		PathPrefix("/auth/v1/api/").
		Name("Api").
		Handler(handler)
	// User supplied service API files
	handler = http.StripPrefix("/auth/v1/user-api/", http.FileServer(http.Dir("./user-api/")))
	router.
		PathPrefix("/auth/v1/user-api/").
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
		"/auth/v1/",
		Index,
	},

	Route{
		"Authenticate",
		strings.ToUpper("Get"),
		"/auth/v1/authenticate",
		Authenticate,
	},

	Route{
		"Authorize",
		strings.ToUpper("Get"),
		"/auth/v1/authorize",
		Authorize,
	},

	Route{
		"Login",
		strings.ToUpper("Get"),
		"/auth/v1/login",
		Login,
	},

	Route{
		"LoginSupported",
		strings.ToUpper("Get"),
		"/auth/v1/loginSupported",
		LoginSupported,
	},

	Route{
		"LoginUser",
		strings.ToUpper("Post"),
		"/auth/v1/login",
		LoginUser,
	},

	Route{
		"Logout",
		strings.ToUpper("Get"),
		"/auth/v1/logout",
		Logout,
	},

	Route{
		"TriggerWatchdog",
		strings.ToUpper("Post"),
		"/auth/v1/watchdog",
		TriggerWatchdog,
	},
}
