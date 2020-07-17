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
/**
 * 
 * @export
 * @interface V1K8sResourceSchema
 */
export interface V1K8sResourceSchema {
    /**
     * 
     * @type {string}
     * @memberof V1K8sResourceSchema
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof V1K8sResourceSchema
     */
    mount_path?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof V1K8sResourceSchema
     */
    items?: Array<string>;
}

export function V1K8sResourceSchemaFromJSON(json: any): V1K8sResourceSchema {
    return V1K8sResourceSchemaFromJSONTyped(json, false);
}

export function V1K8sResourceSchemaFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1K8sResourceSchema {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'name': !exists(json, 'name') ? undefined : json['name'],
        'mount_path': !exists(json, 'mount_path') ? undefined : json['mount_path'],
        'items': !exists(json, 'items') ? undefined : json['items'],
    };
}

export function V1K8sResourceSchemaToJSON(value?: V1K8sResourceSchema | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'name': value.name,
        'mount_path': value.mount_path,
        'items': value.items,
    };
}


