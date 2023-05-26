<script lang="ts">
	import { onMount } from 'svelte';
	import type { PageData } from './$types';

	export let data: PageData;

	onMount(() => {});
</script>

<div>
	<h1>HI FROM DASHBOARD</h1>
    <p>Current User is: {data.currentUser.username}</p>
    <p>RAW DATA: {JSON.stringify(data)}</p>
</div>
