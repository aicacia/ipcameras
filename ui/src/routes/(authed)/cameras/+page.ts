import { camerasApi } from '$lib/openapi';
import type { PageLoad } from './$types';

export const load: PageLoad = async (_event) => {
	return {
		cameras: await camerasApi.cameras()
	};
};
