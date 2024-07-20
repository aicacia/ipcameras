import type { LayoutLoad } from './$types';
import { base } from '$app/paths';
import { tryGetCurrentUser } from '$lib/stores/user';
import { redirect } from '@sveltejs/kit';

export const prerender = true;
export const ssr = false;

export const load: LayoutLoad = async (event) => {
	await event.parent();

	const currentUser = await tryGetCurrentUser();

	if (currentUser) {
		return {
			user: currentUser
		};
	} else {
		redirect(302, `${base}/signin`);
	}
};
