<svelte:options immutable />

<script lang="ts">
	import LL from '$lib/i18n/i18n-svelte';
	import P2PComponent from '$lib/components/Signin/P2PAccess.svelte';
	import HTTPAccessComponent from '$lib/components/Signin/HTTPAccess.svelte';
	import { goto } from '$app/navigation';
	import { base } from '$app/paths';
	import { access } from '$lib/stores/access';

	async function onSignIn() {
		await goto(`${base}/`);
	}
</script>

<svelte:head>
	<title>{$LL.auth.signIn()}</title>
</svelte:head>

<div class="flex flex-grow flex-col justify-end md:justify-start">
	<div class="mx-auto flex w-full flex-shrink flex-col p-4 py-10 md:w-72">
		<div class="flex flex-grow flex-col bg-white p-4 shadow dark:bg-gray-800">
			<h1 class="mb-1">{$LL.auth.signIn()}</h1>
			<select name="access" bind:value={$access}>
				<option value="http">{$LL.auth.httpAccess()}</option>
				<option value="p2p">{$LL.auth.p2pAccess()}</option>
			</select>
			{#if $access === 'p2p'}
				<P2PComponent onSignin={onSignIn} />
			{:else if $access === 'http'}
				<HTTPAccessComponent onSignin={onSignIn} />
			{/if}
		</div>
	</div>
</div>
