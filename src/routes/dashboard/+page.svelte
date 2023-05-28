<script lang="ts">
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	import { HoshiiAPI } from '$lib/api/client';

	const client = new HoshiiAPI();

	onMount(async () => {
	});
</script>

<div>
	<h1>HI FROM DASHBOARD</h1>
</div>
