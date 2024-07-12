import { localstorageWritable } from 'svelte-localstorage-writable';
import { get } from 'svelte/store';

export type AccessType = 'http' | 'p2p';

export const access = localstorageWritable<AccessType>('access', 'http');

export function getAccess() {
	return get(access);
}
