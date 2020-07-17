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
    V1Environment,
    V1EnvironmentFromJSON,
    V1EnvironmentFromJSONTyped,
    V1EnvironmentToJSON,
    V1Init,
    V1InitFromJSON,
    V1InitFromJSONTyped,
    V1InitToJSON,
} from './';

/**
 * 
 * @export
 * @interface V1SparkReplica
 */
export interface V1SparkReplica {
    /**
     * 
     * @type {number}
     * @memberof V1SparkReplica
     */
    replicas?: number;
    /**
     * 
     * @type {V1Environment}
     * @memberof V1SparkReplica
     */
    environment?: V1Environment;
    /**
     * 
     * @type {Array<V1Init>}
     * @memberof V1SparkReplica
     */
    init?: Array<V1Init>;
    /**
     * 
     * @type {Array<object>}
     * @memberof V1SparkReplica
     */
    sidecars?: Array<object>;
    /**
     * 
     * @type {object}
     * @memberof V1SparkReplica
     */
    container?: object;
}

export function V1SparkReplicaFromJSON(json: any): V1SparkReplica {
    return V1SparkReplicaFromJSONTyped(json, false);
}

export function V1SparkReplicaFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1SparkReplica {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'replicas': !exists(json, 'replicas') ? undefined : json['replicas'],
        'environment': !exists(json, 'environment') ? undefined : V1EnvironmentFromJSON(json['environment']),
        'init': !exists(json, 'init') ? undefined : ((json['init'] as Array<any>).map(V1InitFromJSON)),
        'sidecars': !exists(json, 'sidecars') ? undefined : json['sidecars'],
        'container': !exists(json, 'container') ? undefined : json['container'],
    };
}

export function V1SparkReplicaToJSON(value?: V1SparkReplica | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'replicas': value.replicas,
        'environment': V1EnvironmentToJSON(value.environment),
        'init': value.init === undefined ? undefined : ((value.init as Array<any>).map(V1InitToJSON)),
        'sidecars': value.sidecars,
        'container': value.container,
    };
}


