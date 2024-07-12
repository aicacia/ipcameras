<svelte:options immutable />

<script lang="ts">
	import Header from '$lib/components/Header.svelte';
</script>

<div class="flex h-full w-full flex-col">
	<Header />
	<div class="flex flex-grow flex-col overflow-y-auto overflow-x-hidden">
		<slot />
	</div>
</div>
