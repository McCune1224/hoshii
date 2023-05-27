<script lang="ts">
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	import { HoshiiAPI } from '$lib/api/client';

	const client = new HoshiiAPI();

	let user: CurrentUserResponse;
	onMount(async () => {
		user = await client.GetMe();
		console.log('USER', user);
	});
</script>

<div>
	<h1>HI FROM DASHBOARD</h1>
	{#if user}
		<h2>Greetings {user.name}</h2>
        <p>Iunno use this to tweak profile settings or somethin</p>
	{:else}
		<h2>Loading...</h2>
	{/if}
</div>
