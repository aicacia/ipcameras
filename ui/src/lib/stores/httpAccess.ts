import { localstorageWritable } from 'svelte-localstorage-writable';
import { derived, get } from 'svelte/store';

export type HTTPAccess = {
	host: string;
	ssl: boolean;
};

const httpAccessWritable = localstorageWritable<HTTPAccess | null>('local-access', null);

export const httpAccess = derived(httpAccessWritable, (state) => state);

export function getHTTPAccess() {
	return get(httpAccess);
}

export function setHTTPAccess(httpAccess: HTTPAccess | null) {
	httpAccessWritable.set(httpAccess);
}
