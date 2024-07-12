<svelte:options immutable />

<script lang="ts">
	import type { SuiteResult } from 'vest';

	export let result: SuiteResult;
	export let name: string;
</script>

<ul>
	{#each result.getErrors(name) as error}
		<li class="text-sm text-red-600">{error}</li>
	{/each}
	{#each result.getWarnings(name) as warning}
		<li class="text-sm text-orange-600">{warning}</li>
	{/each}
</ul>
