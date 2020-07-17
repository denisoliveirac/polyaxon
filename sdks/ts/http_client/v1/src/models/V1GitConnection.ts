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
 * @interface V1GitConnection
 */
export interface V1GitConnection {
    /**
     * 
     * @type {string}
     * @memberof V1GitConnection
     */
    url?: string;
    /**
     * 
     * @type {boolean}
     * @memberof V1GitConnection
     */
    revision?: boolean;
}

export function V1GitConnectionFromJSON(json: any): V1GitConnection {
    return V1GitConnectionFromJSONTyped(json, false);
}

export function V1GitConnectionFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1GitConnection {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'url': !exists(json, 'url') ? undefined : json['url'],
        'revision': !exists(json, 'revision') ? undefined : json['revision'],
    };
}

export function V1GitConnectionToJSON(value?: V1GitConnection | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'url': value.url,
        'revision': value.revision,
    };
}


