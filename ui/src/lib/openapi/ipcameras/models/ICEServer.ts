/* tslint:disable */
/* eslint-disable */
/**
 * IPCameras API
 * IPCameras API API
 *
 * The version of the OpenAPI document: 0.1.0
 * Contact: nathanfaucett@gmail.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ICEServer
 */
export interface ICEServer {
    /**
     * 
     * @type {object}
     * @memberof ICEServer
     */
    credential?: object;
    /**
     * 
     * @type {string}
     * @memberof ICEServer
     */
    credentialType?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof ICEServer
     */
    urls?: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof ICEServer
     */
    username?: string;
}

/**
 * Check if a given object implements the ICEServer interface.
 */
export function instanceOfICEServer(value: object): boolean {
    return true;
}

export function ICEServerFromJSON(json: any): ICEServer {
    return ICEServerFromJSONTyped(json, false);
}

export function ICEServerFromJSONTyped(json: any, ignoreDiscriminator: boolean): ICEServer {
    if (json == null) {
        return json;
    }
    return {
        
        'credential': json['credential'] == null ? undefined : json['credential'],
        'credentialType': json['credentialType'] == null ? undefined : json['credentialType'],
        'urls': json['urls'] == null ? undefined : json['urls'],
        'username': json['username'] == null ? undefined : json['username'],
    };
}

export function ICEServerToJSON(value?: ICEServer | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'credential': value['credential'],
        'credentialType': value['credentialType'],
        'urls': value['urls'],
        'username': value['username'],
    };
}

