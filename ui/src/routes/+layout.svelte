<svelte:options immutable />

<script lang="ts" context="module">
	import '../app.postcss';
	import '$lib/stores/theme';
</script>

<script lang="ts">
	import Notifications from '$lib/components/Notifications.svelte';
	import { onMount } from 'svelte';
	import Header from '$lib/components/Header.svelte';

	onMount(() => {
		// for testing to know that svelte has mounted the app
		document.body.classList.add('hydrated');
	});
</script>

<div class="flex h-full w-full flex-col">
	<Header />
	<div class="flex flex-grow flex-col overflow-y-auto overflow-x-hidden">
		<slot />
	</div>
</div>

<Notifications />
