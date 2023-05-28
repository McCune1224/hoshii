<script lang="ts">
	import type { ActionData } from './$types';
	export let form: ActionData;
</script>

<section class="">
	<div class="mx-auto flex flex-col items-center justify-center px-6 py-8 md:h-screen lg:py-0">
		<div class="w-full rounded-lg sm:max-w-md md:mt-0 xl:p-0">
			<div class="space-y-4 p-6 sm:p-8 md:space-y-6">
				{#if form?.errors}
					<aside class="variant-filled-error alert">
						{#each form?.errors as formError}
							<div class="alert-message">
								<h3 class="h3">ERROR:</h3>
								<p>{formError}</p>
							</div>
						{/each}
					</aside>
				{/if}
				<h1 class="text-center text-xl font-bold leading-tight tracking-tight md:text-2xl">
					Make an Account
				</h1>
				<form class="space-y-4 md:space-y-6" method="POST">
					<label for="username"
						><span>Username</span>
						<input
							class="input"
							type="text"
							name="username"
							id="username"
							placeholder="FooMcSmoo"
						/>
					</label>
					<label for="email"
						><span>Email</span>
						<input class="input" type="email" name="email" id="email" placeholder="email@foo.com" />
					</label>
					<label for="password"
						><span>Password</span>
						<input
							class="input"
							type="password"
							name="password"
							id="password"
							placeholder="••••••••"
						/>
					</label>
					<button
						type="submit"
						class="btn variant-filled-secondary w-full rounded-lg px-5 py-2.5 text-center text-sm font-medium focus:outline-none focus:ring-4"
					>
						Sign Up
					</button>
					<p class="text-sm font-light text-gray-500 dark:text-gray-400">
						Already have an account? <a
							href="/login"
							class="font-medium text-yellow-600 hover:underline dark:text-yellow-500">Log in</a
						>
					</p>
				</form>
			</div>
		</div>
	</div>
</section>
