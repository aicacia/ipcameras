import type { ICEServer } from '$lib/openapi/ipcameras';
import { localstorageWritable } from 'svelte-localstorage-writable';
import { derived, get } from 'svelte/store';

const iceServersWritable = localstorageWritable<ICEServer[]>('ice-servers', []);

export const iceServers = derived(iceServersWritable, (state) => state);

export function getICEServers(): RTCIceServer[] {
	return get(iceServers)
		.filter((server) => server.urls)
		.map(
			(server) =>
				({
					urls: server.urls,
					credential: server.credential,
					username: server.username
				}) as RTCIceServer
		);
}

export function setICEServers(iceServers: ICEServer[]) {
	iceServersWritable.set(iceServers);
}
