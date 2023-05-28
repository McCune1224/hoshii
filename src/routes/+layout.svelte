<script lang="ts">
	// Your selected Skeleton theme:
	import '@skeletonlabs/skeleton/themes/theme-skeleton.css';

	// This contains the bulk of Skeletons required styles:
	// NOTE: this will be renamed skeleton.css in the v2.x release.
	import '@skeletonlabs/skeleton/styles/skeleton.css';

	// Finally, your application's global stylesheet (sometimes labeled 'app.css')
	import '../app.postcss';

	import { AppBar, Avatar, LightSwitch } from '@skeletonlabs/skeleton';
	import type { LayoutServerData } from './$types';
	import { page } from '$app/stores';
</script>

<AppBar>
	<svelte:fragment slot="lead">
		<a class="flex flex-row gap-4" href="/">
			<p class="text-4xl sm:text-6xl">🌟</p>
			<h4 class="text-4xl sm:text-6xl">Hoshii</h4></a
		></svelte:fragment
	>
	<svelte:fragment slot="trail">
		<button type="button" class="btn variant-filled">Expore</button>

		{#if !$page.data.activeUser}
			<a href="/login">
				<button type="button" class="btn variant-filled">Log In</button>
			</a>
			<a href="/signup">
				<button type="button" class="btn variant-filled-secondary">Sign Up</button></a
			>
		{:else}
			<a href="/logout"> <button type="button" class="btn variant-filled">Logout</button></a>
			<a href={`/${$page.data.activeUser}`}
				><Avatar src="invalid-image.jpg" initials={$page.data.activeUser.username[0]} /></a
			>
		{/if}

		<LightSwitch />
	</svelte:fragment>
</AppBar>
<slot />
