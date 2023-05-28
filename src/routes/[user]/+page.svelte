<script lang="ts">
	import { HoshiiAPI } from '$lib/api/client';
	import type { PageServerData } from './$types';
	import { page } from '$app/stores';

	export let data: PageServerData;
</script>

<div class="flex flex-col">
	<section class="">
		<h1 class="text-6xl sm:text-8xl">Greetings {data.user.username}</h1>
	</section>
	<section>Wishlists</section>
</div>
