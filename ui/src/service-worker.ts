/// <reference types="@sveltejs/kit" />
/// <reference no-default-lib="true"/>
/// <reference lib="esnext" />
/// <reference lib="webworker" />

import { build, files, version } from '$service-worker';

const worker = self as unknown as ServiceWorkerGlobalScope;
const CACHE = `cache-${version}`;

const ASSETS = [...build, ...files];

worker.addEventListener('install', (event) => {
	async function addFilesToCache() {
		const cache = await caches.open(CACHE);
		await cache.addAll(ASSETS);
	}

	event.waitUntil(addFilesToCache());
});

worker.addEventListener('activate', (event) => {
	async function deleteOldCaches() {
		for (const key of await caches.keys()) {
			if (key !== CACHE) await caches.delete(key);
		}
	}

	event.waitUntil(deleteOldCaches());
});

worker.addEventListener('fetch', (event) => {
	if (event.request.method === 'GET') {
		const response = (async () => {
			const url = new URL(event.request.url);
			const cache = await caches.open(CACHE);

			if (ASSETS.includes(url.pathname)) {
				return cache.match(url.pathname);
			}

			try {
				const response = await fetch(event.request);

				if (response.status === 200 && !event.request.url.startsWith('chrome-extension://')) {
					cache.put(event.request, response.clone());
				}

				return response;
			} catch {
				return cache.match(event.request);
			}
		})();

		event.respondWith(response as Promise<Response>);
	}
});
