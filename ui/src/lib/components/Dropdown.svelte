<svelte:options immutable />

<script lang="ts">
	import Popup, { type Position } from './Popup.svelte';

	export let name = '';
	export let anchorPosition: Position | undefined = undefined;
	export let position: Position | undefined = undefined;
	export let open = false;

	function toggle() {
		open = !open;
	}

	let anchor: HTMLDivElement;
</script>

<div class="static flex min-w-fit flex-col" bind:this={anchor}>
	<button
		type="button"
		aria-label={name}
		class="btn icon primary min-w-fit flex-grow"
		on:pointerdown|stopPropagation|preventDefault={toggle}
	>
		<slot name="button" />
	</button>
	<Popup {anchor} {anchorPosition} {position} bind:open>
		<slot />
	</Popup>
</div>
