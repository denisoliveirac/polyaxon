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
import V1Cloning from './V1Cloning';
import V1Pipeline from './V1Pipeline';
import V1RunKind from './V1RunKind';
import V1RunSettings from './V1RunSettings';
import V1StatusCondition from './V1StatusCondition';
import V1Statuses from './V1Statuses';

/**
 * The V1Run model module.
 * @module model/V1Run
 * @version 1.1.4-rc8
 */
class V1Run {
    /**
     * Constructs a new <code>V1Run</code>.
     * @alias module:model/V1Run
     */
    constructor() { 
        
        V1Run.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1Run</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1Run} obj Optional instance to populate.
     * @return {module:model/V1Run} The populated <code>V1Run</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1Run();

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
            if (data.hasOwnProperty('deleted')) {
                obj['deleted'] = ApiClient.convertToType(data['deleted'], 'Boolean');
            }
            if (data.hasOwnProperty('user')) {
                obj['user'] = ApiClient.convertToType(data['user'], 'String');
            }
            if (data.hasOwnProperty('owner')) {
                obj['owner'] = ApiClient.convertToType(data['owner'], 'String');
            }
            if (data.hasOwnProperty('project')) {
                obj['project'] = ApiClient.convertToType(data['project'], 'String');
            }
            if (data.hasOwnProperty('created_at')) {
                obj['created_at'] = ApiClient.convertToType(data['created_at'], 'Date');
            }
            if (data.hasOwnProperty('updated_at')) {
                obj['updated_at'] = ApiClient.convertToType(data['updated_at'], 'Date');
            }
            if (data.hasOwnProperty('started_at')) {
                obj['started_at'] = ApiClient.convertToType(data['started_at'], 'Date');
            }
            if (data.hasOwnProperty('finished_at')) {
                obj['finished_at'] = ApiClient.convertToType(data['finished_at'], 'Date');
            }
            if (data.hasOwnProperty('run_time')) {
                obj['run_time'] = ApiClient.convertToType(data['run_time'], 'Number');
            }
            if (data.hasOwnProperty('is_managed')) {
                obj['is_managed'] = ApiClient.convertToType(data['is_managed'], 'String');
            }
            if (data.hasOwnProperty('content')) {
                obj['content'] = ApiClient.convertToType(data['content'], 'String');
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = V1Statuses.constructFromObject(data['status']);
            }
            if (data.hasOwnProperty('bookmarked')) {
                obj['bookmarked'] = ApiClient.convertToType(data['bookmarked'], 'Boolean');
            }
            if (data.hasOwnProperty('meta_info')) {
                obj['meta_info'] = ApiClient.convertToType(data['meta_info'], Object);
            }
            if (data.hasOwnProperty('is_helper')) {
                obj['is_helper'] = ApiClient.convertToType(data['is_helper'], 'Boolean');
            }
            if (data.hasOwnProperty('kind')) {
                obj['kind'] = V1RunKind.constructFromObject(data['kind']);
            }
            if (data.hasOwnProperty('meta_kind')) {
                obj['meta_kind'] = V1RunKind.constructFromObject(data['meta_kind']);
            }
            if (data.hasOwnProperty('hub')) {
                obj['hub'] = ApiClient.convertToType(data['hub'], 'String');
            }
            if (data.hasOwnProperty('inputs')) {
                obj['inputs'] = ApiClient.convertToType(data['inputs'], Object);
            }
            if (data.hasOwnProperty('outputs')) {
                obj['outputs'] = ApiClient.convertToType(data['outputs'], Object);
            }
            if (data.hasOwnProperty('original')) {
                obj['original'] = V1Cloning.constructFromObject(data['original']);
            }
            if (data.hasOwnProperty('pipeline')) {
                obj['pipeline'] = V1Pipeline.constructFromObject(data['pipeline']);
            }
            if (data.hasOwnProperty('status_conditions')) {
                obj['status_conditions'] = ApiClient.convertToType(data['status_conditions'], [V1StatusCondition]);
            }
            if (data.hasOwnProperty('settings')) {
                obj['settings'] = V1RunSettings.constructFromObject(data['settings']);
            }
            if (data.hasOwnProperty('role')) {
                obj['role'] = ApiClient.convertToType(data['role'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} uuid
 */
V1Run.prototype['uuid'] = undefined;

/**
 * @member {String} name
 */
V1Run.prototype['name'] = undefined;

/**
 * @member {String} description
 */
V1Run.prototype['description'] = undefined;

/**
 * @member {Array.<String>} tags
 */
V1Run.prototype['tags'] = undefined;

/**
 * @member {Boolean} deleted
 */
V1Run.prototype['deleted'] = undefined;

/**
 * @member {String} user
 */
V1Run.prototype['user'] = undefined;

/**
 * @member {String} owner
 */
V1Run.prototype['owner'] = undefined;

/**
 * @member {String} project
 */
V1Run.prototype['project'] = undefined;

/**
 * @member {Date} created_at
 */
V1Run.prototype['created_at'] = undefined;

/**
 * @member {Date} updated_at
 */
V1Run.prototype['updated_at'] = undefined;

/**
 * @member {Date} started_at
 */
V1Run.prototype['started_at'] = undefined;

/**
 * @member {Date} finished_at
 */
V1Run.prototype['finished_at'] = undefined;

/**
 * @member {Number} run_time
 */
V1Run.prototype['run_time'] = undefined;

/**
 * @member {String} is_managed
 */
V1Run.prototype['is_managed'] = undefined;

/**
 * @member {String} content
 */
V1Run.prototype['content'] = undefined;

/**
 * @member {module:model/V1Statuses} status
 */
V1Run.prototype['status'] = undefined;

/**
 * @member {Boolean} bookmarked
 */
V1Run.prototype['bookmarked'] = undefined;

/**
 * @member {Object} meta_info
 */
V1Run.prototype['meta_info'] = undefined;

/**
 * @member {Boolean} is_helper
 */
V1Run.prototype['is_helper'] = undefined;

/**
 * @member {module:model/V1RunKind} kind
 */
V1Run.prototype['kind'] = undefined;

/**
 * @member {module:model/V1RunKind} meta_kind
 */
V1Run.prototype['meta_kind'] = undefined;

/**
 * @member {String} hub
 */
V1Run.prototype['hub'] = undefined;

/**
 * @member {Object} inputs
 */
V1Run.prototype['inputs'] = undefined;

/**
 * @member {Object} outputs
 */
V1Run.prototype['outputs'] = undefined;

/**
 * @member {module:model/V1Cloning} original
 */
V1Run.prototype['original'] = undefined;

/**
 * @member {module:model/V1Pipeline} pipeline
 */
V1Run.prototype['pipeline'] = undefined;

/**
 * @member {Array.<module:model/V1StatusCondition>} status_conditions
 */
V1Run.prototype['status_conditions'] = undefined;

/**
 * @member {module:model/V1RunSettings} settings
 */
V1Run.prototype['settings'] = undefined;

/**
 * @member {String} role
 */
V1Run.prototype['role'] = undefined;






export default V1Run;

