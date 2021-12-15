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
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"

	couch "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-couch"
	dkm "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-data-key-mgr"
	dataModel "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-data-model"
	gc "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-gc"
	log "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-logger"
	mod "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-model"
	mq "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-mq"
	redis "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-redis"
	ss "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-sandbox-store"
	sam "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-swagger-api-mgr"
)

type Scenario struct {
	Name string `json:"name,omitempty"`
}

type PlatformCtrl struct {
	scenarioStore    *couch.Connector
	rc               *redis.Connector
	sandboxStore     *ss.SandboxStore
	mqGlobal         *mq.MsgQueue
	apiMgr           *sam.SwaggerApiMgr
	garbageCollector *gc.GarbageCollector
}

const scenarioDBName = "scenarios"
const redisTable = 0
const serviceName = "Platform Controller"
const moduleName = "meep-platform-ctrl"
const moduleNamespace = "default"

// MQ payload fields
const fieldSandboxName = "sandbox-name"
const fieldScenarioName = "scenario-name"

// Declare as variables to enable overwrite in test
var couchDBAddr = "http://meep-couchdb-svc-couchdb:5984/"
var redisDBAddr = "meep-redis-master:6379"
var influxDBAddr = "http://meep-influxdb.default.svc.cluster.local:8086"

// Platform Controller
var pfmCtrl *PlatformCtrl

// Init Initializes the Platform Controller
func Init() (err error) {
	log.Debug("Init")

	// Seed rand
	rand.Seed(time.Now().UnixNano())

	// Create new Platform Controller
	pfmCtrl = new(PlatformCtrl)

	// Retrieve Garbage collection configuration from environment variables
	gcEnabledStr := strings.TrimSpace(os.Getenv("MEEP_GC_ENABLED"))
	gcEnabled, err := strconv.ParseBool(gcEnabledStr)
	if err != nil {
		gcEnabled = false
	}
	log.Info("MEEP_GC_ENABLED: ", gcEnabledStr)
	if gcEnabled {

		// Retrieve & parse GC Interval from environment variable
		interval := strings.TrimSpace(os.Getenv("MEEP_GC_INTERVAL"))
		if interval == "" {
			err = errors.New("MEEP_GC_INTERVAL variable not set")
			log.Error(err.Error())
			return err
		}
		duration, err := time.ParseDuration(interval)
		if err != nil {
			log.Error(err.Error())
			return err
		}
		log.Info("MEEP_GC_INTERVAL: ", interval)

		// Retrieve run-on-start flag from environment variable
		runOnStartStr := strings.TrimSpace(os.Getenv("MEEP_GC_RUN_ON_START"))
		runOnStart, err := strconv.ParseBool(runOnStartStr)
		if err != nil {
			runOnStart = false
		}
		log.Info("MEEP_GC_RUN_ON_START: ", runOnStartStr)

		// Retrieve Redis enabled flag from environment variable
		redisEnabledStr := strings.TrimSpace(os.Getenv("MEEP_GC_REDIS_ENABLED"))
		redisEnabled, err := strconv.ParseBool(redisEnabledStr)
		if err != nil {
			redisEnabled = false
		}
		log.Info("MEEP_GC_REDIS_ENABLED: ", redisEnabledStr)

		// Retrieve Redis enabled flag from environment variable
		influxEnabledStr := strings.TrimSpace(os.Getenv("MEEP_GC_INFLUX_ENABLED"))
		influxEnabled, err := strconv.ParseBool(influxEnabledStr)
		if err != nil {
			influxEnabled = false
		}
		log.Info("MEEP_GC_INFLUX_ENABLED: ", influxEnabledStr)

		// Retrieve Influx exceptions list from environment variable
		influxExceptions := make([]string, 0)
		influxExceptionsStr := strings.TrimSpace(os.Getenv("MEEP_GC_INFLUX_EXCEPTIONS"))
		log.Info("MEEP_GC_INFLUX_EXCEPTIONS: ", influxExceptionsStr)
		if influxExceptionsStr != "" {
			influxExceptions = strings.Split(influxExceptionsStr, ",")
		}

		// Retrieve Redis enabled flag from environment variable
		postgisEnabledStr := strings.TrimSpace(os.Getenv("MEEP_GC_POSTGIS_ENABLED"))
		postgisEnabled, err := strconv.ParseBool(postgisEnabledStr)
		if err != nil {
			postgisEnabled = false
		}
		log.Info("MEEP_GC_POSTGIS_ENABLED: ", postgisEnabledStr)

		// Create Garbage Collector
		gcCfg := gc.GarbageCollectorCfg{
			Redis:            redisEnabled,
			RedisAddr:        redisDBAddr,
			RedisTable:       0,
			Influx:           influxEnabled,
			InfluxAddr:       influxDBAddr,
			InfluxExceptions: influxExceptions,
			Postgis:          postgisEnabled,
			RunOnStart:       runOnStart,
			Interval:         duration,
		}
		pfmCtrl.garbageCollector, err = gc.NewGarbageCollector(gcCfg)
		if err != nil {
			log.Error("Failed to create Garbage Collector with error: ", err.Error())
			return err
		}
		log.Info("Garbage Collector created")
	}

	// Create message queue
	pfmCtrl.mqGlobal, err = mq.NewMsgQueue(mq.GetGlobalName(), moduleName, moduleNamespace, redisDBAddr)
	if err != nil {
		log.Error("Failed to create Message Queue with error: ", err)
		return err
	}
	log.Info("Message Queue created")

	// Connect to Redis DB
	pfmCtrl.rc, err = redis.NewConnector(redisDBAddr, redisTable)
	if err != nil {
		log.Error("Failed connection to Redis DB. Error: ", err)
		return err
	}
	log.Info("Connected to Redis DB")

	// Create Swagger API Manager
	pfmCtrl.apiMgr, err = sam.NewSwaggerApiMgr(moduleName, "", "", pfmCtrl.mqGlobal)
	if err != nil {
		log.Error("Failed to create Swagger API Manager. Error: ", err)
		return err
	}
	log.Info("Swagger API Manager created")

	// Connect to Scenario Store
	pfmCtrl.scenarioStore, err = couch.NewConnector(couchDBAddr, scenarioDBName)
	if err != nil {
		log.Error("Failed connection to Scenario Store. Error: ", err)
		return err
	}
	log.Info("Connected to Scenario Store")

	// Retrieve scenario list from DB
	_, scenarioList, err := pfmCtrl.scenarioStore.GetDocList()
	if err != nil {
		log.Error(err.Error())
		return err
	}

	// Validate DB scenarios & upgrade them if compatible
	for _, scenario := range scenarioList {
		validScenario, status, err := mod.ValidateScenario(scenario, "")
		if err == nil && status == mod.ValidatorStatusUpdated {
			// Retrieve scenario name
			s := new(Scenario)
			err = json.Unmarshal(validScenario, s)
			if err != nil || s.Name == "" {
				return errors.New("Failed to get scenario name from valid scenario")
			}

			// Update scenario in DB
			rev, err := pfmCtrl.scenarioStore.UpdateDoc(s.Name, validScenario)
			if err != nil {
				return errors.New("Failed to update scenario with error: " + err.Error())
			}
			log.Debug("Scenario updated with rev: ", rev)
		}
	}

	// Connect to Sandbox Store
	pfmCtrl.sandboxStore, err = ss.NewSandboxStore(redisDBAddr)
	if err != nil {
		log.Error("Failed connection to Sandbox Store: ", err.Error())
		return err
	}
	log.Info("Connected to Sandbox Store")

	log.Info("Platform Controller initialized")
	return nil
}

// Run Starts the Platform Controller
func Run() (err error) {

	// Start Garbage Collector
	if pfmCtrl.garbageCollector != nil {
		err = pfmCtrl.garbageCollector.Start()
		if err != nil {
			log.Error("Failed to start Garbage Collector with error: ", err.Error())
			return err
		}
		log.Info("Garbage Collector started")
	}

	// Start Swagger API Manager (provider & aggregator)
	err = pfmCtrl.apiMgr.Start(true, true)
	if err != nil {
		log.Error("Failed to start Swagger API Manager with error: ", err.Error())
		return err
	}
	log.Info("Swagger API Manager started")

	log.Info("Platform Controller started")
	return nil
}

func Stop() {
	if pfmCtrl == nil {
		return
	}

	// Stop Swagger API Manager
	if pfmCtrl.apiMgr != nil {
		_ = pfmCtrl.apiMgr.Stop()
	}

	// Stop Garbage Collector
	if pfmCtrl.garbageCollector != nil {
		_ = pfmCtrl.garbageCollector.Stop()
	}
}

// Create a new scenario in the scenario store
// POST /scenario/{name}
func pcCreateScenario(w http.ResponseWriter, r *http.Request) {
	log.Debug("pcCreateScenario")

	// Get scenario name from request parameters
	vars := mux.Vars(r)
	scenarioName := vars["name"]
	log.Debug("Scenario name: ", scenarioName)

	// Retrieve scenario from request body
	if r.Body == nil {
		err := errors.New("Request body is missing")
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate scenario
	validScenario, _, err := mod.ValidateScenario(b, scenarioName)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Add new scenario to DB
	rev, err := pfmCtrl.scenarioStore.AddDoc(scenarioName, validScenario)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	log.Debug("Scenario added with rev: ", rev)

	// OK
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// Delete scenario from scenario store
// DELETE /scenarios/{name}
func pcDeleteScenario(w http.ResponseWriter, r *http.Request) {
	log.Debug("pcDeleteScenario")

	// Get scenario name from request parameters
	vars := mux.Vars(r)
	scenarioName := vars["name"]
	log.Debug("Scenario name: ", scenarioName)

	// Remove scenario from DB
	err := pfmCtrl.scenarioStore.DeleteDoc(scenarioName)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// OK
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// Remove all scenarios from sceanrio store
// DELETE /scenarios
func pcDeleteScenarioList(w http.ResponseWriter, r *http.Request) {
	log.Debug("pcDeleteScenarioList")

	// Remove all scenario from DB
	err := pfmCtrl.scenarioStore.DeleteAllDocs()
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// Retrieve scenario from scenario store
// GET /scenarios/{name}
func pcGetScenario(w http.ResponseWriter, r *http.Request) {
	log.Debug("pcGetScenario")

	// Get scenario name from request parameters
	vars := mux.Vars(r)
	scenarioName := vars["name"]
	log.Debug("Scenario name: ", scenarioName)

	// Validate scenario name
	if scenarioName == "" {
		log.Debug("Invalid scenario name")
		http.Error(w, "Invalid scenario name "+scenarioName, http.StatusBadRequest)
		return
	}

	// Retrieve scenario from DB
	var scenario []byte
	scenario, err := pfmCtrl.scenarioStore.GetDoc(false, scenarioName)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	s, err := mod.JSONMarshallScenario(scenario)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, s)
}

// Retrieve all scenarios from scenario store
// GET /scenarios
func pcGetScenarioList(w http.ResponseWriter, r *http.Request) {
	log.Debug("pcGetScenarioList")

	// Retrieve scenario list from DB
	_, scenarioList, err := pfmCtrl.scenarioStore.GetDocList()
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	sl, err := mod.JSONMarshallScenarioList(scenarioList)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, sl)
}

// Update scenario in scenario store
// PUT /scenarios/{name}
func pcSetScenario(w http.ResponseWriter, r *http.Request) {
	log.Debug("pcSetScenario")

	// Get scenario name from request parameters
	vars := mux.Vars(r)
	scenarioName := vars["name"]
	log.Debug("Scenario name: ", scenarioName)

	// Retrieve scenario from request body
	if r.Body == nil {
		err := errors.New("Request body is missing")
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate scenario
	validScenario, _, err := mod.ValidateScenario(b, scenarioName)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update scenario in DB
	rev, err := pfmCtrl.scenarioStore.UpdateDoc(scenarioName, validScenario)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	log.Debug("Scenario updated with rev: ", rev)

	// OK
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// Create new Sandbox
// POST /sandboxes
func pcCreateSandbox(w http.ResponseWriter, r *http.Request) {
	log.Debug("pcCreateSandbox")

	// Retrieve sandbox config from request body
	var sandboxConfig dataModel.SandboxConfig
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&sandboxConfig)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get unique sandbox name
	sandboxName := getUniqueSandboxName()
	if sandboxName == "" {
		err = errors.New("Failed to generate a unique sandbox name")
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create sandbox in DB
	err = createSandbox(sandboxName, &sandboxConfig)
	if err != nil {
		log.Error("Failed to create sandbox with error: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Prepare response
	var sandbox dataModel.Sandbox
	sandbox.Name = sandboxName

	// Format response
	jsonResponse, err := json.Marshal(sandbox)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonResponse))
}

// Create new Sandbox with provided name
// POST /sandboxes/{name}
func pcCreateSandboxWithName(w http.ResponseWriter, r *http.Request) {
	log.Debug("pcCreateSandboxWithName")

	// Get sandbox name from request parameters
	vars := mux.Vars(r)
	sandboxName := vars["name"]
	log.Debug("Sandbox to create: ", sandboxName)

	// Retrieve sandbox config from request body
	var sandboxConfig dataModel.SandboxConfig
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&sandboxConfig)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Make sure sandbox does not already exist
	if sbox, _ := pfmCtrl.sandboxStore.Get(sandboxName); sbox != nil {
		err = errors.New("Sandbox already exists")
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	// Create sandbox in DB
	err = createSandbox(sandboxName, &sandboxConfig)
	if err != nil {
		log.Error("Failed to create sandbox with error: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Prepare response
	var sandbox dataModel.Sandbox
	sandbox.Name = sandboxName

	// Format response
	jsonResponse, err := json.Marshal(sandbox)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonResponse))
}

// Delete Sandbox with provided name
// DELETE /sandboxes/{name}
func pcDeleteSandbox(w http.ResponseWriter, r *http.Request) {
	log.Debug("pcDeleteSandbox")

	// Get sandbox name from request parameters
	vars := mux.Vars(r)
	sandboxName := vars["name"]
	log.Debug("Sandbox to delete: ", sandboxName)

	// Make sure sandbox exists
	if sbox, _ := pfmCtrl.sandboxStore.Get(sandboxName); sbox == nil {
		err := errors.New("Sandbox not found")
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Delete sandbox
	deleteSandbox(sandboxName)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// Delete all Sandboxes
// DELETE /sandboxes
func pcDeleteSandboxList(w http.ResponseWriter, r *http.Request) {
	log.Debug("pcDeleteSandboxList")

	// Get all sandboxes
	sboxMap, err := pfmCtrl.sandboxStore.GetAll()
	if err != nil || sboxMap == nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Delete all sandboxes
	for _, sbox := range sboxMap {
		deleteSandbox(sbox.Name)
	}

	// Flush sandbox store
	pfmCtrl.sandboxStore.Flush()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// Retrieve Sandbox with provided name
// GET /sandboxes/{name}
func pcGetSandbox(w http.ResponseWriter, r *http.Request) {
	log.Debug("pcGetSandbox")

	// Get sandbox name from request parameters
	vars := mux.Vars(r)
	sandboxName := vars["name"]
	log.Debug("Sandbox to retrieve: ", sandboxName)

	// Get sandbox
	sbox, _ := pfmCtrl.sandboxStore.Get(sandboxName)
	if sbox == nil {
		err := errors.New("Sandbox not found")
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Prepare response
	var sandbox dataModel.Sandbox
	sandbox.Name = sbox.Name

	// Format response
	jsonResponse, err := json.Marshal(sandbox)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonResponse))
}

// Retrieve all Sandboxes
// GET /sandboxes
func pcGetSandboxList(w http.ResponseWriter, r *http.Request) {
	log.Debug("pcGetSandboxList")

	// Get all sandboxes
	sboxMap, err := pfmCtrl.sandboxStore.GetAll()
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Update sandbox list
	var sandboxList dataModel.SandboxList
	for _, sbox := range sboxMap {
		var sandbox dataModel.Sandbox
		sandbox.Name = sbox.Name
		sandboxList.Sandboxes = append(sandboxList.Sandboxes, sandbox)
	}

	// Format response
	jsonResponse, err := json.Marshal(sandboxList)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonResponse))
}

// Create new sandbox in store and publish updagte
func createSandbox(sandboxName string, sandboxConfig *dataModel.SandboxConfig) (err error) {

	// Flush sandbox data
	_ = pfmCtrl.rc.DBFlush(dkm.GetKeyRoot(sandboxName))

	// Create sandbox in DB
	sbox := new(ss.Sandbox)
	sbox.Name = sandboxName
	sbox.ScenarioName = sandboxConfig.ScenarioName
	err = pfmCtrl.sandboxStore.Set(sbox)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	// Send message to create sandbox
	msg := pfmCtrl.mqGlobal.CreateMsg(mq.MsgSandboxCreate, mq.TargetAll, mq.TargetAll)
	msg.Payload[fieldSandboxName] = sandboxName
	msg.Payload[fieldScenarioName] = sandboxConfig.ScenarioName
	log.Debug("TX MSG: ", mq.PrintMsg(msg))
	err = pfmCtrl.mqGlobal.SendMsg(msg)
	if err != nil {
		log.Error("Failed to send message. Error: ", err.Error())
		return err
	}

	return nil
}

func deleteSandbox(sandboxName string) {
	if sandboxName == "" {
		return
	}

	// Remove sandbox from store
	pfmCtrl.sandboxStore.Del(sandboxName)

	// Send message to destroy sandbox
	msg := pfmCtrl.mqGlobal.CreateMsg(mq.MsgSandboxDestroy, mq.TargetAll, mq.TargetAll)
	msg.Payload[fieldSandboxName] = sandboxName
	log.Debug("TX MSG: ", mq.PrintMsg(msg))
	err := pfmCtrl.mqGlobal.SendMsg(msg)
	if err != nil {
		log.Error("Failed to send message. Error: ", err.Error())
	}
}

func getUniqueSandboxName() (name string) {
	retryCount := 3
	for i := 0; i < retryCount; i++ {
		// sandboxName = "sbox-" + xid.New().String()
		randName := "sbx" + randSeq(7)
		if sbox, _ := pfmCtrl.sandboxStore.Get(randName); sbox == nil {
			name = randName
			break
		}
	}
	return name
}

var charset = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
