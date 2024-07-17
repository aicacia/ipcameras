<svelte:options immutable />

<script lang="ts">
	import { createPeer } from '$lib/peer';
	import type { Peer } from '@aicacia/peer';
	import { onMount } from 'svelte';

	export let rtspUrl: string;
	export let peer: Peer | undefined = undefined;
	export let video: HTMLVideoElement = undefined as never as HTMLVideoElement;

	let mounted = false;

	async function init(peer: Peer) {
		peer.ready().then(() => {
			peer.send(
				JSON.stringify({
					type: 'rtsp',
					rtspUrl
				})
			);
		});
		let mediaStream = new MediaStream();
		video.srcObject = mediaStream;
		peer.on('track', (event) => {
			mediaStream.addTrack(event.track);
		});
		peer.on('close', () => {
			if (mounted) {
				peer = createPeer('webrtcrtsp');
				init(peer);
			}
		});
	}

	onMount(() => {
		mounted = true;
		peer = createPeer('webrtcrtsp');
		init(peer);

		return () => {
			mounted = false;
			peer?.close();
		};
	});
</script>

<video class="w-full" bind:this={video} controls autoplay>
	<track kind="captions" />
</video>
