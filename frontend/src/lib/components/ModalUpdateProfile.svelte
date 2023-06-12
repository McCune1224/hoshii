<script lang="ts">
	// Props
	/** Exposes parent props to this component. */
	export let parent: any;
	let savePending = false;

	// Stores
	import { ProgressRadial, modalStore } from '@skeletonlabs/skeleton';

	// Form Data
	const formData = {
		display_name: '',
		bio: ''
	};

	// We've created a custom submit function to pass the response and close the modal.
	function onFormSubmit(): void {
		savePending = true;
		if ($modalStore[0].response) $modalStore[0].response(formData);
		console.log(formData);
		savePending = false;
		modalStore.close();
	}

	// Base Classes
	const cBase = 'card p-4 w-modal shadow-xl space-y-4';
	const cHeader = 'text-2xl font-bold';
	const cForm = 'border border-surface-500 p-4 space-y-4 rounded-container-token';
</script>

<!-- @component This example creates a simple form modal. -->

{#if $modalStore[0]}
	<div class="modal-example-form {cBase}">
		<header class={cHeader}>{$modalStore[0].title ?? '(title missing)'}</header>
		<article>{$modalStore[0].body ?? '(body missing)'}</article>
		<!-- Enable for debugging: -->
		<form class="modal-form {cForm}">
			<label class="label">
				<span>Display Name</span>
				<input
					class="input"
					type="text"
					bind:value={formData.display_name}
					placeholder="Enter name..."
				/>
			</label>
			<label class="label">
				<span>Bio</span>
				<textarea
					class="input textarea"
					rows="4"
					bind:value={formData.bio}
					placeholder="Enter Bio..."
				/>
			</label>
		</form>
		<!-- prettier-ignore -->
		<footer class="modal-footer {parent.regionFooter}">
        <button class="btn {parent.buttonNeutral}" on:click={parent.onClose}>{parent.buttonTextCancel}</button>
        <button class="btn {parent.buttonPositive}" on:click={onFormSubmit}>
            {#if savePending}
                <ProgressRadial size="small" />
            {:else}
                    Save
            {/if}
            </button>
        <button class="btn {parent.buttonTextCancel}"></button>
    </footer>
	</div>
{/if}
