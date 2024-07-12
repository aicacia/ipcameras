<svelte:options immutable />

<script lang="ts" context="module">
	import '../app.postcss';
	import '$lib/stores/theme';
</script>

<script lang="ts">
	import Notifications from '$lib/components/Notifications.svelte';
	import { onMount } from 'svelte';

	onMount(() => {
		// for testing to know that svelte has mounted the app
		document.body.classList.add('hydrated');
	});
</script>

<slot />

<Notifications />
