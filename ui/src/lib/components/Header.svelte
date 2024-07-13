<svelte:options immutable />

<script lang="ts">
	import { base } from '$app/paths';
	import Dropdown from '$lib/components/Dropdown.svelte';
	import Menu from 'lucide-svelte/icons/menu';
	import Router from 'lucide-svelte/icons/router';
	import LogOut from 'lucide-svelte/icons/log-out';
	import LogIn from 'lucide-svelte/icons/log-in';
	import User from 'lucide-svelte/icons/user';
	import { page } from '$app/stores';
	import LL from '$lib/i18n/i18n-svelte';
	import { signedIn, signOut } from '$lib/stores/user';
	import { goto } from '$app/navigation';

	let open = false;
	function onGoto() {
		open = false;
	}

	async function onSignOut() {
		await signOut();
		goto(`${base}/signin`);
		onGoto();
	}
</script>

<div class="flex flex-shrink flex-row justify-between bg-white shadow dark:bg-gray-800">
	<div class="ms-2 flex flex-shrink flex-row">
		<a class="btn text-lg" href={`${base}/`}>{$LL.header.title()}</a>
	</div>
	<div class="me-2 flex flex-shrink flex-row">
		<div class="flex flex-col content-center justify-center">
			<Dropdown bind:open>
				<Menu slot="button" />
				{#if $signedIn}
					<a
						href={`${base}/cameras`}
						class="default flex cursor-pointer flex-row justify-between p-2 hover:bg-gray-200 dark:hover:bg-gray-600"
						class:active={$page.route.id === '/(connected)/(authed)/cameras'}
						on:click={onGoto}
					>
						<User /><span class="ms-4">{$LL.cameras.title()}</span>
					</a>
					<a
						class="default flex cursor-pointer flex-row justify-between p-2 hover:bg-gray-200 dark:hover:bg-gray-600"
						on:click={onSignOut}
					>
						<LogOut /><span class="ms-4">{$LL.auth.signOut()}</span>
					</a>
				{:else}
					<a
						href={`${base}/signin`}
						class="default flex cursor-pointer flex-row justify-between p-2 hover:bg-gray-200 dark:hover:bg-gray-600"
						class:active={$page.route.id === '/(connected)/(unauthed)/signin'}
						on:click={onGoto}
					>
						<LogIn /><span class="ms-4">{$LL.auth.signIn()}</span>
					</a>
				{/if}
				<hr class="my-2" />
				<a
					href={`${base}/connect`}
					class="default flex cursor-pointer flex-row justify-between p-2 hover:bg-gray-200 dark:hover:bg-gray-600"
					class:active={$page.route.id === '/(not-connected)/connect'}
					on:click={onGoto}
				>
					<Router /><span class="ms-4">{$LL.connect.connect()}</span>
				</a>
			</Dropdown>
		</div>
	</div>
</div>

<style lang="postcss">
	li.active,
	a.active {
		@apply bg-gray-200 dark:bg-gray-600;
	}
</style>
