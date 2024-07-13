import { getWebRTCFetch } from '$lib/peer';
import { getAccess } from '$lib/stores/access';
import { getHTTPAccess } from '$lib/stores/httpAccess';
import { getLocale } from '$lib/stores/locale';
import { getP2PAccess } from '$lib/stores/p2pAccess';
import { getTimezone } from '$lib/stores/timezone';
import { getHttpOrigin } from '$lib/util';
import {
	AppApi,
	CameraApi,
	Configuration,
	CurrentUserApi,
	TokenApi,
	type ConfigurationParameters,
	type Token
} from './ipcameras';

let authToken: Token | undefined;

export const defaultConfiguration: ConfigurationParameters = {
	middleware: [
		{
			pre: async (context) => ({ ...context, init: { ...context.init, mode: 'cors' } })
		}
	],
	apiKey(name: string) {
		switch (name) {
			case 'X-Timezone':
				return getTimezone();
			case 'X-Locale':
				return getLocale();
			case 'Authorization':
				return `${authToken?.tokenType} ${authToken?.accessToken}`;
			default:
				return '';
		}
	},
	credentials: 'same-origin',
	get fetchApi() {
		const access = getAccess();
		if (access === 'p2p' && !!getP2PAccess()) {
			return getWebRTCFetch();
		} else {
			return fetch;
		}
	},
	get basePath() {
		const access = getAccess();
		if (access === 'http') {
			const httpAccess = getHTTPAccess();
			if (httpAccess) {
				return getHttpOrigin(httpAccess.host, httpAccess.ssl);
			}
		}
		return '';
	}
};

export const ipcamerasConfiguration = new Configuration(defaultConfiguration);

export const tokenApi = new TokenApi(ipcamerasConfiguration);
export const currentUserApi = new CurrentUserApi(ipcamerasConfiguration);
export const appApi = new AppApi(ipcamerasConfiguration);
export const camerasApi = new CameraApi(ipcamerasConfiguration);

export function setAuthToken(newAuthToken?: Token) {
	authToken = newAuthToken;
}
export function getAuthToken() {
	return authToken;
}
