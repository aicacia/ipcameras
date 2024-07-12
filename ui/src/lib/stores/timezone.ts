import { browser } from '$app/environment';
import { derived, writable, get } from 'svelte/store';

const defaultTimezone = 'America/New_York';
const timezoneWritable = writable(browser ? getBrowserTimezoneRequired() : defaultTimezone);

export const timezone = derived(timezoneWritable, (state) => state);

export function setTimezone(timezone: string) {
	timezoneWritable.set(timezone);
}

export function getTimezone() {
	return get(timezoneWritable);
}

export function useBrowserTimezone() {
	timezoneWritable.set(navigator.language);
}

export function getBrowserTimezone() {
	try {
		return Intl.DateTimeFormat().resolvedOptions().timeZone;
	} catch {
		return undefined;
	}
}

export function getBrowserTimezoneRequired() {
	return getBrowserTimezone() || defaultTimezone;
}
