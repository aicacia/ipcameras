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
  Errors,
  ResetPassword,
  User,
} from '../models/index';
import {
    ErrorsFromJSON,
    ErrorsToJSON,
    ResetPasswordFromJSON,
    ResetPasswordToJSON,
    UserFromJSON,
    UserToJSON,
} from '../models/index';

export interface ResetPasswordRequest {
    resetPassword: ResetPassword;
}

/**
 * CurrentUserApi - interface
 * 
 * @export
 * @interface CurrentUserApiInterface
 */
export interface CurrentUserApiInterface {
    /**
     * 
     * @summary Get current user
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CurrentUserApiInterface
     */
    currentUserRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<User>>;

    /**
     * Get current user
     */
    currentUser(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<User>;

    /**
     * 
     * @summary Resets a user\'s password
     * @param {ResetPassword} resetPassword reset user\&#39;s password
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CurrentUserApiInterface
     */
    resetPasswordRaw(requestParameters: ResetPasswordRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>>;

    /**
     * Resets a user\'s password
     */
    resetPassword(resetPassword: ResetPassword, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void>;

}

/**
 * 
 */
export class CurrentUserApi extends runtime.BaseAPI implements CurrentUserApiInterface {

    /**
     * Get current user
     */
    async currentUserRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<User>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // Authorization authentication
        }

        const response = await this.request({
            path: `/user`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => UserFromJSON(jsonValue));
    }

    /**
     * Get current user
     */
    async currentUser(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<User> {
        const response = await this.currentUserRaw(initOverrides);
        return await response.value();
    }

    /**
     * Resets a user\'s password
     */
    async resetPasswordRaw(requestParameters: ResetPasswordRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['resetPassword'] == null) {
            throw new runtime.RequiredError(
                'resetPassword',
                'Required parameter "resetPassword" was null or undefined when calling resetPassword().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // Authorization authentication
        }

        const response = await this.request({
            path: `/user/reset-password`,
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: ResetPasswordToJSON(requestParameters['resetPassword']),
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Resets a user\'s password
     */
    async resetPassword(resetPassword: ResetPassword, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.resetPasswordRaw({ resetPassword: resetPassword }, initOverrides);
    }

}
