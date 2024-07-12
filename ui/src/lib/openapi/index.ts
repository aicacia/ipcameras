import { getWebRTCFetch } from '$lib/peer';
import { getAccess } from '$lib/stores/access';
import { getHTTPAccess } from '$lib/stores/httpAccess';
import { getP2PAccess } from '$lib/stores/p2pAccess';
import { getHttpOrigin } from '$lib/util';
import { AppApi, CameraApi, Configuration, type ConfigurationParameters } from './ipcameras';

export const defaultConfiguration: ConfigurationParameters = {
	middleware: [
		{
			pre: async (context) => ({ ...context, init: { ...context.init, mode: 'cors' } })
		}
	],
	apiKey(name: string) {
		switch (name) {
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

export const appApi = new AppApi(ipcamerasConfiguration);
export const camerasApi = new CameraApi(ipcamerasConfiguration);
