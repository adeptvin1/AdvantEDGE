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
 * AdvantEDGE Sandbox Controller REST API
 * This API is the main Sandbox Controller API for scenario deployment & event injection <p>**Micro-service**<br>[meep-sandbox-ctrl](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-sandbox-ctrl) <p>**Type & Usage**<br>Platform runtime interface to manage active scenarios and inject events in AdvantEDGE platform <p>**Details**<br>API details available at _your-AdvantEDGE-ip-address/api_
 *
 * OpenAPI spec version: 1.0.0
 * Contact: AdvantEDGE@InterDigital.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.9
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/ApplicationState', 'model/ApplicationType'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./ApplicationState'), require('./ApplicationType'));
  } else {
    // Browser globals (root is window)
    if (!root.AdvantEdgeSandboxControllerRestApi) {
      root.AdvantEdgeSandboxControllerRestApi = {};
    }
    root.AdvantEdgeSandboxControllerRestApi.ApplicationInfo = factory(root.AdvantEdgeSandboxControllerRestApi.ApiClient, root.AdvantEdgeSandboxControllerRestApi.ApplicationState, root.AdvantEdgeSandboxControllerRestApi.ApplicationType);
  }
}(this, function(ApiClient, ApplicationState, ApplicationType) {
  'use strict';

  /**
   * The ApplicationInfo model module.
   * @module model/ApplicationInfo
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>ApplicationInfo</code>.
   * MEC Application instance information
   * @alias module:model/ApplicationInfo
   * @class
   * @param name {String} Application name
   * @param state {module:model/ApplicationState} 
   * @param mepName {String} MEP Name where application instance is running
   */
  var exports = function(name, state, mepName) {
    this.name = name;
    this.state = state;
    this.mepName = mepName;
  };

  /**
   * Constructs a <code>ApplicationInfo</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/ApplicationInfo} obj Optional instance to populate.
   * @return {module:model/ApplicationInfo} The populated <code>ApplicationInfo</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();
      if (data.hasOwnProperty('id'))
        obj.id = ApiClient.convertToType(data['id'], 'String');
      if (data.hasOwnProperty('name'))
        obj.name = ApiClient.convertToType(data['name'], 'String');
      if (data.hasOwnProperty('type'))
        obj.type = ApplicationType.constructFromObject(data['type']);
      if (data.hasOwnProperty('state'))
        obj.state = ApplicationState.constructFromObject(data['state']);
      if (data.hasOwnProperty('mepName'))
        obj.mepName = ApiClient.convertToType(data['mepName'], 'String');
      if (data.hasOwnProperty('version'))
        obj.version = ApiClient.convertToType(data['version'], 'String');
    }
    return obj;
  }

  /**
   * Application Instance UUID
   * @member {String} id
   */
  exports.prototype.id = undefined;

  /**
   * Application name
   * @member {String} name
   */
  exports.prototype.name = undefined;

  /**
   * @member {module:model/ApplicationType} type
   */
  exports.prototype.type = undefined;

  /**
   * @member {module:model/ApplicationState} state
   */
  exports.prototype.state = undefined;

  /**
   * MEP Name where application instance is running
   * @member {String} mepName
   */
  exports.prototype.mepName = undefined;

  /**
   * Application Version
   * @member {String} version
   */
  exports.prototype.version = undefined;

  return exports;

}));
