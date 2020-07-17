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
* Enum class V1CloningKind.
* @enum {}
* @readonly
*/
export default class V1CloningKind {
    
        /**
         * value: "copy"
         * @const
         */
        "copy" = "copy";

    
        /**
         * value: "restart"
         * @const
         */
        "restart" = "restart";

    
        /**
         * value: "cache"
         * @const
         */
        "cache" = "cache";

    
        /**
         * value: "schedule"
         * @const
         */
        "schedule" = "schedule";

    

    /**
    * Returns a <code>V1CloningKind</code> enum value from a Javascript object name.
    * @param {Object} data The plain JavaScript object containing the name of the enum value.
    * @return {module:model/V1CloningKind} The enum <code>V1CloningKind</code> value.
    */
    static constructFromObject(object) {
        return object;
    }
}

