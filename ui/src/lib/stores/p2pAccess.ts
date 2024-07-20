import type { P2PAccess } from '$lib/openapi/ipcameras';
import { localstorageWritable } from 'svelte-localstorage-writable';
import { derived, get } from 'svelte/store';

const p2pAccessWritable = localstorageWritable<P2PAccess | null>('p2p-access', null);

export const p2pAccess = derived(p2pAccessWritable, (state) => state);

export function getP2PAccess() {
	return get(p2pAccess);
}

export function setP2PAccess(p2pAccess: P2PAccess | null) {
	p2pAccessWritable.set(p2pAccess);
}
