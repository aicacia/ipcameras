/* eslint-disable @typescript-eslint/ban-ts-comment */

import { defineConfig, loadEnv } from 'vite';
import { sveltekit } from '@sveltejs/kit/vite';
import { networkInterfaces } from 'node:os';
import { readFileSync } from 'node:fs';

// https://vitejs.dev/config/
export default defineConfig(async ({ mode }) => {
	const isProd = mode === 'production';

	process.env = { ...process.env, ...loadEnv(mode, process.cwd(), '') };

	const packageJSON = JSON.parse(readFileSync(`${__dirname}/package.json`).toString('utf8'));

	const host = await getInternalIpV4();

	const define = {
		__VERSION__: JSON.stringify(packageJSON.version)
	};

	/** @type {import('vite').UserConfig} */
	const config = {
		clearScreen: false,
		server: {
			host: '0.0.0.0',
			port: 5173,
			strictPort: true,
			hmr: {
				protocol: 'ws',
				host,
				port: 5183
			},
			watch: {
				ignored: ['**/src-tauri/**']
			}
		},
		preview: {
			port: 5173,
			strictPort: true
		},
		plugins: [sveltekit()],
		define
	};

	return config;
});

function getInternalIpV4() {
	const nets = networkInterfaces();
	for (const networks of Object.values(nets)) {
		if (networks) {
			for (const net of networks) {
				const familyV4Value = typeof net.family === 'string' ? 'IPv4' : 4;
				if (net.family === familyV4Value && !net.internal) {
					return net.address;
				}
			}
		}
	}
	return '0.0.0.0';
}
