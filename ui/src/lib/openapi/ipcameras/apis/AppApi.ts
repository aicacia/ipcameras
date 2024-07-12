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


import * as runtime from '../runtime';
import type {
  Health,
  P2PAccess,
  Version,
} from '../models/index';
import {
    HealthFromJSON,
    HealthToJSON,
    P2PAccessFromJSON,
    P2PAccessToJSON,
    VersionFromJSON,
    VersionToJSON,
} from '../models/index';

/**
 * AppApi - interface
 * 
 * @export
 * @interface AppApiInterface
 */
export interface AppApiInterface {
    /**
     * 
     * @summary Get Health Check
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof AppApiInterface
     */
    healthCheckRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Health>>;

    /**
     * Get Health Check
     */
    healthCheck(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Health>;

    /**
     * 
     * @summary Get p2p access info
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof AppApiInterface
     */
    p2pAccessRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<P2PAccess>>;

    /**
     * Get p2p access info
     */
    p2pAccess(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<P2PAccess>;

    /**
     * 
     * @summary Get Version
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof AppApiInterface
     */
    versionRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Version>>;

    /**
     * Get Version
     */
    version(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Version>;

}

/**
 * 
 */
export class AppApi extends runtime.BaseAPI implements AppApiInterface {

    /**
     * Get Health Check
     */
    async healthCheckRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Health>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/health`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => HealthFromJSON(jsonValue));
    }

    /**
     * Get Health Check
     */
    async healthCheck(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Health> {
        const response = await this.healthCheckRaw(initOverrides);
        return await response.value();
    }

    /**
     * Get p2p access info
     */
    async p2pAccessRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<P2PAccess>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/p2p-access`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => P2PAccessFromJSON(jsonValue));
    }

    /**
     * Get p2p access info
     */
    async p2pAccess(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<P2PAccess> {
        const response = await this.p2pAccessRaw(initOverrides);
        return await response.value();
    }

    /**
     * Get Version
     */
    async versionRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Version>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/version`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => VersionFromJSON(jsonValue));
    }

    /**
     * Get Version
     */
    async version(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Version> {
        const response = await this.versionRaw(initOverrides);
        return await response.value();
    }

}
