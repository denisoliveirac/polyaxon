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
 * @interface V1Environment
 */
export interface V1Environment {
    /**
     * 
     * @type {{ [key: string]: string; }}
     * @memberof V1Environment
     */
    labels?: { [key: string]: string; };
    /**
     * 
     * @type {{ [key: string]: string; }}
     * @memberof V1Environment
     */
    annotations?: { [key: string]: string; };
    /**
     * 
     * @type {{ [key: string]: string; }}
     * @memberof V1Environment
     */
    node_selector?: { [key: string]: string; };
    /**
     * 
     * @type {object}
     * @memberof V1Environment
     */
    affinity?: object;
    /**
     * Optional Tolerations to apply.
     * @type {Array<object>}
     * @memberof V1Environment
     */
    tolerations?: Array<object>;
    /**
     * Optional NodeName is a request to schedule this pod onto a specific node. If it is non-empty,
     * the scheduler simply schedules this pod onto that node, assuming that it fits resource
     * requirements.
     * @type {string}
     * @memberof V1Environment
     */
    node_name?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Environment
     */
    service_account_name?: string;
    /**
     * Optional HostAliases is an optional list of hosts and IPs that will be injected into the pod spec.
     * @type {Array<object>}
     * @memberof V1Environment
     */
    host_aliases?: Array<object>;
    /**
     * 
     * @type {object}
     * @memberof V1Environment
     */
    security_context?: object;
    /**
     * 
     * @type {Array<string>}
     * @memberof V1Environment
     */
    image_pull_secrets?: Array<string>;
    /**
     * Host networking requested for this workflow pod. Default to false.
     * @type {boolean}
     * @memberof V1Environment
     */
    host_network?: boolean;
    /**
     * Set DNS policy for the pod.
     * Defaults to "ClusterFirst".
     * Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'.
     * DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy.
     * To have DNS options set along with hostNetwork, you have to specify DNS policy
     * explicitly to 'ClusterFirstWithHostNet'.
     * @type {string}
     * @memberof V1Environment
     */
    dns_policy?: string;
    /**
     * 
     * @type {object}
     * @memberof V1Environment
     */
    dns_config?: object;
    /**
     * 
     * @type {string}
     * @memberof V1Environment
     */
    scheduler_name?: string;
    /**
     * If specified, indicates the pod's priority. "system-node-critical" and
     * "system-cluster-critical" are two special keywords which indicate the
     * highest priorities with the former being the highest priority. Any other
     * name must be defined by creating a PriorityClass object with that name.
     * If not specified, the pod priority will be default or zero if there is no
     * default.
     * @type {string}
     * @memberof V1Environment
     */
    priority_class_name?: string;
    /**
     * The priority value. Various system components use this field to find the
     * priority of the pod. When Priority Admission Controller is enabled, it
     * prevents users from setting this field. The admission controller populates
     * this field from PriorityClassName.
     * The higher the value, the higher the priority.
     * @type {number}
     * @memberof V1Environment
     */
    priority?: number;
    /**
     * 
     * @type {string}
     * @memberof V1Environment
     */
    restart_policy?: string;
}

export function V1EnvironmentFromJSON(json: any): V1Environment {
    return V1EnvironmentFromJSONTyped(json, false);
}

export function V1EnvironmentFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1Environment {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'labels': !exists(json, 'labels') ? undefined : json['labels'],
        'annotations': !exists(json, 'annotations') ? undefined : json['annotations'],
        'node_selector': !exists(json, 'node_selector') ? undefined : json['node_selector'],
        'affinity': !exists(json, 'affinity') ? undefined : json['affinity'],
        'tolerations': !exists(json, 'tolerations') ? undefined : json['tolerations'],
        'node_name': !exists(json, 'node_name') ? undefined : json['node_name'],
        'service_account_name': !exists(json, 'service_account_name') ? undefined : json['service_account_name'],
        'host_aliases': !exists(json, 'host_aliases') ? undefined : json['host_aliases'],
        'security_context': !exists(json, 'security_context') ? undefined : json['security_context'],
        'image_pull_secrets': !exists(json, 'image_pull_secrets') ? undefined : json['image_pull_secrets'],
        'host_network': !exists(json, 'host_network') ? undefined : json['host_network'],
        'dns_policy': !exists(json, 'dns_policy') ? undefined : json['dns_policy'],
        'dns_config': !exists(json, 'dns_config') ? undefined : json['dns_config'],
        'scheduler_name': !exists(json, 'scheduler_name') ? undefined : json['scheduler_name'],
        'priority_class_name': !exists(json, 'priority_class_name') ? undefined : json['priority_class_name'],
        'priority': !exists(json, 'priority') ? undefined : json['priority'],
        'restart_policy': !exists(json, 'restart_policy') ? undefined : json['restart_policy'],
    };
}

export function V1EnvironmentToJSON(value?: V1Environment | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'labels': value.labels,
        'annotations': value.annotations,
        'node_selector': value.node_selector,
        'affinity': value.affinity,
        'tolerations': value.tolerations,
        'node_name': value.node_name,
        'service_account_name': value.service_account_name,
        'host_aliases': value.host_aliases,
        'security_context': value.security_context,
        'image_pull_secrets': value.image_pull_secrets,
        'host_network': value.host_network,
        'dns_policy': value.dns_policy,
        'dns_config': value.dns_config,
        'scheduler_name': value.scheduler_name,
        'priority_class_name': value.priority_class_name,
        'priority': value.priority,
        'restart_policy': value.restart_policy,
    };
}


