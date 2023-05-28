<script lang="ts">
	import { HoshiiAPI } from '$lib/api/client';
	import type { PageServerData } from './$types';

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
		<h1 class="text-6xl sm:text-8xl">{data.props.userEndpoint}</h1>
	</section>
	<section>
		Wishlists
		{#if data.props.sameUser}
			<button
				on:click={async () => {
					console.log('CLICKED');
					const res = await hoshiiClient.GetWishlist();
					console.log(res);
				}}
				class=" rounded bg-blue-300 px-4 py-2 font-bold text-white hover:bg-blue-400"
				>GET WISHLIST</button
			>
		{/if}
	</section>
</div>
