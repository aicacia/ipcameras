<svelte:options immutable />

<script lang="ts">
	import LL from '$lib/i18n/i18n-svelte';
	import { goto } from '$app/navigation';
	import { base } from '$app/paths';
	import Signin from '$lib/components/Signin.svelte';

	async function onSignin() {
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
			<Signin {onSignin} />
		</div>
	</div>
</div>
