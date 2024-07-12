import type { LayoutLoad } from './$types';
import { loadLocaleAsync } from '$lib/i18n/i18n-util.async';
import { setLocale } from '$lib/i18n/i18n-svelte';
import { get } from 'svelte/store';
import { locale } from '$lib/stores/locale';
import { detectLocale } from '$lib/i18n/i18n-util';

export const prerender = true;
export const ssr = true;

export const load: LayoutLoad = async (_event) => {
	const currentLocale = detectLocale(() => [get(locale)]);
	await loadLocaleAsync(currentLocale);
	setLocale(currentLocale);
};
