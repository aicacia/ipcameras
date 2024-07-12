<svelte:options immutable />

<script lang="ts">
	import Header from '$lib/components/Header.svelte';
	import { page } from '$app/stores';
</script>

<div class="flex h-full w-full flex-col">
	<Header />
	<div class="mx-4 mb-16 mt-4 flex w-full flex-grow flex-col sm:container sm:mx-auto">
		<h3>{$page.status}</h3>
		{#if $page.error?.message}
			<p>{$page.error.message}</p>
		{/if}
	</div>
</div>
