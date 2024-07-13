<svelte:options immutable />

<script lang="ts">
	import LL from '$lib/i18n/i18n-svelte';
	import Play from 'lucide-svelte/icons/play';
	import Pencil from 'lucide-svelte/icons/pencil';
	import type { PageData } from './$types';
	import Modal from '$lib/components/Modal.svelte';
	import type { Camera } from '$lib/openapi/ipcameras';
	import Stream from '$lib/components/Stream.svelte';
	import Dropdown from '$lib/components/Dropdown.svelte';
	import CameraEditor, { type CameraEditorForm } from '$lib/components/CameraEditor.svelte';
	import { handleError } from '$lib/errors';
	import { camerasApi } from '$lib/openapi';
	import { invalidateAll } from '$app/navigation';

	export let data: PageData;

	let playOpen = false;
	let playCamera: Camera | undefined;
	let playName: string | undefined;
	function createOnPlay(camera: Camera, name: string) {
		return () => {
			playOpen = true;
			playCamera = camera;
			playName = name;
		};
	}

	let editOpen = false;
	let editCamera: Camera | undefined;
	function createOnEdit(camera: Camera) {
		return () => {
			editOpen = true;
			editCamera = camera;
		};
	}

	async function onCameraUpdate(updates: CameraEditorForm) {
		if (!editCamera) {
			return;
		}
		try {
			await camerasApi.updateCameraByHardwareId(editCamera.hardwareId, updates);
			editCamera = undefined;
			editOpen = false;
			await invalidateAll();
		} catch (error) {
			await handleError(error);
		}
	}
</script>

<svelte:head>
	<title>{$LL.cameras.title()}</title>
</svelte:head>

<div class="flex flex-col justify-end px-4 md:justify-start">
	<div
		class="mx-auto mt-4 flex w-full max-w-6xl flex-shrink flex-col bg-white p-4 shadow dark:bg-gray-800"
	>
		<div class="flex flex-col">
			{#each data.cameras as camera (camera.hardwareId)}
				<div class="flex flex-row justify-between">
					<div class="flex flex-row">
						<p>{camera.name}</p>
					</div>
					<div class="flex flex-row">
						<button class="btn primary icon me-2" on:click={createOnEdit(camera)}>
							<Pencil />
						</button>
						<Dropdown>
							<Play slot="button" />
							{#each Object.keys(camera.mediaUris) as name (name)}
								<li
									class="flex cursor-pointer flex-row justify-between p-2 hover:bg-gray-200 dark:hover:bg-gray-600"
									on:click={createOnPlay(camera, name)}
								>
									{name}
								</li>
							{/each}
						</Dropdown>
					</div>
				</div>
			{/each}
		</div>
	</div>
</div>

<Modal bind:open={editOpen} small>
	{#if editCamera}
		<CameraEditor
			discovered={editCamera.discovered}
			saved={editCamera.saved}
			hardwareId={editCamera.hardwareId}
			name={editCamera.name}
			mediaUris={editCamera.mediaUris}
			record={editCamera.record}
			recordWindow={editCamera.recordWindow}
			onUpdate={onCameraUpdate}
		/>
	{/if}
</Modal>

<Modal bind:open={playOpen}>
	<h4>{playCamera?.name || playCamera?.hardwareId}</h4>
	{#if playCamera && playName}
		<Stream rtspUrl={playCamera.mediaUris[playName]} />
	{/if}
</Modal>
