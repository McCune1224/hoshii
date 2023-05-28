<script lang="ts">
	import { HoshiiAPI } from '$lib/api/client';
	import type { PageServerData } from './$types';
	import { page } from '$app/stores';

	export let data: PageServerData;
	const hoshiiClient = new HoshiiAPI();
</script>

<div class="flex flex-col">
	<section class="">
		<img
			class="h-32 w-32 rounded-full border-4 border-black sm:h-48 sm:w-48"
			src="https://pbs.twimg.com/profile_images/1589027795962765312/NzQmfiV9_400x400.jpg"
			alt="user pfp"
		/>
		<h1 class="text-6xl sm:text-8xl">Greetings {data.user.username}</h1>
	</section>
	<section>Wishlists</section>
</div>
