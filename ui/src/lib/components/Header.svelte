<svelte:options immutable />

<script lang="ts">
	import { base } from '$app/paths';
	import Dropdown from '$lib/components/Dropdown.svelte';
	import Menu from 'lucide-svelte/icons/menu';
	import AppWindow from 'lucide-svelte/icons/app-window';
	import LogOut from 'lucide-svelte/icons/log-out';
	import User from 'lucide-svelte/icons/user';
	import { page } from '$app/stores';
	import LL from '$lib/i18n/i18n-svelte';

	let open = false;
	function onGoto() {
		open = false;
	}
</script>

<div class="flex flex-shrink flex-row justify-between bg-white shadow dark:bg-gray-800">
	<div class="ms-2 flex flex-shrink flex-row">
		<a class="btn text-lg" href={`${base}/`}>{$LL.header.title()}</a>
	</div>
	<div class="me-2 flex flex-shrink flex-row">
		<div class="flex flex-col content-center justify-center">
			{#if true}
				<Dropdown bind:open>
					<Menu slot="button" />
					<a
						href={`${base}/dashboard`}
						class="default flex cursor-pointer flex-row justify-between p-2 hover:bg-gray-200 dark:hover:bg-gray-600"
						class:active={$page.route.id === '/(authed)/dashboard'}
						on:click={onGoto}
					>
						<AppWindow /><span class="ms-4">{$LL.dashboard.title()}</span>
					</a>
					<a
						href={`${base}/cameras`}
						class="default flex cursor-pointer flex-row justify-between p-2 hover:bg-gray-200 dark:hover:bg-gray-600"
						class:active={$page.route.id === '/(authed)/cameras'}
						on:click={onGoto}
					>
						<User /><span class="ms-4">{$LL.cameras.title()}</span>
					</a>
					<a
						href={`${base}/signin`}
						class="default flex cursor-pointer flex-row justify-between p-2 hover:bg-gray-200 dark:hover:bg-gray-600"
						class:active={$page.route.id === '/(unauthed)/signin'}
						on:click={onGoto}
					>
						<LogOut /><span class="ms-4">{$LL.auth.signIn()}</span>
					</a>
				</Dropdown>
			{/if}
		</div>
	</div>
</div>

<style lang="postcss">
	li.active,
	a.active {
		@apply bg-gray-200 dark:bg-gray-600;
	}
</style>
