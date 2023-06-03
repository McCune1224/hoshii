<script lang="ts">
	import { HoshiiAPI } from '$lib/api/client';
	import type { PageServerData } from './$types';
	import { modalStore, type ModalSettings, ProgressRadial } from '@skeletonlabs/skeleton';
	import { onMount } from 'svelte';

	export let data: PageServerData;
	let hoshiiClient = new HoshiiAPI(data.user.token);
	let userWishlistResponse: MeWishlistsResponse;
	const profileModal: ModalSettings = {
		title: 'Update Profile',
		body: '',
		type: 'component',
		component: 'modalUpdateProfile'
	};

	onMount(async () => {
		try {
			userWishlistResponse = await hoshiiClient.GetMeWishlists();
			console.log(userWishlistResponse);
		} catch (e) {
			console.log(e);
		}
	});
</script>

<div class="flex flex-col">
	<section class="">
		<h1 class="text-6xl sm:text-8xl">Greetings {data.user.username}</h1>
	</section>
	<section>
		Wishlists
		<button
			on:click={async () => {
				modalStore.trigger(profileModal);
			}}
			class="btn variant-filled-secondary"
		>
			Update Profile
		</button>
		{#if userWishlistResponse}
			<div class="border-2 border-primary-200">
				{#each userWishlistResponse.wishlists as wishlist}
					<div class="flex flex-col">
						<h1 class="text-4xl sm:text-6xl">{wishlist.id}</h1>
						<h3 class="text-4xl sm:text-6xl">{wishlist.name}</h3>
						<button
							class="btn variant-soft"
							on:click={() => {
								console.log('Item add/update invoke for wishlist: ' + wishlist.name);
							}}
						>
							Update Wishlist
						</button>
					</div>
				{/each}
			</div>
		{:else}
			<ProgressRadial />
		{/if}
	</section>
</div>
