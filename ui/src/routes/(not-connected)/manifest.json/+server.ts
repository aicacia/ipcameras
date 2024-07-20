import { base } from '$app/paths';
import { json } from '@sveltejs/kit';

export const prerender = true;

export async function GET() {
	return json({
    "name": "IPCameras",
    "short_name": "IPCameras",
    "description": "IPCameras: The IP Camera App",
    "version": "1.0",
    "manifest_version": 3,
    "icons": [
      {
        "src": `${base}/icon256x256.png`,
        "sizes": "256x256",
        "type": "image/png"
      }
    ],
    "id": `${base}/?source=pwa`,
    "start_url": `${base}/?source=pwa`,
    "scope": `${base}`,
    "display": "standalone",
    "background_color": "#111827",
    "theme_color": "#3A82F7"
  });
}
