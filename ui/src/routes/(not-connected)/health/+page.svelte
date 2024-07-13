<svelte:options immutable />

<script lang="ts">
	import { base } from '$app/paths';
	import LL from '$lib/i18n/i18n-svelte';
</script>

<svelte:head>
	<title>{$LL.health.title()}</title>
</svelte:head>

<div class="flex flex-grow flex-col justify-end md:justify-start">
	<div
		class="mx-auto my-10 flex w-full flex-shrink flex-col bg-white p-4 shadow md:w-80 dark:bg-gray-800"
	>
		<h1 class="text-center">{$LL.health.header()}</h1>
		<p class="py-2 text-center">
			<a href={`${base}/signin`}>{$LL.auth.signIn()}</a>
		</p>
	</div>
</div>
