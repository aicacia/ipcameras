<svelte:options immutable />

<script lang="ts" context="module">
	import { create, test, enforce, only, omitWhen } from 'vest';

	export type CameraEditorForm = {
		name: string;
		mediaUris: { [key: string]: string };
		hardwareId: string;
		record: boolean;
		recordWindow?: number;
	};

	const createSuite = (LL: TranslationFunctions) =>
		create((data: Partial<CameraEditorForm> = {}, fields: string[]) => {
			if (!fields.length) {
				return;
			}
			only(fields);

			test('name', LL.errors.message.required(), () => {
				enforce(data.name).isNotBlank();
			});
			test('hardwareId', LL.errors.message.required(), () => {
				enforce(data.hardwareId).isNotBlank();
			});
			test('record', LL.errors.message.required(), () => {
				enforce(data.record).isNotBlank();
			});
			omitWhen(!data.recordWindow, () => {
				test('recordWindow', LL.errors.message.required(), () => {
					enforce(data.recordWindow).isNotBlank();
				});
			});
		});
</script>

<script lang="ts">
	import LL from '$lib/i18n/i18n-svelte';
	import classNames from 'vest/classnames';
	import Spinner from '$lib/components/Spinner.svelte';
	import { handleError } from '$lib/errors';
	import { debounce } from '@aicacia/debounce';
	import InputResults from '$lib/components/InputResults.svelte';
	import type { TranslationFunctions } from '$lib/i18n/i18n-types';
	import deepEqual from 'deep-equal';

	export let discovered = false;
	export let saved = false;
	export let hardwareId: string;
	export let name: string;
	export let mediaUris: { [key: string]: string };
	export let record = false;
	export let recordWindow: number | undefined;
	export let onUpdate: (data: CameraEditorForm) => void;

	const original = {
		hardwareId,
		name,
		mediaUris: { ...mediaUris },
		record,
		recordWindow
	};
	$: updates = {
		hardwareId,
		name,
		mediaUris,
		record,
		recordWindow
	};

	$: suite = createSuite($LL);
	$: result = suite.get();
	$: hasUpdates = !deepEqual(updates, original);
	$: disabled = loading;
	$: cn = classNames(result, {
		untested: 'untested',
		tested: 'tested',
		invalid: 'invalid',
		valid: 'valid',
		warning: 'warning'
	});

  let recordWindowString: string | undefined;
  $: if (recordWindowString) {
    const value = parseInt(recordWindowString);
    if (value) {
      recordWindow = value;
    }
  }

	const fields = new Set<string>();
	const validate = debounce(() => {
		suite(updates, Array.from(fields)).done((r) => {
			result = r;
		});
		fields.clear();
		hasUpdates = !deepEqual(updates, original);
	}, 300);
	function validateAll() {
		for (const field of Object.keys(updates)) {
			fields.add(field);
		}
		validate();
		validate.flush();
	}
	function onChange(e: Event & { currentTarget: HTMLInputElement | HTMLSelectElement }) {
		fields.add(e.currentTarget.name);
		validate();
	}

	let loading = false;
	async function onSubmit(e: SubmitEvent) {
		try {
			loading = true;
			validateAll();
			if (result.isValid()) {
				await onUpdate(updates);
				suite.reset();
				result = suite.get();
			}
		} catch (error) {
			await handleError(error);
		} finally {
			loading = false;
		}
	}
</script>

<form on:submit|preventDefault={onSubmit}>
	<div class="mb-2">
		<label for="hardwareId">{$LL.camera.hardwareIdLabel()}</label>
		<input
			class="w-full {cn('hardwareId')}"
			type="text"
			name="hardwareId"
			disabled={discovered}
			placeholder={$LL.camera.hardwareIdPlaceholder()}
			bind:value={hardwareId}
			on:input={onChange}
		/>
		<InputResults name="hardwareId" {result} />
	</div>
	<div class="mb-2">
		<label for="name">{$LL.camera.nameLabel()}</label>
		<input
			class="w-full {cn('name')}"
			type="text"
			name="name"
			placeholder={$LL.camera.namePlaceholder()}
			bind:value={name}
			on:input={onChange}
		/>
		<InputResults name="name" {result} />
	</div>
	<div class="mb-2">
		<label for="record">{$LL.camera.recordLabel()}</label>
		<input
			class={cn('record')}
			type="checkbox"
			name="record"
			bind:checked={record}
			on:input={onChange}
		/>
		<InputResults name="record" {result} />
	</div>
	{#if record}
		<div class="mb-2">
			<label for="recordWindow">{$LL.camera.recordWindowLabel()}</label>
			<input
				class="w-full {cn('recordWindow')}"
				type="text"
				name="recordWindow"
				placeholder={$LL.camera.recordWindowPlaceholder()}
				bind:value={recordWindowString}
				on:input={onChange}
			/>
			<InputResults name="recordWindow" {result} />
		</div>
	{/if}
	<div class="flex flex-row justify-end">
		{#if !saved || hasUpdates}
			<button type="submit" class="btn primary flex flex-shrink" {disabled}>
				{#if loading}<div class="mr-2 flex flex-row justify-center">
						<div class="inline-block h-6 w-6"><Spinner /></div>
					</div>{/if}
				{$LL.camera.save()}
			</button>
		{/if}
	</div>
</form>
