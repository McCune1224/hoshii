<script lang="ts">
	// Your selected Skeleton theme:
	import '@skeletonlabs/skeleton/themes/theme-skeleton.css';
	/* import '../theme.postcss'; */

	// This contains the bulk of Skeletons required styles:
	// NOTE: this will be renamed skeleton.css in the v2.x release.
	import '@skeletonlabs/skeleton/styles/skeleton.css';

	// Finally, your application's global stylesheet (sometimes labeled 'app.css')
	import '../app.postcss';

	import {
		AppBar,
		Avatar,
		LightSwitch,
		Modal,
		popup,
		type ModalComponent,
		type PopupSettings,
		ListBox,
		ListBoxItem
	} from '@skeletonlabs/skeleton';
	import type { LayoutServerData } from './$types';
	import { page } from '$app/stores';
	import hoshii_logo from '../assets/hoshii_logo.png';
	import ModalUpdateProfile from '$lib/components/ModalUpdateProfile.svelte';
	import Container from '$lib/components/Container.svelte';
	import { computePosition, autoUpdate, offset, shift, flip, arrow } from '@floating-ui/dom';
	import { storePopup } from '@skeletonlabs/skeleton';

	storePopup.set({ computePosition, autoUpdate, offset, shift, flip, arrow });
	const modalComponentRegistry: Record<string, ModalComponent> = {
		modalUpdateProfile: {
			ref: ModalUpdateProfile,
			props: {},
			slot: '<p>Modal Update Profile Slot</p>'
		}
	};

	let comboboxValue: string;
	let avatarInitials: string;

	export let data: LayoutServerData;

	const { activeUser } = data;
	if (activeUser) {
		avatarInitials = activeUser.username[0];
	} else {
		avatarInitials = '';
	}

	const popupCombobox: PopupSettings = {
		event: 'focus-click',
		target: 'popupCombobox',
		placement: 'bottom',
		closeQuery: '.listbox-item, .listbox-item a, a[href], button'
	};
</script>

<Modal components={modalComponentRegistry} />
<Container>
	<AppBar>
		<svelte:fragment slot="lead">
			<a class="flex flex-row items-center gap-4" href="/">
				<!-- <p class="text-4xl sm:text-6xl">🌟</p> -->
				<img alt="The project logo" height="100" width="100" src={hoshii_logo} />
				<h4 class="text-4xl sm:text-6xl">Hoshii</h4></a
			></svelte:fragment
		>
		<svelte:fragment slot="trail">
			{#if !activeUser}
				<a href="/login">
					<button type="button" class="btn variant-filled">Log In</button>
				</a>
				<a href="/signup">
					<button type="button" class="btn variant-filled-secondary">Sign Up</button></a
				>
			{:else}
				<button class="" use:popup={popupCombobox}>
					<Avatar src="invalid-image.jpg" initials={avatarInitials} />
					<span>
						<div class="card w-48 py-2 shadow-xl" data-popup="popupCombobox">
							<ListBox rounded="rounded-none">
								<ListBoxItem bind:group={comboboxValue} name="profile" value="profile">
									<a href={`/${activeUser.username}`}>Profile</a></ListBoxItem
								>
								<ListBoxItem bind:group={comboboxValue} name="logout" value="logout">
									<a href="/logout">Logout</a></ListBoxItem
								>
							</ListBox>
						</div>
					</span>
				</button>
			{/if}
			<LightSwitch />
		</svelte:fragment>
	</AppBar>
	<slot />
</Container>
