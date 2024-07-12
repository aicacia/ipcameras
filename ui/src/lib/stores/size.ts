import { browser } from '$app/environment';
import { derived, writable } from 'svelte/store';

const sizeWritable = writable<[width: number, height: number]>([
	browser ? window.innerWidth : 0,
	browser ? window.innerHeight : 0
]);
export const size = derived(sizeWritable, (size) => size);

function onResize() {
	sizeWritable.set([window.innerWidth, window.innerHeight]);
}

if (browser) {
	window.addEventListener('resize', onResize);
}
