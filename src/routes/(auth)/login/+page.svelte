<script lang="ts">
	import { page } from '$app/stores';
	import type { ActionData } from './$types';
	export let form: ActionData;

	let redirectPath = $page.url.searchParams.get('redirect');
</script>

<section class="border-gray-100 bg-gradient-to-r from-yellow-300 to-yellow-200">
	{#if redirectPath}
		<div class="bg-yellow-600 p-2 text-white">
			<p class="text-center">Login to return to last page.</p>
		</div>
	{/if}
	<div class="mx-auto flex flex-col items-center justify-center px-6 py-8 md:h-screen lg:py-0">
		<p class="text-5xl text-yellow-600">LOGO HERE</p>
		<div
			class="w-full rounded-lg bg-white shadow dark:border dark:border-gray-700 dark:bg-gray-800 sm:max-w-md md:mt-0 xl:p-0"
		>
			<div class="space-y-4 p-6 sm:p-8 md:space-y-6">
				<h1
					class="text-center text-xl font-bold leading-tight tracking-tight text-gray-900 dark:text-white md:text-2xl"
				>
					Login to Hoshii
				</h1>
				{#if form?.errors}
					<div>
						{#each form?.errors as formError}
							<p class="text-center text-red-500">{formError}</p>
						{/each}
					</div>
				{/if}
				<form class="space-y-4 md:space-y-6" method="POST">
					<label for="email" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
						>Email
						<input
							type="email"
							name="email"
							id="email"
							class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-gray-900 focus:border-red-500 focus:ring-red-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500 sm:text-sm"
							placeholder="name@company.com"
						/>
					</label>
					<label for="password" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
						>Password <input
							type="password"
							name="password"
							id="password"
							placeholder="••••••••"
							class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-gray-900 focus:border-red-500 focus:ring-red-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500 sm:text-sm"
						/>
					</label>

					<button
						type="submit"
						class="w-full rounded-lg bg-yellow-600 px-5 py-2.5 text-center text-sm font-medium text-white hover:bg-yellow-700 focus:outline-none focus:ring-4 focus:ring-yellow-300 dark:bg-yellow-600 dark:hover:bg-yellow-700 dark:focus:ring-yellow-800"
					>
						Log in
					</button>
					<p class="text-sm font-light text-gray-500 dark:text-gray-400">
						Forgot Password? <a
							href="/recovery"
							class="font-medium text-yellow-600 hover:underline dark:text-yellow-500">Help</a
						>
					</p>
				</form>
			</div>
		</div>
	</div>
</section>
