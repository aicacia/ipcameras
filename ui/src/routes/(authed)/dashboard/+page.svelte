<svelte:options immutable />

<script lang="ts">
	import LL from '$lib/i18n/i18n-svelte';
</script>

<svelte:head>
	<title>{$LL.dashboard.title()}</title>
</svelte:head>

<div class="flex flex-col justify-end px-4 md:justify-start">
	<div
		class="mx-auto mt-4 flex w-full max-w-6xl flex-shrink flex-col bg-white p-4 shadow dark:bg-gray-800"
	></div>
</div>
