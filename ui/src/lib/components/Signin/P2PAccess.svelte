<svelte:options immutable />

<script lang="ts" context="module">
	import { create, test, enforce, only } from 'vest';

	type SignInForm = {
		host: string;
		ssl: boolean;
		id: string;
		password: string;
	};

	const createSuite = (LL: TranslationFunctions) =>
		create((data: Partial<SignInForm> = {}, fields: string[]) => {
			if (!fields.length) {
				return;
			}
			only(fields);

			test('host', LL.errors.message.required(), () => {
				enforce(data.host).isNotBlank();
			});
			test('ssl', LL.errors.message.required(), () => {
				enforce(data.ssl).isNotBlank();
			});
			test('id', LL.errors.message.required(), () => {
				enforce(data.id).isNotBlank();
			});
			test('password', LL.errors.message.required(), () => {
				enforce(data.password).isNotBlank();
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
	import type { P2P } from '$lib/openapi/ipcameras';
	import type { MaybePromise } from '@sveltejs/kit';
	import { createAccessJWT } from '$lib/peer';
	import { p2pAccess, setP2PAccess } from '$lib/stores/p2pAccess';

	export let onSignin: (p2pAccess: P2P) => MaybePromise<void>;

	$: host = $p2pAccess?.host || 'p2p.aicacia.com';
	$: ssl = $p2pAccess?.ssl === false ? false : true;
	$: id = $p2pAccess?.id || '';
	$: password = $p2pAccess?.password || '';

	$: suite = createSuite($LL);
	$: result = suite.get();
	$: disabled = loading;
	$: cn = classNames(result, {
		untested: 'untested',
		tested: 'tested',
		invalid: 'invalid',
		valid: 'valid',
		warning: 'warning'
	});

	const fields = new Set<string>();
	const validate = debounce(() => {
		suite({ host, ssl, id, password }, Array.from(fields)).done((r) => {
			result = r;
		});
		fields.clear();
	}, 300);
	function validateAll() {
		fields.add('host');
		fields.add('ssl');
		fields.add('id');
		fields.add('password');
		validate();
		validate.flush();
	}
	function onChange(e: Event & { currentTarget: HTMLInputElement | HTMLSelectElement }) {
		e.currentTarget.value = e.currentTarget.value;
		fields.add(e.currentTarget.name);
		validate();
	}

	let loading = false;
	async function onSubmit(e: SubmitEvent) {
		try {
			loading = true;
			validateAll();
			if (result.isValid()) {
				const p2pAccess = { host, ssl, id, password };
				await createAccessJWT(p2pAccess);
				setP2PAccess(p2pAccess);
				await onSignin(p2pAccess);
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
		<label for="host">{$LL.auth.hostLabel()}</label>
		<input
			class="w-full {cn('host')}"
			type="text"
			name="host"
			placeholder={$LL.auth.hostPlaceholder()}
			bind:value={host}
			on:input={onChange}
		/>
		<InputResults name="host" {result} />
	</div>
	<div class="mb-2">
		<label for="ssl">{$LL.auth.sslLabel()}</label>
		<input class={cn('ssl')} type="checkbox" name="ssl" bind:checked={ssl} on:input={onChange} />
		<InputResults name="ssl" {result} />
	</div>
	<div class="mb-2">
		<label for="host">{$LL.auth.idLabel()}</label>
		<input
			class="w-full {cn('id')}"
			type="text"
			name="id"
			autocomplete="username"
			placeholder={$LL.auth.idPlaceholder()}
			bind:value={id}
			on:input={onChange}
		/>
		<InputResults name="id" {result} />
	</div>
	<div class="mb-2">
		<label for="host">{$LL.auth.passwordLabel()}</label>
		<input
			class="w-full {cn('password')}"
			type="password"
			name="password"
			autocomplete="current-password"
			placeholder={$LL.auth.passwordPlaceholder()}
			bind:value={password}
			on:input={onChange}
		/>
		<InputResults name="password" {result} />
	</div>
	<div class="flex flex-row justify-end">
		<button type="submit" class="btn primary flex flex-shrink" {disabled}>
			{#if loading}<div class="mr-2 flex flex-row justify-center">
					<div class="inline-block h-6 w-6"><Spinner /></div>
				</div>{/if}
			{$LL.auth.signIn()}
		</button>
	</div>
</form>
