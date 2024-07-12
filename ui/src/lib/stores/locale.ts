import { browser } from '$app/environment';
import { derived, writable, get } from 'svelte/store';

const localeWritable = writable(browser ? navigator.language : 'en-US');

export const locale = derived(localeWritable, (state) => state);

export function setLocale(locale: string) {
	localeWritable.set(locale);
}

export function getLocale() {
	return get(localeWritable);
}

export function useBrowserLocale() {
	localeWritable.set(navigator.language);
}

if (browser) {
	window.addEventListener('languagechange', () => {
		setLocale(navigator.language);
	});
}
