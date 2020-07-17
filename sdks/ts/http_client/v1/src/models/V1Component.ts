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

/* tslint:disable */
/* eslint-disable */
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
 */

import { exists, mapValues } from '../runtime';
import {
    V1Cache,
    V1CacheFromJSON,
    V1CacheFromJSONTyped,
    V1CacheToJSON,
    V1IO,
    V1IOFromJSON,
    V1IOFromJSONTyped,
    V1IOToJSON,
    V1Plugins,
    V1PluginsFromJSON,
    V1PluginsFromJSONTyped,
    V1PluginsToJSON,
    V1Termination,
    V1TerminationFromJSON,
    V1TerminationFromJSONTyped,
    V1TerminationToJSON,
} from './';

/**
 * 
 * @export
 * @interface V1Component
 */
export interface V1Component {
    /**
     * 
     * @type {number}
     * @memberof V1Component
     */
    version?: number;
    /**
     * 
     * @type {string}
     * @memberof V1Component
     */
    kind?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Component
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Component
     */
    description?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof V1Component
     */
    tags?: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof V1Component
     */
    profile?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Component
     */
    queue?: string;
    /**
     * 
     * @type {V1Cache}
     * @memberof V1Component
     */
    cache?: V1Cache;
    /**
     * 
     * @type {V1Termination}
     * @memberof V1Component
     */
    termination?: V1Termination;
    /**
     * 
     * @type {V1Plugins}
     * @memberof V1Component
     */
    plugins?: V1Plugins;
    /**
     * 
     * @type {Array<V1IO>}
     * @memberof V1Component
     */
    inputs?: Array<V1IO>;
    /**
     * 
     * @type {Array<V1IO>}
     * @memberof V1Component
     */
    outputs?: Array<V1IO>;
    /**
     * 
     * @type {object}
     * @memberof V1Component
     */
    run?: object;
}

export function V1ComponentFromJSON(json: any): V1Component {
    return V1ComponentFromJSONTyped(json, false);
}

export function V1ComponentFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1Component {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'version': !exists(json, 'version') ? undefined : json['version'],
        'kind': !exists(json, 'kind') ? undefined : json['kind'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'tags': !exists(json, 'tags') ? undefined : json['tags'],
        'profile': !exists(json, 'profile') ? undefined : json['profile'],
        'queue': !exists(json, 'queue') ? undefined : json['queue'],
        'cache': !exists(json, 'cache') ? undefined : V1CacheFromJSON(json['cache']),
        'termination': !exists(json, 'termination') ? undefined : V1TerminationFromJSON(json['termination']),
        'plugins': !exists(json, 'plugins') ? undefined : V1PluginsFromJSON(json['plugins']),
        'inputs': !exists(json, 'inputs') ? undefined : ((json['inputs'] as Array<any>).map(V1IOFromJSON)),
        'outputs': !exists(json, 'outputs') ? undefined : ((json['outputs'] as Array<any>).map(V1IOFromJSON)),
        'run': !exists(json, 'run') ? undefined : json['run'],
    };
}

export function V1ComponentToJSON(value?: V1Component | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'version': value.version,
        'kind': value.kind,
        'name': value.name,
        'description': value.description,
        'tags': value.tags,
        'profile': value.profile,
        'queue': value.queue,
        'cache': V1CacheToJSON(value.cache),
        'termination': V1TerminationToJSON(value.termination),
        'plugins': V1PluginsToJSON(value.plugins),
        'inputs': value.inputs === undefined ? undefined : ((value.inputs as Array<any>).map(V1IOToJSON)),
        'outputs': value.outputs === undefined ? undefined : ((value.outputs as Array<any>).map(V1IOToJSON)),
        'run': value.run,
    };
}


