<script lang="ts">
	import { onMount } from 'svelte';
	import type { PageData } from './$types';

	export let data: PageData;

	let user: any;
	onMount(async () => {
		const userFetch = await fetch('/api/users/me');
		user = await userFetch.json();
	});
</script>

<div>
	<h1>HI FROM DASHBOARD</h1>
	{#if user}
		<h2>{JSON.stringify(user)}</h2>
	{:else}
		<h2>Loading...</h2>
	{/if}
</div>
