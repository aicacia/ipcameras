<svelte:options immutable />

<script lang="ts">
	import LL from '$lib/i18n/i18n-svelte';
	import P2PComponent from '$lib/components/Connect/P2PAccess.svelte';
	import HTTPAccessComponent from '$lib/components/Connect/HTTPAccess.svelte';
	import { goto } from '$app/navigation';
	import { base } from '$app/paths';
	import { access } from '$lib/stores/access';

	async function onConnect() {
		await goto(`${base}/`);
	}
</script>

<svelte:head>
	<title>{$LL.connect.connect()}</title>
</svelte:head>

<div class="flex flex-grow flex-col justify-end md:justify-start">
	<div class="mx-auto flex w-full flex-shrink flex-col p-4 py-10 md:w-72">
		<div class="flex flex-grow flex-col bg-white p-4 shadow dark:bg-gray-800">
			<h1 class="mb-1">{$LL.connect.connect()}</h1>
			<select name="access" bind:value={$access}>
				<option value="http">{$LL.connect.httpAccess()}</option>
				<option value="p2p">{$LL.connect.p2pAccess()}</option>
			</select>
			{#if $access === 'p2p'}
				<P2PComponent {onConnect} />
			{:else if $access === 'http'}
				<HTTPAccessComponent {onConnect} />
			{/if}
		</div>
	</div>
</div>
