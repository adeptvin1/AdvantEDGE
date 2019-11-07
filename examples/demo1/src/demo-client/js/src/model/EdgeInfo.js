/*
 * MEEP Demo App API
 * Copyright (c) 2019  InterDigital Communications, Inc Licensed under the Apache License, Version 2.0 (the \"License\"); you may not use this file except in compliance with the License. You may obtain a copy of the License at      http://www.apache.org/licenses/LICENSE-2.0  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an \"AS IS\" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License. 
 *
 * OpenAPI spec version: 0.0.1
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
    define(['ApiClient'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'));
  } else {
    // Browser globals (root is window)
    if (!root.MeepDemoAppApi) {
      root.MeepDemoAppApi = {};
    }
    root.MeepDemoAppApi.EdgeInfo = factory(root.MeepDemoAppApi.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';

  /**
   * The EdgeInfo model module.
   * @module model/EdgeInfo
   * @version 0.0.1
   */

  /**
   * Constructs a new <code>EdgeInfo</code>.
   * Edge app basic information object
   * @alias module:model/EdgeInfo
   * @class
   */
  var exports = function() {
  };

  /**
   * Constructs a <code>EdgeInfo</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/EdgeInfo} obj Optional instance to populate.
   * @return {module:model/EdgeInfo} The populated <code>EdgeInfo</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();
      if (data.hasOwnProperty('svc'))
        obj.svc = ApiClient.convertToType(data['svc'], 'String');
      if (data.hasOwnProperty('name'))
        obj.name = ApiClient.convertToType(data['name'], 'String');
      if (data.hasOwnProperty('ip'))
        obj.ip = ApiClient.convertToType(data['ip'], 'String');
    }
    return obj;
  }

  /**
   * Edge app service
   * @member {String} svc
   */
  exports.prototype.svc = undefined;

  /**
   * Edge app local name
   * @member {String} name
   */
  exports.prototype.name = undefined;

  /**
   * IP address where the local edge app reside
   * @member {String} ip
   */
  exports.prototype.ip = undefined;

  return exports;

}));
