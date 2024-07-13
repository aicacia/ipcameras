import { getP2PAccess } from '$lib/stores/p2pAccess';
import { redirect } from '@sveltejs/kit';
import type { LayoutLoad } from './$types';
import { getHTTPAccess } from '$lib/stores/httpAccess';
import { base } from '$app/paths';

export const prerender = true;
export const ssr = false;

export const load: LayoutLoad = async (event) => {
	await event.parent();

	if (!getP2PAccess() && !getHTTPAccess()) {
		redirect(302, `${base}/connect`);
	}
};
