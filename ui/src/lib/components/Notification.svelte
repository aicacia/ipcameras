<svelte:options immutable />

<script lang="ts">
	import { removeNotification, type INotification } from '$lib/stores/notifications';
	import CircleX from 'lucide-svelte/icons/circle-x';
	import CircleCheck from 'lucide-svelte/icons/circle-check';
	import Info from 'lucide-svelte/icons/info';
	import TriangleAlert from 'lucide-svelte/icons/triangle-alert';

	export let notification: INotification;

	function onDelete() {
		removeNotification(notification.id);
	}
</script>

<button
	class="m-1 flex flex-grow cursor-pointer flex-row items-center px-3 py-2 shadow"
	class:bg-green-600={notification.type === 'success'}
	class:bg-red-600={notification.type === 'error'}
	class:bg-blue-600={notification.type === 'info'}
	class:bg-yellow-600={notification.type === 'warning'}
	on:click={onDelete}
>
	<div class="mr-2 h-6 w-6 text-white">
		{#if notification.type === 'error'}
			<CircleX />
		{:else if notification.type === 'success'}
			<CircleCheck />
		{:else if notification.type === 'info'}
			<Info />
		{:else if notification.type === 'warning'}
			<TriangleAlert />
		{/if}
	</div>
	<div class="flex-grow text-left text-white">{notification.message}</div>
</button>
