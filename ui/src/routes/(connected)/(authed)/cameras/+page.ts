import { camerasApi } from '$lib/openapi';
import type { PageLoad } from './$types';

export const load: PageLoad = async (event) => {
	await event.parent();
	return {
		cameras: await camerasApi.cameras()
	};
};
