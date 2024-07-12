import { derived, writable } from 'svelte/store';

const nowWritable = writable(Date.now());
export const now = derived(nowWritable, (now) => now);

let lastMS = Date.now();
let offsetMS = 0;

function onTimeout() {
	const nowMS = Date.now();
	offsetMS = 1000 - (nowMS - lastMS);
	lastMS = nowMS;
	nowWritable.set(Date.now());
	setTimeout(onTimeout, 1000 + offsetMS);
}

setTimeout(onTimeout, 1000);
