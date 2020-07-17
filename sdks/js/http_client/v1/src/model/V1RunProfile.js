// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.1.4-rc8
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The V1RunProfile model module.
 * @module model/V1RunProfile
 * @version 1.1.4-rc8
 */
class V1RunProfile {
    /**
     * Constructs a new <code>V1RunProfile</code>.
     * @alias module:model/V1RunProfile
     */
    constructor() { 
        
        V1RunProfile.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1RunProfile</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1RunProfile} obj Optional instance to populate.
     * @return {module:model/V1RunProfile} The populated <code>V1RunProfile</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1RunProfile();

            if (data.hasOwnProperty('uuid')) {
                obj['uuid'] = ApiClient.convertToType(data['uuid'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('tags')) {
                obj['tags'] = ApiClient.convertToType(data['tags'], ['String']);
            }
            if (data.hasOwnProperty('created_at')) {
                obj['created_at'] = ApiClient.convertToType(data['created_at'], 'Date');
            }
            if (data.hasOwnProperty('updated_at')) {
                obj['updated_at'] = ApiClient.convertToType(data['updated_at'], 'Date');
            }
            if (data.hasOwnProperty('frozen')) {
                obj['frozen'] = ApiClient.convertToType(data['frozen'], 'Boolean');
            }
            if (data.hasOwnProperty('disabled')) {
                obj['disabled'] = ApiClient.convertToType(data['disabled'], 'Boolean');
            }
            if (data.hasOwnProperty('agent')) {
                obj['agent'] = ApiClient.convertToType(data['agent'], 'String');
            }
            if (data.hasOwnProperty('queue')) {
                obj['queue'] = ApiClient.convertToType(data['queue'], 'String');
            }
            if (data.hasOwnProperty('namespace')) {
                obj['namespace'] = ApiClient.convertToType(data['namespace'], 'String');
            }
            if (data.hasOwnProperty('termination')) {
                obj['termination'] = ApiClient.convertToType(data['termination'], Object);
            }
            if (data.hasOwnProperty('environment')) {
                obj['environment'] = ApiClient.convertToType(data['environment'], Object);
            }
        }
        return obj;
    }


}

/**
 * @member {String} uuid
 */
V1RunProfile.prototype['uuid'] = undefined;

/**
 * @member {String} name
 */
V1RunProfile.prototype['name'] = undefined;

/**
 * @member {String} description
 */
V1RunProfile.prototype['description'] = undefined;

/**
 * @member {Array.<String>} tags
 */
V1RunProfile.prototype['tags'] = undefined;

/**
 * @member {Date} created_at
 */
V1RunProfile.prototype['created_at'] = undefined;

/**
 * @member {Date} updated_at
 */
V1RunProfile.prototype['updated_at'] = undefined;

/**
 * @member {Boolean} frozen
 */
V1RunProfile.prototype['frozen'] = undefined;

/**
 * @member {Boolean} disabled
 */
V1RunProfile.prototype['disabled'] = undefined;

/**
 * @member {String} agent
 */
V1RunProfile.prototype['agent'] = undefined;

/**
 * @member {String} queue
 */
V1RunProfile.prototype['queue'] = undefined;

/**
 * @member {String} namespace
 */
V1RunProfile.prototype['namespace'] = undefined;

/**
 * @member {Object} termination
 */
V1RunProfile.prototype['termination'] = undefined;

/**
 * @member {Object} environment
 */
V1RunProfile.prototype['environment'] = undefined;






export default V1RunProfile;

