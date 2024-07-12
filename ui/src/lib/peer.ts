import { getHttpOrigin, getWSOrigin, parseJWTClaims } from './util';
import { getP2PAccess } from './stores/p2pAccess';
import { type WebRTCFetch, createWebRTCFetch } from '@aicacia/webrtchttp';
import { Peer } from '@aicacia/peer';
import { KeepAliveWebSocket } from '@aicacia/keepalivewebsocket';
import type { P2PAccess } from './openapi/ipcameras';

const peers = new Map<string, Peer>();
const websocket = new KeepAliveWebSocket({
	url: getClientWSUrl
});
websocket.on('message', (raw) => {
	const data = JSON.parse(raw as string);
	const peer = peers.get(data.peerId);
	if (!peer) {
		return;
	}
	peer.signal(data.message);
});

type Fetch = typeof fetch;

let webrtcFetch: Fetch;
export function getWebRTCFetch() {
	if (webrtcFetch) {
		return webrtcFetch;
	} else {
		let webrtcFetchPromise: Promise<WebRTCFetch> | undefined;
		webrtcFetch = (input, init) => {
			if (!webrtcFetchPromise) {
				const peer = createPeer('webrtchttp');
				webrtcFetchPromise = peer.ready().then(() => {
					console.debug(`${peer.getId()}: ready`);
					return createWebRTCFetch(peer.getChannel()!);
				});
				peer.on('close', () => {
					webrtcFetchPromise = undefined;
				});
			}
			return webrtcFetchPromise.then((webrtcFetch) => webrtcFetch(input, init));
		};
		return webrtcFetch;
	}
}

export function createPeer(type: string) {
	websocket.connect();
	const peer = new Peer({
		config: {
			iceServers: [
				{
					urls: ['stun:stun.l.google.com:19302']
				}
			]
		},
		channelConfig: {
			ordered: true
		}
	});
	peer.on('error', (err) => console.error(err));
	peer.on('signal', (message) => {
		websocket.send(
			JSON.stringify({
				type: 'signal',
				peerId: peer.getId(),
				message: message
			})
		);
	});

	websocket.ready().then(() => {
		websocket.send(
			JSON.stringify({
				type: 'init',
				peerId: peer.getId(),
				peerType: type
			})
		);
	});

	peer.on('close', () => {
		peers.delete(peer.getId());
	});
	peers.set(peer.getId(), peer);

	return peer;
}

async function getClientWSUrl() {
	const p2pAccess = getP2PAccess();
	if (!p2pAccess) {
		throw new Error('No remote access configured');
	}
	const token = await getClientWSToken(p2pAccess);
	return `${getWSOrigin(p2pAccess.host, p2pAccess.ssl)}/client/websocket?token=${token}`;
}

let token: string | undefined;
let tokenExpiration: number | undefined;

async function getClientWSToken(p2pAccess: P2PAccess) {
	if (!tokenExpiration || tokenExpiration <= Date.now()) {
		token = await createAccessJWT(p2pAccess);
		tokenExpiration = parseJWTClaims(token).exp * 1000;
	}
	return token as string;
}

export async function createAccessJWT(p2pAccess: P2PAccess) {
	const body: { id?: string; password: string } = {
		id: p2pAccess.id,
		password: p2pAccess.password
	};
	const headers: HeadersInit = {
		'Content-Type': 'application/json'
	};
	const res = await fetch(`${getHttpOrigin(p2pAccess.host, p2pAccess.ssl)}/client`, {
		method: 'POST',
		headers,
		credentials: 'same-origin',
		mode: 'cors',
		body: JSON.stringify(body)
	});
	if (res.status >= 400) {
		throw new Error('failed to authenticate');
	}
	return await res.text();
}
