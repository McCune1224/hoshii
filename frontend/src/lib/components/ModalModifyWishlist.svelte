<script lang="ts">
	import { modalStore } from '@skeletonlabs/skeleton/index';

	export let parent: any;
</script>

{#if $modalStore[0]}
	<div>
		<h1>Wishlist Modal</h1>
	</div>
{/if}
